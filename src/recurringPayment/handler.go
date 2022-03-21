package recurringpayment

import (
	"MyApp/datastore/model"
	"MyApp/helper"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Recurringpayment(c *gin.Context) {
	userID := helper.Authorization(c)

	plannedPaymentID := c.Param("id")
	var plannedPayment model.PlannedPayment
	var balances model.Balance
	var category model.Category

	db := model.SetupDB()
	tx := db.Begin()

	tx.Where("user_id = ? AND id = ?", userID, plannedPaymentID).Find(&plannedPayment)
	tx.Where("id = ?", plannedPayment.CategoryID).Find(&category)

	var income float64
	var expense float64

	if category.Type == 1 {
		expense = 0.00
		income = plannedPayment.Price
	} else if category.Type == 0 {
		income = 0.00
		expense = plannedPayment.Price
	}

	transaction := model.Transaction{
		Description: plannedPayment.Description,
		UserID:      plannedPayment.UserID,
		CategoryID:  plannedPayment.CategoryID,
		Income:      income,
		Expense:     expense,
	}

	if result := tx.Create(&transaction); result.Error != nil {
		panic(result)
	}

	tx.Where("user_id = ?", userID).First(&balances)

	tx.First(&category, "id = ?", plannedPayment.CategoryID)

	balances.Balance += income - expense
	balances.UpdatedAt = time.Now()
	category.Total += income - expense

	tx.Save(&balances)
	tx.Save(&category)

	data := gin.H{
		"descriotion": plannedPayment.Description,
		"category":    category.Category,
		"price":       plannedPayment.Price,
	}

	meta := gin.H{
		"message":    "Data successfully retrieved/transmitted!",
		"statusCode": http.StatusOK,
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"version": "v1", "data": data, "meta": meta})
}
