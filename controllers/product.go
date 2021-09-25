package controller

import (
	"net/http"

	service "m9-backstore-service/services"

	"github.com/gin-gonic/gin"
)

func GetProductsHandler(c *gin.Context) {
	productService := service.NewProductService()
	products, err := productService.GetProductsService()
	if err != nil {
		c.JSON(err.GetHttpCode(), CreateErrorResponse(err))
	}
	c.JSON(http.StatusOK, CreateSuccessResponse(products))
}
