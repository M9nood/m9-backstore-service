package route

import (
	"net/http"
	"os"

	controller "m9-backstore-service/controllers"
	transport "m9-backstore-service/controllers/http"

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
	pdRoute := transport.NewProductHttpRoute(pd)
	pdRoute.Route(e)

	lb := controller.NewLineBotController(db)
	lbRoute := transport.NewLineBotHttpRoute(lb)
	lbRoute.Route(e)

	auth := controller.NewAuthController(db)
	authRoute := transport.NewAuthHttpRoute(auth)
	authRoute.Route(e)

	return e

}
