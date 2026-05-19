package schemas


type UpdateUserInput struct {
	ID uint `path:"id"`
	
	Body struct {
		Email    *string `json:"email,omitempty"`
		Username *string `json:"username,omitempty"`
		Password *string `json:"password,omitempty"`
		Role     *string `json:"role,omitempty"`
	}
}