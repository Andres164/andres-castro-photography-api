package main

import (
	"andres_castro_photography_api/internal/database"
	"andres_castro_photography_api/internal/models"
	"andres_castro_photography_api/internal/schemas"
	"context"
	"fmt"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()
	api := humagin.New(r, huma.DefaultConfig("Andres Castro photography API", "0.1.0"))

	huma.Get(api, "/photos", func(ctx context.Context, input *struct{}) (*schemas.GetPhotosResponse, error) {

	var photos []models.Photo

	if err := database.DB.Find(&photos).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch photos: %w", err)
	}

	return &schemas.GetPhotosResponse{
		Body: photos,
	}, nil
})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}