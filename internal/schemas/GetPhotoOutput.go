package schemas

type GetPhotoResponse struct {
	ID       	uint `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUri  	string `json:"image_uri"`
}

type GetPhotoOutput struct {
	Body GetPhotoResponse
}