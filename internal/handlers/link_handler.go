package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/polishedfeedback/shorten/internal/models"
	"github.com/polishedfeedback/shorten/internal/services"
	"gorm.io/gorm"
)

type LinkRequest struct {
	URL string `json:"url" binding:"required"`
}

type LinkResponse struct {
	ID          uint      `json:"id"`
	ShortCode   string    `json:"short_code"`
	OriginalURL string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
	Clicks      int64     `json:"clicks"`
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

		var links []models.Link
		db.Where("user_id = ?", uint(userID.(float64))).Find(&links)

		type ClickCount struct {
			LinkID uint
			Count  int64
		}
		var clickCounts []ClickCount
		db.Model(&models.Click{}).
			Select("link_id, count(*) as count").
			Where("link_id IN ?", func() []uint {
				ids := make([]uint, len(links))
				for i, l := range links {
					ids[i] = l.ID
				}
				return ids
			}()).
			Group("link_id").
			Scan(&clickCounts)

		countMap := make(map[uint]int64)
		for _, cc := range clickCounts {
			countMap[cc.LinkID] = cc.Count
		}

		response := make([]LinkResponse, 0)
		for _, link := range links {
			response = append(response, LinkResponse{
				ID:          link.ID,
				ShortCode:   link.ShortCode,
				OriginalURL: link.OriginalURL,
				CreatedAt:   link.CreatedAt,
				Clicks:      countMap[link.ID],
			})
		}

		c.JSON(http.StatusOK, gin.H{"data": response})
	}
}
