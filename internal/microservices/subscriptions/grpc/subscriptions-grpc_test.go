package grpcSubscriptions

import (
	"context"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscriptions/protobuf"
	userProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	mock_domain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertToModel(t *testing.T) {
	input := &protobuf.AuthorSubscription{
		ID:           1,
		AuthorID:     1,
		Img:          "img",
		Tier:         1,
		Title:        "title",
		Text:         "text",
		Price:        1,
		AuthorName:   "authorName",
		AuthorAvatar: "authorAvatar",
	}

	expected := models.AuthorSubscription{
		ID:           1,
		AuthorID:     1,
		Img:          "img",
		Tier:         1,
		Title:        "title",
		Text:         "text",
		Price:        1,
		AuthorName:   "authorName",
		AuthorAvatar: "authorAvatar",
	}

	actual := ConvertToModel(input)

	assert.Equal(t, expected, actual)
}

func TestConvertToProto(t *testing.T) {
	input := models.AuthorSubscription{
		ID:           1,
		AuthorID:     1,
		Img:          "img",
		Tier:         1,
		Title:        "title",
		Text:         "text",
		Price:        1,
		AuthorName:   "authorName",
		AuthorAvatar: "authorAvatar",
	}

	expected := &protobuf.AuthorSubscription{
		ID:           1,
		AuthorID:     1,
		Img:          "img",
		Tier:         1,
		Title:        "title",
		Text:         "text",
		Price:        1,
		AuthorName:   "authorName",
		AuthorAvatar: "authorAvatar",
	}

	actual := ConvertToProto(input)

	assert.Equal(t, expected, actual)
}

func TestSubscriptionsService_AddSubscription(t *testing.T) {
	type mockBehaviorAddSubscription func(r *mock_domain.MockSubscriptionsRepository, input models.AuthorSubscription)

	tests := []struct {
		name         string
		input        models.AuthorSubscription
		mockBehavior mockBehaviorAddSubscription
		expectedErr  string
	}{
		{
			name: "OK",
			input: models.AuthorSubscription{
				ID:           1,
				AuthorID:     1,
				Img:          "img",
				Tier:         1,
				Title:        "title",
				Text:         "text",
				Price:        1,
				AuthorName:   "authorName",
				AuthorAvatar: "authorAvatar",
			},
			mockBehavior: func(r *mock_domain.MockSubscriptionsRepository, input models.AuthorSubscription) {
				r.EXPECT().AddSubscription(input).Return(input.ID, nil)
			},
		},
		{
			name: "Error",
			input: models.AuthorSubscription{
				ID:           1,
				AuthorID:     1,
				Img:          "img",
				Tier:         1,
				Title:        "title",
				Text:         "text",
				Price:        1,
				AuthorName:   "authorName",
				AuthorAvatar: "authorAvatar",
			},
			mockBehavior: func(r *mock_domain.MockSubscriptionsRepository, input models.AuthorSubscription) {
				r.EXPECT().AddSubscription(input).Return(input.ID, domain.ErrInternal)
			},
			expectedErr: domain.ErrInternal.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockSubscriptionsRepository(ctrl)
			test.mockBehavior(repo, test.input)

			s := New(repo)

			_, err := s.AddSubscription(context.Background(), ConvertToProto(test.input))

			if err != nil {
				assert.Equal(t, test.expectedErr, err.Error())
			}
		})
	}
}

func TestSubscriptionsService_DeleteSubscription(t *testing.T) {
	type mockBehaviorDeleteSubscription func(r *mock_domain.MockSubscriptionsRepository, id uint64)

	tests := []struct {
		name         string
		id           uint64
		mockBehavior mockBehaviorDeleteSubscription
		expectedErr  string
	}{
		{
			name: "OK",
			id:   1,
			mockBehavior: func(r *mock_domain.MockSubscriptionsRepository, id uint64) {
				r.EXPECT().DeleteSubscription(id).Return(nil)
			},
		},
		{
			name: "Error",
			id:   1,
			mockBehavior: func(r *mock_domain.MockSubscriptionsRepository, id uint64) {
				r.EXPECT().DeleteSubscription(id).Return(domain.ErrInternal)
			},
			expectedErr: domain.ErrInternal.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockSubscriptionsRepository(ctrl)
			test.mockBehavior(repo, test.id)

			s := New(repo)

			_, err := s.DeleteSubscription(context.Background(), &protobuf.AuthorSubscriptionID{ID: test.id})

			if err != nil {
				assert.Equal(t, test.expectedErr, err.Error())
			}
		})
	}
}

