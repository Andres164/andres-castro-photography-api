package main

import (
	"andres_castro_photography_api/internal/database"
	"andres_castro_photography_api/internal/handlers"
	"andres_castro_photography_api/internal/middleware"
	"net/http"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()
	config := huma.DefaultConfig("Andres Castro photography API", "0.1.0")
	config.Components.SecuritySchemes = map[string]*huma.SecurityScheme {
		"bearerAuth": {
			Type: "http",
			Name: "Authorization",
			In: "header",
			Scheme: "bearer",
			BearerFormat: "JWT",
		},
	}
	config.Security = []map[string][]string{
		{
			"bearerAuth": {},
		},
	}
	
	api := humagin.New(r, config)

	protectedGroup := huma.NewGroup(api)
	protectedGroup.UseMiddleware(middleware.AuthMiddleware(api))

	adminGroup := huma.NewGroup(api)
	adminGroup.UseMiddleware(
		middleware.AuthMiddleware(api),
		middleware.RequireAdmin(api),
	)

	// Users
	huma.Register(adminGroup, huma.Operation{
		OperationID: "get-users",
		Method: http.MethodGet,
		Path: "/users",
	}, handlers.GetUsers)

	// Photos
	huma.Register(api, huma.Operation{
		OperationID: "get-photos",
		Method: http.MethodGet,
		Path: "/photos",
	}, handlers.GetPhotos)

	huma.Register(protectedGroup, huma.Operation{
		OperationID: "get-photo-by-id",
		Method: http.MethodGet,
		Path: "/photos/{id}",
	}, handlers.GetPhotoById)

	huma.Post(adminGroup, "/photos", handlers.CreatePhoto)
	huma.Delete(adminGroup, "/photos/{id}", handlers.DeletePhoto)
	huma.Patch(adminGroup, "/photos/{id}", handlers.UpdatePhoto)

	// USERS
	huma.Post(api, "/users/login", handlers.LogIn)
	huma.Post(api, "/users", handlers.CreateUser)
	huma.Patch(adminGroup, "/users/{id}", handlers.UpdateUser)
	huma.Delete(adminGroup, "/users/{id}", handlers.DeleteUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}