package transaction

import (
	"MyApp/datastore/model"
	"MyApp/helper"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var income float64
var expense float64

// var balance float64

type Response struct {
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Income      float64 `json:"income"`
	Expense     float64 `json:"expense"`
	Balance     float64 `json:"balance"`
}

func GetListTransaction(c *gin.Context) {
	userID := helper.Authorization(c)
	var transactions []model.Transaction
	db := model.SetupDB()

	db.Where("user_id = ?", userID).Find(&transactions)

	meta := gin.H{
		"message":    "Data successfully retrieved/transmitted!",
		"statusCode": http.StatusOK,
	}

	c.JSON(http.StatusOK, gin.H{"data": transactions, "meta": meta})
}

func CreateTransaction(c *gin.Context) {

	userID := helper.Authorization(c)

	var input InputUser
	var balances model.Balance
	var category model.Category

	db := model.SetupDB()
	tx := db.Begin()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx.First(&category, "id = ?", input.CategoryID)

	if category.Type == 1 {
		expense = 0.00
		income = input.Price
	}

	if category.Type == 0 {
		income = 0.00
		expense = input.Price
	}

	transaction := model.Transaction{
		Description: input.Description,
		UserID:      userID,
		CategoryID:  input.CategoryID,
		Income:      income,
		Expense:     expense,
	}

	if result := tx.Create(&transaction); result.Error != nil {
		panic(result)
	}

	tx.Where("user_id = ?", userID).First(&balances)

	balances.Balance += income - expense
	balances.UpdatedAt = time.Now()
	category.Total += income - expense

	tx.Save(&balances)
	tx.Save(&category)

	response := Response{
		Description: input.Description,
		Category:    category.Category,
		Income:      income,
		Expense:     expense,
		Balance:     balances.Balance,
	}

	data := gin.H{
		"descriotion": response.Description,
		"category":    response.Category,
		"income":      response.Income,
		"expense":     response.Expense,
		"balance":     response.Balance,
	}

	meta := gin.H{
		"message":    "Data successfully retrieved/transmitted!",
		"statusCode": http.StatusOK,
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"data": data, "meta": meta})
}
