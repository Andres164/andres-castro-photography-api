package schemas

type LogInInput struct {
	Body struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}
}