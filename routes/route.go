package route

import (
	"net/http"
	"os"

	controller "m9-backstore-service/controllers"

	productHttp "m9-backstore-service/controllers/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func RouterSetup(e *echo.Echo, db *gorm.DB) *echo.Echo {
	port := os.Getenv("PORT")
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Service OK",
			"port":    port,
			"env":     os.Getenv("APP_ENV"),
		})
	})

	pd := controller.NewProductController(db)
	pdRoute := productHttp.NewProductHttpRoute(pd)
	pdRoute.Route(e)

	return e

}
