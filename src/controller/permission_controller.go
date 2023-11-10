package controller

import (
	"net/http"
	"typograph_back/src/dto"
	"typograph_back/src/service"
	"typograph_back/src/service/service_interface"

	"github.com/labstack/echo/v4"
)

type PermissionController struct {
	*BaseController
	service service_interface.PermissionServiceInterface
}

func NewPermissionController() *PermissionController {
	return &PermissionController{
		BaseController: NewBaseController(),
		service:        service.NewPermissionService(),
	}
}

// Index
// @title Index
// @description List of permissions
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags permission
// @success 200 {object} dto.JSONResult{data=[]dto.PermissionResponse}
// @router /permissions [get]
func (this *PermissionController) Index(c echo.Context) error {
	roleID, err := this.parseToUint(c.QueryParam("role_id"))
	if err != nil {
		return err
	}

	permissions, err := this.service.GetAllByRole(roleID)
	if err != nil {
		return err
	}

	response := make([]*dto.PermissionResponse, len(permissions))
	for ind, permission := range permissions {
		response[ind] = dto.NewPermissionResponse(permission)
	}

	return this.json(http.StatusOK, response, c)
}

// Store
// @title Store
// @description Create a permission
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags permission
// @param permissionStoreRequest body dto.PermissionStoreRequest true "Permission store request"
// @success 200 {object} dto.JSONResult{data=dto.PermissionResponse}
// @router /permissions [post]
func (this *PermissionController) Store(c echo.Context) error {
	request := dto.PermissionStoreRequest{}
	if err := this.handleRequest(&request, c); err != nil {
		return err
	}

	permission, err := this.service.Create(&request)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewPermissionResponse(permission), c)
}

// Show
// @title Show
// @description Get a permission
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags permission
// @param id path int true "Permission ID"
// @success 200 {object} dto.JSONResult{data=dto.PermissionResponse}
// @router /permissions/{id} [get]
func (this *PermissionController) Show(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"))
	if err != nil {
		return err
	}

	permission, err := this.service.GetById(id)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewPermissionResponse(permission), c)
}

// Update
// @title Update
// @description Update a permission
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags permission
// @param permissionUpdateRequest body dto.PermissionUpdateRequest true "Permission update request"
// @param id path int true "Permission ID"
// @success 200 {object} dto.JSONResult{data=dto.PermissionResponse}
// @router /permissions/{id} [patch]
func (this *PermissionController) Update(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"))
	if err != nil {
		return err
	}

	request := dto.PermissionUpdateRequest{}
	if err = this.handleRequest(&request, c); err != nil {
		return err
	}

	permission, err := this.service.Update(id, &request)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewPermissionResponse(permission), c)
}

// Delete
// @title Delete
// @description Delete a permission
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags permission
// @param id path int true "Permission ID"
// @success 200 {object} dto.JSONResult
// @router /permissions/{id} [delete]
func (this *PermissionController) Delete(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"))
	if err != nil {
		return err
	}

	if err = this.service.Delete(id); err != nil {
		return err
	}

	return this.json(http.StatusOK, nil, c)
}
