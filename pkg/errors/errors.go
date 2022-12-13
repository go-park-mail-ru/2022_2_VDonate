package errorHandling

import (
	"net/http"

	"google.golang.org/grpc/status"

	"github.com/ztrue/tracerr"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/labstack/echo/v4"
)

func WrapEcho(errHTTP, errInternal error) error {
	if grpcError, ok := status.FromError(tracerr.Unwrap(errInternal)); ok {
		switch grpcError.Message() {
		case domain.ErrUsernameOrEmailNotExist.Error():
			return echo.NewHTTPError(http.StatusNotFound, grpcError.Message()).SetInternal(errInternal)
		case domain.ErrPasswordsNotEqual.Error(), domain.ErrBadRequest.Error():
			return echo.NewHTTPError(http.StatusBadRequest, grpcError.Message()).SetInternal(errInternal)
		case domain.ErrEmailExist.Error(), domain.ErrUsernameExist.Error(), domain.ErrUnknownFormat.Error():
			return echo.NewHTTPError(http.StatusConflict, grpcError.Message()).SetInternal(errInternal)
		case domain.ErrNoLikes.Error():
			return echo.NewHTTPError(http.StatusNoContent, grpcError.Message()).SetInternal(errInternal)
		}
	} else {
		switch errInternal.Error() {
		case domain.ErrUsernameOrEmailNotExist.Error():
			return echo.NewHTTPError(http.StatusNotFound, errInternal.Error()).SetInternal(errInternal)
		case domain.ErrPasswordsNotEqual.Error(), domain.ErrBadRequest.Error():
			return echo.NewHTTPError(http.StatusBadRequest, errInternal.Error()).SetInternal(errInternal)
		case domain.ErrEmailExist.Error(), domain.ErrUsernameExist.Error(), domain.ErrUnknownFormat.Error():
			return echo.NewHTTPError(http.StatusConflict, errInternal.Error()).SetInternal(errInternal)
		case domain.ErrNoLikes.Error():
			return echo.NewHTTPError(http.StatusNoContent, errInternal.Error()).SetInternal(errInternal)
		}
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
	case domain.ErrUserOrEmailAlreadyExist, domain.ErrConflict:
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
