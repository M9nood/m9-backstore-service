package jwtpkg

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

var signingKey = getSecretKey()

func getSecretKey() []byte {
	secret := os.Getenv("AUTH_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return []byte(secret)
}

func KeyFunc(t *jwt.Token) (interface{}, error) {
	if t.Method.Alg() != "HS256" {
		return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
	}
	return signingKey, nil
}

func ParseWithClaims(tokenStr string) (map[string]interface{}, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		return KeyFunc(t)
	}
	token, err := jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return nil, nil
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, nil
	}
}
