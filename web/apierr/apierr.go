package apierr

import (
	"encoding/json"
	"errors"
	"net/http"
)

type apiErr string

const (
	errNotFound   apiErr = "not found"
	errBadRequest apiErr = "bad request"
	errInternal   apiErr = "internal"
	errConflict   apiErr = "conflict"
)

type apiError struct {
	Message apiErr `json:"message"`
	Status  int    `json:"status"`
}

func (e apiError) Error() string {
	return string(e.Message)
}

func NewApiErr(oldErr error, w http.ResponseWriter) {
	var berr apiError
	switch {
	case errors.As(oldErr, &berr):
		berr.build(w)
	default:
		createBasicErr(errInternal, http.StatusInternalServerError).build(w)
	}
}

func (e apiError) build(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(e.Status)
	json.NewEncoder(w).Encode(apiError{Message: e.Message, Status: e.Status})
}

type apiErrorOptions func(*apiError)

func WithStatus(status int) apiErrorOptions {
	return func(e *apiError) {
		e.Status = status
	}
}

func WithMessage(message apiErr) apiErrorOptions {
	return func(e *apiError) {
		e.Message = message
	}
}

func NewBadRequestApiErr(opt ...apiErrorOptions) apiError {
	return createBasicErr(errBadRequest, http.StatusBadRequest, opt...)
}

func NewNotFoundApiErr(opt ...apiErrorOptions) apiError {
	return createBasicErr(errNotFound, http.StatusNotFound, opt...)
}

func NewConflictApiErr(opt ...apiErrorOptions) apiError {
	return createBasicErr(errConflict, http.StatusConflict, opt...)
}
func createBasicErr(message apiErr, status int, opt ...apiErrorOptions) apiError {
	v := apiError{
		Message: message,
		Status:  status,
	}
	for _, o := range opt {
		o(&v)
	}
	return v
}
