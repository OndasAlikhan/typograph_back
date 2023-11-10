package middleware

import (
	"typograph_back/src/exception"
	"typograph_back/src/service"
	"typograph_back/src/service/service_interface"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JwtMiddleware struct {
	userService service_interface.UserServiceInterface
	tokenPrefix string
	secret      []byte
}

func NewJwtMiddleware(secret []byte, tokenPrefix string) *JwtMiddleware {
	return &JwtMiddleware{
		userService: service.NewUserService(),
		tokenPrefix: tokenPrefix,
		secret:      secret,
	}
}

func (this *JwtMiddleware) Run() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeaderValue, err := this.checkHeader(c)
			if err != nil {
				return err
			}

			token, err := this.parseJwt(authHeaderValue)
			if err != nil {
				return err
			}

			authID, err := this.processClaims(token)
			if err != nil {
				return err
			}

			user, err := this.userService.GetById(authID)
			if err != nil {
				return err
			}

			c.Set("auth_user", user)

			return next(c)
		}
	}
}

func (this *JwtMiddleware) checkHeader(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if !(len(authHeader) > len(this.tokenPrefix) && authHeader[:len(this.tokenPrefix)] == this.tokenPrefix) {
		return "", exception.ErrUnauthorized
	}

	return authHeader[len(this.tokenPrefix):], nil
}

func (this *JwtMiddleware) parseJwt(authHeaderValue string) (*jwt.Token, error) {
	token, err := jwt.Parse(authHeaderValue, this.parseToken)
	if err != nil || !token.Valid {
		return nil, exception.ErrUnauthorized
	}

	return token, nil
}

func (this *JwtMiddleware) parseToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, exception.ErrUnauthorized
	}

	return this.secret, nil
}

func (this *JwtMiddleware) processClaims(token *jwt.Token) (uint, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return 0, exception.ErrUnauthorized
	}

	authID, ok := claims["auth_id"].(float64)
	if !ok {
		return 0, exception.ErrUnauthorized
	}

	return uint(authID), nil
}
