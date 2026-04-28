package handlers

import (
	"context"
	"fmt"
	"net/http"

	"andres_castro_photography_api/internal/database"
	"andres_castro_photography_api/internal/models"
	"andres_castro_photography_api/internal/schemas"
	"github.com/danielgtaylor/huma/v2"

	"github.com/gin-gonic/gin"
)

func CreatePhoto(ctx context.Context, input *schemas.CreatePhotoRequest) (*schemas.GetPhotoByIdResponse, error) {
	var photo models.Photo
	newPhoto := input.Body

	photo.Title = newPhoto.Title
	photo.Description = newPhoto.Description
	photo.Url = newPhoto.Url

	database.DB.Create(&photo)
	return &schemas.GetPhotoByIdResponse{
		Body: photo,
	}, nil
}

func GetPhotoById(ctx context.Context, input *schemas.GetPhotoByIdRequest) (*schemas.GetPhotoByIdResponse, error) {

	var photo models.Photo

	if err := database.DB.First(&photo, input.ID).Error; err != nil {
		return nil, huma.Error404NotFound("Photo not found")
	}

	return &schemas.GetPhotoByIdResponse{
		Body: photo,
	}, nil
}

func GetPhotos(ctx context.Context, input *struct{}) (*schemas.GetPhotosResponse, error) {
	var photos []models.Photo
	if err := database.DB.Find(&photos).Error; err != nil {
		return nil, fmt.Errorf("Error al buscar fotos: %w", err)
	}

	return &schemas.GetPhotosResponse{
		Body: photos,
	}, nil
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