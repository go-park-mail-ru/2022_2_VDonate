package usersErrors

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type errJSON struct {
	Message string `json:"message"`
}

func responceError(err error) errJSON {
	return errJSON{
		Message: err.Error(),
	}
}

var (
	ErrConvertID     = errors.New("unable to convert id")
	ErrUserNotFound  = errors.New("user not found")
	ErrJSONMarshal   = errors.New("failed to marshal json")
	ErrJSONUnmarshal = errors.New("failed to unmarshal json")
	ErrResponse      = errors.New("failed to response")
	ErrUpdate        = errors.New("failed to update")
	ErrBadRequest    = errors.New("bad request")
)

func Wrap(c echo.Context, errCode, errLog error) error {
	c.Logger().Error(errLog)
	switch errCode {
	case ErrUserNotFound:
		return c.JSON(http.StatusNotFound, responceError(errCode))
	case ErrBadRequest:
		return c.JSON(http.StatusBadRequest, responceError(errCode))
	case ErrConvertID, ErrJSONMarshal, ErrJSONUnmarshal, ErrResponse, ErrUpdate:
		return c.JSON(http.StatusInternalServerError, responceError(errCode))
	default:
		return c.JSON(http.StatusInternalServerError, responceError(errCode))
	}
}
