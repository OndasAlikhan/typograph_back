package controller

import (
	"net/http"
	"typograph_back/src/dto"
	"typograph_back/src/exception"
	"typograph_back/src/service"
	"typograph_back/src/service/service_interface"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	*BaseController
	service service_interface.UserServiceInterface
}

func NewUserController() *UserController {
	return &UserController{
		BaseController: NewBaseController(),
		service:        service.NewUserService(),
	}
}

// Index
// @title Index
// @description List of users
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags user
// @success 200 {object} dto.JSONResult{data=[]dto.UserResponse}
// @router /users [get]
func (this *UserController) Index(c echo.Context) error {
	users, err := this.service.GetAll()
	if err != nil {
		return err
	}

	response := make([]*dto.UserResponse, len(users))
	for ind, user := range users {
		response[ind] = dto.NewUserResponse(user)
	}

	return this.json(http.StatusOK, response, c)
}

// Store
// @title Store
// @description Create a user
// @accept json
// @produce json
// @tags user
// @param userStoreRequest body dto.UserStoreRequest true "User store request"
// @success 200 {object} dto.JSONResult{data=dto.UserResponse}
// @router /users [post]
func (this *UserController) Store(c echo.Context) error {
	request := dto.UserStoreRequest{}
	if err := this.handleRequest(&request, c); err != nil {
		return err
	}

	user, err := this.service.Create(&request)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewUserResponse(user), c)
}

// Show
// @title Show
// @description Get a user
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags user
// @param id path int true "User ID"
// @success 200 {object} dto.JSONResult{data=dto.UserResponse}
// @router /users/{id} [get]
func (this *UserController) Show(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"))
	if err != nil {
		return err
	}

	authUser, err := this.authUser(c)
	if err != nil {
		return err
	}

	if authUser.ID != id {
		return exception.ErrNotPermission
	}

	user, err := this.service.GetById(id)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewUserResponse(user), c)
}

// UpdatePassword
// @title UpdatePassword
// @description Update password a user
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags user
// @param userUpdatePasswordRequest body dto.UserUpdatePasswordRequest true "User update password request"
// @param id path int true "User ID"
// @success 200 {object} dto.JSONResult{data=dto.UserResponse}
// @router /users/{id} [patch]
func (this *UserController) UpdatePassword(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"))
	if err != nil {
		return err
	}

	authUser, err := this.authUser(c)
	if err != nil {
		return err
	}

	if authUser.ID != id {
		return exception.ErrNotPermission
	}

	request := dto.UserUpdatePasswordRequest{}
	if err = this.handleRequest(&request, c); err != nil {
		return err
	}

	user, err := this.service.UpdatePassword(id, &request)
	if err != nil {
		return err
	}

	return this.json(http.StatusOK, dto.NewUserResponse(user), c)
}

// Delete
// @title Delete
// @description Delete a user
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags user
// @param id path int true "User ID"
// @success 200 {object} dto.JSONResult
// @router /users/{id} [delete]
func (this *UserController) Delete(c echo.Context) error {
	id, err := this.parseToUint(c.Param("id"))
	if err != nil {
		return nil
	}

	authUser, err := this.authUser(c)
	if err != nil {
		return err
	}

	if authUser.ID != id {
		return exception.ErrNotPermission
	}

	if err = this.service.Delete(id); err != nil {
		return err
	}

	return this.json(http.StatusOK, nil, c)
}
