package main

import (
	"andres_castro_photography_api/internal/database"
	"andres_castro_photography_api/internal/handlers"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()
	api := humagin.New(r, huma.DefaultConfig("Andres Castro photography API", "0.1.0"))

	huma.Get(api, "/photos", handlers.GetPhotos)
	huma.Get(api, "/photos{id}", handlers.GetPhotoById)
	huma.Post(api, "/photos", handlers.CreatePhoto)
	huma.Delete(api, "/photos{id}", handlers.DeletePhoto)
	huma.Patch(api, "/photos{id}", handlers.UpdatePhoto)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}