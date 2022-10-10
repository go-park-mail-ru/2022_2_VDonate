package httpAuth

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/errors"
	auth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/users/delivery"
	users "github.com/go-park-mail-ru/2022_2_VDonate/internal/users/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
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
		SameSite: http.SameSiteLaxMode,
	}
}

type Handler struct {
	authUseCase  auth.UseCase
	usersUseCase users.UseCase
}

func NewHandler(authUseCase auth.UseCase, usersUseCase users.UseCase) *Handler {
	return &Handler{authUseCase: authUseCase, usersUseCase: usersUseCase}
}

func (h *Handler) Auth(c echo.Context) error {
	cookie, err := GetCookie(c)
	if err != nil {
		return authErrors.Wrap(c, authErrors.ErrNoSession, err)
	}

	isAuth, err := h.authUseCase.Auth(cookie.Value)
	if !isAuth {
		return authErrors.Wrap(c, authErrors.ErrNoSession, err)
	}

	user, err := h.usersUseCase.GetBySessionID(cookie.Value)
	if err != nil {
		return authErrors.Wrap(c, authErrors.ErrNoSession, err)
	}

	return httpUsers.UserResponse(user, c)
}

func (h *Handler) Login(c echo.Context) error {
	var data models.AuthUser
	err := c.Bind(&data)
	if err != nil {
		return authErrors.Wrap(c, authErrors.ErrBadRequestLogin, err)
	}

	sessionID, err := h.authUseCase.Login(data.Username, data.Password)
	if err != nil {
		return authErrors.Wrap(c, authErrors.ErrBadRequestLogin, err)
	}

	c.SetCookie(models.MakeHTTPCookieFromValue(sessionID))

	user, err := h.usersUseCase.GetBySessionID(sessionID)
	if err != nil {
		return authErrors.Wrap(c, authErrors.ErrUserNotFound, err)
	}

	return httpUsers.UserResponse(user, c)
}

func (h *Handler) Logout(c echo.Context) error {
	session, err := GetCookie(c)
	if err != nil {
		return authErrors.Wrap(c, authErrors.ErrNoSession, err)
	}

	isLogout, err := h.authUseCase.Logout(session.Value)
	if !isLogout {
		return authErrors.Wrap(c, authErrors.ErrBadSession, err)
	}

	session.Expires = time.Now().AddDate(
		deleteExpire["year"],
		deleteExpire["month"],
		deleteExpire["day"],
	)
	c.SetCookie(MakeHTTPCookie(session))

	return c.JSON(http.StatusOK, struct{}{})
}

func (h *Handler) SignUp(c echo.Context) error {
	newUser := models.User{}

	err := c.Bind(&newUser)
	if err != nil {
		return authErrors.Wrap(c, authErrors.ErrBadRequestSignUp, err)
	}

	sessionID, err := h.authUseCase.SignUp(&newUser)
	if err != nil {
		return authErrors.Wrap(c, authErrors.ErrBadRequestSignUp, err)
	}

	c.SetCookie(models.MakeHTTPCookieFromValue(sessionID))
	return httpUsers.UserResponse(&newUser, c)
}
