package route

import (
	"net/http"
	"os"

	controller "m9-backstore-service/controllers"

	"github.com/gin-gonic/gin"
)

func RouterSetup() {
	port := os.Getenv("PORT")
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "service OK.")
	})
	router.GET("/products", controller.GetProductsHandler)

	router.Run(":" + port)
}
