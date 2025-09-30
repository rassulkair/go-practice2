package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetUser(w, r)
	case http.MethodPost:
		handlePostUser(w, r)
	default:
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || idStr == "" {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}
	respondJSON(w, http.StatusOK, map[string]int{"user_id": id})
}

func handlePostUser(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Name string `json:"name"`
	}
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&payload); err != nil {
		respondError(w, http.StatusBadRequest, "invalid name")
		return
	}
	if strings.TrimSpace(payload.Name) == "" {
		respondError(w, http.StatusBadRequest, "invalid name")
		return
	}
	respondJSON(w, http.StatusCreated, map[string]string{"created": payload.Name})
}

func respondJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func respondError(w http.ResponseWriter, status int, msg string) {
	respondJSON(w, status, map[string]string{"error": msg})
}
