package helpers

import (
	"errors"
	"regexp"
)

// ValidateEmail checks if the provided email has a valid format.
func ValidateEmail(email string) error {
	// Regex for validating email
	const emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

// ValidatePassword checks if the provided password meets the criteria.
func ValidatePassword(password string) error {
	// Basic length check
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	// Regex for checking if the password contains at least one number
	numberRegex := regexp.MustCompile(`[0-9]`)
	if !numberRegex.MatchString(password) {
		return errors.New("password must include at least one number")
	}

	// Regex for checking if the password contains at least one uppercase letter
	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	if !uppercaseRegex.MatchString(password) {
		return errors.New("password must include at least one uppercase letter")
	}

	// Regex for checking if the password contains at least one lowercase letter
	lowercaseRegex := regexp.MustCompile(`[a-z]`)
	if !lowercaseRegex.MatchString(password) {
		return errors.New("password must include at least one lowercase letter")
	}

	return nil
}

// ValidateUsername checks if the provided username meets the criteria.
func ValidateUsername(username string) error {
	// Username should be alphanumeric and between 3 to 20 characters
	const usernameRegex = `^[a-zA-Z0-9]{3,20}$`
	re := regexp.MustCompile(usernameRegex)
	if !re.MatchString(username) {
		return errors.New("username must be alphanumeric and between 3 to 20 characters")
	}
	return nil
}
