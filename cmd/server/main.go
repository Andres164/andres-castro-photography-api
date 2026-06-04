package main

import (
	"andres_castro_photography_api/internal/database"
	"andres_castro_photography_api/internal/handlers"
	"andres_castro_photography_api/internal/middleware"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()
	api := humagin.New(r, huma.DefaultConfig("Andres Castro photography API", "0.1.0"))

	protectedGroup := huma.NewGroup(api)
	protectedGroup.UseMiddleware(middleware.AuthMiddleware(api))

	adminGroup := huma.NewGroup(api)
	adminGroup.UseMiddleware(
		middleware.AuthMiddleware(api),
		middleware.RequireAdmin(api),
	)

	huma.Get(api, "/photos", handlers.GetPhotos)
	huma.Get(protectedGroup, "/photos{id}", handlers.GetPhotoById)
	huma.Post(adminGroup, "/photos", handlers.CreatePhoto)
	huma.Delete(adminGroup, "/photos{id}", handlers.DeletePhoto)
	huma.Patch(adminGroup, "/photos{id}", handlers.UpdatePhoto)

	// USERS
	huma.Post(api, "/users/login", handlers.LogIn)
	huma.Get(protectedGroup, "/users", handlers.GetUsers)
	huma.Post(api, "/users", handlers.CreateUser)
	huma.Patch(adminGroup, "/users/{id}", handlers.UpdateUser)
	huma.Delete(adminGroup, "/users/{id}", handlers.DeleteUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}