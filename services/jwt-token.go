package service

import "github.com/golang-jwt/jwt"

type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}
