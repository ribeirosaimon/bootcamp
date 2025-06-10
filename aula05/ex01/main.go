package main

import (
	"github.com/ribeirosaimon/bootcamp/aula05/ex01/web/router"

	"net/http"
)

type server struct {
	btServer router.BootcampServer
}

func NewServer(b router.BootcampServer) *server {
	return &server{
		btServer: b,
	}
}

func main() {
	routers := router.New()
	NewServer(routers)

	if err := http.ListenAndServe(":8080", routers.MapRoutes()); err != nil {
		panic(err)
	}

}
