package httpSubscriptions

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"

	"github.com/stretchr/testify/assert"

	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
)

func TestHandler_GetSubscriptions(t *testing.T) {
	type mockBehaviorUserGet func(u *mockDomain.MockUsersUseCase, userID uint64)
	type mockBehaviorSubscriptions func(u *mockDomain.MockSubscriptionsUseCase, userID uint64)
	type mockBehaviorImage func(u *mockDomain.MockImageUseCase, bucket, name string)

	tests := []struct {
		name                      string
		userID                    int
		mockBehaviorSubscriptions mockBehaviorSubscriptions
		mockBehaviorImage         mockBehaviorImage
		mockBehaviorUserGet       mockBehaviorUserGet
		expectedResponseBody      string
		expectedErrorMessage      string
	}{
		{
			name:   "OK",
			userID: 1,
			mockBehaviorSubscriptions: func(u *mockDomain.MockSubscriptionsUseCase, userID uint64) {
				u.EXPECT().GetSubscriptionsByUserID(userID).Return([]models.AuthorSubscription{
					{Img: "filename.jpg"},
				}, nil)
			},
			mockBehaviorImage: func(u *mockDomain.MockImageUseCase, bucket, name string) {
				u.EXPECT().GetImage(bucket, name).Return("path/to/img", nil)
			},
			mockBehaviorUserGet: func(u *mockDomain.MockUsersUseCase, userID uint64) {
				u.EXPECT().GetByID(userID).Return(models.User{ID: 12}, nil)
			},
			expectedResponseBody: "[{\"id\":0,\"authorID\":0,\"img\":\"path/to/img\",\"tier\":0,\"title\":\"\",\"text\":\"\",\"price\":0,\"author\":{\"username\":\"\",\"email\":\"\",\"isAuthor\":true,\"about\":\"\",\"countSubscriptions\":0,\"countSubscribers\":0}}]",
		},
		{
			name:   "OK-Empty",
			userID: 1,
			mockBehaviorSubscriptions: func(u *mockDomain.MockSubscriptionsUseCase, userID uint64) {
				u.EXPECT().GetSubscriptionsByUserID(userID).Return([]models.AuthorSubscription{}, nil)
			},
			mockBehaviorImage:    func(u *mockDomain.MockImageUseCase, bucket, name string) {},
			mockBehaviorUserGet:  func(u *mockDomain.MockUsersUseCase, userID uint64) {},
			expectedResponseBody: "{}",
		},
		{
			name:   "ErrInternal-subscriptions",
			userID: 1,
			mockBehaviorSubscriptions: func(u *mockDomain.MockSubscriptionsUseCase, userID uint64) {
				u.EXPECT().GetSubscriptionsByUserID(userID).Return(nil, domain.ErrInternal)
			},
			mockBehaviorImage:    func(u *mockDomain.MockImageUseCase, bucket, name string) {},
			mockBehaviorUserGet:  func(u *mockDomain.MockUsersUseCase, userID uint64) {},
			expectedErrorMessage: "code=500, message=server error, internal=server error",
		},
		{
			name:   "ErrInternal-images",
			userID: 1,
			mockBehaviorSubscriptions: func(u *mockDomain.MockSubscriptionsUseCase, userID uint64) {
				u.EXPECT().GetSubscriptionsByUserID(userID).Return([]models.AuthorSubscription{
					{Img: "filename.jpg"},
				}, nil)
			},
			mockBehaviorImage: func(u *mockDomain.MockImageUseCase, bucket, name string) {
				u.EXPECT().GetImage(bucket, name).Return("", domain.ErrInternal)
			},
			mockBehaviorUserGet:  func(u *mockDomain.MockUsersUseCase, userID uint64) {},
			expectedErrorMessage: "code=500, message=server error, internal=server error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			user := mockDomain.NewMockUsersUseCase(ctrl)
			subscription := mockDomain.NewMockSubscriptionsUseCase(ctrl)
			image := mockDomain.NewMockImageUseCase(ctrl)

			handler := NewHandler(subscription, user, image)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/", nil)

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1")
			c.Set("bucket", "image")
			c.SetParamNames("user_id")
			c.SetParamValues(fmt.Sprint(test.userID))

			test.mockBehaviorSubscriptions(subscription, uint64(test.userID))
			test.mockBehaviorImage(image, "image", "filename.jpg")
			test.mockBehaviorUserGet(user, uint64(test.userID))

			err := handler.GetSubscriptions(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			} else {
				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				assert.Equal(t, test.expectedResponseBody, strings.Trim(string(body), "\n"))
			}
		})
	}
}

