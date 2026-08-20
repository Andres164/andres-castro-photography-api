package schemas

type GetPhotoOutput struct {
	ID       string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUri  string `json:"image_uri"`
}