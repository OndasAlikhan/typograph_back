package dto

import "typograph_back/src/model"

type RoleStoreRequest struct {
	Name string `json:"name" validate:"required,min=3,max=125"`
}

type RoleUpdateRequest struct {
	Name string `json:"name" validate:"required,min=3,max=125"`
}

type RoleAddPermissionsRequest struct {
	PermissionsID []uint `json:"permissions_id" validate:"required,dive,gt=0"`
}

type RoleRemovePermissionsRequest struct {
	PermissionsID []uint `json:"permissions_id" validate:"required,dive,gt=0"`
}

type RoleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func NewRoleResponse(role *model.Role) *RoleResponse {
	return &RoleResponse{
		ID:   role.ID,
		Name: role.Name,
		Slug: role.Slug,
	}
}
