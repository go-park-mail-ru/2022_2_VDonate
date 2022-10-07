package authMiddlewares

import (
	authErrors "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/errors"
	sessionRepository "github.com/go-park-mail-ru/2022_2_VDonate/internal/session/repository"
	"github.com/labstack/echo/v4"
)

type Middlewares struct {
	sessionAPI sessionRepository.API
}

func New(sessionAPI sessionRepository.API) *Middlewares {
	return &Middlewares{sessionAPI: sessionAPI}
}

func (m *Middlewares) LoginRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		cookie, err := c.Cookie("session_id")
		if err != nil {
			return authErrors.Wrap(c, authErrors.ErrNoSession, err)
		}
		if _, err = m.sessionAPI.GetByValue(cookie.Value); err != nil {
			return authErrors.Wrap(c, authErrors.ErrNoSession, err)
		}
		return next(c)
	}
}
