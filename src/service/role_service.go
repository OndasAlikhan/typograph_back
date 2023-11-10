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

func (this *RoleService) GetAll() ([]*model.Role, error) {
	return this.repository.GetAll()
}

func (this *RoleService) GetById(id uint) (*model.Role, error) {
	return this.repository.GetById(id)
}

func (this *RoleService) GetBySlug(slug string) (*model.Role, error) {
	return this.repository.GetBySlug(slug)
}

func (this *RoleService) Create(request *dto.RoleStoreRequest) (*model.Role, error) {
	role := model.Role{Name: request.Name, Slug: slug.Make(request.Name)}

	return this.repository.Save(role)
}

func (this *RoleService) Update(id uint, request *dto.RoleUpdateRequest) (*model.Role, error) {
	role, err := this.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	role.Name = request.Name
	role.Slug = slug.Make(request.Name)

	return this.repository.Save(*role)
}

func (this *RoleService) Delete(id uint) error {
	return this.repository.Delete(id)
}

func (this *RoleService) AddPermissions(id uint, request *dto.RoleAddPermissionsRequest) error {
	return this.repository.AddPermissions(id, request.PermissionsID)
}

func (this *RoleService) RemovePermissions(id uint, request *dto.RoleRemovePermissionsRequest) error {
	return this.repository.RemovePermissions(id, request.PermissionsID)
}
