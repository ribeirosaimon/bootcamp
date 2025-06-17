package web

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ribeirosaimon/bootcamp/web/handler"
	"log"
	"net/http"
	"time"
)

type Router interface {
	Start()
}

type router struct {
	handlers []handler.EngineHandler
}

func NewRouters(
	routes ...handler.EngineHandler,
) *router {
	return &router{
		handlers: routes,
	}
}

func (r *router) Start(chiRouter *chi.Mux) {

	chiRouter.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.StripSlashes,
		middleware.Timeout(5*time.Second),
	)

	// this mount only works once in the life cycle
	chiRouter.Route("/api/v1", func(chiRouter chi.Router) {
		for indexHandlers := range r.handlers {
			basicHandler := r.handlers[indexHandlers]
			chiRouter.Route(basicHandler.GetGroup(), func(groupHandler chi.Router) {
				for index := range basicHandler.GetHandlers() {
					singleBasicHandler := basicHandler.GetHandlers()[index]
					groupHandler.Method(singleBasicHandler.Method, singleBasicHandler.Path, singleBasicHandler.Handler)
				}
			})
		}
	})

	chi.Walk(chiRouter, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s", method, route)
		return nil
	})
}
