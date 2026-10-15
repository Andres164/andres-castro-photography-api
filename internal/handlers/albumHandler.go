package handlers

import (
	"andres_castro_photography_api/internal/models"
	"andres_castro_photography_api/internal/schemas"
	"andres_castro_photography_api/internal/services"
	"context"
	"errors"

	"github.com/danielgtaylor/huma/v2"
)

type AlbumHandler struct {
	service *services.AlbumService
}

func NewAlbumHandler(albumService *services.AlbumService) *AlbumHandler {
	return &AlbumHandler{
		service: albumService,
	}
}

func (h *AlbumHandler) CreateAlbum(ctx context.Context, input *schemas.CreateAlbumInput) (*schemas.CreateAlbumOutput, error) {
	newAlbum := &models.Album{
		Title: input.Body.Title,
		Description: input.Body.Description,
	}

	createdAlbum, err := h.service.CreateAlbum(ctx, newAlbum)
	if err != nil {
		return nil, huma.Error500InternalServerError("Error al crear album", err)
	}

	albumResponse := &schemas.AlbumSummaryResponse{
		ID: createdAlbum.ID,
		Title: createdAlbum.Title,
		Description: createdAlbum.Description,
		CoverImageURI: "",
	}

	return &schemas.CreateAlbumOutput{
		Body: *albumResponse,
	}, nil
}

func (h *AlbumHandler) AddPhotos(ctx context.Context, input *schemas.AddAlbumPhotosInput) error {
	err := h.service.AddPhotos(ctx, input.ID, input.Body.PhotoIDs)

	if err != nil {
		if errors.Is(err, services.ErrPhotosNotFound) {
			return huma.Error404NotFound("Una o mas fotos no existen", err)
		} 
		if errors.Is(err, services.ErrAlbumNotFound) {
			return huma.Error404NotFound("El album que quiere actualizar no existe", err)
		}

		return huma.Error500InternalServerError("Error al agregar fotos al album", err)
	}

	return nil
}