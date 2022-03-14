package helper

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authorization(c *gin.Context) string {
	secretKey := os.Getenv("JWT_SECRET")
	var SECRET_KEY = []byte(secretKey)
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		c.String(http.StatusForbidden, "No Authorization header provided")
		c.Abort()
	}
	tokenString := strings.TrimPrefix(auth, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return SECRET_KEY, nil
	})

	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusUnauthorized, err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	var userID string

	if ok && token.Valid {
		email := claims["email"]
		userID = claims["user_id"].(string)
		fmt.Println(claims, email, "claims")
	} else {
		fmt.Println(err, "error")
		c.JSON(http.StatusBadGateway, err)
	}

	if !token.Valid {
		c.JSON(http.StatusInternalServerError, "token invalid")
	}

	return userID
}
