package httpsubscribers

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	subscribersDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/subscribers"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	subscribersUsecase subscribersDomain.UseCase
	userUsecase        domain.UsersUseCase
}

func New(subscribersUsecase subscribersDomain.UseCase, userUsecase domain.UsersUseCase) *Handler {
	return &Handler{
		subscribersUsecase: subscribersUsecase,
		userUsecase:        userUsecase,
	}
}

func (h *Handler) GetSubscribers(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("author_id"), 10, 64)
	if err != nil {
		return err
	}
	s, err := h.subscribersUsecase.GetSubscribers(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, s)
}

func (h *Handler) CreateSubscriber(c echo.Context) error {
	var s models.Subscription
	if err := c.Bind(&s); err != nil {
		return err
	}

	err := h.subscribersUsecase.Subscribe(s)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, s)
}

func (h *Handler) DeleteSubscriber(c echo.Context) error {
	var s models.Subscription
	if err := c.Bind(&s); err != nil {
		return err
	}
	err := h.subscribersUsecase.Unsubscribe(s.SubscriberID, s.AuthorID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, struct{}{})
}
