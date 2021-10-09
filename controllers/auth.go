package controller

import (
	"log"
	"net/http"

	"m9-backstore-service/models/auth"
	"m9-backstore-service/pkg"
	service "m9-backstore-service/services"

	iterror "github.com/M9nood/go-iterror"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService service.AuthServiceInterface
}

func NewAuthController(db *gorm.DB) AuthHandler {
	svc := service.NewAuthService(db)
	return AuthHandler{
		authService: svc,
	}
}

func (h *AuthHandler) RegisterHandler(c echo.Context) error {
	requestData := auth.RegisterRequest{}
	if err := c.Bind(&requestData); err != nil {
		log.Println("register err binding: ")
		return c.JSON(http.StatusUnprocessableEntity, CreateErrorResponse(iterror.ErrorBadRequest("Invalid request data")))
	}
	if err := c.Validate(requestData); err != nil {
		log.Println("validate err binding: ", err)
		return c.JSON(http.StatusUnprocessableEntity, CreateErrorResponse(iterror.ErrorBadRequest("Invalid request data"), pkg.ParseValidationError(err)))
	}
	products, err := h.authService.RegisterService(requestData)
	if err != nil {
		return c.JSON(err.GetHttpCode(), CreateErrorResponse(err))
	}
	return c.JSON(http.StatusOK, CreateSuccessResponse(products))
}

func (h *AuthHandler) LoginHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, CreateSuccessResponse(""))
}
