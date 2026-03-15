package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title string `json:"title"`
	Description string `json:"description"`
	Url string `json:"url"`
}