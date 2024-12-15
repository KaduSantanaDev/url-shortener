package migrations

import (
	"gorm.io/gorm"
	"url-shortner/database/models"
)

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&models.Link{})
}
