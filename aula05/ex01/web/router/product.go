package router

import (
	"github.com/ribeirosaimon/bootcamp/aula05/ex01/product"
	"github.com/ribeirosaimon/bootcamp/aula05/ex01/web/handler"
	"net/http"
)

func buildProductRouter(eg BootcampRouter) http.Handler {
	repository := product.NewRepository()
	service := product.NewService(repository)
	controller := handler.NewProduct(service)

	eg.Get("/", controller.GetProducts)
	eg.Get("/{id}", controller.GetProductById)
	eg.Get("/search", controller.SearchProducts)
	eg.Post("/", controller.SaveProduct)
	eg.Put("/", controller.UpdateProduct)
	eg.Patch("/{id}/is-published", controller.IsPublished)
	eg.Delete("/{id}", controller.DeleteProduct)

	return eg
}
