package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    int `json:"email"`
	Username int `json:"username"`
	Password int `json:"password"`
	Role     int `json:"role"`
}