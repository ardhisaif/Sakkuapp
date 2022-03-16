package category

import (
	"MyApp/datastore/model"
	"MyApp/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListCategory(c *gin.Context) {
	userID := helper.Authorization(c)
	var category []model.Category
	db := model.SetupDB()

	db.Where("user_id = ?", userID).Find(&category)

	meta := gin.H{
		"message":    "Data successfully retrieved/transmitted!",
		"statusCode": http.StatusOK,
	}

	c.JSON(http.StatusOK, gin.H{"data": category, "meta": meta})
}

func CreateCategory(c *gin.Context) {
	userID := helper.Authorization(c)

	var input InputUser

	db := model.SetupDB()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := model.Category{
		UserID:   userID,
		Category: input.Category,
		Type:     input.Type,
	}

	if result := db.Create(&category); result.Error != nil {
		panic(result)
	}

	data := gin.H{
		"id":       category.ID,
		"category": category.Category,
		"type":     category.Type,
		"total": category.Total,
	}

	meta := gin.H{
		"message":    "Data successfully retrieved/transmitted!",
		"statusCode": http.StatusOK,
	}

	c.JSON(http.StatusOK, gin.H{"data": data, "meta": meta})

}
