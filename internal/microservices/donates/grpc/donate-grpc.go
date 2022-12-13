package grpcDonate

import (
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/donates/protobuf"
	userProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
)

type DonateService struct {
	donateRepo domain.DonatesRepository
	protobuf.UnimplementedDonatesServer
}

func New(r domain.DonatesRepository) protobuf.DonatesServer {
	return &DonateService{
		donateRepo: r,
	}
}

func ConvertToModel(d *protobuf.Donate) models.Donate {
	return models.Donate{
		ID:       d.GetId(),
		UserID:   d.GetUserId(),
		AuthorID: d.GetAuthorId(),
		Price:    d.GetPrice(),
	}
}

func ConvertToProto(d models.Donate) *protobuf.Donate {
	return &protobuf.Donate{
		Id:       d.ID,
		UserId:   d.UserID,
		AuthorId: d.AuthorID,
		Price:    d.Price,
	}
}

func (s DonateService) SendDonate(_ context.Context, donate *protobuf.Donate) (*protobuf.Donate, error) {
	d, err := s.donateRepo.SendDonate(ConvertToModel(donate))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return ConvertToProto(d), nil
}

func (s DonateService) GetDonatesByUserID(_ context.Context, userID *userProto.UserID) (*protobuf.DonateArray, error) {
	d, err := s.donateRepo.GetDonatesByUserID(userID.GetUserId())
	if err != nil && err != sql.ErrNoRows {
		return nil, status.Error(codes.Internal, err.Error())
	}

	result := make([]*protobuf.Donate, 0)

	for _, donate := range d {
		result = append(result, ConvertToProto(donate))
	}

	return &protobuf.DonateArray{
		Donates: result,
	}, nil
}

func (s DonateService) GetDonateByID(_ context.Context, donateID *protobuf.DonateID) (*protobuf.Donate, error) {
	d, err := s.donateRepo.GetDonateByID(donateID.GetId())
	if err != nil && err != sql.ErrNoRows {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return ConvertToProto(d), nil
}
