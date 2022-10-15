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
	ErrInternal                = errors.New("server error")
	ErrForbidden               = errors.New("you are not supposed to be here")
	ErrAuth                    = errors.New("failed to authenticate")
	ErrNoContent               = errors.New("no content was found")
)

func Wrap(c echo.Context, errHTTP, errLog error) error {
	c.Logger().Error(errLog)
	switch errHTTP {
	case ErrNoContent:
		return echo.NewHTTPError(http.StatusNoContent, errHTTP)
	case ErrNoSession, ErrAuth:
		return echo.NewHTTPError(http.StatusUnauthorized, errHTTP)
	case ErrUserNotFound:
		return echo.NewHTTPError(http.StatusNotFound, errHTTP)
	case ErrForbidden:
		return echo.NewHTTPError(http.StatusForbidden, errHTTP)
	case ErrInvalidLoginOrPassword, ErrBadRequest:
		return echo.NewHTTPError(http.StatusBadRequest, errHTTP)
	case ErrUserOrEmailAlreadyExist:
		return echo.NewHTTPError(http.StatusConflict, errHTTP)
	case ErrJSONMarshal, ErrResponse, ErrJSONUnmarshal, ErrCreateUser, ErrCopy, ErrBadSession, ErrInternal, ErrDeleteSession:
		return echo.NewHTTPError(http.StatusInternalServerError, errHTTP)
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, errHTTP)
	}
}
