package services

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/polishedfeedback/shorten/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CreateUser takes username and password and stores it in the database
func CreateUser(db *gorm.DB, username, password string) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, fmt.Errorf("error while hashing password: %v", err)
	}

	user := models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	result := db.Create(&user)
	if result.Error != nil {
		return models.User{}, fmt.Errorf("error saving user: %v", result.Error)
	}

	return user, nil
}

// Login User takes in username and password and logs the user in
func LoginUser(db *gorm.DB, username, password string) (string, error) {
	var user models.User

	result := db.First(&user, "username = ?", username)
	if result.Error == gorm.ErrRecordNotFound {
		return "", fmt.Errorf("user not found: %v", result.Error)
	}
	if result.Error != nil {
		return "", fmt.Errorf("error finding user: %v", result.Error)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", fmt.Errorf("error while generating token: %v", err)
	}

	return tokenString, nil
}
