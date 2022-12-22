package httpSubscribers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ztrue/tracerr"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"

	mock_domain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
)

func TestHandler_CreateSubscriber(t *testing.T) {
	type mockSubscribe func(u *mock_domain.MockSubscribersUseCase, s models.Subscription, userID uint64, a models.AuthorSubscription)
	type mockUser func(u *mock_domain.MockUsersUseCase, sessionID string, userID uint64)
	type mockAuthor func(u *mock_domain.MockSubscriptionsUseCase, authorID uint64)

	tests := []struct {
		name          string
		sub           models.Subscription
		mockSubscribe mockSubscribe
		mockUser      mockUser
		mockAuthor    mockAuthor
		authorSub     models.AuthorSubscription
		inputBody     string
		userID        uint64
		cookie        *http.Cookie
		responseError string
		response      bool
	}{
		{
			name:      "OK",
			inputBody: `{"authorID":1,"authorSubscriptionID":5}`,
			userID:    2,
			cookie: &http.Cookie{
				Name:  "session_id",
				Value: "some cookie",
			},
			sub: models.Subscription{
				AuthorID:             1,
				AuthorSubscriptionID: 5,
			},
			authorSub: models.AuthorSubscription{
				AuthorID: 1,
			},
			mockUser: func(u *mock_domain.MockUsersUseCase, sessionID string, userID uint64) {
				u.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: userID,
				}, nil)
			},
			mockAuthor: func(u *mock_domain.MockSubscriptionsUseCase, authorID uint64) {
				u.EXPECT().GetAuthorSubscriptionByID(authorID).Return(models.AuthorSubscription{
					AuthorID: 1,
				}, nil)
			},
			mockSubscribe: func(u *mock_domain.MockSubscribersUseCase, s models.Subscription, userID uint64, a models.AuthorSubscription) {
				u.EXPECT().Subscribe(s, userID, a).Return(true, nil)
			},
			response: true,
		},
		{
			name:      "ErrorBind",
			inputBody: `{"authorID":-1,"authorSubscriptionID":5}`,
			userID:    2,
			cookie: &http.Cookie{
				Name:  "session_id",
				Value: "some cookie",
			},
			mockUser: func(u *mock_domain.MockUsersUseCase, sessionID string, userID uint64) {
				u.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: userID,
				}, nil)
			},
			mockAuthor: func(u *mock_domain.MockSubscriptionsUseCase, authorID uint64) {},
			mockSubscribe: func(u *mock_domain.MockSubscribersUseCase, s models.Subscription, userID uint64, a models.AuthorSubscription) {
			},
			responseError: "code=400, message=bad request, internal=code=400, message=Unmarshal type error: expected=uint64, got=number -1, field=authorID, offset=14, internal=json: cannot unmarshal number -1 into Go struct field Subscription.authorID of type uint64",
		},
		{
			name:      "ErrorBadRequest",
			inputBody: `{"authorSubscriptionID":5}`,
			userID:    2,
			cookie: &http.Cookie{
				Name:  "session_id",
				Value: "some cookie",
			},
			sub: models.Subscription{
				AuthorSubscriptionID: 5,
			},
			mockUser: func(u *mock_domain.MockUsersUseCase, sessionID string, userID uint64) {
				u.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: userID,
				}, nil)
			},
			mockAuthor: func(u *mock_domain.MockSubscriptionsUseCase, authorID uint64) {
				u.EXPECT().GetAuthorSubscriptionByID(authorID).Return(models.AuthorSubscription{}, nil)
			},
			mockSubscribe: func(u *mock_domain.MockSubscribersUseCase, s models.Subscription, userID uint64, a models.AuthorSubscription) {
				u.EXPECT().Subscribe(s, userID, a).Return(false, domain.ErrBadRequest)
			},
			responseError: "code=400, message=bad request, internal=bad request",
		},
		{
			name:      "ErrorCreate",
			inputBody: `{"authorID":1,"authorSubscriptionID":5}`,
			userID:    2,
			cookie: &http.Cookie{
				Name:  "session_id",
				Value: "some cookie",
			},
			sub: models.Subscription{
				AuthorID:             1,
				AuthorSubscriptionID: 5,
			},
			mockUser: func(u *mock_domain.MockUsersUseCase, sessionID string, userID uint64) {
				u.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: userID,
				}, nil)
			},
			mockAuthor: func(u *mock_domain.MockSubscriptionsUseCase, authorID uint64) {
				u.EXPECT().GetAuthorSubscriptionByID(authorID).Return(models.AuthorSubscription{}, domain.ErrBadRequest)
			},
			mockSubscribe: func(u *mock_domain.MockSubscribersUseCase, s models.Subscription, userID uint64, a models.AuthorSubscription) {
			},
			responseError: "code=400, message=bad request, internal=bad request",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			subUseCase := mock_domain.NewMockSubscribersUseCase(ctrl)
			userUseCase := mock_domain.NewMockUsersUseCase(ctrl)
			sUseCase := mock_domain.NewMockSubscriptionsUseCase(ctrl)

			test.mockUser(userUseCase, test.cookie.Value, test.userID)
			test.mockSubscribe(subUseCase, test.sub, test.userID, test.authorSub)
			test.mockAuthor(sUseCase, test.sub.AuthorSubscriptionID)

			handler := NewHandler(subUseCase, userUseCase, sUseCase)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/", bytes.NewBufferString(test.inputBody))
			req.Header.Add("Content-Type", "application/json")
			req.AddCookie(test.cookie)

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/auth")

			err := handler.CreateSubscriber(c)
			if err != nil {
				assert.Equal(t, test.responseError, err.Error())
			} else {
				assert.Equal(t, true, test.response)
			}
		})
	}
}

