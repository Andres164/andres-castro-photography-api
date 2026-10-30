package services

import (
	"andres_castro_photography_api/internal/database"
	"andres_castro_photography_api/internal/models"
	"context"
	"errors"
	"fmt"

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
	// TODO: Remove duplicate photo IDs

	// Validate photo IDs exists
	doPhotosExists, err := s.photoService.DoPhotosExists(ctx, photoIDs)
	if err != nil {
		return err
	}
	if !doPhotosExists {
		return ErrPhotosNotFound
	}

	// TODO: Remove Photo IDs which already exist in the Album
	return database.DB.Transaction(func(tx *gorm.DB) error {
		var album models.Album
		
		// TODO: take a look at db.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&users) for use instead of transaction
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

func (s *AlbumService) GetAlbums() ([]models.Album, error) {
	// Get albums from the database
	var albums []models.Album // TODO: optimize Find query, we dont need the photos column
	if err := database.DB.Find(&albums).Error; err != nil {
		return nil, err
	}

	return albums, nil
}

func (s *AlbumService) GetAlbumByID(ID uint) (*models.Album, error) {
	var album *models.Album
	if err := database.DB.
	Preload("Photos").
	First(&album, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAlbumNotFound
		}

		return nil, err
	}

	return album, nil
}

func (s *AlbumService) UpdateCoverPhoto(albumID uint, photoID uint) (error) {
	var album *models.Album
	if err := database.DB.First(&album, albumID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrAlbumNotFound
		}

		return err
	}

	var count int64
	err := database.DB.
        Table("album_photos").
        Where("album_id = ? AND photo_id = ?", albumID, photoID).
        Count(&count).Error

	if err != nil {
		return fmt.Errorf("check album photo association: %w", err)
	}
	if count < 1 {
		return ErrPhotoNotInAlbum
	}

	if err := database.DB.
		Model(&album).
		Update("cover_photo_id", photoID).Error; err != nil {
		return fmt.Errorf("update album %d cover: %w", albumID, err)
	}

	return nil
}