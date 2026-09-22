package schemas

type UpdatePhotoInput struct {
	ID uint `path:"id"`

	Body struct {
		Title       *string `json:"title,omitempty"`
		Description *string `json:"description,omitempty"`
		Url *string `json:"url,omitempty"`
		ContentType *string `json:"content_type,omitempty"`
	} `json:"body"`
}