func TestHandler_GetSubscribers(t *testing.T) {
	type mockSubscribe func(u *mock_domain.MockSubscribersUseCase, id uint64)

	tests := []struct {
		name          string
		authorID      int
		mockSubscribe mockSubscribe
		responseError string
		response      string
	}{
		{
			name:     "OK",
			authorID: 1,
			mockSubscribe: func(u *mock_domain.MockSubscribersUseCase, id uint64) {
				u.EXPECT().GetSubscribers(id).Return([]models.User{
					{ID: 15, Email: "test@test.ru", Password: "*****", IsAuthor: true},
				}, nil)
			},
			response: "[{\"id\":15,\"username\":\"\",\"email\":\"test@test.ru\",\"avatar\":\"\",\"password\":\"*****\",\"isAuthor\":true,\"balance\":0,\"about\":\"\",\"countSubscriptions\":0,\"countSubscribers\":0,\"countPosts\":0}]",
		},
		{
			name:     "OK-Empty",
			authorID: 1,
			mockSubscribe: func(u *mock_domain.MockSubscribersUseCase, id uint64) {
				u.EXPECT().GetSubscribers(id).Return([]models.User{}, nil)
			},
			response: "[]",
		},
		{
			name:     "ErrorBadRequest-PathParam",
			authorID: -1,
			mockSubscribe: func(u *mock_domain.MockSubscribersUseCase, id uint64) {
			},
			responseError: "code=400, message=bad request, internal=strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
		{
			name:     "ErrorNotFound",
			authorID: 1,
			mockSubscribe: func(u *mock_domain.MockSubscribersUseCase, id uint64) {
				u.EXPECT().GetSubscribers(id).Return(nil, domain.ErrNotFound)
			},
			responseError: "code=404, message=failed to find item, internal=failed to find item",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			subUseCase := mock_domain.NewMockSubscribersUseCase(ctrl)
			userUseCase := mock_domain.NewMockUsersUseCase(ctrl)
			sUseCase := mock_domain.NewMockSubscriptionsUseCase(ctrl)

			test.mockSubscribe(subUseCase, uint64(test.authorID))

			handler := NewHandler(subUseCase, userUseCase, sUseCase)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/", nil)
			req.Header.Add("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/auth")
			c.SetParamNames("author_id")
			c.SetParamValues(fmt.Sprint(test.authorID))

			err := handler.GetSubscribers(c)
			if err != nil {
				assert.Equal(t, err.Error(), test.responseError)
			} else {
				body, _ := io.ReadAll(rec.Body)
				assert.Equal(t, test.response, strings.Trim(string(body), "\n"))
			}
		})
	}
}

