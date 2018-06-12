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

// Response defines the structure of a HTTP JSON response to client.
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// NewImageRecognitionHandler returns a HTTP handler that will handle a request to perform image
// recognition.
func NewImageRecognitionHandler(labels map[int]string, modelPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		imgFile, header, err := r.FormFile("image")
		if err != nil {
			response := &Response{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			}

			if resBytes, err := json.Marshal(response); err == nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write(resBytes)
			}
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
			classList := make([]Class, 0, len(softmaxScore[0]))
			for idx, prob := range softmaxScore[0] {
				classList = append(classList, Class{Prob: prob, Index: idx})
			}

			// Perform sorting
			Sort(classList, 0, len(classList)-1)

			message := fmt.Sprintf("Most probable classes: ")
			for i := len(classList) - 1; i > len(classList)-6; i-- {
				message += fmt.Sprintf(" %s ", labels[classList[i].Index])
			}

			renderResponse(w, http.StatusOK, message)
			return
		}

		renderResponse(w, http.StatusInternalServerError, "failed to run TF model")
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

func renderResponse(w http.ResponseWriter, status int, msg string) {
	res := &Response{
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
