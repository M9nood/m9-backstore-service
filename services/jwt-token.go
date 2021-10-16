package service

import (
	"encoding/json"
	"fmt"
	"m9-backstore-service/models/auth"
	"os"
	"strings"
	"time"

	jwtpkg "m9-backstore-service/pkg/jwt"

	"github.com/golang-jwt/jwt"
)

type AuthCustomClaims struct {
	Id       int    `json:"id"`
	Username string `json:"name"`
	Email    string `json:"email"`
	StoreId  *int   `json:"store_id"`
	jwt.StandardClaims
}

type RefreshCustomClaims struct {
	UserId int `json:"id"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

type JWTServiceInterface interface {
	GenerateToken(payload auth.LoginResponse) string
	ParseToken(tokenStr string) (*AuthCustomClaims, error)
	GenerateRefreshToken(userId int) string
	ParseRefreshToken(tokenStr string) (*RefreshCustomClaims, error)
}

func NewJWTAuthService() JWTServiceInterface {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "m9",
	}
}

func getSecretKey() string {
	secret := os.Getenv("AUTH_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(payload auth.LoginResponse) string {
	claims := &AuthCustomClaims{
		payload.Id,
		payload.UserName,
		payload.Email,
		payload.StoreId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
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

func (service *jwtServices) GenerateRefreshToken(userId int) string {
	rtClaims := RefreshCustomClaims{}
	rtClaims.Issuer = service.issure
	rtClaims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	rtClaims.IssuedAt = time.Now().Unix()
	rtClaims.UserId = userId
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	rt, err := refreshToken.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return rt
}

func (service *jwtServices) ParseToken(tokenStr string) (*AuthCustomClaims, error) {
	tokenStr = service.splitToken(tokenStr)
	claims, err := jwtpkg.ParseWithClaims(tokenStr)
	if err != nil {
		return nil, nil
	}
	data := AuthCustomClaims{}
	claimsJson, err := json.Marshal(claims)
	if err != nil {
		// do error check
		fmt.Println(err)
		return nil, nil
	}
	if err := json.Unmarshal(claimsJson, &data); err != nil {
		return nil, nil
	}
	return &data, nil
}

func (service *jwtServices) ParseRefreshToken(tokenStr string) (*RefreshCustomClaims, error) {
	claims, err := jwtpkg.ParseWithClaims(tokenStr)
	if err != nil {
		return nil, nil
	}
	data := RefreshCustomClaims{}
	claimsJson, err := json.Marshal(claims)
	if err != nil {
		// do error check
		fmt.Println(err)
		return nil, nil
	}
	if err := json.Unmarshal(claimsJson, &data); err != nil {
		return nil, nil
	}
	return &data, nil
}

func (service *jwtServices) splitToken(reqToken string) string {
	splitToken := strings.Split(reqToken, "Bearer ")
	token := splitToken[1]
	return token
}
