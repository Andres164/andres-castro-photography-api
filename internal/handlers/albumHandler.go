package handlers

import (
	"andres_castro_photography_api/internal/models"
	"andres_castro_photography_api/internal/schemas"
	"andres_castro_photography_api/internal/services"
	"context"
	"errors"
	"fmt"

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

func (h *AlbumHandler) AddPhotos(ctx context.Context, input *schemas.AddAlbumPhotosInput) (*struct{}, error) {
	err := h.service.AddPhotos(ctx, input.ID, input.Body.PhotoIDs)

	if err != nil {
		if errors.Is(err, services.ErrPhotosNotFound) {
			return nil, huma.Error404NotFound("Una o mas fotos no existen", err)
		} 
		if errors.Is(err, services.ErrAlbumNotFound) {
			return nil, huma.Error404NotFound("El album que quiere actualizar no existe", err)
		}

		return nil, huma.Error500InternalServerError("Error al agregar fotos al album", err)
	}

	return nil, nil
}

func (h *AlbumHandler) GetAlbums(ctx context.Context, input *struct{}) (*schemas.GetAlbumsOutput, error) {
	albums, err := h.service.GetAlbums()
	if err != nil {
		return nil, huma.Error500InternalServerError("Error al obtener albumes")
	}

	var albumResponseCollection []schemas.AlbumSummaryResponse
	for _, a := range albums {
		coverImageURI := ""
		if a.CoverPhotoID != nil {
			coverImageURI = fmt.Sprintf("photo/%d/image", a.CoverPhotoID)
		}
		albumResponseCollection = append(albumResponseCollection, schemas.AlbumSummaryResponse{
			ID: a.ID,
			Title: a.Title,
			Description: a.Description,
			CoverImageURI: coverImageURI,
		})
	}

	return &schemas.GetAlbumsOutput{
		Body: albumResponseCollection,
	}, nil
}

func (h *AlbumHandler) GetAlbumByID(ctx context.Context, input *schemas.GetAlbumInput) (*schemas.GetAlbumOutput, error) {
	album, err := h.service.GetAlbumByID(input.ID)
	if err != nil {
		if errors.Is(err, services.ErrAlbumNotFound) {
			return nil, huma.Error404NotFound("No se encontro un album con el ID proporcionado")
		}
		return nil, err
	}

	var albumPhotosResCollection []schemas.GetPhotoResponse
	for _, p := range album.Photos {
		imageURI := fmt.Sprintf("photo/%d/image", p.ID)
		photoResponse := &schemas.GetPhotoResponse{
			ID: p.ID,
			Title: p.Title,
			Description: p.Description,
			ImageUri: imageURI,
			ContentType: p.ContentType,
		}
		albumPhotosResCollection = append(albumPhotosResCollection, *photoResponse)
	}

	coverImageURI := ""
	if album.CoverPhotoID != nil {
		coverImageURI = fmt.Sprintf("photo/%d/image", album.CoverPhotoID)
	}
	albumResponse := &schemas.AlbumDetailedResponse{
		ID: album.ID,
		Title: album.Title,
		Description: album.Description,
		CoverImageURI: coverImageURI,
		Photos: albumPhotosResCollection,
	}

	return &schemas.GetAlbumOutput{
		Body: *albumResponse,
	}, nil
}