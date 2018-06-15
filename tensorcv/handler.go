// Package tensorcv provides computer vision handlers for handling image recognition requests.
package tensorcv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// ErrorRes defines the structure of a HTTP error JSON response to client.
type ErrorRes struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// SuccessRes defines the structures of a HTTP success JSON response to client.
type SuccessRes struct {
	Status  int      `json:"status"`
	Results []*Class `json:"results"`
}

// NewImageRecognitionHandler returns a HTTP handler that will handle a request to perform image
// recognition.
func NewImageRecognitionHandler(labels map[int]string, modelPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		imgFile, header, err := r.FormFile("image")
		if err != nil {
			renderErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		imgName := strings.Split(header.Filename, ".")

		var imgBuffer bytes.Buffer
		io.Copy(&imgBuffer, imgFile)
		fmt.Printf("Received image %s which as %d bytes\n", imgName, len(imgBuffer.Bytes()))

		var imgFormat string
		if imgName[1] == "jpeg" || imgName[1] == "jpg" {
			imgFormat = "jpeg"
		} else {
			imgFormat = "png"
		}

		imgTensor, err := GetTensorFromImageBuffer(imgBuffer, imgFormat, 3)
		fmt.Println("Image tensor is loaded:", imgTensor.Shape())

		softmaxScore := RunResNetModel(imgTensor, modelPath)
		if softmaxScore != nil {
			classList := make([]*Class, 0, len(softmaxScore[0]))
			for idx, prob := range softmaxScore[0] {
				classList = append(classList, &Class{Prob: prob, Index: idx})
			}

			// Perform sorting
			Sort(classList, 0, len(classList)-1)

			top5Results := classList[len(classList)-6 : len(classList)-1]

			for _, result := range top5Results {
				result.Name = labels[result.Index]
			}

			renderResult(w, top5Results)
			return
		}

		renderErrorResponse(w, http.StatusInternalServerError, "failed to run TF model")
	}
}

// NewHelloWorldHandler returns a HTTP handler that will return a hello world message from
// tensorflow to client.
func NewHelloWorldHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		msg := HelloWorldFromTF()

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(msg))
	}
}

func renderResult(w http.ResponseWriter, results []*Class) {
	res := &SuccessRes{
		Status:  http.StatusOK,
		Results: results,
	}

	if resBytes, err := json.Marshal(res); err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(resBytes)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func renderErrorResponse(w http.ResponseWriter, status int, msg string) {
	res := &ErrorRes{
		Status:  status,
		Message: msg,
	}

	if resBytes, err := json.Marshal(res); err == nil {
		w.WriteHeader(status)
		w.Write(resBytes)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
