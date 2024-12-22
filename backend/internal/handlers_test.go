package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestPingHandler tests the PingHandler logic.
func TestPingHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()

	PingHandler(w, req)

	// Validate status code
	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}

	// Validate response body
	body, _ := io.ReadAll(w.Body)
	expected := `{"message":"pong"}`
	if string(body) != expected {
		t.Errorf("Expected body %s, got %s", expected, string(body))
	}
}

// TestProjectsHandler tests the ProjectsHandler logic.
func TestProjectsHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/projects", nil)
	w := httptest.NewRecorder()

	// Call the ProjectsHandler directly since we're testing the handler logic
	ProjectsHandler(w, req)

	resp := w.Result()

	// Validate status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d for /projects, got %d", http.StatusOK, resp.StatusCode)
	}

	// Validate response body structure
	var projects []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		t.Fatalf("Failed to decode JSON response for /projects: %v", err)
	}

	// Validate response content
	if len(projects) == 0 {
		t.Errorf("Expected at least one project, but got none")
	} else if projects[0].Name != "Example Project" {
		t.Errorf("Expected first project name to be 'Example Project', got '%s'", projects[0].Name)
	}
}
