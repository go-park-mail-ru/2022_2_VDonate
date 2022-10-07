package httpAuth

import (
	"errors"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/errors"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/middlewares"
	model "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/session/repository"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/users/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Handler struct {
	userAPI     userAPI.UseCase
	sessionRepo sessionRepository.API
}

func NewHandler(userAPI userAPI.UseCase, sessionRepo sessionRepository.API) *Handler {
	return &Handler{userAPI: userAPI, sessionRepo: sessionRepo}
}

func (h *Handler) getCookie(c echo.Context) (*http.Cookie, error) {
	return c.Cookie("session_id")
}

func (h *Handler) createCookie(id uint) *model.Cookie {
	return &model.Cookie{
		UserID:  id,
		Value:   utils.RandStringRunes(32),
		Expires: time.Now().AddDate(0, 1, 0),
	}
}

func (h *Handler) makeHTTPCookie(c *http.Cookie) *http.Cookie {
	return &http.Cookie{
		Name:     "session_id",
		Value:    c.Value,
		Expires:  c.Expires,
		HttpOnly: true,
		Secure:   true,
	}
}

func (h *Handler) httpCookieFromModel(c *model.Cookie) *http.Cookie {
	return &http.Cookie{
		Name:     "session_id",
		Value:    c.Value,
		Expires:  c.Expires,
		HttpOnly: true,
		Secure:   true,
	}
}

func (h *Handler) Auth(c echo.Context) error {
	httpCookie, err := h.getCookie(c)
	if err == http.ErrNoCookie {
		return authErrors.Wrap(c, authErrors.ErrNoSession, err)
	}

	cookie, err := h.sessionRepo.GetByValue(httpCookie.Value)
	if err != nil {
		return authErrors.Wrap(c, authErrors.ErrNoSession, errors.New("failed to get session"))
	}
	user, err := h.userAPI.GetByID(cookie.UserID)
	if err != nil {
		return authErrors.Wrap(c, authErrors.ErrUserNotFound, err)
	}

	return middlewares.UserResponse(user, c)
}

func (h *Handler) Login(c echo.Context) error {
	var data model.AuthUser
	err := c.Bind(&data)
	if err != nil {
		return authErrors.Wrap(c, authErrors.ErrBadRequest, err)
	}

	user, err := h.userAPI.GetByUsername(data.Username)
	if err != nil {
		return authErrors.Wrap(c, authErrors.ErrUserNotFound, err)
	}

	if data.Password != user.Password {
		return authErrors.Wrap(c, authErrors.ErrInvalidLoginOrPassword, errors.New("passwords not the same"))
	}

	s, err := h.sessionRepo.Create(h.createCookie(user.ID))
	if err != nil {
		return authErrors.Wrap(c, authErrors.ErrBadSession, err)
	}

	c.SetCookie(h.httpCookieFromModel(s))
	return middlewares.UserResponse(user, c)
}

func (h *Handler) Logout(c echo.Context) error {
	session, err := h.getCookie(c)
	if err == http.ErrNoCookie {
		return authErrors.Wrap(c, authErrors.ErrNoSession, err)
	}

	if err = h.sessionRepo.DeleteByValue(session.Value); err != nil {
		return authErrors.Wrap(c, authErrors.ErrDeleteSession, err)
	}

	session.Expires = time.Now().AddDate(0, 0, -1)
	c.SetCookie(h.makeHTTPCookie(session))

	return c.JSON(http.StatusOK, struct{}{})
}

func (h *Handler) SignUp(c echo.Context) error {
	newUser := model.UserDB{}

	err := c.Bind(&newUser)
	if err != nil {
		return authErrors.Wrap(c, authErrors.ErrBadRequest, err)
	}

	if h.userAPI.IsExistUsernameAndEmail(newUser.Username, newUser.Email) {
		return authErrors.Wrap(c, authErrors.ErrUserOrEmailAlreadyExist, authErrors.ErrUserOrEmailAlreadyExist)
	}

	user, err := h.userAPI.Create(&newUser)
	if err != nil {
		return authErrors.Wrap(c, authErrors.ErrCreateUser, err)
	}

	s, err := h.sessionRepo.Create(h.createCookie(user.ID))
	if err != nil {
		return authErrors.Wrap(c, authErrors.ErrBadSession, err)
	}

	c.SetCookie(h.httpCookieFromModel(s))
	return middlewares.UserResponse(user, c)
}
