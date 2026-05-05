package handlers

import (
	"andres_castro_photography_api/internal/database"
	"andres_castro_photography_api/internal/models"
	"andres_castro_photography_api/internal/schemas"
	"context"
	"fmt"
)

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
