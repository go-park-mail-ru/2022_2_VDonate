package httpAuth

import (
	"net/http"
	"time"

	"github.com/ztrue/tracerr"

	errorHandling "github.com/go-park-mail-ru/2022_2_VDonate/pkg/errors"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
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
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return cookie, nil
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

// Auth godoc
// @Summary     User authentication request
// @Description Check authentication of `User` by cookies
// @ID          auth
// @Tags        auth
// @Produce     json
// @Success     200 {object} models.UserID  "Session was successfully found"
// @Failure     401 {object} echo.HTTPError "User is unauthorized"
// @Failure     404 {object} echo.HTTPError "User was not found"
// @Router      /auth [get]
func (h Handler) Auth(c echo.Context) error {
	cookie, err := GetCookie(c)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}

	isAuth, err := h.authUseCase.Auth(cookie.Value)
	if !isAuth {
		return errorHandling.WrapEcho(domain.ErrAuth, err)
	}

	user, err := h.usersUseCase.GetBySessionID(cookie.Value)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}

	return c.JSON(http.StatusOK, models.UserID{
		ID: user.ID,
	})
}

// Login godoc
// @Summary     User login request
// @Description Authorization of `User`
// @ID          login
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       authData body     models.AuthUser true "username and password"
// @Success     200      {object} models.UserID   "Session was successfully found"
// @Failure     400      {object} echo.HTTPError  "Wrong login or password or bad data was received"
// @Failure     404      {object} echo.HTTPError  "User was not found"
// @Router      /login [post]
func (h Handler) Login(c echo.Context) error {
	var data models.AuthUser
	if err := c.Bind(&data); err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	sessionID, err := h.authUseCase.Login(data.Username, data.Password)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}

	c.SetCookie(makeHTTPCookieFromValue(sessionID))

	user, err := h.usersUseCase.GetBySessionID(sessionID)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}

	return c.JSON(http.StatusOK, models.UserID{
		ID: user.ID,
	})
}

// Logout godoc
// @Summary     User logout
// @Description Get request for user logout
// @ID          logout
// @Tags        auth
// @Produce     json
// @Success     200 {object} models.EmptyStruct "Successfully logout"
// @Failure     400 {object} echo.HTTPError     "Bad session / request"
// @Failure     401 {object} echo.HTTPError     "No session provided"
// @Security    ApiKeyAuth
// @Router      /logout [delete]
func (h Handler) Logout(c echo.Context) error {
	session, err := GetCookie(c)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}

	isLogout, err := h.authUseCase.Logout(session.Value)
	if !isLogout {
		return errorHandling.WrapEcho(domain.ErrBadSession, err)
	}

	session.Expires = time.Now().AddDate(
		deleteExpire["year"],
		deleteExpire["month"],
		deleteExpire["day"],
	)
	c.SetCookie(MakeHTTPCookie(session))

	return c.JSON(http.StatusOK, models.EmptyStruct{})
}

// SignUp godoc
// @Summary     Creates a User
// @Description Request to server for `User` creation
// @ID          signup
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       data body     models.UserMpfd true "POST request of all information about `User`"
// @Success     200  {object} models.UserID   "User was successfully created"
// @Failure     400  {object} echo.HTTPError  "Bad request"
// @Failure     409  {object} echo.HTTPError  "Username or email is already exists"
// @Failure     500  {object} echo.HTTPError  "Internal error"
// @Router      /users [post]
func (h Handler) SignUp(c echo.Context) error {
	var newUser models.User

	if err := c.Bind(&newUser); err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	sessionID, err := h.authUseCase.SignUp(newUser)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrInternal, err)
	}

	c.SetCookie(makeHTTPCookieFromValue(sessionID))

	user, err := h.usersUseCase.GetBySessionID(sessionID)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}

	return c.JSON(http.StatusOK, models.UserID{
		ID: user.ID,
	})
}
