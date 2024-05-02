package controller

import (
	"fmt"
	"net/http"
	"typograph_back/src/dto"
	"typograph_back/src/service/service_interface"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	*BaseController
	service service_interface.AuthServiceInterface
}

func NewAuthController(authService service_interface.AuthServiceInterface) *AuthController {
	return &AuthController{
		BaseController: NewBaseController(),
		service:        authService,
	}
}

// Login
// @title Login
// @description Login a user
// @accept json
// @produce json
// @tags auth
// @param loginRequest body dto.LoginRequest true "Login Request"
// @success	200 {object} dto.JSONResult{data=dto.LoginResponse}
// @router /login [post]
func (ac *AuthController) Login(c echo.Context) error {
	request := dto.LoginRequest{}
	if err := ac.handleRequest(&request, c); err != nil {
		return err
	}

	token, err := ac.service.Login(&request)
	if err != nil {
		return err
	}

	return ac.json(http.StatusOK, dto.NewLoginResponse(token), c)
}

// Register
// @title Register
// @description Register a user
// @accept json
// @produce json
// @tags auth
// @param registerRequest body dto.RegisterRequest true "Register Request"
// @success	200 {object} dto.JSONResult{data=dto.RegisterResponse}
// @router /register [post]
func (ac *AuthController) Register(c echo.Context) error {
	request := dto.RegisterRequest{}
	if err := ac.handleRequest(&request, c); err != nil {
		return err
	}
	fmt.Printf("request %s", request)

	token, err := ac.service.Register(&request)
	if err != nil {
		return err
	}

	return ac.json(http.StatusOK, dto.NewRegisterResponse(token), c)
}

// Me
// @title Me
// @description	Get user info
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags auth
// @success 200	{object} dto.JSONResult{data=dto.UserResponse}
// @router /me [get]
func (ac *AuthController) Me(c echo.Context) error {
	user, err := ac.authUser(c)
	if err != nil {
		return err
	}

	return ac.json(http.StatusOK, dto.NewUserResponse(user), c)
}
