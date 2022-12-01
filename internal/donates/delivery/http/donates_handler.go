package httpDonates

import (
	"net/http"
	"strconv"

	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery/http"

	errorHandling "github.com/go-park-mail-ru/2022_2_VDonate/pkg/errors"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	donateUsecase domain.DonatesUseCase
	usersUsecase  domain.UsersUseCase
}

func NewHandler(d domain.DonatesUseCase, u domain.UsersUseCase) *Handler {
	return &Handler{
		donateUsecase: d,
		usersUsecase:  u,
	}
}

// CreateDonate godoc
// @Summary     Create donate
// @Description Send donate to author
// @ID          create_donate
// @Tags        donates
// @Param       post body models.DonateMpfd true "Donate Fields"
// @Produce     json
// @Success     200 {object} models.Donate  "Donate was successfully created"
// @Failure     400 {object} echo.HTTPError "Bad request"
// @Failure     401 {object} echo.HTTPError "No session provided"
// @Failure     404 {object} echo.HTTPError "Post not found"
// @Failure     500 {object} echo.HTTPError "Internal error"
// @Security    ApiKeyAuth
// @Router      /donate [post]
func (h Handler) CreateDonate(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}
	user, err := h.usersUsecase.GetBySessionID(cookie.Value)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}

	var d models.Donate
	if err = c.Bind(&d); err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	donate, err := h.donateUsecase.SendDonate(user.ID, d.AuthorID, d.Price)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrCreate, err)
	}

	return c.JSON(http.StatusOK, donate)
}

// GetDonate    godoc
// @Summary     Get donate
// @Description Get donate by id
// @ID          get_donate
// @Tags        donates
// @Param       id path integer true "Post ID"
// @Produce     json
// @Success     200 {object} models.Donate  "Donate was successfully create"
// @Failure     400 {object} echo.HTTPError "Bad request"
// @Failure     401 {object} echo.HTTPError "No session provided"
// @Failure     500 {object} echo.HTTPError "Internal error"
// @Security    ApiKeyAuth
// @Router      /donate/{id} [get]
func (h Handler) GetDonate(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	donate, err := h.donateUsecase.GetDonateByID(id)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}

	return c.JSON(http.StatusOK, donate)
}

// GetDonates   godoc
// @Summary     Get donates
// @Description Get donates of user
// @ID          get_donates
// @Tags        donates
// @Produce     json
// @Success     200 {object} []models.Donate "Donates were successfully recieved"
// @Failure     400 {object} echo.HTTPError  "Bad request"
// @Failure     401 {object} echo.HTTPError  "No session provided"
// @Failure     500 {object} echo.HTTPError  "Internal error"
// @Security    ApiKeyAuth
// @Router      /donates [get]
func (h Handler) GetDonates(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}
	user, err := h.usersUsecase.GetBySessionID(cookie.Value)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}

	donates, err := h.donateUsecase.GetDonatesByUserID(user.ID)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}

	return c.JSON(http.StatusOK, donates)
}
