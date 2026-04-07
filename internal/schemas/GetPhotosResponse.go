package schemas

import "andres_castro_photography_api/internal/models"

type GetPhotosResponse struct {
	Body []models.Photo
}