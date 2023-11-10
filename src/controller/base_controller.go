package controller

import (
	"net/http"
	"strconv"
	"typograph_back/src/dto"
	"typograph_back/src/exception"
	"typograph_back/src/model"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type BaseController struct {
	validate *validator.Validate
}

func NewBaseController() *BaseController {
	return &BaseController{validate: validator.New()}
}

func (this *BaseController) authUser(c echo.Context) (*model.User, error) {
	user, ok := c.Get("auth_user").(*model.User)
	if !ok {
		return nil, exception.ErrUnauthorized
	}

	return user, nil
}

func (this *BaseController) handleRequest(request interface{}, c echo.Context) error {
	if err := c.Bind(&request); err != nil {
		return err
	}

	if err := this.validate.Struct(request); err != nil {
		return err
	}

	return nil
}

func (this *BaseController) parseToUint(value string) (uint, error) {
	id, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return uint(id), exception.ErrInvalidParam
	}

	return uint(id), nil
}

func (this *BaseController) json(code int, response interface{}, c echo.Context) error {
	return c.JSON(code, dto.NewJSONResult(http.StatusText(code), response))
}
