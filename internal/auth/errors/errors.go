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

type errJSON struct {
	Message string `json:"message"`
}

func responceError(err error) errJSON {
	return errJSON{
		Message: err.Error(),
	}
}

func Wrap(c echo.Context, errHTTP, errLog error) error {
	c.Logger().Error(errLog)
	switch errHTTP {
	case ErrNoSession:
		return c.JSON(http.StatusUnauthorized, responceError(errHTTP))
	case ErrUserNotFound:
		return c.JSON(http.StatusNotFound, responceError(errHTTP))
	case ErrInvalidLoginOrPassword, ErrBadRequest:
		return c.JSON(http.StatusBadRequest, responceError(errHTTP))
	case ErrUserOrEmailAlreadyExist:
		return c.JSON(http.StatusConflict, responceError(errHTTP))
	case ErrJSONMarshal, ErrResponse, ErrJSONUnmarshal, ErrCreateUser, ErrCopy, ErrBadSession:
		return c.JSON(http.StatusInternalServerError, responceError(errHTTP))
	default:
		return c.JSON(http.StatusInternalServerError, responceError(errHTTP))
	}
}
