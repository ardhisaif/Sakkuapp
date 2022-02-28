package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"primary_key"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4()

	return
}

type InputUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(c *gin.Context) {
	var input InputUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createUser := User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	db := c.MustGet("db").(*gorm.DB)
	if result := db.Create(&createUser); result.Error != nil {
		panic(result)
	} else {
		fmt.Println(createUser.ID)
		c.JSON(http.StatusOK, gin.H{"data": createUser})
	}
}

func ReadHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func SetupDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/myapp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("database connected")
	return db
}

func main() {
	db := SetupDB()
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/ping", ReadHello)
	r.POST("/user", CreateUser)
	r.Run(":9000")
}
