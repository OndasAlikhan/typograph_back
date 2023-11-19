package service

import (
	"typograph_back/src/dto"
	"typograph_back/src/model"
	"typograph_back/src/repository"
	"typograph_back/src/repository/repository_interface"

	"github.com/gosimple/slug"
)

type PermissionService struct {
	repository repository_interface.PermissionRepositoryInterface
}

func NewPermissionService() *PermissionService {
	return &PermissionService{repository: repository.NewPermissionRepository()}
}

func (ps *PermissionService) GetAllByRole(roleID uint) ([]*model.Permission, error) {
	return ps.repository.GetAllByRole(roleID)
}

func (ps *PermissionService) GetById(id uint) (*model.Permission, error) {
	return ps.repository.GetById(id)
}

func (ps *PermissionService) Create(request *dto.PermissionStoreRequest) (*model.Permission, error) {
	permission := model.Permission{Name: request.Name, Slug: slug.Make(request.Name)}

	return ps.repository.Save(permission)
}

func (ps *PermissionService) Update(id uint, request *dto.PermissionUpdateRequest) (*model.Permission, error) {
	permission, err := ps.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	permission.Name = request.Name
	permission.Slug = slug.Make(request.Name)

	return ps.repository.Save(*permission)
}

func (ps *PermissionService) Delete(id uint) error {
	return ps.repository.Delete(id)
}
