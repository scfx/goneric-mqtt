package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

type Counter struct {
	Count float64 `json:"count"`
}

// Standard Health handler, that returns a 200 status code and a json with status: up
func (c *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := c.Encode(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Counter) Encode(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(c)
}

func (c *Counter) AddOne() {
	c.Count++
}

func NewCounter() *Counter {
	return &Counter{
		Count: 0,
	}
}
