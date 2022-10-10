package authMiddlewares

import (
	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery"
	authErrors "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/errors"
	auth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/usecase"
	"github.com/labstack/echo/v4"
	"strconv"
)

type Middlewares struct {
	authUseCase auth.UseCase
}

func New(authUseCase auth.UseCase) *Middlewares {
	return &Middlewares{authUseCase: authUseCase}
}

func (m *Middlewares) LoginRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		cookie, err := httpAuth.GetCookie(c)
		if err != nil {
			return authErrors.Wrap(c, authErrors.ErrNoSession, err)
		}
		isAuth, err := m.authUseCase.Auth(cookie.Value)
		if !isAuth {
			return authErrors.Wrap(c, authErrors.ErrNoSession, err)
		}

		return next(c)
	}
}

func (m *Middlewares) SameSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := httpAuth.GetCookie(c)
		if err != nil {
			return authErrors.Wrap(c, authErrors.ErrNoSession, err)
		}
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return authErrors.Wrap(c, authErrors.ErrInternal, err)
		}
		if !m.authUseCase.IsSameSession(cookie.Value, id) {
			return authErrors.Wrap(c, authErrors.ErrForbidden, err)
		}

		return next(c)
	}
}
