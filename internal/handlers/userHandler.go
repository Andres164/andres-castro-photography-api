package handlers

import (
	"andres_castro_photography_api/internal/database"
	"andres_castro_photography_api/internal/models"
	"andres_castro_photography_api/internal/schemas"
	"context"
	"fmt"

	"github.com/danielgtaylor/huma/v2"
)

// TEST
func CreateUser(ctx context.Context, input *schemas.CreateUserInput) (*schemas.UserOutput, error) {
	var user models.User
	newUser := input.Body

	user.Email = newUser.Email
	user.Username = newUser.Username
	user.Password = newUser.Password
	user.Role = newUser.Role

	if err := database.DB.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("Error al crear usuario: %w", err)
	}
	
	return &schemas.UserOutput{
		Body: schemas.UserResponse{
			Email:    user.Email,
			Username: user.Username,
			Role:     user.Role,
		},
	}, nil
}

func GetUsers(ctx context.Context, input *struct{}) (*schemas.GetUsersOutput, error) {
    var users []models.User

    if err := database.DB.Find(&users).Error; err != nil {
        return nil, huma.Error500InternalServerError("Error al buscar los usuarios: %w", err)
    }

    responses := make([]schemas.UserResponse, len(users))
    for i, user := range users {
        responses[i] = schemas.UserResponse{
            Email:    user.Email,
            Username: user.Username,
            Role:     user.Role,
        }
    }

    return &schemas.GetUsersOutput{
        Body: responses,
    }, nil
}