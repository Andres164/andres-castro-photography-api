package schemas

import "time"

type AlbumSummaryResponse struct  {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CoverImageURI string `json:"coverImageUri"`
	TakenAt time.Time `json:"takenAt"`
}

type AlbumDetailedResponse struct  {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CoverImageURI string `json:"coverImageUri"`
	TakenAt time.Time `json:"takenAt"`
	Photos []PhotoThumbnailResponse `json:"photos"`
}

type GetAlbumInput struct {
	ID uint `path:"id"`
}

type GetAlbumOutput struct {
	Body AlbumDetailedResponse
}

type GetAlbumsOutput struct {
	Body []AlbumSummaryResponse
}

type CreateAlbumRequest struct {
	Title string `json:"title" minlength:"1" required:"true"`
	Description string `json:"description"`
	TakenAt string `json:"takenAt"`
}

type CreateAlbumInput struct {
	Body CreateAlbumRequest
}

type CreateAlbumOutput struct {
	Body AlbumSummaryResponse
}

type AddAlbumPhotosRequest struct {
    PhotoIDs []uint `json:"photoIds" minItems:"1" required:"true"`
}

type AddAlbumPhotosInput struct {
	ID uint `path:"id"`
	Body AddAlbumPhotosRequest
}

type SetAlbumCoverRequest struct {
    PhotoID uint `json:"photoId" required:"true"`
}

type SetAlbumCoverInput struct {
    ID   uint `path:"id"`
    Body SetAlbumCoverRequest
}
