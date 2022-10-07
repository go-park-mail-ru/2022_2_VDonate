package postsErrors

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	ErrJSONMarshal   = errors.New("failed to marshal json")
	ErrJSONUnmarshal = errors.New("failed to unmarshal json")
	ErrResponse      = errors.New("failed to response")
	ErrBadRequest    = errors.New("bar request")
	ErrInternal      = errors.New("server error")
)

func Wrap(c echo.Context, errCode, errLog error) error {
	c.Logger().Error(errLog)
	switch errCode {
	case ErrBadRequest:
		return c.JSON(http.StatusBadRequest, errCode.Error())
	case ErrJSONMarshal, ErrJSONUnmarshal, ErrResponse, ErrInternal:
		return c.JSON(http.StatusInternalServerError, errCode.Error())
	default:
		return c.JSON(http.StatusInternalServerError, errCode.Error())
	}
}
