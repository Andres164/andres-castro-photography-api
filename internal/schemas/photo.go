package schemas

import "github.com/danielgtaylor/huma/v2"

type PhotoIdInput struct {
	ID uint `path:"id"`
}

type GetPhotoResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUri    string `json:"image_uri"`
	ContentType string `json:"content_type"`
}

type GetPhotoOutput struct {
	Body GetPhotoResponse
}

type GetPhotosOutput struct {
	Body []GetPhotoResponse
}

type CreatePhotoInput struct {
	RawBody huma.MultipartFormFiles[ struct {
		File        huma.FormFile `form:"file" required:"true"`
		Title       string        `form:"title" required:"true"`
		Description string        `form:"description"`
		ContentType string        `form:"content_type" required:"true"`
	}]
}

type UpdatePhotoInput struct {
	ID uint `path:"id"`

	Body struct {
		Title       *string `json:"title,omitempty"`
		Description *string `json:"description,omitempty"`
		Url *string `json:"url,omitempty"`
		ContentType *string `json:"content_type,omitempty"`
	} `json:"body"`
}

type PhotoThumbnailResponse struct {
	ID       uint   `json:"id"`
	ImageURI string `json:"imageUri"`
}

type PhotoImageResponse struct {
	ID        uint   `json:"id"`
	SignedUrl string `json:"signedUrl"`
}

type GetPhotoImageOutput struct {
	Body PhotoImageResponse
}
