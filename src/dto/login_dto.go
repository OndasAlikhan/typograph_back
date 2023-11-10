package dto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=125"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func NewLoginResponse(token string) *LoginResponse {
	return &LoginResponse{Token: token}
}
