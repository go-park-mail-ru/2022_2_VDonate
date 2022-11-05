package donatesUsecase

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
)

type usecase struct {
	donatesRepo domain.DonatesRepository
	userRepo    domain.UsersRepository
}

func New(d domain.DonatesRepository, u domain.UsersRepository) domain.DonatesUseCase {
	return usecase{
		donatesRepo: d,
		userRepo:    u,
	}
}

func (u usecase) SendDonate(userID, authorID, price uint64) (models.Donate, error) {
	donate := models.Donate{
		UserID:   userID,
		AuthorID: authorID,
		Price:    price,
	}
	return u.donatesRepo.SendDonate(donate)
}

func (u usecase) GetDonateByID(ID uint64) (models.Donate, error) {
	return u.donatesRepo.GetDonateByID(ID)
}

func (u usecase) GetDonatesByUserID(userID uint64) ([]models.Donate, error) {
	return u.donatesRepo.GetDonatesByUserID(userID)
}
