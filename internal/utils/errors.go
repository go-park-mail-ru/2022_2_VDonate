package utils

import (
	"net/http"
	"strings"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/labstack/echo/v4"
)

func WrapEchoError(errHTTP, errInternal error) error {
	switch errInternal {
	case domain.ErrUsernameOrEmailNotExist:
		return echo.NewHTTPError(http.StatusNotFound, errInternal.Error()).SetInternal(errInternal)
	case domain.ErrPasswordsNotEqual:
		return echo.NewHTTPError(http.StatusBadRequest, errInternal.Error()).SetInternal(errInternal)
	case domain.ErrEmailExist, domain.ErrUsernameExist:
		return echo.NewHTTPError(http.StatusConflict, errInternal.Error()).SetInternal(errInternal)
	}

	switch errHTTP {
	case domain.ErrNoContent:
		return echo.NewHTTPError(http.StatusNoContent, errHTTP.Error()).SetInternal(errInternal)
	case domain.ErrNoSession, domain.ErrAuth:
		return echo.NewHTTPError(http.StatusUnauthorized, errHTTP.Error()).SetInternal(errInternal)
	case domain.ErrNotFound:
		return echo.NewHTTPError(http.StatusNotFound, errHTTP.Error()).SetInternal(errInternal)
	case domain.ErrForbidden:
		return echo.NewHTTPError(http.StatusForbidden, errHTTP.Error()).SetInternal(errInternal)
	case domain.ErrInvalidLoginOrPassword,
		domain.ErrBadRequest, domain.ErrBadSession:
		return echo.NewHTTPError(http.StatusBadRequest, errHTTP.Error()).SetInternal(errInternal)
	case domain.ErrUserOrEmailAlreadyExist:
		return echo.NewHTTPError(http.StatusConflict, errHTTP.Error()).SetInternal(errInternal)
	case domain.ErrJSONMarshal,
		domain.ErrResponse,
		domain.ErrJSONUnmarshal,
		domain.ErrCreate,
		domain.ErrCopy,
		domain.ErrInternal,
		domain.ErrDelete:
		return echo.NewHTTPError(http.StatusInternalServerError, errHTTP.Error()).SetInternal(errInternal)
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, errHTTP.Error()).SetInternal(errInternal)
	}
}

func CutCodeFromError(err error) string {
	if err == nil {
		return ""
	}

	return err.Error()[strings.Index(err.Error(), " ")+1:]
}
