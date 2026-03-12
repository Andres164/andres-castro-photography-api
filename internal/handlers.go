package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := DB.AutoMigrate(&Photo{}); err != nil {
		log.Fatal("Failed to migrate schema:", err)
	}
}

func CreatePhoto(c *gin.Context) {
	var photo Photo

	if err := c.ShouldBindJSON(&photo); err != nil {
		ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}
	DB.Create(&photo)
	ResponseJSON(c, http.StatusOK, "Photo created succesfully", photo)
}

func GetPhotoById(c *gin.Context) {
	var photo Photo
	if err := DB.First(&photo, c.Param("id")).Error; err != nil {
		ResponseJSON(c, http.StatusNotFound, "Photo not found", nil)
		return
	}
	ResponseJSON(c, http.StatusOK, "Photo retrieved successfully", photo)
}

func GetPhotos(c* gin.Context) {
	var photos []Photo
	if err := DB.Find(&photos).Error; err != nil {
		ResponseJSON(c, http.StatusInternalServerError, "Error while fetching photos " + err.Error(), nil)
		return
	}
	ResponseJSON(c, http.StatusOK, "Photos fetched succesfully", photos)
}

func DeletePhoto(c* gin.Context) {
	photoId := c.Param("id") // TODO: handle Bad request when id is not provided
	var photo Photo
	if err := DB.Delete(photo, photoId).Error; err != nil {
		ResponseJSON(c, http.StatusInternalServerError, "Error while fetching photos " + err.Error(), nil)
		return
	}
	ResponseJSON(c, http.StatusOK, "Photo deleted succesfully", photo)
}

func UpdatePhoto(c* gin.Context) {
	photoId := c.Param("id") // TODO: handle Bad request when id is not provided
	var photo Photo

	if err := DB.First(&photo, photoId).Error; err != nil {
		ResponseJSON(c, http.StatusNotFound, "Photo not found", nil)
		return
	}

	if err := c.ShouldBindJSON(&photo); err != nil {
		ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	DB.Save(&photo)
	ResponseJSON(c, http.StatusOK, "Photo updated succesfully", photo)
}