func TestHandler_DeleteSubscriber(t *testing.T) {
	type mockSubscribe func(u *mock_domain.MockSubscribersUseCase, s models.Subscription, userID uint64)
	type mockUser func(u *mock_domain.MockUsersUseCase, sessionID string, userID uint64)

	tests := []struct {
		name          string
		sub           models.Subscription
		mockSubscribe mockSubscribe
		mockUser      mockUser
		userID        uint64
		cookie        *http.Cookie
		inputBody     string
		responseError string
		response      bool
	}{
		{
			name:      "OK",
			inputBody: `{"authorID":1,"authorSubscriptionID":5}`,
			userID:    2,
			cookie: &http.Cookie{
				Name:  "session_id",
				Value: "some cookie",
			},
			sub: models.Subscription{
				AuthorID:             1,
				AuthorSubscriptionID: 5,
			},
			mockUser: func(u *mock_domain.MockUsersUseCase, sessionID string, userID uint64) {
				u.EXPECT().GetBySessionID(sessionID).Return(models.User{ID: userID}, nil)
			},
			mockSubscribe: func(u *mock_domain.MockSubscribersUseCase, s models.Subscription, userID uint64) {
				u.EXPECT().Unsubscribe(userID, s.AuthorID).Return(nil)
			},
			response: true,
		},
		{
			name:      "ErrorBind",
			inputBody: `{"authorID":-1,"authorSubscriptionID":5}`,
			userID:    2,
			cookie: &http.Cookie{
				Name:  "session_id",
				Value: "some cookie",
			},
			mockUser: func(u *mock_domain.MockUsersUseCase, sessionID string, userID uint64) {
				u.EXPECT().GetBySessionID(sessionID).Return(models.User{ID: userID}, nil)
			},
			mockSubscribe: func(u *mock_domain.MockSubscribersUseCase, s models.Subscription, userID uint64) {},
			responseError: "code=400, message=bad request, internal=code=400, message=Unmarshal type error: expected=uint64, got=number -1, field=authorID, offset=14, internal=json: cannot unmarshal number -1 into Go struct field Subscription.authorID of type uint64",
		},
		{
			name:      "ErrorBadRequest",
			inputBody: `{"authorSubscriptionID":5}`,
			userID:    2,
			cookie: &http.Cookie{
				Name:  "session_id",
				Value: "some cookie",
			},
			sub: models.Subscription{
				AuthorSubscriptionID: 5,
			},
			mockUser: func(u *mock_domain.MockUsersUseCase, sessionID string, userID uint64) {
				u.EXPECT().GetBySessionID(sessionID).Return(models.User{ID: userID}, nil)
			},
			mockSubscribe: func(u *mock_domain.MockSubscribersUseCase, s models.Subscription, userID uint64) {
				u.EXPECT().Unsubscribe(userID, s.AuthorID).Return(domain.ErrBadRequest)
			},
			responseError: "code=400, message=bad request, internal=bad request",
		},
		{
			name:      "ErrorDelete",
			inputBody: `{"authorID":1,"authorSubscriptionID":5}`,
			userID:    2,
			cookie: &http.Cookie{
				Name:  "session_id",
				Value: "some cookie",
			},
			sub: models.Subscription{
				AuthorID:             1,
				AuthorSubscriptionID: 5,
			},
			mockUser: func(u *mock_domain.MockUsersUseCase, sessionID string, userID uint64) {
				u.EXPECT().GetBySessionID(sessionID).Return(models.User{ID: userID}, nil)
			},
			mockSubscribe: func(u *mock_domain.MockSubscribersUseCase, s models.Subscription, userID uint64) {
				u.EXPECT().Unsubscribe(userID, s.AuthorID).Return(domain.ErrDelete)
			},
			responseError: "code=500, message=failed to delete item, internal=failed to delete item",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			subUseCase := mock_domain.NewMockSubscribersUseCase(ctrl)
			userUseCase := mock_domain.NewMockUsersUseCase(ctrl)
			sUseCase := mock_domain.NewMockSubscriptionsUseCase(ctrl)

			test.mockSubscribe(subUseCase, test.sub, test.userID)
			test.mockUser(userUseCase, test.cookie.Value, test.userID)

			handler := NewHandler(subUseCase, userUseCase, sUseCase)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/", bytes.NewBufferString(test.inputBody))
			req.Header.Add("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/auth")
			req.AddCookie(test.cookie)

			err := handler.DeleteSubscriber(c)
			if err != nil {
				assert.Equal(t, test.responseError, err.Error())
			} else {
				assert.Equal(t, true, test.response)
			}
		})
	}
}

