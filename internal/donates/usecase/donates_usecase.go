package donatesUsecase

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
)

type usecase struct {
	donatesMicroservice domain.DonatesMicroservice
	userMicroservice    domain.UsersMicroservice
}

func New(d domain.DonatesMicroservice, u domain.UsersMicroservice) domain.DonatesUseCase {
	return usecase{
		donatesMicroservice: d,
		userMicroservice:    u,
	}
}

func (u usecase) SendDonate(userID, authorID, price uint64) (models.Donate, error) {
	donate := models.Donate{
		UserID:   userID,
		AuthorID: authorID,
		Price:    price,
	}
	return u.donatesMicroservice.SendDonate(donate)
}

func (u usecase) GetDonateByID(id uint64) (models.Donate, error) {
	return u.donatesMicroservice.GetDonateByID(id)
}

func (u usecase) GetDonatesByUserID(userID uint64) ([]models.Donate, error) {
	return u.donatesMicroservice.GetDonatesByUserID(userID)
}
