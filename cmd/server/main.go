package main

import (
	"andres_castro_photography_api/internal/database"
	"andres_castro_photography_api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()

	r.GET("/photos", handlers.GetPhotos)
	r.GET("/photos/:id", handlers.GetPhotoById)
	r.POST("/photos", handlers.CreatePhoto)
	r.PATCH("/photos/:id", handlers.UpdatePhoto)
	r.DELETE("/photos/:id", handlers.DeletePhoto)

	r.Run(":8080")
}