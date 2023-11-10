package dto

import "typograph_back/src/model"

type PermissionStoreRequest struct {
	Name string `json:"name" validate:"required,min=3,max=125"`
}

type PermissionUpdateRequest struct {
	Name string `json:"name" validate:"required,min=3,max=125"`
}

type PermissionResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func NewPermissionResponse(permission *model.Permission) *PermissionResponse {
	return &PermissionResponse{
		ID:   permission.ID,
		Name: permission.Name,
		Slug: permission.Slug,
	}
}
