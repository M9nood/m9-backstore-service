package controller

import (
	"log"
	"net/http"

	model "m9-backstore-service/models/auth"
	"m9-backstore-service/pkg"
	service "m9-backstore-service/services"

	iterror "github.com/M9nood/go-iterror"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService     service.AuthServiceInterface
	jwtTokenService service.JWTServiceInterface
}

func NewAuthController(db *gorm.DB) AuthHandler {
	jwtSvc := service.NewJWTAuthService()
	svc := service.NewAuthService(db)
	return AuthHandler{
		authService:     svc,
		jwtTokenService: jwtSvc,
	}
}

func (h *AuthHandler) RegisterHandler(c echo.Context) error {
	requestData := model.RegisterRequest{}
	if err := c.Bind(&requestData); err != nil {
		log.Println("register err binding: ", err)
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
	requestData := model.LoginRequest{}
	if err := c.Bind(&requestData); err != nil {
		log.Println("login err binding: ", err)
		return c.JSON(http.StatusUnprocessableEntity, CreateErrorResponse(iterror.ErrorBadRequest("Invalid request data")))
	}
	if err := c.Validate(requestData); err != nil {
		log.Println("validate err binding: ", err)
		return c.JSON(http.StatusUnprocessableEntity, CreateErrorResponse(iterror.ErrorBadRequest("Invalid request data"), pkg.ParseValidationError(err)))
	}
	user, err := h.authService.LoginService(requestData)
	if err != nil {
		return c.JSON(err.GetHttpCode(), CreateErrorResponse(err))
	}
	return c.JSON(http.StatusOK, CreateSuccessResponse(user))
}

func (h *AuthHandler) RefreshTokenHandler(c echo.Context) error {
	requestData := model.RefreshTokenRequest{}
	if err := c.Bind(&requestData); err != nil {
		log.Println("refresh token request err binding: ", err)
		return c.JSON(http.StatusUnprocessableEntity, CreateErrorResponse(iterror.ErrorBadRequest("Invalid request data")))
	}
	rfToken := requestData.RefreshToken
	dataToken, errParse := h.jwtTokenService.ParseRefreshToken(rfToken)
	if errParse != nil {
		return c.JSON(400, CreateErrorResponse(iterror.ErrorBadRequest("Invalid token")))
	}
	token, err := h.authService.RefreshTokenService(dataToken.UserId)
	if err != nil {
		return c.JSON(err.GetHttpCode(), CreateErrorResponse(err))
	}
	return c.JSON(http.StatusOK, CreateSuccessResponse(token))
}
