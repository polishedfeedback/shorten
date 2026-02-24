package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mssola/useragent"
	"github.com/polishedfeedback/shorten/internal/models"
	"gorm.io/gorm"
)

type RedirectRequest struct {
	OriginalURL string `json:"original_url" binding:"required"`
	ShortURL    string `json:"short_url" binding:"required"`
}

type IPResponse struct {
	Country string `json:"country"`
}

func Redirect(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		shortCode := c.Param("code")
		var link models.Link
		result := db.First(&link, "short_code = ?", shortCode)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "link not found"})
			return
		}
		referrer := c.GetHeader("Referer")
		if referrer == "" {
			referrer = "direct"
		}
		ua := useragent.New(c.GetHeader("User-Agent"))
		device := "Desktop"
		if ua.Mobile() {
			device = "Mobile"
		}
		country := "Unknown"
		ip := c.ClientIP()
		resp, err := http.Get("http://ip-api.com/json/" + ip)
		if err == nil {
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				var ipResponse IPResponse
				body, err := io.ReadAll(resp.Body)
				if err == nil {
					if err := json.Unmarshal(body, &ipResponse); err == nil {
						country = ipResponse.Country
					}
				}
			}
		}
		click := models.Click{
			LinkID:   link.ID,
			Country:  country,
			Referrer: referrer,
			Device:   device,
		}
		db.Create(&click)
		c.Redirect(http.StatusMovedPermanently, link.OriginalURL)
	}
}
