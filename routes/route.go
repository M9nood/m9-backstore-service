package route

import (
	"net/http"
	"os"

	controller "m9-backstore-service/controllers"

	"github.com/labstack/echo/v4"
)

func RouterSetup() {
	port := os.Getenv("PORT")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "bkst, OK!")
	})

	products := e.Group("/products")
	products.GET("", controller.GetProductsHandler)

	e.Logger.Fatal(e.Start(":" + port))
}
