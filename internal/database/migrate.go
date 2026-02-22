package database

import (
	"log"

	"github.com/polishedfeedback/shorten/internal/config"
	"github.com/polishedfeedback/shorten/internal/models"
)

func AutoMigrate() {
	db := config.Connect()

	err := db.AutoMigrate(&models.User{}, &models.Link{}, &models.Click{})
	if err != nil {
		log.Fatalf("error running migrations: %v", err)
	}
	log.Println("database migrated successfully")
}
