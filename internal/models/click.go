package models

import "gorm.io/gorm"

type Click struct {
	gorm.Model
	LinkID   uint   `gorm:"not null"`
	Country  string `gorm:"not null"`
	Referrer string `gorm:"not null"`
	Device   string `gorm:"not null"`
}
