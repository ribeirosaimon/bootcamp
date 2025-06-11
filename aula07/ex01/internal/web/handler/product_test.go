package handler

import (
	"bytes"
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/ribeirosaimon/bootcamp/aula07/ex01/internal/product"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProduct(t *testing.T) {
	repository := product.NewRepository(
		product.WithPath("products_test.json"),
	)
	service := product.NewService(repository)
	testProductHandler := NewProduct(service)

	for _, tt := range []struct {
		handler        http.HandlerFunc
		name           string
		bodyReceived   string
		bodyToSend     string
		httpMethod     string
		responseStatus int
		urlParams      map[string]string
		withError      bool
		auxFunc        func(*testing.T)
	}{
		{
			name:           "Get Product",
			responseStatus: http.StatusOK,
			httpMethod:     http.MethodGet,
			handler:        testProductHandler.GetProducts,
			bodyReceived:   getResp,
		},
		{
			name:           "Get Product",
			responseStatus: http.StatusCreated,
			httpMethod:     http.MethodGet,
			handler:        testProductHandler.GetProducts,
			withError:      true,
			bodyReceived:   getResp,
		},
		{
			name:           "Get Product by Id",
			responseStatus: http.StatusOK,
			httpMethod:     http.MethodGet,
			handler:        testProductHandler.GetProductById,
			urlParams:      map[string]string{"id": "1"},
			bodyReceived:   getSingleId,
		},
		{
			name:           "Get Product by Id",
			responseStatus: http.StatusCreated,
			httpMethod:     http.MethodGet,
			withError:      true,
			handler:        testProductHandler.GetProductById,
			urlParams:      map[string]string{"id": "1"},
			bodyReceived:   getSingleId,
		},
		{
			name:           "Post Product",
			responseStatus: http.StatusCreated,
			httpMethod:     http.MethodPost,
			handler:        testProductHandler.SaveProduct,
			bodyToSend:     "{\n        \"name\": \"Meu produto\",\n        \"quantity\": 1000,\n        \"code_value\": \"S89049F\",\n        \"is_published\": false,\n        \"expiration\": \"09/12/2021\",\n        \"price\": 1000.00\n    }",
			bodyReceived:   postResp,
		},
		{
			name:           "Post Product",
			responseStatus: http.StatusOK,
			httpMethod:     http.MethodPost,
			withError:      true,
			handler:        testProductHandler.SaveProduct,
			bodyToSend:     "{\n        \"name\": \"Meu produto\",\n        \"quantity\": 1000,\n        \"code_value\": \"S89049F\",\n        \"is_published\": false,\n        \"expiration\": \"09/12/2021\",\n        \"price\": 1000.00\n    }",
			bodyReceived:   postResp2,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			rctx := chi.NewRouteContext()
			for k, v := range tt.urlParams {
				rctx.URLParams.Add(k, v)
			}

			req := httptest.NewRequest(tt.httpMethod, "/", bytes.NewReader([]byte(tt.bodyToSend)))
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			req.Header.Set("Accept", "application/json")
			w := httptest.NewRecorder()

			tt.handler(w, req)

			assert.JSONEq(t, tt.bodyReceived, w.Body.String())
			if tt.withError {
				assert.NotEqual(t, tt.responseStatus, w.Code)
			} else {
				assert.Equal(t, tt.responseStatus, w.Code)
			}
			if tt.auxFunc != nil {
				tt.auxFunc(t)
			}
		})
	}
}

var getSingleId = `
{
  "message" : "success",
  "data" : {
    "id" : 1,
    "name" : "Oil - Margarine",
    "quantity" : 439,
    "code_value" : "S82254D",
    "is_published" : true,
    "expiration" : "15/12/2021",
    "price" : 71.42
  },
  "status" : 200
}`

var getResp = `{
  "message" : "success",
  "data" : [ {
    "id" : 1,
    "name" : "Oil - Margarine",
    "quantity" : 439,
    "code_value" : "S82254D",
    "is_published" : true,
    "expiration" : "15/12/2021",
    "price" : 71.42
  }],
  "status" : 200
}`

var postResp = `
{
  "message" : "success",
  "data" : {
    "id" : 2,
    "name" : "Meu produto",
    "quantity" : 1000,
    "code_value" : "S89049F",
    "is_published" : false,
    "expiration" : "09/12/2021",
    "price" : 1000
  },
  "status" : 201
}`

var postResp2 = `
{
  "message" : "success",
  "data" : {
    "id" : 3,
    "name" : "Meu produto",
    "quantity" : 1000,
    "code_value" : "S89049F",
    "is_published" : false,
    "expiration" : "09/12/2021",
    "price" : 1000
  },
  "status" : 201
}`
