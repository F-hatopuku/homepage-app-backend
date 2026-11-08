package main

import (
	"log"

	"github.com/F-hatopuku/homepage-app/handlers"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	router := gin.Default()
	router.POST("/api/contact", handlers.ContactHandler)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
