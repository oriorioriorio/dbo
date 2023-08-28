package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/marioheryanto/dbo/dtos"
)

func Auth(c *gin.Context) {
	var response dtos.Response

	tokenString, err := c.Cookie("jwt")
	if err != nil {
		response.Message = "please login again"
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("PRIVATE_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp, _ := claims["exp"].(float64)

		if float64(time.Now().Unix()) > exp {
			response.Message = "token expired, please login again"
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		email, _ := claims["sub"].(string)
		if email == "" {
			response.Message = "please login again"
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("email", email)
		c.Next()
	} else {
		response.Message = "please login again"
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
}
