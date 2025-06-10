package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

type BootcampServer interface {
	MapRoutes() http.Handler
}
type router struct {
}

func New() *router {
	return &router{}
}

func (rout *router) MapRoutes() http.Handler {
	chiRouter := chi.NewRouter()

	chiRouter.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.StripSlashes,
		middleware.Timeout(5*time.Second),
		middleware.Heartbeat("/ping"),
	)

	chiRouter.Route("/api/v1", func(r chi.Router) {
		r.Route("/products", func(rp chi.Router) {
			rp.Mount("/", buildProductRouter(chiRouter))
		})
		r.Route("/ping", func(rp chi.Router) {
			rp.Mount("/", buildHealth(chiRouter))
		})
	})
	return chiRouter
}
