package repository

import (
	"typograph_back/src/application"
	"typograph_back/src/model"

	"gorm.io/gorm"
)

type RoleRepository struct {
	connection *gorm.DB
}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{connection: application.GlobalDB}
}

func (this *RoleRepository) GetAll() ([]*model.Role, error) {
	var roles []*model.Role
	err := this.connection.Find(&roles).Error

	return roles, err
}

func (this *RoleRepository) GetById(id uint) (*model.Role, error) {
	var role *model.Role
	err := this.connection.First(&role, id).Error

	return role, err
}

func (this *RoleRepository) GetBySlug(slug string) (*model.Role, error) {
	var role *model.Role
	err := this.connection.Where("slug LIKE ?", slug+"%").First(&role).Error

	return role, err
}

func (this *RoleRepository) Save(role model.Role) (*model.Role, error) {
	result := this.connection.Save(&role)

	return &role, result.Error
}

func (this *RoleRepository) Delete(id uint) error {
	var role model.Role
	if err := this.connection.First(&role, id).Error; err != nil {
		return err
	}

	return this.connection.Delete(&role).Error
}

func (this *RoleRepository) AddPermissions(id uint, permissionsID []uint) error {
	var role model.Role
	if err := this.connection.Preload("Permissions").First(&role, id).Error; err != nil {
		return err
	}

	var permissionsToAdd []*model.Permission
	if err := this.connection.Find(&permissionsToAdd, "id IN ?", permissionsID).Error; err != nil {
		return err
	}

	role.Permissions = append(role.Permissions, permissionsToAdd...)

	if err := this.connection.Save(&role).Error; err != nil {
		return err
	}

	return nil
}

func (repo *RoleRepository) RemovePermissions(id uint, permissionsID []uint) error {
	var role model.Role
	if err := repo.connection.Preload("Permissions").First(&role, id).Error; err != nil {
		return err
	}

	var permissionsToRemove []*model.Permission
	if err := repo.connection.Find(&permissionsToRemove, "id IN ?", permissionsID).Error; err != nil {
		return err
	}

	association := repo.connection.Model(&role).Association("Permissions")
	if err := association.Delete(permissionsToRemove); err != nil {
		return err
	}

	return nil
}
