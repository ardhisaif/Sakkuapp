package engine

import (
	"MyApp/src/auth/login"
	"MyApp/src/auth/register"
	Transaction "MyApp/src/transaction"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Router() *gin.Engine {
	r := gin.New()
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))

	auth := r.Group("/auth")
	{
		auth.POST("/register", register.Register)
		auth.POST("/login", login.Login)
	}

	r.Use(AuthToken)
	transaction := r.Group("/transaction")
	transaction.Use(AuthToken)
	{
		transaction.POST("/", Transaction.CreateTransaction)
		transaction.GET("/", Transaction.GetListTransaction)
	}

	return r
}

func AuthToken(c *gin.Context) {
	secretKey := os.Getenv("JWT_SECRET")
	var SECRET_KEY = []byte(secretKey)
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		c.String(http.StatusForbidden, "No Authorization header provided")
		c.Abort()
		return
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
		return 
	}

	if !token.Valid {
		fmt.Println("errorr")
		c.JSON(http.StatusInternalServerError,"token invalid")
	}
}