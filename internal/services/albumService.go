package services

import (
	"andres_castro_photography_api/internal/database"
	"andres_castro_photography_api/internal/models"
	"context"
	"errors"

	"gorm.io/gorm"
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
	albumID uint,
	photoIDs []uint,
) (error) {
	// Validate photo IDs
	doPhotosExists, err := s.photoService.DoPhotosExists(ctx, photoIDs)
	if err != nil {
		return err
	}
	if !doPhotosExists {
		return ErrPhotosNotFound
	}
	
	return database.DB.Transaction(func(tx *gorm.DB) error {
        var album models.Album

        if err := tx.First(&album, albumID).Error; err != nil {
            if errors.Is(err, gorm.ErrRecordNotFound) {
                return ErrAlbumNotFound
            }
            return err
        }

        photos := make([]models.Photo, len(photoIDs))
        for i, id := range photoIDs {
            photos[i].ID = id
        }

        association := tx.Model(&album).Association("Photos")
        if err := association.Error; err != nil {
            return err
        }

        if err := association.Append(&photos); err != nil {
            return err
        }

        return nil
    })
}