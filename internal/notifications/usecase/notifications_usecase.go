package notificationsUsecase

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
)

type usecase struct {
	repo domain.NotificationsRepository
}

func New(repo domain.NotificationsRepository) domain.NotificationsUseCase {
	return &usecase{
		repo: repo,
	}
}

func (u usecase) GetNotifications(userID uint64) ([]models.Notification, error) {
	return u.repo.GetNotificationsByUserID(userID)
}

func (u usecase) DeleteNotifications(userID uint64) error {
	return u.repo.DeleteNotificationByUserID(userID)
}
