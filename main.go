package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"url-shortner/database/config"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("cannot load .env file")
	}

	db, err := config.NewPostgres()
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	log.Println("connected to db successfully:", db)

	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if err := router.Run(":8005"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
