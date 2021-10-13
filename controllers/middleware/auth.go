package middleware

import (
	"errors"
	controller "m9-backstore-service/controllers"

	iterror "github.com/M9nood/go-iterror"

	jwtpkg "m9-backstore-service/pkg/jwt"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var config = middleware.JWTConfig{
	ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
		keyFunc := func(t *jwt.Token) (interface{}, error) {
			return jwtpkg.KeyFunc(t)
		}
		// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
		token, err := jwt.Parse(auth, keyFunc)
		if err != nil {
			return nil, err
		}
		if !token.Valid {
			return nil, errors.New("invalid token")
		}
		return token, nil
	},
	ErrorHandlerWithContext: func(err error, c echo.Context) error {
		return c.JSON(400, controller.CreateErrorResponse(iterror.ErrorBadRequest("Invalid token")))
	},
}

var IsLoggedIn = middleware.JWTWithConfig(config)
