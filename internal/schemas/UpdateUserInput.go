package schemas


type UpdateUserInput struct {
	ID uint `path:"id"`
	
	Body struct {
		Email    *string `json:"email" required:"false"`
		Username *string `json:"username"`
		Password *string `json:"password"`
		Role     *string `json:"role"`
	}
}