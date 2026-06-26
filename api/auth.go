package api

import (
	"encoding/json"
	"net/http"
)

type authHandler struct {
	store *Store
}

// RegisterAuthRoutes mounts auth endpoints on the given mux.
func RegisterAuthRoutes(mux *http.ServeMux, store *Store) {
	h := &authHandler{store: store}
	mux.HandleFunc("/api/auth/anonymous", h.anonymous)
	mux.HandleFunc("/api/auth/register", h.register)
	mux.HandleFunc("/api/auth/login", h.login)
	mux.HandleFunc("/api/auth/logout", h.logout)
}

func (h *authHandler) anonymous(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	user := h.store.CreateUser("")
	writeJSON(w, http.StatusOK, user)
}

func (h *authHandler) register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Email == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"message": "email is required"})
		return
	}
	if existing := h.store.FindUserByEmail(req.Email); existing != nil {
		writeJSON(w, http.StatusConflict, map[string]string{"message": "email already registered"})
		return
	}
	user := h.store.CreateUser(req.Email)
	writeJSON(w, http.StatusCreated, user)
}

func (h *authHandler) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Email == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"message": "email is required"})
		return
	}
	user := h.store.FindUserByEmail(req.Email)
	if user == nil {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"message": "invalid credentials"})
		return
	}
	writeJSON(w, http.StatusOK, user)
}

func (h *authHandler) logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

