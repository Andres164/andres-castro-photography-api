package schemas

type UpdatePhotoInput struct {
	ID uint `path:"id"`

	Body struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		Url *string `json:"url"`
	} `json:"body"`
}