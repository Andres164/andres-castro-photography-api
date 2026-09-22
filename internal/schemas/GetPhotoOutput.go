package schemas

type GetPhotoResponse struct {
	ID       	uint `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUri  	string `json:"image_uri"`
	ContentType string `json:"content_type"`
}

type GetPhotoOutput struct {
	Body GetPhotoResponse
}