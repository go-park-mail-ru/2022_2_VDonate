package httpAuth

import (
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	httpUsers "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/delivery"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"github.com/labstack/echo/v4"
)

const (
	cookieName = "session_id"
)

var deleteExpire = map[string]int{
	"year":  0,
	"month": -1,
	"day":   0,
}

func GetCookie(c echo.Context) (*http.Cookie, error) {
	return c.Cookie(cookieName)
}

func MakeHTTPCookie(c *http.Cookie) *http.Cookie {
	return &http.Cookie{
		Name:     cookieName,
		Value:    c.Value,
		Expires:  c.Expires,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}
}

func makeHTTPCookieFromValue(value string) *http.Cookie {
	return &http.Cookie{
		Name:     cookieName,
		Value:    value,
		Expires:  time.Now().AddDate(0, 1, 0),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}
}

type Handler struct {
	authUseCase  domain.AuthUseCase
	usersUseCase domain.UsersUseCase
}

func NewHandler(authUseCase domain.AuthUseCase, usersUseCase domain.UsersUseCase) *Handler {
	return &Handler{
		authUseCase:  authUseCase,
		usersUseCase: usersUseCase,
	}
}

func (h Handler) Auth(c echo.Context) error {
	cookie, err := GetCookie(c)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNoSession, err)
	}

	isAuth, err := h.authUseCase.Auth(cookie.Value)
	if !isAuth {
		return utils.WrapEchoError(domain.ErrAuth, err)
	}

	user, err := h.usersUseCase.GetBySessionID(cookie.Value)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNotFound, err)
	}

	return httpUsers.UserResponse(c, user)
}

func (h Handler) Login(c echo.Context) error {
	var data models.AuthUser
	if err := c.Bind(&data); err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	sessionID, err := h.authUseCase.Login(data.Username, data.Password)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNoSession, err)
	}

	c.SetCookie(makeHTTPCookieFromValue(sessionID))

	user, err := h.usersUseCase.GetBySessionID(sessionID)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNotFound, err)
	}

	return httpUsers.UserResponse(c, user)
}

func (h Handler) Logout(c echo.Context) error {
	session, err := GetCookie(c)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNoSession, err)
	}

	isLogout, err := h.authUseCase.Logout(session.Value)
	if !isLogout {
		return utils.WrapEchoError(domain.ErrBadSession, err)
	}

	session.Expires = time.Now().AddDate(
		deleteExpire["year"],
		deleteExpire["month"],
		deleteExpire["day"],
	)
	c.SetCookie(MakeHTTPCookie(session))

	return c.JSON(http.StatusOK, struct{}{})
}

func (h Handler) SignUp(c echo.Context) error {
	var newUser models.User

	if err := c.Bind(&newUser); err != nil {
		return utils.WrapEchoError(domain.ErrBadRequest, err)
	}

	sessionID, err := h.authUseCase.SignUp(&newUser)
	if err != nil {
		return utils.WrapEchoError(domain.ErrNoSession, err)
	}

	c.SetCookie(makeHTTPCookieFromValue(sessionID))

	return httpUsers.UserResponse(c, &newUser)
}
