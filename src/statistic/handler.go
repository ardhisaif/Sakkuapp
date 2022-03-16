package home

import (
	"MyApp/datastore/model"
	"MyApp/helper"

	"github.com/gin-gonic/gin"
)

func Statistic(c *gin.Context) {
	var userID = helper.Authorization(c)
	var category []model.Category
	var balance model.Balance

	db := model.SetupDB()

	db.Where("user_id = ? AND type = ?", userID, 1).Find(&category)

	income := category

	db.Where("user_id = ? AND type = ?", userID, 0).Find(&category)

	expense := category

	db.Where("user_id = ?", userID).First(&balance)

	response := gin.H{
		"balance": balance.Balance,
		"income":  income,
		"expense": expense,
	}

	c.JSON(200, gin.H{"data": response})
}
