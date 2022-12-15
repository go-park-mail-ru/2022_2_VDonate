package donatesMicroservice

import (
	"context"
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/donates/protobuf"
	userProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestDonatesClient_SendDonate(t *testing.T) {
	type mockSend func(s *mockDomain.MockDonatesClient, c context.Context, in *protobuf.Donate, opts ...grpc.CallOption)

	tests := []struct {
		name     string
		donate   models.Donate
		mock     mockSend
		response models.Donate
		err      error
	}{
		{
			name: "OK",
			mock: func(s *mockDomain.MockDonatesClient, c context.Context, in *protobuf.Donate, opts ...grpc.CallOption) {
				s.EXPECT().SendDonate(c, in).Return(&protobuf.Donate{
					Id:       1,
					UserId:   1,
					AuthorId: 1,
				}, nil)
			},
			response: models.Donate{
				ID:       1,
				UserID:   1,
				AuthorID: 1,
			},
			err: nil,
		},
		{
			name: "Error",
			mock: func(s *mockDomain.MockDonatesClient, c context.Context, in *protobuf.Donate, opts ...grpc.CallOption) {
				s.EXPECT().SendDonate(c, in).Return(&protobuf.Donate{}, errors.New("error"))
			},
			response: models.Donate{},
			err:      errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockDonatesClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.Donate{
				Id:       test.donate.ID,
				UserId:   test.donate.UserID,
				AuthorId: test.donate.AuthorID,
			})
			donatesClient := DonatesMicroservice{
				client: mock,
			}
			donate, err := donatesClient.SendDonate(test.donate)

			assert.Equal(t, donate, test.response)
			if err != nil {
				assert.Equal(t, err.Error(), test.err.Error())
			}
		})
	}
}

func TestDonatesCLient_GetDonatesByUserID(t *testing.T) {
	type MockGetDonates func(s *mockDomain.MockDonatesClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption)

	tests := []struct {
		name     string
		userID   uint64
		mock     MockGetDonates
		response []models.Donate
		err      error
	}{
		{
			name: "OK",
			mock: func(s *mockDomain.MockDonatesClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption) {
				s.EXPECT().GetDonatesByUserID(c, in).Return(&protobuf.DonateArray{
					Donates: []*protobuf.Donate{
						{
							Id:       1,
							UserId:   1,
							AuthorId: 1,
						},
					},
				}, nil)
			},
			response: []models.Donate{
				{
					ID:       1,
					UserID:   1,
					AuthorID: 1,
				},
			},
			err: nil,
		},
		{
			name: "OK",
			mock: func(s *mockDomain.MockDonatesClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption) {
				s.EXPECT().GetDonatesByUserID(c, in).Return(&protobuf.DonateArray{}, errors.New("error"))
			},
			err: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockDonatesClient(ctrl)

			test.mock(mock, context.Background(), &userProto.UserID{
				UserId: test.userID,
			})
			donatesClient := DonatesMicroservice{
				client: mock,
			}
			donates, err := donatesClient.GetDonatesByUserID(test.userID)

			assert.Equal(t, donates, test.response)
			if err != nil {
				assert.Equal(t, err.Error(), test.err.Error())
			}
		})
	}
}

func TestDonatesClient_GetDonateByID(t *testing.T) {
	type MockGetDonate func(s *mockDomain.MockDonatesClient, c context.Context, in *protobuf.DonateID, opts ...grpc.CallOption)

	tests := []struct {
		name     string
		donateID uint64
		mock     MockGetDonate
		response models.Donate
		err      error
	}{
		{
			name: "OK",
			mock: func(s *mockDomain.MockDonatesClient, c context.Context, in *protobuf.DonateID, opts ...grpc.CallOption) {
				s.EXPECT().GetDonateByID(c, in).Return(&protobuf.Donate{
					Id:       1,
					UserId:   1,
					AuthorId: 1,
				}, nil)
			},
			response: models.Donate{
				ID:       1,
				UserID:   1,
				AuthorID: 1,
			},
			err: nil,
		},
		{
			name: "Error",
			mock: func(s *mockDomain.MockDonatesClient, c context.Context, in *protobuf.DonateID, opts ...grpc.CallOption) {
				s.EXPECT().GetDonateByID(c, in).Return(&protobuf.Donate{}, errors.New("error"))
			},
			response: models.Donate{},
			err:      errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockDonatesClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.DonateID{
				Id: test.donateID,
			})
			donatesClient := DonatesMicroservice{
				client: mock,
			}
			donate, err := donatesClient.GetDonateByID(test.donateID)

			assert.Equal(t, donate, test.response)
			if err != nil {
				assert.Equal(t, err.Error(), test.err.Error())
			}
		})
	}
}
