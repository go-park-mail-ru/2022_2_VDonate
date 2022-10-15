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
	ErrCreate        = errors.New("failed to create post")
	ErrNoPosts       = errors.New("no posts with such user ID")
	ErrNoSession     = errors.New("no session provided")
)

func Wrap(c echo.Context, errHTTP, errLog error) error {
	c.Logger().Error(errLog)
	switch errHTTP {
	case ErrNoSession:
		return echo.NewHTTPError(http.StatusUnauthorized, errHTTP)
	case ErrBadRequest, ErrNoPosts:
		return echo.NewHTTPError(http.StatusBadRequest, errHTTP)
	case ErrJSONMarshal, ErrJSONUnmarshal, ErrResponse, ErrInternal, ErrCreate:
		return echo.NewHTTPError(http.StatusInternalServerError, errHTTP)
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, errHTTP)
	}
}
