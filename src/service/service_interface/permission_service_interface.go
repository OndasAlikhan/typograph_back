package service_interface

import (
	"typograph_back/src/dto"
	"typograph_back/src/model"
)

type PermissionServiceInterface interface {
	GetAllByRole(roleID uint) ([]*model.Permission, error)
	GetById(id uint) (*model.Permission, error)
	Create(request *dto.PermissionStoreRequest) (*model.Permission, error)
	Update(id uint, request *dto.PermissionUpdateRequest) (*model.Permission, error)
	Delete(id uint) error
}
