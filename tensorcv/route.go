// Package tensorcv provides computer vision handlers for handling image recognition requests.
package tensorcv

import (
	"net/http"

	"github.com/gorilla/mux"
)

// LoadRoutes returns a http.Handler as a multiplexer to various routes.
func LoadRoutes(labels map[int]string, modelPath string) http.Handler {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.Handle("/tf/recognition/", NewImageRecognitionHandler(labels, modelPath)).Methods("POST")
	api.Handle("/tf/hello/", NewHelloWorldHandler()).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))

	return r
}
