package models

import (
	"time"

	"gorm.io/gorm"
)

type Album struct {
	gorm.Model;
	Title string;
	Description string;
	CoverPhotoID *uint
	TakenAt *time.Time

	Photos []Photo `gorm:"many2many:album_photos;"`;
}