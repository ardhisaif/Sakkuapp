package main

import (
	"MyApp/datastore/model"
	"MyApp/engine"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	
	db := model.SetupDB() 
	r := engine.Router()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	if err := engine.Router().Run(port); err != nil {
		log.Fatal("Unable to start:", err)
	}
	r.Run(port)
}
