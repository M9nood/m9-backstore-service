package controller

import (
	"log"
	"net/http"

	"m9-backstore-service/pkg"
	service "m9-backstore-service/services"

	model "m9-backstore-service/models/product"

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

func (h *ProductHandler) CreateProductHandler(c echo.Context) error {
	authToken := c.Request().Header.Get("Authorization")
	dataToken, errParse := h.jwtTokenService.ParseToken(authToken)
	if errParse != nil {
		return c.JSON(400, CreateErrorResponse(iterror.ErrorBadRequest("Invalid token")))
	}

	requestData := model.ProductCreateRequest{}
	if err := c.Bind(&requestData); err != nil {
		log.Println("create product err binding: ", err)
		return c.JSON(http.StatusUnprocessableEntity, CreateErrorResponse(iterror.ErrorBadRequest("Invalid request data")))
	}
	if err := c.Validate(requestData); err != nil {
		log.Println("validate err binding: ", err)
		return c.JSON(http.StatusUnprocessableEntity, CreateErrorResponse(iterror.ErrorBadRequest("Invalid request data"), pkg.ParseValidationError(err)))
	}
	result, err := h.productService.CreateProductService(dataToken.StoreId, requestData)
	if err != nil {
		return c.JSON(err.GetHttpCode(), CreateErrorResponse(err))
	}
	return c.JSON(http.StatusOK, CreateSuccessResponse(result))
}
