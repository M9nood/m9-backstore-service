package http

import (
	controller "m9-backstore-service/controllers"
	auth "m9-backstore-service/controllers/middleware"

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
	productRoute := apiV1.Group("/products", auth.IsAuth)
	productRoute.GET("", h.handler.GetProductsHandler)
	productRoute.POST("", h.handler.CreateProductHandler)

}
