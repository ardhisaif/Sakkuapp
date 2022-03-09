package register

import (
	"MyApp/datastore/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func Register(c *gin.Context) {
	var input InputUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	hashPassword, _ := HashPassword(input.Password)

	Register := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashPassword,
	}

	db := model.SetupDB()
	if result := db.Create(&Register); result.Error != nil {
		panic(result)
	} 

	data := gin.H{
		"id": Register.ID,
		"name": Register.Name,
		"email": Register.Email,
	}

	meta := gin.H{
		"message": "Register Success!",
		"statusCode": http.StatusOK,
	}


	response := gin.H{
		"data": data,
		"meta": meta,
	}
	
	c.JSON(http.StatusOK, response)
	
}
