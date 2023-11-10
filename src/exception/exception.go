package exception

import "errors"

var (
	ErrUnauthorized  = errors.New("unauthorized")
	ErrInvalidLogin  = errors.New("invalid login")
	ErrNotPermission = errors.New("permission denied")
	ErrInvalidParam  = errors.New("invalid param")
)
