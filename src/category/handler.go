package category

import (
	"MyApp/datastore/model"
	"MyApp/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListCategory(c *gin.Context) {
	var category []model.Category
	db := model.SetupDB()

	db.Find(&category)

	meta := gin.H{
		"message":    "Data successfully retrieved/transmitted!",
		"statusCode": http.StatusOK,
	}

	c.JSON(http.StatusOK, gin.H{"data": category, "meta": meta})
}

func CreateCategory(c *gin.Context) {
	userID := helper.Authorization(c)
	fmt.Println(userID, "userID..............")

	var input InputUser

	db := model.SetupDB()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := model.Category{
		Category: input.Category,
		Type: input.Type,
	}

	if result := db.Create(&category); result.Error != nil {
		panic(result)
	}

	data := gin.H{
		"id": category.ID,
		"category": category.Category,
		"type": category.Type,
	}

	meta := gin.H{
		"message":    "Data successfully retrieved/transmitted!",
		"statusCode": http.StatusOK,
	}

	c.JSON(http.StatusOK, gin.H{"data": data, "meta": meta})

}
