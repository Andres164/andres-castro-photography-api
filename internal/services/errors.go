package services

import "errors"

var ErrAlbumNotFound = errors.New("Album not found")
var ErrPhotosNotFound = errors.New("Photos not found")
var ErrPhotoNotInAlbum = errors.New("photo does not belong to album")

// PhotoService
var ErrPhotoNotFound = errors.New("Photo not found")