package category

import (
	"MyApp/datastore/model"
	"MyApp/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		c.JSON(http.StatusBadRequest, gin.H{"version": "v1", "message": result.Error})
		return
	}

	data := gin.H{
		"id":       category.ID,
		"category": category.Category,
		"type":     category.Type,
		"total":    category.Total,
	}

	meta := gin.H{
		"message":    "Data successfully retrieved/transmitted!",
		"statusCode": http.StatusOK,
	}

	c.JSON(http.StatusOK, gin.H{"version": "v1", "data": data, "meta": meta})

}

func GetListCategory(c *gin.Context) {
	userID := helper.Authorization(c)
	var category []model.Category
	var balance model.Balance
	db := model.SetupDB()

	var data []interface{}

	meta := gin.H{
		"message":    "Data successfully retrieved/transmitted!",
		"statusCode": http.StatusOK,
	}

	if err := db.Where("user_id = ?", userID).Find(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"version": "v1", "data": data, "meta": meta})
		return
	}

	if err := db.First(&balance, "user_id = ?", userID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"version": "v1", "message": err.Error()})
		return
	}

	for _, v := range category {
		Type := ""
		if v.Type == 0 {
			Type += "expense"
		}else{
			Type += "income"
		}

		response := gin.H{
			"id":       v.ID,
			"category": v.Category,
			"type":     Type,
			"total":    v.Total,
		}

		data = append(data, response)
		fmt.Println(data)
	}

	response := gin.H{
		"balance": balance.Balance,
		"category": data,
	}

	c.JSON(http.StatusOK, gin.H{"version": "v1", "data": response, "meta": meta})
}

func GetCategoryByID(c *gin.Context){
	userID := helper.Authorization(c)
	categoryID := c.Param("id")
	var category model.Category
	db := model.SetupDB()

	if err := db.Where("user_id = ? AND id = ?", userID, categoryID).Find(&category).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"version": "v1", "message": err})
		return
	}

	data := gin.H{
		"id":       categoryID,
		"category": category.Category,
		"type":     category.Type,
		"total":    category.Total,
	}

	meta := gin.H{
		"message":    "Data successfully retrieved/transmitted!",
		"statusCode": http.StatusOK,
	}

	c.JSON(http.StatusOK, gin.H{"version": "v1", "data": data, "meta": meta})
}

func UpdateCategory(c *gin.Context) {
	db := model.SetupDB()
	userID := helper.Authorization(c)
	var category model.Category
	categoryID := c.Param("id")

	db.Where("id = ? AND user_id = ?", categoryID, userID).First(&category)

	var input InputUser


	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.Category = input.Category
	category.Type = input.Type

	if result := db.Save(&category); result.Error != nil {
		panic(result)
	}

	data := gin.H{
		"id":       category.ID,
		"category": category.Category,
		"type":     category.Type,
		"total":    category.Total,
	}

	meta := gin.H{
		"message":    "Data successfully retrieved/transmitted!",
		"statusCode": http.StatusOK,
	}

	c.JSON(http.StatusOK, gin.H{"version": "v1", "data": data, "meta": meta})

}

func DeleteCategory(c *gin.Context) {
	db := model.SetupDB()
	userID := helper.Authorization(c)
	var category model.Category
	categoryID := c.Param("id")

	db.Where("id = ? AND user_id = ?", categoryID, userID).First(&category)

	if result := db.Delete(&category); result.Error != nil {
		panic(result)
	}

	data := gin.H{
		"id":       category.ID,
		"category": category.Category,
	}

	meta := gin.H{
		"message":    "Data successfully deleted",
		"statusCode": http.StatusOK,
	}

	c.JSON(http.StatusOK, gin.H{"version": "v1", "data": data, "meta": meta})

}
