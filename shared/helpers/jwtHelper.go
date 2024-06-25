package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWTToken(username, role, email string, userID int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &jwt.MapClaims{
		"sub":      userID,
		"username": username, // Ensure this matches with the extraction in middleware
		"role":     role,
		"exp":      expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET"))) // Use the same key here
}
