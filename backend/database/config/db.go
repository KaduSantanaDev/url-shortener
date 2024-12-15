package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"url-shortner/database/models"
)

func NewPostgres() (*gorm.DB, error) {
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)
	log.Println("start connecting to db:")
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect do db: %w", err)
	}

	log.Println("Migration started...")
	if err := db.AutoMigrate(&models.Link{}); err != nil {
		log.Printf("Migration failed for Link model: %v\n", err)
		return nil, fmt.Errorf("migration failed: %w", err)
	}
	log.Println("Migration completed successfully.")
	return db, nil
}
