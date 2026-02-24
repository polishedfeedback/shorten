package services

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/polishedfeedback/shorten/internal/models"
	"gorm.io/gorm"
)

func GenerateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	for range 6 {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func CreateLink(db *gorm.DB, userID uint, originalURL string) (models.Link, error) {
	shortCode := GenerateShortCode()
	link := models.Link{
		OriginalURL: originalURL,
		ShortCode:   shortCode,
		UserID:      userID,
	}
	result := db.Create(&link)
	if result.Error != nil {
		return models.Link{}, fmt.Errorf("error creating a link: %v", result.Error)
	}
	return link, nil
}

func GetLinks(db *gorm.DB, userID uint) ([]models.Link, error) {
	var links []models.Link

	result := db.Where("user_id = ?", userID).Find(&links)
	if result.Error != nil {
		return []models.Link{}, fmt.Errorf("error querying links: %v", result.Error)
	}
	return links, nil
}

func DeleteLink(db *gorm.DB, id uint, userID uint) error {
	var link models.Link
	result := db.Where("id = ? AND user_id = ?", id, userID).Delete(&link)
	if result.Error != nil {
		return fmt.Errorf("error deleting the link: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("link not found or unauthorized")
	}
	return nil
}
