package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	h := NewHealth()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	h.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("health route should return 200")
	}
	if w.Body.String() != "{\"status\":\"up\"}" {
		t.Errorf("health route should return a json with status: up")
	}
}