func TestSubscriptionsService_GetSubscriptionByID(t *testing.T) {
	type mockBehaviorGetSubscriptionByID func(r *mock_domain.MockSubscriptionsRepository, id uint64)

	tests := []struct {
		name         string
		id           uint64
		mockBehavior mockBehaviorGetSubscriptionByID
		expectedErr  string
	}{
		{
			name: "OK",
			id:   1,
			mockBehavior: func(r *mock_domain.MockSubscriptionsRepository, id uint64) {
				r.EXPECT().GetSubscriptionByID(id).Return(models.AuthorSubscription{
					ID:           1,
					AuthorID:     1,
					Img:          "img",
					Tier:         1,
					Title:        "title",
					Text:         "text",
					Price:        1,
					AuthorName:   "authorName",
					AuthorAvatar: "authorAvatar",
				}, nil)
			},
		},
		{
			name: "Error",
			id:   1,
			mockBehavior: func(r *mock_domain.MockSubscriptionsRepository, id uint64) {
				r.EXPECT().GetSubscriptionByID(id).Return(models.AuthorSubscription{}, domain.ErrInternal)
			},
			expectedErr: domain.ErrInternal.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockSubscriptionsRepository(ctrl)
			test.mockBehavior(repo, test.id)

			s := New(repo)

			_, err := s.GetSubscriptionByID(context.Background(), &protobuf.AuthorSubscriptionID{ID: test.id})

			if err != nil {
				assert.Equal(t, test.expectedErr, err.Error())
			}
		})
	}
}

func TestSubscriptionsService_GetSubscriptionByUserAndAuthorID(t *testing.T) {
	type mockBehaviorGetSubscriptionByUserAndAuthorID func(r *mock_domain.MockSubscriptionsRepository, userID, authorID uint64)

	tests := []struct {
		name         string
		userID       uint64
		authorID     uint64
		mockBehavior mockBehaviorGetSubscriptionByUserAndAuthorID
		expectedErr  string
	}{
		{
			name:     "OK",
			userID:   1,
			authorID: 1,
			mockBehavior: func(r *mock_domain.MockSubscriptionsRepository, userID, authorID uint64) {
				r.EXPECT().GetSubscriptionByUserAndAuthorID(userID, authorID).Return(models.AuthorSubscription{
					ID:           1,
					AuthorID:     1,
					Img:          "img",
					Tier:         1,
					Title:        "title",
					Text:         "text",
					Price:        1,
					AuthorName:   "authorName",
					AuthorAvatar: "authorAvatar",
				}, nil)
			},
		},
		{
			name:     "Error",
			userID:   1,
			authorID: 1,
			mockBehavior: func(r *mock_domain.MockSubscriptionsRepository, userID, authorID uint64) {
				r.EXPECT().GetSubscriptionByUserAndAuthorID(userID, authorID).Return(models.AuthorSubscription{}, domain.ErrInternal)
			},
			expectedErr: domain.ErrInternal.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockSubscriptionsRepository(ctrl)
			test.mockBehavior(repo, test.userID, test.authorID)

			s := New(repo)

			_, err := s.GetSubscriptionByUserAndAuthorID(context.Background(), &userProto.UserAuthorPair{
				UserId:   test.userID,
				AuthorId: test.authorID,
			})

			if err != nil {
				assert.Equal(t, test.expectedErr, err.Error())
			}
		})
	}
}

