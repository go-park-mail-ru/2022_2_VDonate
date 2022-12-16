package subscribers

import (
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"

	"github.com/stretchr/testify/assert"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/golang/mock/gomock"
)

func TestUsecase_GetSubscribers(t *testing.T) {
	type mockBehaviourGetSub func(r *mockDomain.MockSubscribersRepository, authorID uint64)
	type mockBehaviourGetUser func(r *mockDomain.MockUsersRepository, userID uint64)

	tests := []struct {
		name                 string
		authorID             int
		userIDs              []uint64
		mockBehaviourGetSub  mockBehaviourGetSub
		mockBehaviourGetUser mockBehaviourGetUser
		responseError        string
	}{
		{
			name:     "OK",
			authorID: 1,
			userIDs:  []uint64{12, 13},
			mockBehaviourGetSub: func(r *mockDomain.MockSubscribersRepository, authorID uint64) {
				r.EXPECT().GetSubscribers(authorID).Return([]uint64{12, 13}, nil)
			},
			mockBehaviourGetUser: func(r *mockDomain.MockUsersRepository, userID uint64) {
				r.EXPECT().GetByID(userID).Return(models.User{ID: userID}, nil)
			},
		},
		{
			name:     "OK-EmptyUsers",
			authorID: 1,
			userIDs:  []uint64{},
			mockBehaviourGetSub: func(r *mockDomain.MockSubscribersRepository, authorID uint64) {
				r.EXPECT().GetSubscribers(authorID).Return([]uint64{}, nil)
			},
			mockBehaviourGetUser: func(r *mockDomain.MockUsersRepository, userID uint64) {},
		},
		{
			name:     "ErrNotFound",
			authorID: 1,
			userIDs:  []uint64{12, 13},
			mockBehaviourGetSub: func(r *mockDomain.MockSubscribersRepository, authorID uint64) {
				r.EXPECT().GetSubscribers(authorID).Return([]uint64{}, domain.ErrNotFound)
			},
			mockBehaviourGetUser: func(r *mockDomain.MockUsersRepository, userID uint64) {},
			responseError:        "failed to find item",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			subMock := mockDomain.NewMockSubscribersRepository(ctrl)
			userMock := mockDomain.NewMockUsersRepository(ctrl)

			test.mockBehaviourGetSub(subMock, uint64(test.authorID))
			for _, id := range test.userIDs {
				test.mockBehaviourGetUser(userMock, id)
			}

			usecase := New(subMock, userMock)
			_, err := usecase.GetSubscribers(uint64(test.authorID))
			if err != nil {
				assert.Equal(t, err.Error(), test.responseError)
			}
		})
	}
}

func TestUsecase_Subscribe(t *testing.T) {
	type mockBehaviourSubscribe func(r *mockDomain.MockSubscribersRepository, s models.Subscription)

	tests := []struct {
		name                   string
		sub                    models.Subscription
		mockBehaviourSubscribe mockBehaviourSubscribe
		responseError          string
	}{
		{
			name: "OK",
			sub: models.Subscription{
				SubscriberID:         1,
				AuthorID:             2,
				AuthorSubscriptionID: 3,
			},
			mockBehaviourSubscribe: func(r *mockDomain.MockSubscribersRepository, s models.Subscription) {
				r.EXPECT().Subscribe(s).Return(nil)
			},
		},
		{
			name: "ErrSubscribe",
			sub: models.Subscription{
				SubscriberID:         1,
				AuthorID:             2,
				AuthorSubscriptionID: 3,
			},
			mockBehaviourSubscribe: func(r *mockDomain.MockSubscribersRepository, s models.Subscription) {
				r.EXPECT().Subscribe(s).Return(domain.ErrInternal)
			},
			responseError: "server error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			subMock := mockDomain.NewMockSubscribersRepository(ctrl)
			userMock := mockDomain.NewMockUsersRepository(ctrl)

			test.mockBehaviourSubscribe(subMock, test.sub)

			usecase := New(subMock, userMock)
			err := usecase.Subscribe(test.sub, test.sub.SubscriberID)
			if err != nil {
				assert.Equal(t, err.Error(), test.responseError)
			}
		})
	}
}

