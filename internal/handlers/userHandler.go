package handlers

import (
	"context"
	"andres_castro_photography_api/internal/database"
	"andres_castro_photography_api/internal/models"
	"andres_castro_photography_api/internal/schemas"
)

func CreateUser(ctx context.Context, input *schemas.CreateUserInput) (*schemas.UserOutput, error) {
	var user models.User
	newUser := input.Body

	user.Email = newUser.Email
	user.Username = newUser.Username
	user.Password = newUser.Password
	user.Role = newUser.Role

	database.DB.Create(&user)
	return &schemas.UserOutput{
		Body: struct{
			Emai
		},
	}, nil
}
