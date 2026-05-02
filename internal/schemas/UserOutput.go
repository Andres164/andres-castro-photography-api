package schemas

type UserOutput struct {
	Body struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Role     string `json:"role"`
	}
}