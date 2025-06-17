package handler

import (
	"github.com/ribeirosaimon/bootcamp/web/response"
	"net/http"
)

type health struct {
	group    string
	handlers []BasicHandler
}

func (h *health) GetHandlers() []BasicHandler {
	return h.handlers
}

func (h *health) GetGroup() string {
	return h.group
}

func NewHealth() *health {
	var h health
	return &health{
		group: "/health",
		handlers: []BasicHandler{
			{Method: http.MethodGet, Path: "/", Handler: h.Check()},
		},
	}
}

func (h *health) Check() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, struct {
			Status string
		}{
			Status: "ok",
		})
		return
	}
}
