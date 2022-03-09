package login

import (
	"MyApp/datastore/model"
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(email string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	var SECRET_KEY = []byte(secretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(SECRET_KEY)
	// fmt.Println(tokenString,token)
	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func Login(c *gin.Context) {
	db := model.SetupDB()
	var user model.User
	var input InputUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if isValidPassword := CheckPasswordHash(input.Password, user.Password); !isValidPassword {
		fmt.Println("wrong password")
		return
	}

	validToken, err := GenerateJWT(user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Failed to generate token")
		return
	}

	var token Token
	token.Email = user.Email
	token.TokenString = validToken

	data := gin.H{
		"id": user.ID,
		"name": user.Name,
		"email": user.Email,
		"token": token.TokenString,	
	}

	meta := gin.H{
		"message": "Login Success!",
		"statusCode": http.StatusOK,
	}

	response := gin.H{
		"data": data,
		"meta": meta,
	}


	c.JSON(http.StatusOK, response)
}
