package schemas

import "github.com/danielgtaylor/huma/v2"

type CreatePhotoInput struct {
	RawBody huma.MultipartFormFiles[struct {
		File        huma.FormFile `form:"file" required:"true"`
		Title       string        `form:"title" required:"true"`
		Description string        `form:"description"`
	}]
}
