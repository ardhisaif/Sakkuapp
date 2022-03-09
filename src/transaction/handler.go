package transaction

import (
	"MyApp/datastore/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var income float64
var expense float64
var balance float64

type Response struct {
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Income      float64 `json:"income"`
	Expense     float64 `json:"expense"`
	Balance     float64 `json:"balance"`
}

func GetListTransaction(c *gin.Context) {
	var transactions []model.Transaction
	db := model.SetupDB()

	db.Find(&transactions)

	meta := gin.H{
		"message": "Data successfully retrieved/transmitted!",
		"statusCode": http.StatusOK,
	}

	c.JSON(http.StatusOK, gin.H{"data": transactions, "meta": meta})
}

func CreateTransaction(c *gin.Context) {
	var input InputUser

	db := model.SetupDB()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Type == "Income" {
		expense = 0.00
		income = input.Price
	} else if input.Type == "Expense" {
		income = 0.00
		expense = input.Price
	}

	Transaction := model.Transaction{
		Description: input.Description,
		Category:    input.Category,
		Income:      income,
		Expense:     expense,
	}

	if result := db.Create(&Transaction); result.Error != nil {
		panic(result)
	}

	row := db.Table("transactions").Select("sum(income - expense)").Row()
	row.Scan(&balance)

	response := Response{
		Description: input.Description,
		Category:    input.Category,
		Income:      income,
		Expense:     expense,
		Balance:     balance,
	}

	data := gin.H{
		"descriotion": response.Description,
		"category":    response.Category,
		"income":      response.Income,
		"expense":     response.Expense,
		"balance":     response.Balance,
	}

	meta := gin.H{
		"message": "Data successfully retrieved/transmitted!",
		"statusCode": http.StatusOK,
	}

	c.JSON(http.StatusOK, gin.H{"data": data, "meta": meta})
	balance = 0.00
}
