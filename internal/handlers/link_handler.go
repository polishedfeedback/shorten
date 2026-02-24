package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/polishedfeedback/shorten/internal/services"
	"gorm.io/gorm"
)

type LinkRequest struct {
	URL string `json:"url" binding:"required"`
}

func CreateLink(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		var req LinkRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "url is required"})
			return
		}

		link, err := services.CreateLink(db, uint(userID.(float64)), req.URL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating link"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"data": link})
	}
}

func DeleteLink(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		err = services.DeleteLink(db, uint(id), uint(userID.(float64)))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting the link"})
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func GetLinks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		links, err := services.GetLinks(db, uint(userID.(float64)))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting links"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": links})
	}
}
