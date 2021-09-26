package controller

import (
	"net/http"

	service "m9-backstore-service/services"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductController(db *gorm.DB) ProductHandler {
	svc := service.NewProductService(db)
	return ProductHandler{
		productService: svc,
	}
}

func (h *ProductHandler) GetProductsHandler(c echo.Context) error {
	products, err := h.productService.GetProductsService()
	if err != nil {
		return c.JSON(err.GetHttpCode(), CreateErrorResponse(err))
	}
	return c.JSON(http.StatusOK, CreateSuccessResponse(products))
}
