package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// RegisterChatRoutes mounts /api/chats/ — handles both messages and audio uploads.
func RegisterChatRoutes(mux *http.ServeMux, store *Store, dataDir string) {
	h := &chatHandler{store: store, dataDir: dataDir}
	mux.HandleFunc("/api/chats/", h.route)
}

type chatHandler struct {
	store   *Store
	dataDir string
}

func (h *chatHandler) route(w http.ResponseWriter, r *http.Request) {
	// /api/chats/{chatID}/messages  or  /api/chats/{chatID}/audio/{msgID}
	path := strings.TrimPrefix(r.URL.Path, "/api/chats/")
	parts := strings.Split(path, "/")

	if len(parts) < 2 {
		http.NotFound(w, r)
		return
	}

	chatID := parts[0]
	segment := parts[1]

	switch {
	case segment == "messages":
		h.handleMessages(w, r, chatID)
	case segment == "audio" && len(parts) >= 3:
		h.handleAudio(w, r, chatID, parts[2])
	default:
		http.NotFound(w, r)
	}
}

// --- Messages ---

func (h *chatHandler) handleMessages(w http.ResponseWriter, r *http.Request, chatID string) {
	switch r.Method {
	case http.MethodGet:
		h.getMessages(w, r, chatID)
	case http.MethodPost:
		h.postMessage(w, r, chatID)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *chatHandler) getMessages(w http.ResponseWriter, r *http.Request, chatID string) {
	limit := 10
	if l := r.URL.Query().Get("limit"); l != "" {
		if n, err := strconv.Atoi(l); err == nil && n > 0 {
			limit = n
		}
	}
	msgs := h.store.GetMessages(chatID, limit)
	if msgs == nil {
		msgs = []Message{}
	}
	writeJSON(w, http.StatusOK, msgs)
}

func (h *chatHandler) postMessage(w http.ResponseWriter, r *http.Request, chatID string) {
	var req struct {
		Text      string `json:"text"`
		Sender    string `json:"sender"`
		CreatedAt int64  `json:"createdAt"`
		AudioURL  string `json:"audioURL"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"message": "invalid request body"})
		return
	}
	msg := h.store.AddMessage(chatID, req.Text, req.Sender, req.AudioURL)
	writeJSON(w, http.StatusCreated, msg)
}

// --- Audio upload ---

func (h *chatHandler) handleAudio(w http.ResponseWriter, r *http.Request, chatID, msgID string) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(50 << 20); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"message": "failed to parse upload"})
		return
	}

	file, header, err := r.FormFile("audio")
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"message": "missing audio field"})
		return
	}
	defer file.Close()

	dir := filepath.Join(h.dataDir, "audio", chatID)
	if err := os.MkdirAll(dir, 0755); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"message": "failed to create storage dir"})
		return
	}

	filename := header.Filename
	if filename == "" {
		filename = msgID + ".wav"
	}
	dstPath := filepath.Join(dir, filename)

	dst, err := os.Create(dstPath)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"message": "failed to save file"})
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"message": "failed to write file"})
		return
	}

	url := fmt.Sprintf("/files/audio/%s/%s", chatID, filename)
	writeJSON(w, http.StatusOK, map[string]string{"url": url})
}

// --- helpers ---

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
