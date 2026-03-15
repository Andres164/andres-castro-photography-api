package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"andres_castro_photography_api/internal/models"
	"andres_castro_photography_api/internal/database"
)

func CreatePhoto(c *gin.Context) {
	var photo models.Photo

	if err := c.ShouldBindJSON(&photo); err != nil {
		models.ResponseJSON(c, http.StatusBadRequest, "Invalid input: " + err.Error(), nil)
		return
	}
	database.DB.Create(&photo)
	models.ResponseJSON(c, http.StatusOK, "Photo created succesfully", photo)
}

func GetPhotoById(c *gin.Context) {
	var photo models.Photo
	if err := database.DB.First(&photo, c.Param("id")).Error; err != nil {
		models.ResponseJSON(c, http.StatusNotFound, "Photo not found", nil)
		return
	}
	models.ResponseJSON(c, http.StatusOK, "Photo retrieved successfully", photo)
}

func GetPhotos(c* gin.Context) {
	var photos []models.Photo
	if err := database.DB.Find(&photos).Error; err != nil {
		models.ResponseJSON(c, http.StatusInternalServerError, "Error while fetching photos " + err.Error(), nil)
		return
	}
	models.ResponseJSON(c, http.StatusOK, "Photos fetched succesfully", photos)
}

func DeletePhoto(c* gin.Context) {
	photoId := c.Param("id") // TODO: handle Bad request when id is not provided
	var photo models.Photo
	if err := database.DB.Delete(&photo, photoId).Error; err != nil {
		models.ResponseJSON(c, http.StatusInternalServerError, "Error while fetching photos " + err.Error(), nil)
		return
	}
	models.ResponseJSON(c, http.StatusOK, "Photo deleted succesfully", photo)
}

func UpdatePhoto(c* gin.Context) {
	photoId := c.Param("id") // TODO: handle Bad request when id is not provided
	var photo models.Photo

	if err := database.DB.First(&photo, photoId).Error; err != nil {
		models.ResponseJSON(c, http.StatusNotFound, "Photo not found", nil)
		return
	}

	if err := c.ShouldBindJSON(&photo); err != nil {
		models.ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	database.DB.Save(&photo)
	models.ResponseJSON(c, http.StatusOK, "Photo updated succesfully", photo)
}