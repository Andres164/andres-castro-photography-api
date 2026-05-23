package schemas

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginOutput struct {
	Body LoginResponse
}