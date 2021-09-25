package route

import (
	"net/http"

	controller "m9-backstore-service/controllers"

	"github.com/labstack/echo/v4"
)

func RouterSetup() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "bkst, OK!")
	})

	products := e.Group("/products")
	products.GET("", controller.GetProductsHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
