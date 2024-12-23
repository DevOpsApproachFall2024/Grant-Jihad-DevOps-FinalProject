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
func PingHandler(w http.ResponseWriter, r *http.Request) {
	response := PingResponse{Message: "pong"}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

// Project represents a project in the /projects endpoint response.
type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ProjectsHandler handles requests to the /projects endpoint.
func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	projects := []Project{
		{ID: 1, Name: "Example Project"},
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(projects)
}
