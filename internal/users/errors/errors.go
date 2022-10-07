package usersErrors

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type errJSON struct {
	message string
}

func responceError(err error) errJSON {
	return errJSON{
		message: err.Error(),
	}
}

var (
	ErrConvertID     = errors.New("unable to convert id")
	ErrUserNotFound  = errors.New("user not found")
	ErrJSONMarshal   = errors.New("failed to marshal json")
	ErrJSONUnmarshal = errors.New("failed to unmarshal json")
	ErrResponse      = errors.New("failed to response")
)

func Wrap(c echo.Context, errCode, errLog error) error {
	c.Logger().Error(errLog)
	switch errCode {
	case ErrUserNotFound:
		return c.JSON(http.StatusNotFound, responceError(errCode))
	case ErrConvertID, ErrJSONMarshal, ErrJSONUnmarshal, ErrResponse:
		return c.JSON(http.StatusInternalServerError, responceError(errCode))
	default:
		return c.JSON(http.StatusInternalServerError, responceError(errCode))
	}
}
