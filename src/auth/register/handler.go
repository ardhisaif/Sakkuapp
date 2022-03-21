package register

import (
	"MyApp/datastore/model"
	"fmt"
	"net/http"
	"time"

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

	fmt.Println(Register.ID)

	Balance := model.Balance{
		UserID:    Register.ID,
		Balance:   0.00,
		CreatedAt: time.Now(),
	}

	if result := db.Create(&Balance); result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"version": "v1", "message": result.Error})
		return
	}

	category := []model.Category{
		{
			UserID:   Register.ID.String(),
			Category: "food",
			Type:     0,
		},
		{
			UserID:   Register.ID.String(),
			Category: "transportation",
			Type:     0,
		},
		{
			UserID:   Register.ID.String(),
			Category: "salary",
			Type:     1,
		},
		{
			UserID:   Register.ID.String(),
			Category: "charity, infaq, shodaqoh",
			Type:     0,
		},
		{
			UserID:   Register.ID.String(),
			Category: "sale",
			Type:     1,
		},
		{
			UserID:   Register.ID.String(),
			Category: "clothes",
			Type:     0,
		},
	}

	if result := db.Create(&category); result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"version": "v1", "message": result.Error})
	}

	data := gin.H{
		"id":      Register.ID,
		"name":    Register.Name,
		"email":   Register.Email,
		"balance": Balance.Balance,
	}

	meta := gin.H{
		"message":    "Register Success!",
		"statusCode": http.StatusOK,
	}

	response := gin.H{
		"version": "v1",
		"data":    data,
		"meta":    meta,
	}

	c.JSON(http.StatusOK, response)

}
