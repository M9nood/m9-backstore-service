package http

import (
	controller "m9-backstore-service/controllers"

	"github.com/labstack/echo/v4"
)

type httpRoute struct {
	handler controller.Handler
}

func NewProductHttpRoute(productHandler controller.Handler) httpRoute {
	return httpRoute{
		handler: productHandler,
	}
}

func (h httpRoute) Route(e *echo.Echo) {
	apiV1 := e.Group("/api/v1")
	productRoute := apiV1.Group("/products")
	productRoute.GET("", h.handler.GetProductsHandler)
}
