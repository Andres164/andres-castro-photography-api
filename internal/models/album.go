package models

import (
	"gorm.io/gorm"
)

type Album struct {
	gorm.Model;
	Title string;
	Description string;
	CoverPhotoID *uint

	Photos []Photo `gorm:"many2many:album_photos;"`;
}