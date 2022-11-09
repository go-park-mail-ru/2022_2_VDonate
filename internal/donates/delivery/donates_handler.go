package httpdonates

import (
	"net/http"
	"strconv"

	errorHandling "github.com/go-park-mail-ru/2022_2_VDonate/pkg/errors"

	httpAuth "github.com/go-park-mail-ru/2022_2_VDonate/internal/auth/delivery"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	handlerUsecase domain.DonatesUseCase
	usersUsecase   domain.UsersUseCase
}

func New(d domain.DonatesUseCase, u domain.UsersUseCase) *Handler {
	return &Handler{
		handlerUsecase: d,
		usersUsecase:   u,
	}
}

func (h *Handler) CreateDonate(c echo.Context) error {
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

	donate, err := h.handlerUsecase.SendDonate(user.ID, d.AuthorID, d.Price)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrCreate, err)
	}

	return c.JSON(http.StatusOK, donate)
}

func (h *Handler) GetDonate(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrBadRequest, err)
	}

	donate, err := h.handlerUsecase.GetDonateByID(id)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}

	return c.JSON(http.StatusOK, donate)
}

func (h *Handler) GetDonates(c echo.Context) error {
	cookie, err := httpAuth.GetCookie(c)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}
	user, err := h.usersUsecase.GetBySessionID(cookie.Value)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNoSession, err)
	}

	donates, err := h.handlerUsecase.GetDonatesByUserID(user.ID)
	if err != nil {
		return errorHandling.WrapEcho(domain.ErrNotFound, err)
	}

	return c.JSON(http.StatusOK, donates)
}
