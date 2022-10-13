package usersErrors

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	ErrConvertID     = errors.New("unable to convert id")
	ErrUserNotFound  = errors.New("user not found")
	ErrJSONMarshal   = errors.New("failed to marshal json")
	ErrJSONUnmarshal = errors.New("failed to unmarshal json")
	ErrResponse      = errors.New("failed to response")
	ErrUpdate        = errors.New("failed to update")
	ErrBadRequest    = errors.New("bad request")
)

func Wrap(c echo.Context, errHTTP, errLog error) error {
	c.Logger().Error(errLog)
	switch errHTTP {
	case ErrUserNotFound:
		return echo.NewHTTPError(http.StatusNotFound, errHTTP)
	case ErrBadRequest:
		return echo.NewHTTPError(http.StatusBadRequest, errHTTP)
	case ErrConvertID, ErrJSONMarshal, ErrJSONUnmarshal, ErrResponse, ErrUpdate:
		return echo.NewHTTPError(http.StatusInternalServerError, errHTTP)
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, errHTTP)
	}
}
