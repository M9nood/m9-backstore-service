package controller

import (
	"net/http"

	service "m9-backstore-service/services"

	"github.com/labstack/echo/v4"
)

func GetProductsHandler(c echo.Context) error {
	productService := service.NewProductService()
	products, err := productService.GetProductsService()
	if err != nil {
		return c.JSON(err.GetHttpCode(), CreateErrorResponse(err))
	}
	return c.JSON(http.StatusOK, CreateSuccessResponse(products))
}
