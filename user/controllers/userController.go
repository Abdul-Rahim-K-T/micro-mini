package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	"micro-mini/shared/database"
	"micro-mini/shared/helpers"
	"micro-mini/user/models"
)

var validate = validator.New()

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func compareHashPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func Signup(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}

	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	database.DB.Where("user_name = ?", user.User_Name).First(&existingUser)
	if existingUser.User_Name == user.User_Name {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already taken"})
		return
	}

	database.DB.Where("phone_number = ?", user.Phone_Number).First(&existingUser)
	if existingUser.Phone_Number == user.Phone_Number {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number already taken"})
		return
	}

	if len(user.Phone_Number) != 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone number"})
		return
	}

	password, err := hashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	otp := helpers.GenerateOtp()
	fmt.Println(otp)
	helpers.SendOtp(strconv.Itoa(otp), user.Email)

	referalCode := helpers.RandomStringGenerator()

	database.DB.Create(&models.User{
		First_Name:   user.First_Name,
		Last_Name:    user.Last_Name,
		User_Name:    user.User_Name,
		Password:     password,
		Email:        user.Email,
		Otp:          uint(otp),
		Phone_Number: user.Phone_Number,
		Created_at:   time.Now(),
		Referal_Code: referalCode,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Signup successful, validate OTP"})
}

func ValidateOtp(c *gin.Context) {
	var data struct {
		Email string `json:"email"`
		Otp   int    `json:"otp"`
	}

	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}

	var user models.User
	database.DB.Where("email = ?", data.Email).First(&user)

	if user.Otp == uint(data.Otp) {
		database.DB.Model(&models.User{}).Where("email = ?", data.Email).Update("validate", true)
		c.JSON(http.StatusOK, gin.H{"message": "Account validated successfully"})
	} else {
		database.DB.Where("validate = ?", false).Delete(&models.User{})
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OTP or email"})
	}
}

func Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}

	var user models.User
	database.DB.Where("user_name = ?", credentials.Username).First(&user)

	if user.IsBlocked {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User is blocked"})
		return
	}

	if !compareHashPassword(user.Password, credentials.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := helpers.GenerateJWTToken(user.User_Name, "user", user.Email, int(user.User_ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("jwt_token", token, 3600*24, "", "", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func Logout(c *gin.Context) {
	user, _ := c.Get("user")
	c.SetCookie("jwt_token", "", -1, "", "", false, false)
	c.JSON(http.StatusOK, gin.H{"message": user.(models.User).User_Name + " successfully logged out"})
}
