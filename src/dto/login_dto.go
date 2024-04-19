package dto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=125"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Email          string `json:"email" validate:"required,email,max=125"`
	Name           string `json:"name" validate:"required,max=125"`
	Password       string `json:"password" validate:"required,min=8"`
	PasswordRepeat string `json:"password_repeat" validate:"required,min=8"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

type Letter struct {
	Char  string `json:"char"`
	Color string `json:"color"`
}

func NewLoginResponse(token string) *LoginResponse {
	return &LoginResponse{Token: token}
}
func NewRegisterResponse(token string) *RegisterResponse {
	return &RegisterResponse{Token: token}
}
