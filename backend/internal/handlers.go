package internal

import (
	"encoding/json"
	"net/http"
)

// PingResponse represents the response structure for the /ping endpoint.
type PingResponse struct {
	Message string `json:"message"`
}

// PingHandler handles requests to the /ping endpoint.
func PingHandler(w http.ResponseWriter, _ *http.Request) {
	response := PingResponse{Message: "pong"}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// ProjectsHandler handles requests to the /projects endpoint.
func ProjectsHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("Projects endpoint")); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
