package notificationsWS

import (
	"database/sql"
	"net/http"
	"reflect"
	"strconv"

	"github.com/ztrue/tracerr"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

	"github.com/sirupsen/logrus"

	"github.com/go-park-mail-ru/2022_2_VDonate/pkg/logger"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/gorilla/websocket"
)

const (
	exit = 1
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

func (h Handler) GetNewNotifications(newN, oldN []models.Notification) []models.Notification {
	if reflect.DeepEqual(newN, oldN) {
		return []models.Notification{}
	}
	if len(newN) > len(oldN) {
		return newN[len(oldN):]
	}

	return []models.Notification{}
}

type status struct {
	Remove      chan interface{}
	CloseReader chan interface{}
	CloseWriter chan interface{}
}

func handleError(err error, c *websocket.Conn, l *logrus.Logger) {
	if err != nil {
		l.WithFields(
			logrus.Fields{
				"error": err,
			},
		).Error(err)
		if err != websocket.ErrCloseSent {
			err = c.Close()
			l.Warn("connection closed: ", err)
		}
	}

	return
}

func (h Handler) Reader(conn *websocket.Conn, s *status) {
	var cancel models.NotificationCancel
	for {
		err := conn.ReadJSON(&cancel)
		if err != nil {
			s.CloseWriter <- exit
			return
		}
		if cancel.Cancel {
			s.Remove <- exit
			cancel.Cancel = false
		}
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
		).Error("Bad Request")
		return
	}

	closeReaderSignal := make(chan interface{})
	closeWriterSignal := make(chan interface{})
	removeSignal := make(chan interface{})

	s := &status{
		Remove:      removeSignal,
		CloseReader: closeReaderSignal,
		CloseWriter: closeWriterSignal,
	}

	go h.Reader(c, s)

	var oldN []models.Notification
	for {
		select {
		case <-s.CloseWriter:
			l.Warn("writer done")
			return
		case <-s.Remove:
			if err = h.u.DeleteNotifications(userID); err == nil {
				handleError(tracerr.Wrap(err), c, l)
			}
		default:
			notifications, err := h.u.GetNotifications(userID)
			if err != nil && err != sql.ErrNoRows {
				handleError(tracerr.Wrap(err), c, l)
			}

			toSend := h.GetNewNotifications(notifications, oldN)

			if len(toSend) > 0 {
				err = c.WriteJSON(toSend)
				handleError(tracerr.Wrap(err), c, l)
			}

			oldN = notifications
		}
	}
}
