package models

import (
	"time"
)

type Link struct {
	ID        uint   `gorm:"primary_key;auto_increment"`
	Url       string `gorm:"not null"`
	Slug      string `gorm:"unique;not null"`
	ShortUrl  string
	CreatedAt *time.Time `gorm:"not null;default:current_timestamp"`
}

func (Link) TableName() string {
	return "links"
}
