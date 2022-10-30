package authMiddlewares

import (
	"strconv"

	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"github.com/labstack/echo/v4"
)

type Middlewares struct {
	authUseCase  domain.AuthUseCase
	postsUseCase domain.PostsUseCase
	usersUseCase domain.UsersUseCase
}

func New(a domain.AuthUseCase, u domain.UsersUseCase, p domain.PostsUseCase) *Middlewares {
	return &Middlewares{
		authUseCase:  a,
		usersUseCase: u,
		postsUseCase: p,
	}
}

func (m *Middlewares) LoginRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		cookie, err := httpAuth.GetCookie(c)
		if err != nil {
			return utils.WrapEchoError(domain.ErrNoSession, err)
		}
		isAuth, err := m.authUseCase.Auth(cookie.Value)
		if !isAuth {
			return utils.WrapEchoError(domain.ErrAuth, err)
		}

		return next(c)
	}
}

func (m *Middlewares) PostSameSessionByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := httpAuth.GetCookie(c)
		if err != nil {
			return utils.WrapEchoError(domain.ErrNoSession, err)
		}
		postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return utils.WrapEchoError(domain.ErrBadRequest, err)
		}
		post, err := m.postsUseCase.GetPostByID(postID)
		if err != nil {
			return utils.WrapEchoError(domain.ErrNoContent, err)
		}
		if !m.authUseCase.IsSameSession(cookie.Value, post.UserID) {
			return utils.WrapEchoError(domain.ErrForbidden, domain.ErrForbidden)
		}

		return next(c)
	}
}
