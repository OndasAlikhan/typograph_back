package service_interface

import (
	"typograph_back/src/dto"
	"typograph_back/src/model"
)

type UserServiceInterface interface {
	GetAll() ([]*model.User, error)
	GetById(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(request *dto.UserStoreRequest) (*model.User, error)
	UpdatePassword(id uint, request *dto.UserUpdatePasswordRequest) (*model.User, error)
	Delete(id uint) error
	HasPermission(id uint, permissionSlug string) bool
}
