package handlers

import (
	"context"
	"errors"

	"andres_castro_photography_api/internal/database"
	"andres_castro_photography_api/internal/models"
	"andres_castro_photography_api/internal/schemas"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
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

func GetPhotoById(ctx context.Context, input *schemas.PhotoIdInput) (*schemas.GetPhotoByIdResponse, error) {

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
		return nil, huma.Error500InternalServerError("Error al buscar fotos: %w", err)
	}

	return &schemas.GetPhotosResponse{
		Body: photos,
	}, nil
}

func DeletePhoto(ctx context.Context, input *schemas.PhotoIdInput) (*schemas.GetPhotoByIdResponse, error) {
	var photo models.Photo
	if err := database.DB.Delete(&photo, input.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, huma.Error404NotFound("Foto no encontrada");
		}

		return nil, huma.Error500InternalServerError("Error al eliminar la foto: %w", err)
	}

	response := &schemas.GetPhotoByIdResponse{
		Body: photo,
	}
	return response, nil;
}

func UpdatePhoto(ctx context.Context, input *schemas.UpdatePhotoInput) (*schemas.GetPhotoByIdResponse, error) {
	var photo models.Photo

	if err := database.DB.First(&photo, input.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, huma.Error404NotFound("Foto no encontrada");
		}

		return nil, huma.Error500InternalServerError("Error al actualizar foto: %w", err)
	}

	if input.Body.Title != nil {
		photo.Title = *input.Body.Title
	}
	if input.Body.Description != nil {
		photo.Description = *input.Body.Description
	}
	if input.Body.Url != nil {
		photo.Url = *input.Body.Url
	}

	if err := database.DB.Save(&photo).Error; err != nil {
		return nil, huma.Error500InternalServerError("Error al actualizar foto: %w", err)
	}

	return &schemas.GetPhotoByIdResponse{
		Body: photo,
	}, nil
}
