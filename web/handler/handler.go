package handler

import "net/http"

type EngineHandler interface {
	GetGroup() string
	GetHandlers() []BasicHandler
}
type BasicHandler struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

func NewBasicHandler(method, path string, handlerFunc http.HandlerFunc) *BasicHandler {
	return &BasicHandler{
		Path:    path,
		Method:  method,
		Handler: handlerFunc,
	}
}
