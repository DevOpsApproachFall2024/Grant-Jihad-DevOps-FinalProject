package internal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestPingHandler tests the /ping endpoint
func TestPingHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()

	// Call the PingHandler directly
	PingHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	// Check for the correct status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Check for the correct Content-Type header
	if contentType := resp.Header.Get("Content-Type"); contentType != "application/json" {
		t.Errorf("Expected Content-Type 'application/json', got '%s'", contentType)
	}

	// Check the response body
	var pingResponse PingResponse
	if err := json.NewDecoder(resp.Body).Decode(&pingResponse); err != nil {
		t.Fatalf("Failed to decode JSON response for /ping: %v", err)
	}

	if pingResponse.Message != "pong" {
		t.Errorf("Expected message 'pong', got '%s'", pingResponse.Message)
	}
}

// TestProjectsHandler tests the /projects endpoint
func TestProjectsHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/projects", nil)
	w := httptest.NewRecorder()

	// Call the ProjectsHandler directly
	ProjectsHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	// Check for the correct status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d for /projects, got %d", http.StatusOK, resp.StatusCode)
	}

	// Check for the correct Content-Type header
	if contentType := resp.Header.Get("Content-Type"); contentType != "application/json" {
		t.Errorf("Expected Content-Type 'application/json', got '%s'", contentType)
	}

	// Check the response body
	var projects []Project
	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		t.Fatalf("Failed to decode JSON response for /projects: %v", err)
	}

	// Validate the project data
	if len(projects) == 0 {
		t.Errorf("Expected at least one project, got none")
	} else if projects[0].Name != "Example Project" {
		t.Errorf("Expected project name 'Example Project', got '%s'", projects[0].Name)
	}

	// Additional check for project ID
	if projects[0].ID <= 0 {
		t.Errorf("Expected positive project ID, got %d", projects[0].ID)
	}
}
