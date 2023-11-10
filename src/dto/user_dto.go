package dto

import (
	"typograph_back/src/model"
)

type UserStoreRequest struct {
	Email    string `json:"email" validate:"required,email,max=125"`
	Name     string `json:"name" validate:"required,min=2,max=125"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserUpdatePasswordRequest struct {
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
}

func NewUserResponse(user *model.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Unix(),
	}
}