func TestUsecase_Unsubscribe(t *testing.T) {
	type mockBehaviourUnsubscribe func(r *mockDomain.MockSubscribersRepository, uID, aID uint64)

	tests := []struct {
		name                     string
		sub                      models.Subscription
		mockBehaviourUnsubscribe mockBehaviourUnsubscribe
		responseError            string
	}{
		{
			name: "OK",
			sub: models.Subscription{
				SubscriberID:         1,
				AuthorID:             2,
				AuthorSubscriptionID: 3,
			},
			mockBehaviourUnsubscribe: func(r *mockDomain.MockSubscribersRepository, uID, aID uint64) {
				r.EXPECT().Unsubscribe(uID, aID).Return(nil)
			},
		},
		{
			name: "ErrSubscribe",
			sub: models.Subscription{
				SubscriberID:         1,
				AuthorID:             2,
				AuthorSubscriptionID: 3,
			},
			mockBehaviourUnsubscribe: func(r *mockDomain.MockSubscribersRepository, uID, aID uint64) {
				r.EXPECT().Unsubscribe(uID, aID).Return(domain.ErrInternal)
			},
			responseError: "server error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			subMock := mockDomain.NewMockSubscribersRepository(ctrl)
			userMock := mockDomain.NewMockUsersRepository(ctrl)

			test.mockBehaviourUnsubscribe(subMock, test.sub.SubscriberID, test.sub.AuthorID)

			usecase := New(subMock, userMock)
			err := usecase.Unsubscribe(test.sub.SubscriberID, test.sub.AuthorID)
			if err != nil {
				assert.Equal(t, err.Error(), test.responseError)
			}
		})
	}
}

func TestUsecase_IsSubscriber(t *testing.T) {
	type mockBehaviourGetSub func(r *mockDomain.MockSubscribersRepository, userID, authorID uint64)

	tests := []struct {
		name                string
		sub                 models.Subscription
		mockBehaviourGetSub mockBehaviourGetSub
		responseError       string
	}{
		{
			name: "OK",
			sub: models.Subscription{
				SubscriberID:         1,
				AuthorID:             2,
				AuthorSubscriptionID: 3,
			},
			mockBehaviourGetSub: func(r *mockDomain.MockSubscribersRepository, userID, authorID uint64) {
				r.EXPECT().GetSubscribers(authorID).Return([]uint64{
					userID,
				}, nil)
			},
		},
		{
			name: "NotFound",
			sub: models.Subscription{
				SubscriberID:         1,
				AuthorID:             2,
				AuthorSubscriptionID: 3,
			},
			mockBehaviourGetSub: func(r *mockDomain.MockSubscribersRepository, userID, authorID uint64) {
				r.EXPECT().GetSubscribers(authorID).Return(nil, domain.ErrNotFound)
			},
			responseError: "failed to find item",
		},
		{
			name: "NotSubscriber",
			sub: models.Subscription{
				SubscriberID:         1,
				AuthorID:             2,
				AuthorSubscriptionID: 3,
			},
			mockBehaviourGetSub: func(r *mockDomain.MockSubscribersRepository, userID, authorID uint64) {
				r.EXPECT().GetSubscribers(authorID).Return([]uint64{
					userID + 1,
				}, nil)
			},
			responseError: "user is not a subscriber",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			subRepo := mockDomain.NewMockSubscribersRepository(ctrl)
			userRepo := mockDomain.NewMockUsersRepository(ctrl)

			test.mockBehaviourGetSub(subRepo, test.sub.SubscriberID, test.sub.AuthorID)

			usecase := New(subRepo, userRepo)
			_, err := usecase.IsSubscriber(test.sub.SubscriberID, test.sub.AuthorID)
			if err != nil {
				assert.Equal(t, err.Error(), test.responseError)
			}
		})
	}
}
