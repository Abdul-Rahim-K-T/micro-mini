package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"micro-mini/shared/helpers"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func compareHashPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

type adminDetail struct {
	Adminname string `json:"username"`
	Password  string `json:"password"`
}

func Login(c *gin.Context) {
	var admin adminDetail
	if err := c.Bind(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}

	// Getting data from env files
	username := os.Getenv("ADMIN")
	password := os.Getenv("ADMIN_PASSWORD")
	fmt.Println("username:", username)
	fmt.Println("password:", password)

	fmt.Println("json username:", admin.Adminname)
	fmt.Println("json password:", admin.Password)
	// Checking username and password
	if username != admin.Adminname || password != admin.Password {
		c.JSON(401, gin.H{
			"error": "Unauthorized access incorrect username or password",
		})
		return
	}

	// Generate token
	token, err := helpers.GenerateJWTToken(username, "admin", "", 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Set token into browser
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("jwt_token", token, 3600*24, "", "", true, true)
	// success message
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func Dashboard(c *gin.Context) {
	admin, _ := c.Get("admin")
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to admin dashboard", "admin": admin})
}

func Logout(c *gin.Context) {
	admin, exists := c.Get("admin")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No admin data found"})
		return
	}

	// Type assert the admin data to the expected type

	c.SetCookie("jwt_token", "", -1, "", "", false, false)
	c.JSON(http.StatusOK, gin.H{"message": admin.(adminDetail).Adminname + " successfully logged out"})
}
