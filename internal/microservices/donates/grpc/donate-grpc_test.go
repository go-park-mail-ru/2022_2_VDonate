package grpcDonate

import (
	"context"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/donates/protobuf"
	userProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	mock_domain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConvertToModel(t *testing.T) {
	input := &protobuf.Donate{
		Id:       1,
		UserId:   2,
		AuthorId: 3,
		Price:    4,
	}

	expected := models.Donate{
		ID:       1,
		UserID:   2,
		AuthorID: 3,
		Price:    4,
	}

	actual := ConvertToModel(input)

	assert.Equal(t, expected, actual)
}

func TestConvertToProto(t *testing.T) {
	input := models.Donate{
		ID:       1,
		UserID:   2,
		AuthorID: 3,
		Price:    4,
	}

	expected := &protobuf.Donate{
		Id:       1,
		UserId:   2,
		AuthorId: 3,
		Price:    4,
	}

	actual := ConvertToProto(input)

	assert.Equal(t, expected, actual)
}

func TestDonate_SendDonate(t *testing.T) {
	type mockBehaviorSendDonate func(r *mock_domain.MockDonatesRepository, d models.Donate)

	tests := []struct {
		name           string
		inputDonate    *protobuf.Donate
		mockBehavior   mockBehaviorSendDonate
		expectedDonate *protobuf.Donate
		expectedError  string
	}{
		{
			name: "OK",
			inputDonate: &protobuf.Donate{
				Id:       1,
				UserId:   2,
				AuthorId: 3,
				Price:    4,
			},
			mockBehavior: func(r *mock_domain.MockDonatesRepository, d models.Donate) {
				r.EXPECT().SendDonate(d).Return(d, nil)
			},
			expectedDonate: &protobuf.Donate{
				Id:       1,
				UserId:   2,
				AuthorId: 3,
				Price:    4,
			},
		},
		{
			name: "Error",
			inputDonate: &protobuf.Donate{
				Id:       1,
				UserId:   2,
				AuthorId: 3,
				Price:    4,
			},
			mockBehavior: func(r *mock_domain.MockDonatesRepository, d models.Donate) {
				r.EXPECT().SendDonate(d).Return(d, domain.ErrInternal)
			},
			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockDonatesRepository(ctrl)
			test.mockBehavior(repo, ConvertToModel(test.inputDonate))

			s := New(repo)

			session, err := s.SendDonate(context.Background(), test.inputDonate)

			if err != nil {
				require.Equal(t, test.expectedError, err.Error())
			} else {
				require.Equal(t, test.expectedDonate, session)
			}
		})
	}
}

func TestDonate_GetDonatesByUserID(t *testing.T) {
	type mockBehaviorGetDonatesByUserID func(r *mock_domain.MockDonatesRepository, userID uint64)

	tests := []struct {
		name            string
		inputUserID     uint64
		mockBehavior    mockBehaviorGetDonatesByUserID
		expectedDonates *protobuf.DonateArray
		expectedError   string
	}{
		{
			name:        "OK",
			inputUserID: 1,
			mockBehavior: func(r *mock_domain.MockDonatesRepository, userID uint64) {
				r.EXPECT().GetDonatesByUserID(userID).Return([]models.Donate{
					{
						ID:       1,
						UserID:   2,
						AuthorID: 3,
						Price:    4,
					},
				}, nil)
			},
			expectedDonates: &protobuf.DonateArray{
				Donates: []*protobuf.Donate{
					{
						Id:       1,
						UserId:   2,
						AuthorId: 3,
						Price:    4,
					},
				},
			},
		},
		{
			name:        "Error",
			inputUserID: 1,
			mockBehavior: func(r *mock_domain.MockDonatesRepository, userID uint64) {
				r.EXPECT().GetDonatesByUserID(userID).Return([]models.Donate{}, domain.ErrInternal)
			},
			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockDonatesRepository(ctrl)
			test.mockBehavior(repo, test.inputUserID)

			s := New(repo)

			donates, err := s.GetDonatesByUserID(context.Background(), &userProto.UserID{UserId: test.inputUserID})

			if err != nil {
				require.Equal(t, test.expectedError, err.Error())
			} else {
				require.Equal(t, test.expectedDonates, donates)
			}
		})
	}
}

func TestDonate_GetDonateByID(t *testing.T) {
	type mockBehaviorGetDonateByID func(r *mock_domain.MockDonatesRepository, donateID uint64)

	tests := []struct {
		name           string
		inputDonateID  uint64
		mockBehavior   mockBehaviorGetDonateByID
		expectedDonate *protobuf.Donate
		expectedError  string
	}{
		{
			name:          "OK",
			inputDonateID: 1,
			mockBehavior: func(r *mock_domain.MockDonatesRepository, donateID uint64) {
				r.EXPECT().GetDonateByID(donateID).Return(models.Donate{
					ID:       1,
					UserID:   2,
					AuthorID: 3,
					Price:    4,
				}, nil)
			},
			expectedDonate: &protobuf.Donate{
				Id:       1,
				UserId:   2,
				AuthorId: 3,
				Price:    4,
			},
		},
		{
			name:          "Error",
			inputDonateID: 1,
			mockBehavior: func(r *mock_domain.MockDonatesRepository, donateID uint64) {
				r.EXPECT().GetDonateByID(donateID).Return(models.Donate{}, domain.ErrInternal)
			},
			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockDonatesRepository(ctrl)
			test.mockBehavior(repo, test.inputDonateID)

			s := New(repo)

			donate, err := s.GetDonateByID(context.Background(), &protobuf.DonateID{Id: test.inputDonateID})

			if err != nil {
				require.Equal(t, test.expectedError, err.Error())
			} else {
				require.Equal(t, test.expectedDonate, donate)
			}
		})
	}
}
