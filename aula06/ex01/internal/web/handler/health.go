package handler

import (
	"encoding/json"
	"net/http"
)

type health struct {
}

type Health interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

func NewHealth() *health {
	return &health{}
}

func (h *health) Ping(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("pong")
	return
}
