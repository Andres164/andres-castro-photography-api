package main

import (
	api "andres_castro_photography_api/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	api.InitDB()
	r := gin.Default()

	r.GET("/photos", api.GetPhotos)
	r.GET("/photos/:id", api.GetPhotoById)
	r.POST("/photos", api.CreatePhoto)
	r.PUT("/photos/:id", api.UpdatePhoto)
	r.DELETE("/photos/:id", api.DeletePhoto)

	r.Run(":8080")
}