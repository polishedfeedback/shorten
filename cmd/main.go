package main

import (
	"os"

	"github.com/polishedfeedback/shorten/internal/database"
	"github.com/polishedfeedback/shorten/internal/router"
)

func main() {
	database.AutoMigrate()

	r := router.SetupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
