package server

import (
	"github.com/ribeirosaimon/bootcamp/aula07/ex01/internal/web/router"
)

type server struct {
	btServer router.BootcampServer
}

func New(b router.BootcampServer) *server {
	return &server{
		btServer: b,
	}
}
