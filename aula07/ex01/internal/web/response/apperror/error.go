package apperror

import (
	"encoding/json"
	"errors"
	"net/http"
)

type bError string

const (
	ErrEntityNotFound        bError = "entity not found"
	ErrResourceNotExists     bError = "resource does not exist"
	ErrInternalError         bError = "internal error"
	ErrValidation            bError = "validation error"
	ErrResourceAlreadyExists bError = "resource already exists"
	ErrUnauthorized          bError = "unauthorized"
	ErrForbidden             bError = "forbidden"
)

type bootcampError struct {
	Message bError `json:"message"`
	Status  int    `json:"code"`
}

func NewAppErrorNotFound() bootcampError {
	return bootcampError{
		Message: ErrEntityNotFound,
		Status:  http.StatusNotFound,
	}
}

func NewGenericError() bootcampError {
	return bootcampError{
		Message: ErrInternalError,
		Status:  http.StatusInternalServerError,
	}
}

func NewAuthenticationError() bootcampError {
	return bootcampError{
		Message: ErrUnauthorized,
		Status:  http.StatusUnauthorized,
	}
}

type bootcampErrorOpt func(b *bootcampError)
type WithMessage func(b *bootcampError)

func NewError(oldError error, opt ...bootcampErrorOpt) *bootcampError {
	var berr bootcampError
	switch {
	case errors.As(oldError, &berr):
		berr = bootcampError{Status: berr.Status, Message: berr.Message}
	default:
		berr = bootcampError{Status: http.StatusInternalServerError, Message: "internal server error"}
	}

	for _, o := range opt {
		o(&berr)
	}
	return &berr
}

func (e bootcampError) Build(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(e.Status)
	json.NewEncoder(w).Encode(bootcampError{Message: e.Message, Status: e.Status})
}

func (e bootcampError) Error() string {
	return string(e.Message)
}
