package donatesMicroservice

import (
	"context"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"

	grpcDonate "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/donates/grpc"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/donates/protobuf"
	userProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
)

type DonatesMicroservice struct {
	client protobuf.DonatesClient
}

func New(c protobuf.DonatesClient) domain.DonatesMicroservice {
	return &DonatesMicroservice{
		client: c,
	}
}

func (m DonatesMicroservice) SendDonate(donate models.Donate) (models.Donate, error) {
	newDonate, err := m.client.SendDonate(context.Background(), grpcDonate.ConvertToProto(donate))
	if err != nil {
		return models.Donate{}, err
	}

	return grpcDonate.ConvertToModel(newDonate), nil
}

func (m DonatesMicroservice) GetDonatesByUserID(userID uint64) ([]models.Donate, error) {
	donates, err := m.client.GetDonatesByUserID(context.Background(), &userProto.UserID{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}

	result := make([]models.Donate, 0)

	for _, d := range donates.GetDonates() {
		result = append(result, grpcDonate.ConvertToModel(d))
	}

	return result, nil
}

func (m DonatesMicroservice) GetDonateByID(donateID uint64) (models.Donate, error) {
	d, err := m.client.GetDonateByID(context.Background(), &protobuf.DonateID{
		Id: donateID,
	})
	if err != nil {
		return models.Donate{}, err
	}

	return grpcDonate.ConvertToModel(d), nil
}
