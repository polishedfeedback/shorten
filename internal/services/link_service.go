package services

import (
	"math/rand"
	"strings"
)

func GenerateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	for range 6 {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}
