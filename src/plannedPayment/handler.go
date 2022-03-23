package plannedpayment

import (
	"MyApp/datastore/model"
	"MyApp/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListPlannedPayment(c *gin.Context) {
	userID := helper.Authorization(c)
	var plannedPayment []model.PlannedPayment
	db := model.SetupDB()

	db.Where("user_id = ?", userID).Find(&plannedPayment)

	meta := gin.H{
		"message":    "Data successfully retrieved/transmitted!",
		"statusCode": http.StatusOK,
	}

	c.JSON(http.StatusOK, gin.H{"version": "v1", "data": plannedPayment, "meta": meta})
}

func CreatePlannedPayment(c *gin.Context) {
	userID := helper.Authorization(c)

	var input InputUser
	var category model.Category

	db := model.SetupDB()
	tx := db.Begin()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plannedPayment := model.PlannedPayment{
		Description: input.Description,
		UserID:      userID,
		CategoryID:  input.CategoryID,
		Price:       input.Price,
	}

	if err := tx.Create(&plannedPayment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx.First(&category, "id = ?", input.CategoryID)

	data := gin.H{
		"id":          plannedPayment.ID,
		"descriotion": input.Description,
		"category":    category.Category,
		"price":       input.Price,
	}

	meta := gin.H{
		"message":    "Data successfully retrieved/transmitted!",
		"statusCode": http.StatusOK,
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"version": "v1", "data": data, "meta": meta})
}

func CreatePlannedPaymentByID(c *gin.Context) {
	userID := helper.Authorization(c)

	var category model.Category
	var transaction model.Transaction
	var price float64
	transactionID := c.Param("id")

	db := model.SetupDB()
	tx := db.Begin()

	tx.Where("id = ?", transactionID).First(&transaction)

	if transaction.Expense != 0 {
		price = transaction.Expense
	} else {
		price = transaction.Income
	}

	plannedPayment := model.PlannedPayment{
		Description: transaction.Description,
		UserID:      userID,
		CategoryID:  transaction.CategoryID,
		Price:       price,
	}

	if result := tx.Create(&plannedPayment); result.Error != nil {
		panic(result)
	}

	tx.First(&category, "id = ?", transaction.CategoryID)

	data := gin.H{
		"id":          plannedPayment.ID,
		"descriotion": transaction.Description,
		"category":    category.Category,
		"price":       price,
	}

	meta := gin.H{
		"message":    "Data successfully retrieved/transmitted!",
		"statusCode": http.StatusOK,
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"version": "v1", "data": data, "meta": meta})
}
