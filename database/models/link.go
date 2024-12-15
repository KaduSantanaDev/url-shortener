package models

import "time"

type Link struct {
	ID        string     `gorm:"primary_key"`
	Url       string     `gorm:"unique;not null"`
	Slug      string     `gorm:"unique;not null;size:6"`
	ShortUrl  string     `gorm:"unique"`
	CreatedAt *time.Time `gorm:"not null;default:current_timestamp"`
}

func (Link) TableName() string {
	return "links"
}