func TestSubscriptionsService_GetSubscriptionsByAuthorID(t *testing.T) {
	type mockBehaviorGetSubscriptionsByAuthorID func(r *mock_domain.MockSubscriptionsRepository, authorID uint64)

	tests := []struct {
		name         string
		authorID     uint64
		mockBehavior mockBehaviorGetSubscriptionsByAuthorID
		expectedErr  string
	}{
		{
			name:     "OK",
			authorID: 1,
			mockBehavior: func(r *mock_domain.MockSubscriptionsRepository, authorID uint64) {
				r.EXPECT().GetSubscriptionsByAuthorID(authorID).Return([]models.AuthorSubscription{
					{
						ID:           1,
						AuthorID:     1,
						Img:          "img",
						Tier:         1,
						Title:        "title",
						Text:         "text",
						Price:        1,
						AuthorName:   "authorName",
						AuthorAvatar: "authorAvatar",
					},
				}, nil)
			},
		},
		{
			name:     "Error",
			authorID: 1,
			mockBehavior: func(r *mock_domain.MockSubscriptionsRepository, authorID uint64) {
				r.EXPECT().GetSubscriptionsByAuthorID(authorID).Return([]models.AuthorSubscription{}, domain.ErrInternal)
			},
			expectedErr: domain.ErrInternal.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockSubscriptionsRepository(ctrl)
			test.mockBehavior(repo, test.authorID)

			s := New(repo)

			_, err := s.GetSubscriptionsByAuthorID(context.Background(), &userProto.UserID{UserId: test.authorID})

			if err != nil {
				assert.Equal(t, test.expectedErr, err.Error())
			}
		})
	}
}

func TestSubscriptionsService_GetSubscriptionsByUserID(t *testing.T) {
	type mockBehaviorGetSubscriptionsByUserID func(r *mock_domain.MockSubscriptionsRepository, userID uint64)

	tests := []struct {
		name         string
		userID       uint64
		mockBehavior mockBehaviorGetSubscriptionsByUserID
		expectedErr  string
	}{
		{
			name:   "OK",
			userID: 1,
			mockBehavior: func(r *mock_domain.MockSubscriptionsRepository, userID uint64) {
				r.EXPECT().GetSubscriptionsByUserID(userID).Return([]models.AuthorSubscription{
					{
						ID:           1,
						AuthorID:     1,
						Img:          "img",
						Tier:         1,
						Title:        "title",
						Text:         "text",
						Price:        1,
						AuthorName:   "authorName",
						AuthorAvatar: "authorAvatar",
					},
				}, nil)
			},
		},
		{
			name:   "Error",
			userID: 1,
			mockBehavior: func(r *mock_domain.MockSubscriptionsRepository, userID uint64) {
				r.EXPECT().GetSubscriptionsByUserID(userID).Return([]models.AuthorSubscription{}, domain.ErrInternal)
			},
			expectedErr: domain.ErrInternal.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockSubscriptionsRepository(ctrl)
			test.mockBehavior(repo, test.userID)

			s := New(repo)

			_, err := s.GetSubscriptionsByUserID(context.Background(), &userProto.UserID{UserId: test.userID})

			if err != nil {
				assert.Equal(t, test.expectedErr, err.Error())
			}
		})
	}
}

func TestSubscriptionsService_UpdateSubscription(t *testing.T) {
	type mockBehaviorUpdateSubscription func(r *mock_domain.MockSubscriptionsRepository, subscription models.AuthorSubscription)

	tests := []struct {
		name         string
		subscription models.AuthorSubscription
		mockBehavior mockBehaviorUpdateSubscription
		expectedErr  string
	}{
		{
			name: "OK",
			subscription: models.AuthorSubscription{
				ID:       1,
				AuthorID: 1,
				Img:      "img",
				Tier:     1,
				Title:    "title",
				Text:     "text",
				Price:    1,
			},
			mockBehavior: func(r *mock_domain.MockSubscriptionsRepository, subscription models.AuthorSubscription) {
				r.EXPECT().UpdateSubscription(subscription).Return(nil)
			},
		},
		{
			name: "Error",
			subscription: models.AuthorSubscription{
				ID:       1,
				AuthorID: 1,
				Img:      "img",
				Tier:     1,
				Title:    "title",
				Text:     "text",
				Price:    1,
			},
			mockBehavior: func(r *mock_domain.MockSubscriptionsRepository, subscription models.AuthorSubscription) {
				r.EXPECT().UpdateSubscription(subscription).Return(domain.ErrInternal)
			},
			expectedErr: domain.ErrInternal.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockSubscriptionsRepository(ctrl)
			test.mockBehavior(repo, test.subscription)

			s := New(repo)

			_, err := s.UpdateSubscription(context.Background(), &protobuf.AuthorSubscription{
				ID:       test.subscription.ID,
				AuthorID: test.subscription.AuthorID,
				Img:      test.subscription.Img,
				Tier:     test.subscription.Tier,
				Title:    test.subscription.Title,
				Text:     test.subscription.Text,
				Price:    test.subscription.Price,
			})

			if err != nil {
				assert.Equal(t, test.expectedErr, err.Error())
			}
		})
	}
}
