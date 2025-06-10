package router

import (
	"github.com/ribeirosaimon/bootcamp/aula05/ex01/web/handler"
	"net/http"
)

func buildHealth(eg BootcampRouter) http.Handler {
	eg.Get("/", handler.NewHealth().Ping)
	return eg
}
