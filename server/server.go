package server

import (
	"io"
	"log"
	"net/http"
)

type VersionResponse struct {
	Version string `json:"version"`
}

const (
	// VERSION is the current version for the server.
	VERSION = "0.25.5"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	io.WriteString(w, "{}")
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	res, _ := json.Marshal(&VersionResponse{Version: VERSION})
	io.WriteString(w, string(res))
}

func Serve() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.healthCheck("/health", healthCheck)

	err := http.ListenAndServe(":3333", mux)
	if err != nil {
		log.Fatalf("error starting server: %s\n", err)
	}
}
