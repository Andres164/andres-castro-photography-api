package handlers

import (
	"context"
	"errors"

	"andres_castro_photography_api/internal/database"
	"andres_castro_photography_api/internal/models"
	"andres_castro_photography_api/internal/schemas"
	"andres_castro_photography_api/internal/services"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

type PhotoHandler struct {
	service *services.PhotoService
}

func NewPhotoHandler(photoService *services.PhotoService) *PhotoHandler {
	return &PhotoHandler{
		service: photoService,
	}
}

func (h *PhotoHandler) CreatePhoto(ctx context.Context, input *schemas.CreatePhotoInput) (*schemas.GetPhotoByIdResponse, error) {
	form := input.RawBody.Data()

	photo := &models.Photo{
		Title:       form.Title,
		Description: form.Description,
	}

	createdPhoto, err := h.service.CreatePhoto(ctx, form.File, photo)

	if err != nil {
		return nil, huma.Error500InternalServerError("Error al crear la foto", err)
	}

	return &schemas.GetPhotoByIdResponse{
		Body: *createdPhoto,
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
			return nil, huma.Error404NotFound("Foto no encontrada")
		}

		return nil, huma.Error500InternalServerError("Error al eliminar la foto: %w", err)
	}

	response := &schemas.GetPhotoByIdResponse{
		Body: photo,
	}
	return response, nil
}

func UpdatePhoto(ctx context.Context, input *schemas.UpdatePhotoInput) (*schemas.GetPhotoByIdResponse, error) {
	var photo models.Photo

	if err := database.DB.First(&photo, input.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, huma.Error404NotFound("Foto no encontrada")
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
		photo.StorageKey = *input.Body.Url
	}

	if err := database.DB.Save(&photo).Error; err != nil {
		return nil, huma.Error500InternalServerError("Error al actualizar foto: %w", err)
	}

	return &schemas.GetPhotoByIdResponse{
		Body: photo,
	}, nil
}
