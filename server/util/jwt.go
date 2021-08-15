package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SecretKey = "secret"

func GenerateJwt(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})

	return claims.SignedString([]byte(SecretKey))
}
