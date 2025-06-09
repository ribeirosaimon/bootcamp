package health

import (
	"encoding/json"
	"net/http"
)

type health struct {
}

type Handler interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

func NewHealth() *health {
	return &health{}
}

func (h *health) Ping(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("pong")
	return
}
