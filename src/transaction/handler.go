package transaction

import (
	"MyApp/datastore/model"
	"MyApp/helper"
	"fmt"
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
	fmt.Println(userID, "userID..............")

	var input InputUser

	db := model.SetupDB()
	tx := db.Begin()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Type == 1 {
		expense = 0.00
		income = input.Price
	} else if input.Type == 0 {
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

	var balances model.Balance
	var category model.Category

	tx.Where("user_id = ?", userID).First(&balances)
	fmt.Println("input category", input.CategoryID)
	db.First(&category, "id = ?", input.CategoryID)
	fmt.Println(category.Category, category.Type, "category")

	balances.Balance += income - expense
	balances.UpdatedAt = time.Now()

	tx.Save(&balances)

	// row := db.Table("transactions").Select("sum(income - expense)").Row()
	// row.Scan(&balance)

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
	// balance = 0.00
}
