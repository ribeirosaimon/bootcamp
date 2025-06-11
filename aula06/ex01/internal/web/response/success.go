package response

import (
	"encoding/json"
	"github.com/ribeirosaimon/bootcamp/aula06/ex01/internal/web/response/apperror"
	"net/http"
)

type successResponse struct {
	Message string `json:"message" omitempty`
	Data    any    `json:"data"`
	Status  int    `json:"status"`
}

type optResp func(s *successResponse)

func WithMessage(msg string) optResp {
	return func(s *successResponse) {
		s.Message = msg
	}
}

func WithData(data any) optResp {
	return func(s *successResponse) {
		s.Data = data
	}
}

func WithStatus(status int) optResp {
	return func(s *successResponse) {
		s.Status = status
	}
}

func Success(opts ...optResp) *successResponse {
	resp := &successResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    nil,
	}
	for _, opt := range opts {
		opt(resp)
	}
	return resp
}

func (s *successResponse) Build(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(s.Status)
	if err := json.NewEncoder(w).Encode(successResponse{
		Message: s.Message,
		Data:    s.Data,
		Status:  s.Status,
	}); err != nil {
		apperror.NewGenericError().Build(w)
	}
}
