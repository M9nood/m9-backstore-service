package http

import (
	controller "m9-backstore-service/controllers"

	"github.com/labstack/echo/v4"
)

type productHttpRoute struct {
	handler controller.ProductHandler
}

func NewProductHttpRoute(productHandler controller.ProductHandler) productHttpRoute {
	return productHttpRoute{
		handler: productHandler,
	}
}

func (h productHttpRoute) Route(e *echo.Echo) {
	apiV1 := e.Group("/api/v1")
	productRoute := apiV1.Group("/products")
	productRoute.GET("", h.handler.GetProductsHandler)
}
