package service_interface

import "typograph_back/src/dto"

type AuthServiceInterface interface {
	Login(request *dto.LoginRequest) (string, error)
	Register(request *dto.RegisterRequest) (string, error)
}
