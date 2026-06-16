package services

import "mime/multipart"

type StorageService interface {
	UploadPhoto(
		file multipart.File,
		filename string,
	) (string, error)
}