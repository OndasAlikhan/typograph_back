package service

import (
	"errors"
	"fmt"
	"typograph_back/src/dto"
	"typograph_back/src/exception"
	"typograph_back/src/service/service_interface"

	"gorm.io/gorm"
)

type AuthService struct {
	userService service_interface.UserServiceInterface
	jwtService  service_interface.JWTServiceInterface
}

func NewAuthService() *AuthService {
	return &AuthService{
		userService: NewUserService(),
		jwtService:  NewJWTService(),
	}
}

func (as *AuthService) Login(request *dto.LoginRequest) (string, error) {
	user, err := as.userService.GetByEmail(request.Email)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", exception.ErrInvalidLogin
	} else if err != nil {
		return "", err
	}

	if !as.jwtService.IsEqual(user.Password, request.Password) {
		return "", exception.ErrInvalidLogin
	}

	token, err := as.jwtService.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (as *AuthService) Register(request *dto.RegisterRequest) (string, error) {
	const op = "auth_service.Register"

	if request.Password != request.PasswordRepeat {
		return "", fmt.Errorf("%s: passwords don't match", op)
	}

	user, err := as.userService.Create(&dto.UserStoreRequest{
		Email:    request.Email,
		Name:     request.Name,
		Password: request.Password,
	})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	token, err := as.jwtService.GenerateToken(user.ID)

	return token, nil
}
