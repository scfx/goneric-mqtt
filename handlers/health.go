package handlers

import "net/http"

type Health struct{}

// Standard Health handler, that returns a 200 status code and a json with status: up
func (h Health) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"up"}`))
}

func NewHealth() *Health {
	return &Health{}
}
