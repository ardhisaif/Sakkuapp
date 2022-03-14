package engine

import (
	"MyApp/src/auth/login"
	"MyApp/src/auth/register"
	Transaction "MyApp/src/transaction"
	Category "MyApp/src/category"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.New()
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home page"})
	})

	auth := r.Group("/auth")
	{
		auth.POST("/register", register.Register)
		auth.POST("/login", login.Login)
	}

	transaction := r.Group("/transaction")
	{
		transaction.POST("/", Transaction.CreateTransaction)
		transaction.GET("/", Transaction.GetListTransaction)
	}

	category := r.Group("/category")
	{
		category.POST("/", Category.CreateCategory)
		category.GET("/", Category.GetListCategory)
	}

	return r
}

// func AuthToken(c *gin.Context) *jwt.Token {
// 	secretKey := os.Getenv("JWT_SECRET")
// 	var SECRET_KEY = []byte(secretKey)
// 	auth := c.Request.Header.Get("Authorization")
// 	if auth == "" {
// 		c.String(http.StatusForbidden, "No Authorization header provided")
// 		c.Abort()
// 	}
// 	tokenString := strings.TrimPrefix(auth, "Bearer")

// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return SECRET_KEY, nil
// 	})

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(token, "tokeeeen")

// 	if !token.Valid {
// 		fmt.Println("errorr")
// 		c.JSON(http.StatusInternalServerError, "token invalid")
// 	}

// 	return token
// }
