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

	db.Where("user_id = ? AND type = ? ", userID, 0).Find(&category)

	var totalExpense float64 
	for _, e := range category {
		totalExpense += e.Total
	}
	totalExpense = totalExpense * -1

	db.Where("user_id = ?", userID).First(&balance)

	response := gin.H{
		"balance": balance.Balance,
		"totalExpense": totalExpense,
		"expense": category,
	}

	c.JSON(200, gin.H{"data": response})
}
