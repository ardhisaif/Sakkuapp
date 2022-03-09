package main

import (
	"MyApp/datastore/model"
	"MyApp/engine"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	
	db := model.SetupDB()
	r := engine.Router()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	if err := engine.Router().Run(":9000"); err != nil {
		log.Fatal("Unable to start:", err)
	}
	r.Run(":9000")
}
