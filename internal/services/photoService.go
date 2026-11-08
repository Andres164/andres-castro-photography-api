package services

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	_ "image/jpeg"
	_ "image/png"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"andres_castro_photography_api/internal/database"
	"andres_castro_photography_api/internal/models"
)

type PhotoService struct {
	storage *StorageService
}

func NewPhotoService(storage *StorageService) *PhotoService {
	return &PhotoService{
		storage: storage,
	}
}

const (
    maxUploadSize  = 25 << 20 // 25MB hard cap
    mediumMaxWidth = 1600     // px, long edge; height scales proportionally
    jpegQuality    = 85
)

func (s *PhotoService) CreatePhoto(
    ctx context.Context,
    reader io.Reader,
    photo *models.Photo,
) (*models.Photo, error) {
    limited := io.LimitReader(reader, maxUploadSize+1)
    original, err := io.ReadAll(limited)
    if err != nil {
        return nil, fmt.Errorf("read upload: %w", err)
    }
    if len(original) > maxUploadSize {
        return nil, ErrMaxUploadSizeExceeded // TODO: Inform the max upload size
    }

    img, err := imaging.Decode(bytes.NewReader(original))
    if err != nil {
		log.Printf("failed to decode uploaded image: %v", err)
        return nil, ErrInvalidImageFile
    }

    medium := imaging.Resize(img, mediumMaxWidth, 0, imaging.Lanczos)

    var mediumBuf bytes.Buffer
    if err := imaging.Encode(&mediumBuf, medium, imaging.JPEG, imaging.JPEGQuality(jpegQuality)); err != nil {
        return nil, fmt.Errorf("encode medium image: %w", err)
    }

    storageID := uuid.NewString()
    baseStorageKey := fmt.Sprintf("photos/%s", storageID)
    originalKey := baseStorageKey + "/original"
    mediumKey := baseStorageKey + "/medium.jpg"

    if err := s.storage.Upload(ctx, bytes.NewReader(original), originalKey); err != nil {
        return nil, fmt.Errorf("upload original: %w", err)
    }

    if err := s.storage.Upload(ctx, &mediumBuf, mediumKey); err != nil {
        if delErr := s.storage.Delete(ctx, originalKey); delErr != nil {
            log.Printf("cleanup failed for %s: %v", originalKey, delErr)
        }
        return nil, fmt.Errorf("upload medium: %w", err)
    }

    photo.StorageKey = baseStorageKey

    if err := database.DB.Create(photo).Error; err != nil {
        if delErr := s.storage.Delete(ctx, originalKey); delErr != nil {
            log.Printf("cleanup failed for %s: %v", originalKey, delErr)
        }
        if delErr := s.storage.Delete(ctx, mediumKey); delErr != nil {
            log.Printf("cleanup failed for %s: %v", mediumKey, delErr)
        }
        return nil, fmt.Errorf("save photo record: %w", err)
    }

    return photo, nil
}

func (s *StorageService) Delete(ctx context.Context, objectName string) error {
    if err := s.bucket.Object(objectName).Delete(ctx); err != nil {
        return fmt.Errorf("delete object %q: %w", objectName, err)
    }
    return nil
}

type ImageURLs struct {
	OriginalURL string
	MediumURL 	string
}

func (s *PhotoService) GetPhotoImage(
	ctx context.Context,
	id uint,
) (*ImageURLs, error) {
	var photo models.Photo

	if err := database.DB.First(&photo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPhotoNotFound
		}
		return nil, err
	}

	originalKey := fmt.Sprintf("%s/original", photo.StorageKey)
	mediumKey := fmt.Sprintf("%s/medium.jpg", photo.StorageKey)
	originalURL, err := s.storage.Get(ctx, originalKey)
	if err != nil {
		return nil, fmt.Errorf("get signedURL: %w", err)
	}
	mediumURL, err := s.storage.Get(ctx, mediumKey)
	if err != nil {
		return nil, fmt.Errorf("get signedURL: %w", err)
	}

	return &ImageURLs{
		OriginalURL: originalURL,
		MediumURL: mediumURL,
	}, nil
}

func (s *PhotoService) GetPhotoById(
	ctx context.Context,
	id uint,
) (*models.Photo, error) {
	var photo *models.Photo

	if err := database.DB.First(&photo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPhotoNotFound
		}
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

	photoIDsLenght := int64(len(photoIDs))
	return count == photoIDsLenght, nil
}