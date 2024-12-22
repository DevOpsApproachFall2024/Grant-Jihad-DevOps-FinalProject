package internal

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()

	PingHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	expected := `{"message":"pong"}`
	if string(body) != expected {
		t.Errorf("Expected body %s, got %s", expected, string(body))
	}
}

func TestProjectsHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/projects", nil)
	w := httptest.NewRecorder()

	ProjectsHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	expected := "Projects endpoint"
	if string(body) != expected {
		t.Errorf("Expected body %s, got %s", expected, string(body))
	}
}
