package authMiddlewares

import (
	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery"
	authErrors "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/errors"
	auth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/usecase"
	users "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/usecase"
	"github.com/labstack/echo/v4"
	"strconv"
)

type Middlewares struct {
	authUseCase  auth.UseCase
	usersUseCase users.UseCase
}

func New(authUseCase auth.UseCase, usersUseCase users.UseCase) *Middlewares {
	return &Middlewares{authUseCase: authUseCase, usersUseCase: usersUseCase}
}

func (m *Middlewares) LoginRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		cookie, err := httpAuth.GetCookie(c)
		if err != nil {
			return authErrors.Wrap(c, authErrors.ErrNoSession, err)
		}
		isAuth, err := m.authUseCase.Auth(cookie.Value)
		if !isAuth {
			return authErrors.Wrap(c, authErrors.ErrAuth, err)
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
		id, err := strconv.ParseUint(c.QueryParam("user_id"), 10, 64)
		if err != nil {
			return authErrors.Wrap(c, authErrors.ErrBadRequest, err)
		}
		if !m.authUseCase.IsSameSession(cookie.Value, id) {
			return authErrors.Wrap(c, authErrors.ErrForbidden, authErrors.ErrForbidden)
		}

		return next(c)
	}
}
