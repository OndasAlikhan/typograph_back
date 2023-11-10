package repository_interface

import "typograph_back/src/model"

type RoleRepositoryInterface interface {
	GetAll() ([]*model.Role, error)
	GetById(id uint) (*model.Role, error)
	GetBySlug(slug string) (*model.Role, error)
	Save(role model.Role) (*model.Role, error)
	Delete(id uint) error
	AddPermissions(id uint, permissionsID []uint) error
	RemovePermissions(id uint, permissionsID []uint) error
}
