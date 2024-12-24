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
	_ = json.NewEncoder(w).Encode(response)
}

// ProjectsHandler handles requests to the /projects endpoint.
func ProjectsHandler(w http.ResponseWriter, _ *http.Request) {
	projects := []map[string]interface{}{
		{"id": 1, "name": "Example Project"},
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(projects)
}
