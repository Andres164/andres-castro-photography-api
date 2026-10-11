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

func (s *PhotoService) GetPhotoImage(
	ctx context.Context,
	id uint,
) (io.ReadCloser, string, error) {
	var photo models.Photo

	if err := database.DB.First(&photo, id).Error; err != nil {
		return nil, "", err
	}

	reader, err := s.storage.Get(ctx, photo.StorageKey)
	if err != nil {
		return nil, "", err
	}

	return reader, photo.ContentType, nil
}

func (s *PhotoService) GetPhotoById(
	ctx context.Context,
	id uint,
) (*models.Photo, error) {
	var photo *models.Photo

	if err := database.DB.First(&photo, id).Error; err != nil {
		return nil, err
	}

	return photo, nil
}

func (s *PhotoService) DoPhotosExists(
	ctx context.Context,
	photoIDs []uint,
) (bool, error) {
	if len(photoIDs) == 0 {
        return true, nil
    }

	var count int64
	err := database.DB.
		Model(&models.Photo{}).
		Where("id IN ?", photoIDs).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count == int64(len(photoIDs)), nil
}