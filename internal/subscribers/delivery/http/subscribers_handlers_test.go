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
