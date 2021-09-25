package route

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RouterSetup() {
	port := os.Getenv("PORT")
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "service OK.")
	})

	router.Run(":" + port)
}
