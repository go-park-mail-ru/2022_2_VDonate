package auth

import "errors"

var (
	ErrUserAlreadyExist       = errors.New("user already exist")
	ErrInvalidLoginOrPassword = errors.New("invalid login or password")
	ErrSetSession             = errors.New("error while set session")
	ErrDeleteSession          = errors.New("error while delete session")
)
