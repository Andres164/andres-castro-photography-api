package routes

import "github.com/gin-gonic/gin"

func RegisterPublicEndpoints(router *gin.Engine, photoHandler *handlers.Photo) {
	router.GET("/photos", photoHandler.GetAllPhotos)
	router.GET("/photos/:id", photoHandler.GetPhoto)
	router.POST("/photos", photoHandler.CreatePhoto)
	router.PUT("/photos/:id", photoHandler.UpdatePhoto)
	router.DELETE("/photos/:id", photoHandler.DeletePhoto)
}