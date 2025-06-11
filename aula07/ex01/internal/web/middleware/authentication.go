package middleware

import (
	"github.com/ribeirosaimon/bootcamp/aula07/ex01/internal/config"
	"github.com/ribeirosaimon/bootcamp/aula07/ex01/internal/web/response/apperror"
	"net/http"
	"strings"
)

func AppAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		auth = strings.TrimPrefix(auth, "Bearer ")
		if auth != config.GetSecretKey() {
			apperror.NewError(
				apperror.NewAuthenticationError(),
			).Build(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
