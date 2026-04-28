package schemas

type CreatePhotoRequest struct {
	Body struct {
		Title string `json:"title" doc:"Main title for the photo" minLength:"1"`
		Description string `json:"description,omitempty"`
		Url string `json:"url" minLength:"1"`
	}
}