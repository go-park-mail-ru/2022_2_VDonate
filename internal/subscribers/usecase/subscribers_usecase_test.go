package subscribers

import (
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsecase_GetSubscribers(t *testing.T) {
	type mockBehaviourGetSub func(r *mockDomain.MockSubscribersMicroservice, authorID uint64)
	type mockBehaviourGetUser func(r *mockDomain.MockUsersMicroservice, userID uint64)

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
			mockBehaviourGetSub: func(r *mockDomain.MockSubscribersMicroservice, authorID uint64) {
				r.EXPECT().GetSubscribers(authorID).Return([]uint64{12, 13}, nil)
			},
			mockBehaviourGetUser: func(r *mockDomain.MockUsersMicroservice, userID uint64) {
				r.EXPECT().GetByID(userID).Return(models.User{ID: userID}, nil)
			},
		},
		{
			name:     "OK-EmptyUsers",
			authorID: 1,
			userIDs:  []uint64{},
			mockBehaviourGetSub: func(r *mockDomain.MockSubscribersMicroservice, authorID uint64) {
				r.EXPECT().GetSubscribers(authorID).Return([]uint64{}, nil)
			},
			mockBehaviourGetUser: func(r *mockDomain.MockUsersMicroservice, userID uint64) {},
		},
		{
			name:     "ErrNotFound",
			authorID: 1,
			userIDs:  []uint64{12, 13},
			mockBehaviourGetSub: func(r *mockDomain.MockSubscribersMicroservice, authorID uint64) {
				r.EXPECT().GetSubscribers(authorID).Return([]uint64{}, domain.ErrNotFound)
			},
			mockBehaviourGetUser: func(r *mockDomain.MockUsersMicroservice, userID uint64) {},
			responseError:        "failed to find item",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			subMock := mockDomain.NewMockSubscribersMicroservice(ctrl)
			userMock := mockDomain.NewMockUsersMicroservice(ctrl)

			test.mockBehaviourGetSub(subMock, uint64(test.authorID))
			for _, id := range test.userIDs {
				test.mockBehaviourGetUser(userMock, id)
			}

			usecase := New(subMock, userMock, "123")
			_, err := usecase.GetSubscribers(uint64(test.authorID))
			if err != nil {
				assert.Equal(t, err.Error(), test.responseError)
			}
		})
	}
}

func TestUsecase_Unsubscribe(t *testing.T) {
	type mockBehaviourUnsubscribe func(r *mockDomain.MockSubscribersMicroservice, uID, aID uint64)

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
			mockBehaviourUnsubscribe: func(r *mockDomain.MockSubscribersMicroservice, uID, aID uint64) {
				r.EXPECT().Unsubscribe(models.Subscription{
					AuthorID:     aID,
					SubscriberID: uID,
				}).Return(nil)
			},
		},
		{
			name: "ErrSubscribe",
			sub: models.Subscription{
				SubscriberID:         1,
				AuthorID:             2,
				AuthorSubscriptionID: 3,
			},
			mockBehaviourUnsubscribe: func(r *mockDomain.MockSubscribersMicroservice, uID, aID uint64) {
				r.EXPECT().Unsubscribe(models.Subscription{
					AuthorID:     aID,
					SubscriberID: uID,
				}).Return(domain.ErrInternal)
			},
			responseError: "server error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			subMock := mockDomain.NewMockSubscribersMicroservice(ctrl)
			userMock := mockDomain.NewMockUsersMicroservice(ctrl)

			test.mockBehaviourUnsubscribe(subMock, test.sub.SubscriberID, test.sub.AuthorID)

			usecase := New(subMock, userMock, "123")
			err := usecase.Unsubscribe(test.sub.SubscriberID, test.sub.AuthorID)
			if err != nil {
				assert.Equal(t, err.Error(), test.responseError)
			}
		})
	}
}

func TestUsecase_IsSubscriber(t *testing.T) {
	type mockBehaviourGetSub func(r *mockDomain.MockSubscribersMicroservice, userID, authorID uint64)

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
			mockBehaviourGetSub: func(r *mockDomain.MockSubscribersMicroservice, userID, authorID uint64) {
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
			mockBehaviourGetSub: func(r *mockDomain.MockSubscribersMicroservice, userID, authorID uint64) {
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
			mockBehaviourGetSub: func(r *mockDomain.MockSubscribersMicroservice, userID, authorID uint64) {
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

			subRepo := mockDomain.NewMockSubscribersMicroservice(ctrl)
			userRepo := mockDomain.NewMockUsersMicroservice(ctrl)

			test.mockBehaviourGetSub(subRepo, test.sub.SubscriberID, test.sub.AuthorID)

			usecase := New(subRepo, userRepo, "123")
			_, err := usecase.IsSubscriber(test.sub.SubscriberID, test.sub.AuthorID)
			if err != nil {
				assert.Equal(t, err.Error(), test.responseError)
			}
		})
	}
}

