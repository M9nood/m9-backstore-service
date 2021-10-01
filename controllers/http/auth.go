package http

import (
	controller "m9-backstore-service/controllers"

	"github.com/labstack/echo/v4"
)

type authHttpRoute struct {
	handler controller.AuthHandler
}

func NewAuthHttpRoute(authHandler controller.AuthHandler) authHttpRoute {
	return authHttpRoute{
		handler: authHandler,
	}
}

func (h authHttpRoute) Route(e *echo.Echo) {
	apiV1 := e.Group("/api/v1")
	authRoute := apiV1.Group("/auth")
	authRoute.POST("/register", h.handler.RegisterHandler)
	authRoute.POST("/login", h.handler.LoginHandler)
}
