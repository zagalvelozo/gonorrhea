package server

import (
	"io"
	"log"
	"net/http"
)

const (
	// VERSION is the current version for the server.
	VERSION = "0.25.5"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, VERSION)
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