func TestUsecase_CardValidation(t *testing.T) {
	tests := []struct {
		name          string
		card          string
		expected      models.WithdrawValidation
		responseError string
	}{
		{
			name:          "Fail",
			card:          "1234567890123456",
			responseError: domain.ErrCreatePayment.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			subRepo := mockDomain.NewMockSubscribersMicroservice(ctrl)
			userRepo := mockDomain.NewMockUsersMicroservice(ctrl)

			usecase := New(subRepo, userRepo, "123")
			_, err := usecase.CardValidation(test.card)
			if err != nil {
				assert.Equal(t, err.Error(), test.responseError)
			}
		})
	}
}

func TestUsecase_WithdrawCard(t *testing.T) {
	type mockBehaviorUserGetByID func(r *mockDomain.MockUsersMicroservice, userID uint64)

	tests := []struct {
		name                    string
		userID                  uint64
		card                    string
		provider                string
		mockBehaviorUserGetByID mockBehaviorUserGetByID
		expected                models.WithdrawInfo
		responseError           string
	}{
		{
			name:     "ErrCreatePayment",
			card:     "1234567890123456",
			provider: "123",
			userID:   1,
			mockBehaviorUserGetByID: func(r *mockDomain.MockUsersMicroservice, userID uint64) {
				r.EXPECT().GetByID(userID).Return(models.User{
					ID:       userID,
					Username: "username",
					Email:    "email",
					Balance:  250,
				}, nil)
			},
			responseError: domain.ErrCreatePayment.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			subRepo := mockDomain.NewMockSubscribersMicroservice(ctrl)
			userRepo := mockDomain.NewMockUsersMicroservice(ctrl)

			test.mockBehaviorUserGetByID(userRepo, test.userID)

			usecase := New(subRepo, userRepo, "123")
			_, err := usecase.WithdrawCard(test.userID, test.card, test.provider)
			if err != nil {
				assert.Equal(t, test.responseError, err.Error())
			}
		})
	}
}

func TestUsecase_WithdrawQiwi(t *testing.T) {
	type mockBehaviorUserGetByID func(r *mockDomain.MockUsersMicroservice, userID uint64)

	tests := []struct {
		name                    string
		userID                  uint64
		phone                   string
		mockBehaviorUserGetByID mockBehaviorUserGetByID
		expected                models.WithdrawInfo
		responseError           string
	}{
		{
			name:   "ErrResponse",
			phone:  "1234567890123456",
			userID: 1,
			mockBehaviorUserGetByID: func(r *mockDomain.MockUsersMicroservice, userID uint64) {
				r.EXPECT().GetByID(userID).Return(models.User{
					ID:       userID,
					Username: "username",
					Email:    "email",
				}, nil)
			},
			responseError: "unexpected end of JSON input",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			subRepo := mockDomain.NewMockSubscribersMicroservice(ctrl)
			userRepo := mockDomain.NewMockUsersMicroservice(ctrl)

			test.mockBehaviorUserGetByID(userRepo, test.userID)

			usecase := New(subRepo, userRepo, "123")
			_, err := usecase.WithdrawQiwi(test.userID, test.phone)
			if err != nil {
				assert.Equal(t, test.responseError, err.Error())
			}
		})
	}
}

func TestUsecase_Withdraw(t *testing.T) {
	type mockBehaviorUserGetByID func(r *mockDomain.MockUsersMicroservice, userID uint64)

	tests := []struct {
		name                    string
		userID                  uint64
		phone                   string
		card                    string
		mockBehaviorUserGetByID mockBehaviorUserGetByID
		expected                models.WithdrawInfo
		responseError           string
	}{
		{
			name:   "ErrResponse-Card",
			card:   "1234567890123456",
			userID: 1,
			mockBehaviorUserGetByID: func(r *mockDomain.MockUsersMicroservice, userID uint64) {
			},
			responseError: domain.ErrCreatePayment.Error(),
		},
		{
			name:   "ErrResponse-Qiwi",
			phone:  "1234567890123456",
			card:   "",
			userID: 1,
			mockBehaviorUserGetByID: func(r *mockDomain.MockUsersMicroservice, userID uint64) {
				r.EXPECT().GetByID(userID).Return(models.User{
					ID:       userID,
					Username: "username",
					Email:    "email",
				}, nil)
			},
			responseError: "unexpected end of JSON input",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			subRepo := mockDomain.NewMockSubscribersMicroservice(ctrl)
			userRepo := mockDomain.NewMockUsersMicroservice(ctrl)

			test.mockBehaviorUserGetByID(userRepo, test.userID)

			usecase := New(subRepo, userRepo, "123")
			_, err := usecase.Withdraw(test.userID, test.phone, test.card)
			if err != nil {
				assert.Equal(t, test.responseError, err.Error())
			}
		})
	}
}
