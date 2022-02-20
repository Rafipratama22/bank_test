package dto

type LoginDto struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginDtoResponse struct {
	Token string `json:"token"`
}