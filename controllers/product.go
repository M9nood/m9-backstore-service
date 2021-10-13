package controller

import (
	"net/http"

	service "m9-backstore-service/services"

	"github.com/M9nood/go-iterror"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	jwtTokenService service.JWTServiceInterface
	productService  service.ProductServiceInterface
}

func NewProductController(db *gorm.DB) ProductHandler {
	jwtSvc := service.NewJWTAuthService()
	prodSvc := service.NewProductService(db)
	return ProductHandler{
		jwtTokenService: jwtSvc,
		productService:  prodSvc,
	}
}

func (h *ProductHandler) GetProductsHandler(c echo.Context) error {
	authToken := c.Request().Header.Get("Authorization")
	dataToken, errParse := h.jwtTokenService.ParseToken(authToken)
	if errParse != nil {
		return c.JSON(400, CreateErrorResponse(iterror.ErrorBadRequest("Invalid token")))
	}
	products, err := h.productService.GetProductsInStoreService(dataToken.StoreId)
	if err != nil {
		return c.JSON(err.GetHttpCode(), CreateErrorResponse(err))
	}
	return c.JSON(http.StatusOK, CreateSuccessResponse(products))
}
