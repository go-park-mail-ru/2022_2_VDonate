package domain

import "errors"

var (
	ErrBadRequest              = errors.New("bad request")
	ErrUserOrEmailAlreadyExist = errors.New("username of email is already exist")
	ErrInvalidLoginOrPassword  = errors.New("invalid login or password")
	ErrForbidden               = errors.New("you are not supposed to be here")

	ErrCreate = errors.New("failed to create item")
	ErrUpdate = errors.New("failed to update item")
	ErrDelete = errors.New("failed to delete item")

	ErrResponse  = errors.New("failed to response")
	ErrNotFound  = errors.New("failed to find item")
	ErrNoContent = errors.New("no content was found")

	ErrAuth       = errors.New("failed to authenticate")
	ErrNoSession  = errors.New("no existing session")
	ErrBadSession = errors.New("bad session")

	ErrInternal      = errors.New("server error")
	ErrJSONMarshal   = errors.New("failed to marshal json")
	ErrJSONUnmarshal = errors.New("failed to unmarshal json")
	ErrCopy          = errors.New("failed to copy item")
)

var (
	ErrBucketNotExists         = errors.New("bucket doesn't exist")
	ErrFileOpen                = errors.New("failed to open file")
	ErrUsernameExist           = errors.New("username exists")
	ErrEmailExist              = errors.New("email exists")
	ErrUsernameOrEmailNotExist = errors.New("username or email not exists")
	ErrPasswordsNotEqual       = errors.New("passwords not the same")
)
