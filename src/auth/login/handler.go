package login

import (
	"MyApp/datastore/model"
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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

func GenerateJWT(email string, userID uuid.UUID) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	var SECRET_KEY = []byte(secretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(SECRET_KEY)
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
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Record not found!"})
		return
	}

	if isValidPassword := CheckPasswordHash(input.Password, user.Password); !isValidPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})
		return
	}

	validToken, err := GenerateJWT(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
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
