package application

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"typograph_back/src/dto"
	"typograph_back/src/exception"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

var (
	GlobalApp *echo.Echo
	onceApp   sync.Once
)

func InitializeApp(lvl log.Lvl) {
	onceApp.Do(func() {
		GlobalApp = echo.New()

		GlobalApp.HTTPErrorHandler = errorHandler

		GlobalApp.Logger.SetLevel(lvl)

		GlobalApp.Pre(middleware.RemoveTrailingSlash())

		GlobalApp.Use(middleware.Logger())

		GlobalApp.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
			AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		}))
	})
}

func errorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		messages := make(map[string]string)
		for _, field := range validationErrors {
			fieldName := strings.ToLower(field.Field())
			messages[fieldName] = fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", fieldName, field.Tag())
		}

		c.JSON(http.StatusBadRequest, dto.NewJSONResult(http.StatusText(http.StatusBadRequest), messages))
	} else {
		var he *echo.HTTPError
		if errors.As(err, &he) {
			code = he.Code
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusNotFound
		} else if errors.Is(err, exception.ErrUnauthorized) {
			code = http.StatusUnauthorized
		} else if errors.Is(err, exception.ErrInvalidLogin) ||
			errors.Is(err, exception.ErrInvalidParam) {
			code = http.StatusBadRequest
		} else if errors.Is(err, exception.ErrNotPermission) {
			code = http.StatusForbidden
		}

		c.JSON(code, dto.NewJSONResult(err.Error(), nil))
	}
}
