package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zagalvelozo/gonorrhea/api"
	"github.com/zagalvelozo/gonorrhea/util"
)

const (
	VERSION = "0.25.5"
)

type VersionResponse struct {
	Version string `json:"version"`
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(&VersionResponse{Version: VERSION})
	w.Write(res)
}

func Serve() {
	store := api.NewStore()
	dataDir := util.GetEnv("DATA_DIR", "./data")

	mux := http.NewServeMux()

	// Version endpoint.
	mux.HandleFunc("/", getRoot)

	// Auth routes (anonymous, register, login, logout).
	api.RegisterAuthRoutes(mux, store)

	// Chat routes (messages + audio upload).
	api.RegisterChatRoutes(mux, store, dataDir)

	// Serve uploaded audio files.
	filesDir := http.Dir(dataDir)
	mux.HandleFunc("/files/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}
		http.StripPrefix("/files/", http.FileServer(filesDir)).ServeHTTP(w, r)
	})

	port := util.GetEnv("PORT", "3333")
	log.Printf("server listening on :%s", port)

	handler := api.CORS(mux)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("error starting server: %s\n", err)
	}
}
