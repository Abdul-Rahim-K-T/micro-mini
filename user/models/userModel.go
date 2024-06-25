package models

import "time"

type User struct {
	User_ID      uint      `gorm:"primaryKey" json:"-"`
	First_Name   string    `json:"first_name" validate:"required"`
	Last_Name    string    `json:"last_name" validate:"required"`
	User_Name    string    `json:"user_name" validate:"required"`
	Password     string    `json:"password" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	Otp          uint      `json:"otp"`
	Phone_Number string    `json:"phone_number" validate:"required"`
	Created_at   time.Time `json:"created_at"`
	Referal_Code string    `json:"referal_code"`
	IsBlocked    bool      `json:"is_blocked"`
	Validate     bool      `json:"validate"`
}
