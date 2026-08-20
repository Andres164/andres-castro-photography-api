package services

import (
	"andres_castro_photography_api/internal/database"
	"andres_castro_photography_api/internal/models"
	"context"
	"io"
)

type PhotoService struct {
	storage *StorageService
}

func NewPhotoService(storage *StorageService) *PhotoService {
	return &PhotoService{
		storage: storage,
	}
}

func (s *PhotoService) CreatePhoto(
	ctx context.Context,
	reader io.Reader,
	photo *models.Photo,
) (*models.Photo, error) {
	// Receive uploaded file
	// Upload file using StorageService
	url, err := s.storage.Upload(ctx, reader, photo.Title)

	if err != nil {
		return nil, err
	}
	// Save to PostgreSQL
	photo.StorageKey = url
	// Return Photo
	if err := database.DB.Create(photo).Error; err != nil {
		return nil, err
	}

	return photo, nil
}
