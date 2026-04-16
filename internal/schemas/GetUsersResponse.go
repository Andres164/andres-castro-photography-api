package schemas

import "andres_castro_photography_api/internal/models"

type GetUsersResponse struct {
	Body []models.User
}