func TestSubscribersHandler_Withdraw(t *testing.T) {
	type mockSession func(u *mock_domain.MockUsersUseCase, sessionID string)
	type mockWithdraw func(u *mock_domain.MockSubscribersUseCase, userID uint64, phone, card string)

	tests := []struct {
		name          string
		inputBody     string
		userID        uint64
		withDraw      models.Withdraw
		cookie        *http.Cookie
		sub           models.Subscription
		mockSession   mockSession
		mockWithdraw  mockWithdraw
		response      bool
		responseError string
	}{
		{
			name:   "OK",
			userID: 2,
			cookie: &http.Cookie{
				Name:  "session_id",
				Value: "some cookie",
			},
			inputBody: `{"userID":2,"phone":"123","card":"123"}`,
			withDraw: models.Withdraw{
				UserID: 2,
				Phone:  "123",
				Card:   "123",
			},
			sub: models.Subscription{
				AuthorID:             1,
				AuthorSubscriptionID: 5,
			},
			mockSession: func(u *mock_domain.MockUsersUseCase, sessionID string) {
				u.EXPECT().GetBySessionID(sessionID).Return(models.User{ID: 2}, nil)
			},
			mockWithdraw: func(u *mock_domain.MockSubscribersUseCase, userID uint64, phone, card string) {
				u.EXPECT().Withdraw(userID, phone, card).Return(models.WithdrawInfo{
					Id: "1",
				}, nil)
			},
			response: true,
		},
		{
			name:   "ErrInternal",
			userID: 2,
			cookie: &http.Cookie{
				Name:  "session_id",
				Value: "some cookie",
			},
			inputBody: `{"userID":2,"phone":"123","card":"123"}`,
			withDraw: models.Withdraw{
				UserID: 2,
				Phone:  "123",
				Card:   "123",
			},
			sub: models.Subscription{
				AuthorID:             1,
				AuthorSubscriptionID: 5,
			},
			mockSession: func(u *mock_domain.MockUsersUseCase, sessionID string) {
				u.EXPECT().GetBySessionID(sessionID).Return(models.User{ID: 2}, nil)
			},
			mockWithdraw: func(u *mock_domain.MockSubscribersUseCase, userID uint64, phone, card string) {
				u.EXPECT().Withdraw(userID, phone, card).Return(models.WithdrawInfo{
					Id: "1",
				}, tracerr.Errorf("failed to withdraw"))
			},
			responseError: "code=500, message=server error, internal=failed to withdraw",
		},
		{
			name:   "ErrBadRequest",
			userID: 2,
			cookie: &http.Cookie{
				Name:  "session_id",
				Value: "some cookie",
			},
			inputBody: `{"userID":1,"phone":"123","card":"123"}`,
			withDraw: models.Withdraw{
				UserID: 1,
				Phone:  "123",
				Card:   "123",
			},
			sub: models.Subscription{
				AuthorID:             1,
				AuthorSubscriptionID: 5,
			},
			mockSession: func(u *mock_domain.MockUsersUseCase, sessionID string) {
				u.EXPECT().GetBySessionID(sessionID).Return(models.User{ID: 2}, nil)
			},
			mockWithdraw:  func(u *mock_domain.MockSubscribersUseCase, userID uint64, phone, card string) {},
			responseError: "code=400, message=bad request",
		},
		{
			name:      "ErrorBind",
			inputBody: `{"authorID":-1,"authorSubscriptionID":5}`,
			userID:    2,
			cookie: &http.Cookie{
				Name: "session_id",
			},
			sub: models.Subscription{
				AuthorID:             1,
				AuthorSubscriptionID: 5,
			},
			mockSession: func(u *mock_domain.MockUsersUseCase, sessionID string) {
				u.EXPECT().GetBySessionID(sessionID).Return(models.User{ID: 2}, nil)
			},
			mockWithdraw:  func(u *mock_domain.MockSubscribersUseCase, userID uint64, phone, card string) {},
			responseError: "code=400, message=bad request",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			subUseCase := mock_domain.NewMockSubscribersUseCase(ctrl)
			userUseCase := mock_domain.NewMockUsersUseCase(ctrl)
			sUseCase := mock_domain.NewMockSubscriptionsUseCase(ctrl)

			test.mockSession(userUseCase, test.cookie.Value)
			test.mockWithdraw(subUseCase, test.withDraw.UserID, test.withDraw.Phone, test.withDraw.Card)

			handler := NewHandler(subUseCase, userUseCase, sUseCase)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "https://127.0.0.1/api/v1/", bytes.NewBufferString(test.inputBody))
			req.Header.Add("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/withdraw")
			req.AddCookie(test.cookie)

			err := handler.Withdraw(c)
			if err != nil {
				assert.Equal(t, test.responseError, err.Error())
			} else {
				assert.Equal(t, true, test.response)
			}
		})
	}
}
