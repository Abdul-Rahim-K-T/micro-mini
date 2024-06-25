package helpers

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"

	"gopkg.in/gomail.v2"
)

// GenerateOtp generates a random 6-digit OTP.
func GenerateOtp() int {
	max := big.NewInt(1000000)
	otp, _ := rand.Int(rand.Reader, max)
	return int(otp.Int64())
}

// SendOtp sends the OTP to the specified email address.
func SendOtp(otp string, email string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("EMAIL"))
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Your OTP Code")
	mailer.SetBody("text/plain", fmt.Sprintf("Your OTP code is: %s", otp))

	dialer := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"), os.Getenv("PASSWORD"))

	return dialer.DialAndSend(mailer)
}

// RandomStringGenerator generates a random string of the specified length.
func RandomStringGenerator() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, 8)
	for i := range result {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		result[i] = letters[num.Int64()]
	}
	return string(result)
}
