package main

import (
	"os"

	"github.com/polishedfeedback/shorten/internal/config"
	"github.com/polishedfeedback/shorten/internal/database"
	"github.com/polishedfeedback/shorten/internal/router"
)

func main() {
	database.AutoMigrate()

	db := config.Connect()
	r := router.SetupRouter(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
