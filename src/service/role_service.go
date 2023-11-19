package service

import (
	"typograph_back/src/dto"
	"typograph_back/src/model"
	"typograph_back/src/repository"
	"typograph_back/src/repository/repository_interface"

	"github.com/gosimple/slug"
)

type RoleService struct {
	repository repository_interface.RoleRepositoryInterface
}

func NewRoleService() *RoleService {
	return &RoleService{repository: repository.NewRoleRepository()}
}

func (rs *RoleService) GetAll() ([]*model.Role, error) {
	return rs.repository.GetAll()
}

func (rs *RoleService) GetById(id uint) (*model.Role, error) {
	return rs.repository.GetById(id)
}

func (rs *RoleService) GetBySlug(slug string) (*model.Role, error) {
	return rs.repository.GetBySlug(slug)
}

func (rs *RoleService) Create(request *dto.RoleStoreRequest) (*model.Role, error) {
	role := model.Role{Name: request.Name, Slug: slug.Make(request.Name)}

	return rs.repository.Save(role)
}

func (rs *RoleService) Update(id uint, request *dto.RoleUpdateRequest) (*model.Role, error) {
	role, err := rs.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	role.Name = request.Name
	role.Slug = slug.Make(request.Name)

	return rs.repository.Save(*role)
}

func (rs *RoleService) Delete(id uint) error {
	return rs.repository.Delete(id)
}

func (rs *RoleService) AddPermissions(id uint, request *dto.RoleAddPermissionsRequest) error {
	return rs.repository.AddPermissions(id, request.PermissionsID)
}

func (rs *RoleService) RemovePermissions(id uint, request *dto.RoleRemovePermissionsRequest) error {
	return rs.repository.RemovePermissions(id, request.PermissionsID)
}
