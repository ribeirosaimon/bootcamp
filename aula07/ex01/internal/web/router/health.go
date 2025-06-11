package router

import (
	"github.com/ribeirosaimon/bootcamp/aula07/ex01/internal/web/handler"
	"net/http"
)

func buildHealth(eg BootcampRouter) http.Handler {
	health := handler.NewHealth()
	eg.Get("/", health.Ping)
	return eg
}
