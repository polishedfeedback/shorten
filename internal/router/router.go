package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/polishedfeedback/shorten/internal/handlers"
	"github.com/polishedfeedback/shorten/internal/middleware"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://shorten.krishnasaimarupudi.com", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.POST("/api/auth/register", handlers.Register(db))
	r.POST("/api/auth/login", handlers.Login(db))
	r.GET("/:code", handlers.Redirect(db))

	authorized := r.Group("/api/links")

	authorized.Use(middleware.AuthRequired())
	{
		authorized.GET("", handlers.GetLinks(db))
		authorized.POST("", handlers.CreateLink(db))
		authorized.DELETE("/:id", handlers.DeleteLink(db))
	}
	return r
}
