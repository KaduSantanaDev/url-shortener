package repository

import (
	"gorm.io/gorm"
	"url-shortner/database/models"
)

type LinkRepository interface {
	CreateLink(link *models.Link) error
	FindBySlug(slug string) (*models.Link, error)
}

type GormLinkRepository struct {
	DB *gorm.DB
}

func NewLinkRepository(db *gorm.DB) LinkRepository {
	return &GormLinkRepository{db}
}

func (repository *GormLinkRepository) CreateLink(link *models.Link) error {
	return repository.DB.Create(link).Error
}

func (repository *GormLinkRepository) FindBySlug(slug string) (*models.Link, error) {
	var link models.Link
	if err := repository.DB.Where("slug = ?", slug).First(&link).Error; err != nil {
		return nil, err
	}
	return &link, nil
}
