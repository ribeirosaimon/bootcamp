package main

import (
	"github.com/ribeirosaimon/bootcamp/aula06/ex01/internal/config"
	"github.com/ribeirosaimon/bootcamp/aula06/ex01/internal/web/router"
	"github.com/ribeirosaimon/bootcamp/aula06/ex01/internal/web/server"
	"os"

	"net/http"
)

func main() {
	os.Setenv("SECRET_KEY", "bootcamp")

	config.New()
	routers := router.New()
	server.New(routers)

	if err := http.ListenAndServe(":8080", routers.MapRoutes()); err != nil {
		panic(err)
	}

}
