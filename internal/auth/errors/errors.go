package authErrors

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	ErrUserOrEmailAlreadyExist = errors.New("username of email is already exist")
	ErrCreateUser              = errors.New("failed to create user")
	ErrUserNotFound            = errors.New("user not found")
	ErrInvalidLoginOrPassword  = errors.New("invalid login or password")
	ErrJSONMarshal             = errors.New("failed to marshal json")
	ErrJSONUnmarshal           = errors.New("failed to unmarshal json")
	ErrCopy                    = errors.New("failed to copy")
	ErrResponse                = errors.New("failed to response")
	ErrNoSession               = errors.New("no existing session")
	ErrBadRequest              = errors.New("bad request")
	ErrBadSession              = errors.New("bad session")
	ErrDeleteSession           = errors.New("failed to delete session")
)

func Wrap(c echo.Context, errHTTP, errLog error) error {
	c.Logger().Error(errLog)
	switch errHTTP {
	case ErrNoSession:
		return c.JSON(http.StatusUnauthorized, errHTTP.Error())
	case ErrUserNotFound:
		return c.JSON(http.StatusNotFound, errHTTP.Error())
	case ErrInvalidLoginOrPassword, ErrBadRequest:
		return c.JSON(http.StatusBadRequest, errHTTP.Error())
	case ErrUserOrEmailAlreadyExist:
		return c.JSON(http.StatusConflict, errHTTP.Error())
	case ErrJSONMarshal, ErrResponse, ErrJSONUnmarshal, ErrCreateUser, ErrCopy, ErrBadSession:
		return c.JSON(http.StatusInternalServerError, errHTTP.Error())
	default:
		return c.JSON(http.StatusInternalServerError, errHTTP.Error())
	}
}
