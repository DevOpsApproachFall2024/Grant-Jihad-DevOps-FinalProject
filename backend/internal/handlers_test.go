package internal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/ping", nil)
    w := httptest.NewRecorder()

    PingHandler(w, req)

    if w.Result().StatusCode != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
    }

    body, _ := io.ReadAll(w.Body)
    expected := `{"message":"pong"}`
    if string(body) != expected {
        t.Errorf("Expected body %s, got %s", expected, string(body))
    }
}

// New test for /projects endpoint
func TestProjectsHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/projects", nil)
	w := httptest.NewRecorder()

	// Call the ProjectsHandler directly since we're testing the handler logic
	ProjectsHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d for /projects, got %d", http.StatusOK, resp.StatusCode)
	}

	// Check if response body has expected JSON structure
	var projects []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		t.Fatalf("Failed to decode JSON response for /projects: %v", err)
	}

	if len(projects) == 0 || projects[0].Name != "Example Project" {
		t.Errorf("Expected at least one project named 'Example Project', got %v", projects)
	}
}
