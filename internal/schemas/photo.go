package schemas

type PhotoThumbnailResponse struct {
	ID 		 uint 	`json:"id"`
	ImageURI string `json:"imageUri"`
}

type PhotoImageResponse struct {
	ID 		 uint 	`json:"id"`
	SignedUrl string `json:"signedUrl"`
}

type GetPhotoImageOutput struct {
	Body PhotoImageResponse
}