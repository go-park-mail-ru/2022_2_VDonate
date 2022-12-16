package notificationsWS

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

	"github.com/sirupsen/logrus"

	"github.com/go-park-mail-ru/2022_2_VDonate/pkg/logger"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/gorilla/websocket"
)

type Handler struct {
	u         domain.NotificationsUseCase
	wsUpgrade *websocket.Upgrader
}

func NewHandler(up *websocket.Upgrader, u domain.NotificationsUseCase) *Handler {
	return &Handler{
		u:         u,
		wsUpgrade: up,
	}
}

func (h Handler) Handler(w http.ResponseWriter, r *http.Request) {
	l := logger.GetInstance().Logrus

	c, err := h.wsUpgrade.Upgrade(w, r, nil)
	if err != nil {
		l.Error(err)
		return
	}
	defer c.Close()

	params := r.URL.Query()
	userID, err := strconv.ParseUint(params.Get("userID"), 10, 64)
	if err != nil {
		l.WithFields(
			logrus.Fields{
				"error": err,
			},
		).Error("ERROR in parsing userID")
		return
	}

	cancel := models.NotificationCancel{}

	for {
		time.Sleep(time.Second * 1)
		notifications, err := h.u.GetNotifications(userID)
		if err != nil && err != sql.ErrNoRows {
			l.WithFields(
				logrus.Fields{
					"error": err,
				},
			).Error("ERROR in getting notifications")
			return
		}
		if err == sql.ErrNoRows {
			l.Warn(sql.ErrNoRows.Error())
		}
		if len(notifications) != 0 {
			if err = c.WriteJSON(notifications); err != nil {
				l.WithFields(
					logrus.Fields{
						"error": err,
					},
				).Error("ERROR in writing notifications -- closing connection")
				return
			}

			_ = c.ReadJSON(&cancel)
			if cancel.Cancel {
				if err = h.u.DeleteNotifications(userID); err != nil {
					l.WithFields(
						logrus.Fields{
							"error": err,
						},
					).Error("ERROR in deleting notifications")
					return
				}
				cancel.Cancel = false
			}
		}
	}
}
