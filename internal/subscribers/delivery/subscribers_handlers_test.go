package httpsubscribers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/interface"

	mock_domain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
)

func TestHandler_CreateSubscriber(t *testing.T) {
	type mockSubscribe func(u *mock_domain.MockSubscribersUseCase, s models.Subscription, userID uint64)
	type mockUser func(u *mock_domain.MockUsersUseCase, sessionID string, userID uint64)

	tests := []struct {
		name          string
		sub           models.Subscription
		mockSubscribe mockSubscribe
		mockUser      mockUser
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
			mockUser: func(u *mock_domain.MockUsersUseCase, sessionID string, userID uint64) {
				u.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: userID,
				}, nil)
			},
			mockSubscribe: func(u *mock_domain.MockSubscribersUseCase, s models.Subscription, userID uint64) {
				u.EXPECT().Subscribe(s, userID).Return(nil)
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
			mockSubscribe: func(u *mock_domain.MockSubscribersUseCase, s models.Subscription, userID uint64) {},
			responseError: "code=400, " +
				"message=bad request, " +
				"internal=code=400, " +
				"message=Unmarshal type error: expected=uint64, got=number -1, field=authorID, offset=14, internal=json: cannot unmarshal number -1 into Go struct field Subscription.authorID of type uint64",
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
			mockSubscribe: func(u *mock_domain.MockSubscribersUseCase, s models.Subscription, userID uint64) {
				u.EXPECT().Subscribe(s, userID).Return(domain.ErrBadRequest)
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
			mockSubscribe: func(u *mock_domain.MockSubscribersUseCase, s models.Subscription, userID uint64) {
				u.EXPECT().Subscribe(s, userID).Return(domain.ErrCreate)
			},
			responseError: "code=500, message=failed to create item, internal=failed to create item",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			subUseCase := mock_domain.NewMockSubscribersUseCase(ctrl)
			userUseCase := mock_domain.NewMockUsersUseCase(ctrl)

			test.mockUser(userUseCase, test.cookie.Value, test.userID)
			test.mockSubscribe(subUseCase, test.sub, test.userID)

			handler := NewHandler(subUseCase, userUseCase)

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
					{ID: 15, Email: "test@test.ru", Password: "*****", IsAuthor: false},
				}, nil)
			},
			response: "[{\"id\":15,\"username\":\"\",\"email\":\"test@test.ru\",\"avatar\":\"\",\"password\":\"*****\",\"isAuthor\":false,\"about\":\"\",\"countSubscriptions\":0,\"countSubscribers\":0}]",
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

			test.mockSubscribe(subUseCase, uint64(test.authorID))

			handler := NewHandler(subUseCase, userUseCase)

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
			responseError: "code=400, " +
				"message=bad request, " +
				"internal=code=400, " +
				"message=Unmarshal type error: expected=uint64, got=number -1, field=authorID, offset=14, internal=json: cannot unmarshal number -1 into Go struct field Subscription.authorID of type uint64",
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

			test.mockSubscribe(subUseCase, test.sub, test.userID)
			test.mockUser(userUseCase, test.cookie.Value, test.userID)

			handler := NewHandler(subUseCase, userUseCase)

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
