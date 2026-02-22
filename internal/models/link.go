package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	ShortCode   string  `gorm:"uniqueIndex;not null"`
	OriginalURL string  `gorm:"not null"`
	UserID      uint    `gorm:"not null"`
	Links       []Click `gorm:"foreignKey:LinkID"`
}