func TestHandler_GetAuthorSubscription(t *testing.T) {
	type mockBehaviorSubscriptions func(u *mockDomain.MockSubscriptionsUseCase, subID uint64)
	type mockBehaviorImage func(u *mockDomain.MockImageUseCase, bucket, name string)

	tests := []struct {
		name                      string
		subID                     int
		mockBehaviorSubscriptions mockBehaviorSubscriptions
		mockBehaviorImage         mockBehaviorImage
		expectedResponseBody      string
		expectedErrorMessage      string
	}{
		{
			name:  "OK",
			subID: 1,
			mockBehaviorSubscriptions: func(u *mockDomain.MockSubscriptionsUseCase, subID uint64) {
				u.EXPECT().GetAuthorSubscriptionByID(subID).Return(models.AuthorSubscription{
					Img: "filename.jpg",
				}, nil)
			},
			mockBehaviorImage: func(u *mockDomain.MockImageUseCase, bucket, name string) {
				u.EXPECT().GetImage(bucket, name).Return("path/to/img", nil)
			},
			expectedResponseBody: "{\"id\":0,\"authorID\":0,\"img\":\"path/to/img\",\"tier\":0,\"title\":\"\",\"text\":\"\",\"price\":0,\"author\":{\"username\":\"\",\"email\":\"\",\"isAuthor\":false,\"about\":\"\",\"countSubscriptions\":0,\"countSubscribers\":0}}",
		},
		{
			name:                      "ErrBadRequest",
			subID:                     -1,
			mockBehaviorSubscriptions: func(u *mockDomain.MockSubscriptionsUseCase, subID uint64) {},
			mockBehaviorImage:         func(u *mockDomain.MockImageUseCase, bucket, name string) {},
			expectedErrorMessage:      "code=400, message=bad request, internal=strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
		{
			name:  "ErrNotFound",
			subID: 1,
			mockBehaviorSubscriptions: func(u *mockDomain.MockSubscriptionsUseCase, subID uint64) {
				u.EXPECT().GetAuthorSubscriptionByID(subID).Return(models.AuthorSubscription{}, domain.ErrNotFound)
			},
			mockBehaviorImage:    func(u *mockDomain.MockImageUseCase, bucket, name string) {},
			expectedErrorMessage: "code=404, message=failed to find item, internal=failed to find item",
		},
		{
			name:  "ErrInternal",
			subID: 1,
			mockBehaviorSubscriptions: func(u *mockDomain.MockSubscriptionsUseCase, subID uint64) {
				u.EXPECT().GetAuthorSubscriptionByID(subID).Return(models.AuthorSubscription{
					Img: "filename.jpg",
				}, nil)
			},
			mockBehaviorImage: func(u *mockDomain.MockImageUseCase, bucket, name string) {
				u.EXPECT().GetImage(bucket, name).Return("", domain.ErrInternal)
			},
			expectedErrorMessage: "code=500, message=server error, internal=server error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			user := mockDomain.NewMockUsersUseCase(ctrl)
			subscription := mockDomain.NewMockSubscriptionsUseCase(ctrl)
			image := mockDomain.NewMockImageUseCase(ctrl)

			handler := NewHandler(subscription, user, image)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/", nil)

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1")
			c.Set("bucket", "image")
			c.SetParamNames("id")
			c.SetParamValues(fmt.Sprint(test.subID))

			test.mockBehaviorSubscriptions(subscription, uint64(test.subID))
			test.mockBehaviorImage(image, "image", "filename.jpg")

			err := handler.GetAuthorSubscription(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			} else {
				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				assert.Equal(t, test.expectedResponseBody, strings.Trim(string(body), "\n"))
			}
		})
	}
}

func TestHandler_GetAuthorSubscriptions(t *testing.T) {
}

func TestHandler_CreateAuthorSubscription(t *testing.T) {
}

func TestHandler_UpdateAuthorSubscription(t *testing.T) {
}

func TestHandler_DeleteAuthorSubscription(t *testing.T) {
}
