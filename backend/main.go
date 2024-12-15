package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"url-shortner/database/config"
	"url-shortner/database/repository"
	"url-shortner/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("cannot load .env file")
	}

	db, err := config.NewPostgres()
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	log.Println("connected to db successfully")

	router := gin.Default()
	router.Use(cors.Default())

	linkRepo := repository.NewLinkRepository(db)
	linkHandler := handlers.NewLinkHandler(linkRepo)

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/links", linkHandler.CreateLink)
	router.GET("/:slug", linkHandler.GetLink)

	if err := router.Run(":8005"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
