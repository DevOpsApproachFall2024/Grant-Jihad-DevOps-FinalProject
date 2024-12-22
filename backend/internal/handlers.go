package internal

import (
	"encoding/json"
	"net/http"
)

// PingResponse represents the response structure for the /ping endpoint.
type PingResponse struct {
	Message string `json:"message"`
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
    // Intentional bug: sending a malformed JSON response
    response := "This is not JSON" // This should be a struct encoded as JSON
    w.Header().Set("Content-Type", "application/json")
    _, _ = w.Write([]byte(response))
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
