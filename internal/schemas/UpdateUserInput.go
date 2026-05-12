package schemas


type UpdateUserInput struct {
	ID uint `path:"id"`
	
	Body struct {
		Email    *string `json:"email"`
		Username *string `json:"username"`
		Password *string `json:"password"`
		Role     *string `json:"role"`
	}
}