package router

import (
	"github.com/ribeirosaimon/bootcamp/aula06/ex01/internal/product"
	"github.com/ribeirosaimon/bootcamp/aula06/ex01/internal/web/handler"
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
	eg.Patch("/{id}/is_published", controller.IsPublished)
	eg.Delete("/{id}", controller.DeleteProduct)

	return eg
}
