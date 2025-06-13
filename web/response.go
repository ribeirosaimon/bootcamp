package web

import (
	"encoding/json"
	"net/http"
	"reflect"
)

type response struct {
	Status int  `json:"status"`
	Data   any  `json:"data"`
	Count  uint `json:"count,omitempty"`
}

func (r *response) Build(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	if r.Data != nil {
		json.NewEncoder(w).Encode(r)
	}
}

type responseOptions func(*response)

func WithStatus(status int) responseOptions {
	return func(r *response) {
		r.Status = status
	}
}

func WithData(data any) responseOptions {
	return func(r *response) {
		r.Data = data
	}
}

func NewResponse(opt ...responseOptions) *response {
	resp := response{}

	for _, o := range opt {
		o(&resp)
	}
	if resp.Data != nil && isSliceOrArray(resp.Data) {
		resp.Count = uint(reflect.ValueOf(resp.Data).Len())
	}
	return &resp
}

func isSliceOrArray(x any) bool {
	k := reflect.TypeOf(x).Kind()
	return k == reflect.Slice || k == reflect.Array
}
