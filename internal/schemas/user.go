package schemas

type LogInInput struct {
	Body struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginOutput struct {
	Body LoginResponse
}

type UserResponse struct {
	ID    uint `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type UserOutput struct {
	Body UserResponse
}

type GetUsersOutput struct {
	Body []UserResponse
}

type UserIdInput struct {
	ID uint `path:"id"`
}

type CreateUserInput struct {
	Body struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
}

type UpdateUserInput struct {
	ID uint `path:"id"`
	
	Body struct {
		Email    *string `json:"email,omitempty"`
		Username *string `json:"username,omitempty"`
		Password *string `json:"password,omitempty"`
		Role     *string `json:"role,omitempty"`
	}
}