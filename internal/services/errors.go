package services

import "errors"

var ErrAlbumNotFound = errors.New("Album not found")
var ErrPhotosNotFound = errors.New("Photos not found")