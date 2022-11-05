package domain

import "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

type DonatesUseCase interface {
	SendDonate(userID, authorID, price uint64) (models.Donate, error)
	GetDonateByID(ID uint64) (models.Donate, error)
	GetDonatesByUserID(userID uint64) ([]models.Donate, error)
}