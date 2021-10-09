package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthCustomClaims struct {
	Username string `json:"name"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

type JWTService interface {
	GenerateToken(Username string, Email string) string
}

func NewJWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "Bikash",
	}
}

func getSecretKey() string {
	secret := os.Getenv("AUTH_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(Username string, Email string) string {
	claims := &AuthCustomClaims{
		Username,
		Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}
