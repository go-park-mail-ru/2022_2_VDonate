package authMiddlewares

import (
	authDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth"
	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery"
	authErrors "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/errors"
	postsDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/posts"
	usersDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/users"
	"github.com/labstack/echo/v4"
	"strconv"
)

type Middlewares struct {
	authUseCase  authDomain.UseCase
	postsUseCase postsDomain.UseCase
	usersUseCase usersDomain.UseCase
}

func New(authUseCase authDomain.UseCase, usersUseCase usersDomain.UseCase, postsUseCase postsDomain.UseCase) *Middlewares {
	return &Middlewares{
		authUseCase:  authUseCase,
		usersUseCase: usersUseCase,
		postsUseCase: postsUseCase,
	}
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

func (m *Middlewares) PostSameSessionByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := httpAuth.GetCookie(c)
		if err != nil {
			return authErrors.Wrap(c, authErrors.ErrNoSession, err)
		}
		postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return authErrors.Wrap(c, authErrors.ErrBadRequest, err)
		}
		post, err := m.postsUseCase.GetPostByID(postID)
		if err != nil {
			return authErrors.Wrap(c, authErrors.ErrNoContent, err)
		}
		if !m.authUseCase.IsSameSession(cookie.Value, post.UserID) {
			return authErrors.Wrap(c, authErrors.ErrForbidden, authErrors.ErrForbidden)
		}

		return next(c)
	}
}
