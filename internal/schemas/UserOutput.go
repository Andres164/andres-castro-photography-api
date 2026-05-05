package schemas

type UserResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type UserOutput struct {
	Body UserResponse
}