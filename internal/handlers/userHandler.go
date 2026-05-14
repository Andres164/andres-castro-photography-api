package handlers

import (
	"andres_castro_photography_api/internal/database"
	"andres_castro_photography_api/internal/models"
	"andres_castro_photography_api/internal/schemas"
	"context"
	"errors"
	"fmt"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
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

func DeleteUser(ctx context.Context, input *schemas.UserIdInput) (*schemas.UserOutput, error) {
	var user models.User
	if err := database.DB.Delete(&user, input.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, huma.Error404NotFound("Usuario no encontrado")
		}

		return nil, huma.Error500InternalServerError("Error al eliminar usuario: %w", err)
	}

	deletedUserResponse := schemas.UserResponse{
		Email: user.Email,
		Username: user.Username,
		Role: user.Role,
	}

	return &schemas.UserOutput{
		Body: deletedUserResponse,
	}, nil
}

func UpdateUser(ctx context.Context, input *schemas.UpdateUserInput) (*schemas.UserOutput, error) {
	var user models.User

	if err := database.DB.First(&user, input.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, huma.Error404NotFound("Usuario no encontrado");
		}

		return nil, huma.Error500InternalServerError("Error al actualizar usuario: %w", err)
	}

	if input.Body.Email != nil {
		user.Email = *input.Body.Email
	}
	if input.Body.Password != nil {
		user.Password = *input.Body.Password
	}
	if input.Body.Role != nil {
		user.Role = *input.Body.Role
	}
	if input.Body.Username != nil {
		user.Username = *input.Body.Username
	}

	if err := database.DB.Save(&user).Error; err != nil {
		return nil, huma.Error500InternalServerError("Error al actualizar usuario: %w", err)
	}

	updatedUser := &schemas.UserResponse{
		Email: user.Email,
		Username: user.Username,
		Role: user.Role,
	}

	return &schemas.UserOutput{
		Body: *updatedUser,
	}, nil
}