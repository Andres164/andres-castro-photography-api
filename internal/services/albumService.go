package services

import (
	"andres_castro_photography_api/internal/database"
	"andres_castro_photography_api/internal/models"
	"context"
)

type AlbumService struct {
	photoService *PhotoService
}

func NewAlbumService(photoService *PhotoService) *AlbumService {
	return &AlbumService{
		photoService: photoService,
	}
}

func (s *AlbumService) CreateAlbum(
	ctx context.Context,
	album *models.Album,
) (*models.Album, error) {
	if err := database.DB.Create(album).Error; err != nil {
		return nil, err
	}

	return album, nil;
}

func (s *AlbumService) AddPhotos(
	ctx context.Context,
	photosIDs []uint,
)