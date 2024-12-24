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
// Introduced Bug: The project ID is intentionally set to an invalid value (-1).
// func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
// 	projects := []Project{
// 		{ID: -1, Name: "Bugged Project"}, // Intentional Bug: Negative ID
// 	}

// 	// Forgot to set Content-Type header (Bug!)
// 	_ = json.NewEncoder(w).Encode(projects)
// }

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
    projects := []Project{
        {ID: 1, Name: "Example Project"}, // Correct Project ID and Name
    }

    w.Header().Set("Content-Type", "application/json") // Correct Content-Type
    _ = json.NewEncoder(w).Encode(projects)
}
