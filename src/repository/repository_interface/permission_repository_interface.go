package repository_interface

import "typograph_back/src/model"

type PermissionRepositoryInterface interface {
	GetAllByRole(roleID uint) ([]*model.Permission, error)
	GetById(id uint) (*model.Permission, error)
	Save(permission model.Permission) (*model.Permission, error)
	Delete(id uint) error
}
