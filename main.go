package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/log"
)

type Resource struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var resources []Resource

func main() {
	logger := log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Log("message", "Received request")

		switch r.Method {
		case "GET":
			json.NewEncoder(w).Encode(resources)
		case "POST":
			var resource Resource
			err := json.NewDecoder(r.Body).Decode(&resource)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			resources = append(resources, resource)
			w.WriteHeader(http.StatusCreated)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Listening on http://localhost:7000")
	http.ListenAndServe(":7000", nil)

}
