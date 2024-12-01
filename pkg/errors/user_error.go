package errors

import "errors"

var (
	ErrUsernameExist      = errors.New("username is already exist")
	ErrEmailExist         = errors.New("email is already exist")
	ErrUsernameRequired   = errors.New("username is required")
	ErrEmailRequired      = errors.New("email is required")
	ErrPasswordIsRequired = errors.New("password is required")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidPassword    = errors.New("invalid password")
)
