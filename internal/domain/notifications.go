package domain

import "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

type NotificationsUseCase interface {
	GetNotifications(userID uint64) ([]models.Notification, error)
	DeleteNotifications(userID uint64) error
}
