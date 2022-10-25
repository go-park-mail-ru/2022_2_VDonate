package utils

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func WrapEchoError(errHTTP, errInternal error) error {
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
		domain.ErrBadRequest:
		return echo.NewHTTPError(http.StatusBadRequest, errHTTP.Error()).SetInternal(errInternal)
	case domain.ErrToken:
		return echo.NewHTTPError(http.StatusConflict, errHTTP.Error()).SetInternal(errInternal)
	case domain.ErrUserOrEmailAlreadyExist:
		return echo.NewHTTPError(http.StatusConflict, errHTTP.Error()).SetInternal(errInternal)
	case domain.ErrJSONMarshal,
		domain.ErrResponse,
		domain.ErrJSONUnmarshal,
		domain.ErrCreate,
		domain.ErrCopy,
		domain.ErrBadSession,
		domain.ErrInternal,
		domain.ErrDelete:
		return echo.NewHTTPError(http.StatusInternalServerError, errHTTP.Error())
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, errHTTP.Error())
	}
}

func CutCodeFromError(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()[strings.Index(err.Error(), " ")+1:]
}
