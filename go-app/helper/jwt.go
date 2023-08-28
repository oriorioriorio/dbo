package helper

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("PRIVATE_KEY")))
}
