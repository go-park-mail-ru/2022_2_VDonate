package httpDonates

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_CreateDonate(t *testing.T) {
	type mockBehaviourSendDonate func(s mockDomain.MockDonatesUseCase, userID, authorID, price uint64)
	type mockBehaviourCookie func(s mockDomain.MockUsersUseCase, sessionID string)

	tests := []struct {
		name                    string
		reqBody                 string
		cookie                  string
		userID                  uint64
		authorID                uint64
		price                   uint64
		mockBehaviourCookie     mockBehaviourCookie
		mockBehaviourSendDonate mockBehaviourSendDonate
		expectedResponse        string
		expectedErrorMessage    string
	}{
		{
			name:     "OK",
			reqBody:  `{"userId":22,"authorId":12,"price":3000}`,
			cookie:   "JSAoPdaAsdasjdJNPdapoSAjdasakZcs",
			userID:   22,
			authorID: 12,
			price:    3000,
			mockBehaviourCookie: func(s mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: 22,
				}, nil)
			},
			mockBehaviourSendDonate: func(s mockDomain.MockDonatesUseCase, userID, authorID, price uint64) {
				s.EXPECT().SendDonate(userID, authorID, price).Return(models.Donate{
					UserID:   userID,
					AuthorID: authorID,
					Price:    price,
				}, nil)
			},
			expectedResponse: `{"id":0,"userId":22,"authorId":12,"price":3000}`,
		},
		{
			name: "ErrNoSession-1",
			mockBehaviourCookie: func(s mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{}, errors.New("not valid session id"))
			},
			mockBehaviourSendDonate: func(s mockDomain.MockDonatesUseCase, userID, authorID, price uint64) {},
			expectedErrorMessage:    `code=401, message=no existing session, internal=not valid session id`,
		},
		{
			name:                    "ErrNoSession-2",
			mockBehaviourCookie:     func(s mockDomain.MockUsersUseCase, sessionID string) {},
			mockBehaviourSendDonate: func(s mockDomain.MockDonatesUseCase, userID, authorID, price uint64) {},
			expectedErrorMessage:    `code=401, message=no existing session, internal=http: named cookie not present`,
		},
		{
			name:    "ErrBadRequest",
			reqBody: `ksda[k[askd[aksd[a`,
			cookie:  "JSAoPdaAsdasjdJNPdapoSAjdasakZcs",
			userID:  22,
			mockBehaviourCookie: func(s mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: 22,
				}, nil)
			},
			mockBehaviourSendDonate: func(s mockDomain.MockDonatesUseCase, userID, authorID, price uint64) {},
			expectedErrorMessage:    `code=400, message=bad request, internal=code=400, message=Syntax error: offset=1, error=invalid character 'k' looking for beginning of value, internal=invalid character 'k' looking for beginning of value`,
		},
		{
			name:     "ErrCreate",
			reqBody:  `{"userId":22,"authorId":12,"price":3000}`,
			cookie:   "JSAoPdaAsdasjdJNPdapoSAjdasakZcs",
			userID:   22,
			authorID: 12,
			price:    3000,
			mockBehaviourCookie: func(s mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: 22,
				}, nil)
			},
			mockBehaviourSendDonate: func(s mockDomain.MockDonatesUseCase, userID, authorID, price uint64) {
				s.EXPECT().SendDonate(userID, authorID, price).Return(models.Donate{}, errors.New("create error"))
			},
			expectedErrorMessage: `code=500, message=failed to create item, internal=create error`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			donate := mockDomain.NewMockDonatesUseCase(ctrl)
			users := mockDomain.NewMockUsersUseCase(ctrl)

			test.mockBehaviourCookie(*users, test.cookie)
			test.mockBehaviourSendDonate(*donate, test.userID, test.authorID, test.price)

			handler := NewHandler(donate, users)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "https://127.0.0.1/api/v1/donates", bytes.NewBufferString(test.reqBody))
			req.Header.Set("Content-Type", "application/json")
			if test.name != "ErrNoSession-2" {
				req.Header.Add("Cookie", "session_id="+test.cookie)
			}
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/donates")

			err := handler.CreateDonate(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}

			body, err := io.ReadAll(rec.Body)
			require.NoError(t, err)

			assert.Equal(t, test.expectedResponse, strings.Trim(string(body), "\n"))
		})
	}
}

func TestHandler_GetDonate(t *testing.T) {
	type mockBehaviourGetDonate func(s mockDomain.MockDonatesUseCase, ID uint64)

	tests := []struct {
		name                   string
		id                     int64
		mockBehaviourGetDonate mockBehaviourGetDonate
		expectedResponse       string
		expectedErrorMessage   string
	}{
		{
			name: "OK",
			id:   3,
			mockBehaviourGetDonate: func(s mockDomain.MockDonatesUseCase, ID uint64) {
				s.EXPECT().GetDonateByID(ID).Return(models.Donate{
					ID:       ID,
					UserID:   12,
					AuthorID: 21,
					Price:    3000,
				}, nil)
			},
			expectedResponse: `{"id":3,"userId":12,"authorId":21,"price":3000}`,
		},
		{
			name: "ErrNotFound",
			id:   3,
			mockBehaviourGetDonate: func(s mockDomain.MockDonatesUseCase, ID uint64) {
				s.EXPECT().GetDonateByID(ID).Return(models.Donate{}, errors.New("donate not found"))
			},
			expectedErrorMessage: `code=404, message=failed to find item, internal=donate not found`,
		},
		{
			name:                   "ErrBadRequest",
			id:                     -3,
			mockBehaviourGetDonate: func(s mockDomain.MockDonatesUseCase, ID uint64) {},
			expectedErrorMessage:   `code=400, message=bad request, internal=strconv.ParseUint: parsing "-3": invalid syntax`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			donate := mockDomain.NewMockDonatesUseCase(ctrl)
			users := mockDomain.NewMockUsersUseCase(ctrl)

			test.mockBehaviourGetDonate(*donate, uint64(test.id))

			handler := NewHandler(donate, users)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "https://127.0.0.1/api/v1/donates", nil)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/donates")
			c.SetParamNames("id")
			c.SetParamValues(strconv.FormatInt(test.id, 10))

			err := handler.GetDonate(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}

			body, err := io.ReadAll(rec.Body)
			require.NoError(t, err)

			assert.Equal(t, test.expectedResponse, strings.Trim(string(body), "\n"))
		})
	}
}

func TestHandler_GetDonates(t *testing.T) {
	type mockBehaviourGetDonates func(s mockDomain.MockDonatesUseCase, userID uint64)
	type mockBehaviourGetUser func(s mockDomain.MockUsersUseCase, sessionID string)

	tests := []struct {
		name                    string
		userID                  uint64
		authorID                uint64
		price                   uint64
		cookie                  string
		mockBehaviourGetDonates mockBehaviourGetDonates
		mockBehaviourGetUser    mockBehaviourGetUser
		expectedResponse        string
		expectedErrorMessage    string
	}{
		{
			name:     "OK",
			userID:   23,
			authorID: 12,
			price:    3000,
			cookie:   "JSAoPdaAsdasjdJNPdapoSAjdasakZcs",
			mockBehaviourGetUser: func(s mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: 23,
				}, nil)
			},
			mockBehaviourGetDonates: func(s mockDomain.MockDonatesUseCase, userID uint64) {
				s.EXPECT().GetDonatesByUserID(userID).Return([]models.Donate{
					{
						UserID:   userID,
						AuthorID: 12,
						Price:    3000,
					},
				}, nil)
			},
			expectedResponse: `[{"id":0,"userId":23,"authorId":12,"price":3000}]`,
		},
		{
			name:   "ErrNoSession-1",
			cookie: "JSAoPdaAsdasjdJNPdapoSAjdasakZcs",
			mockBehaviourGetUser: func(s mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{}, errors.New("not valid session"))
			},
			mockBehaviourGetDonates: func(s mockDomain.MockDonatesUseCase, userID uint64) {},
			expectedErrorMessage:    `code=401, message=no existing session, internal=not valid session`,
		},
		{
			name:                    "ErrNoSession-2",
			mockBehaviourGetUser:    func(s mockDomain.MockUsersUseCase, sessionID string) {},
			mockBehaviourGetDonates: func(s mockDomain.MockDonatesUseCase, userID uint64) {},
			expectedErrorMessage:    `code=401, message=no existing session, internal=http: named cookie not present`,
		},
		{
			name:     "ErrNotFound",
			userID:   23,
			authorID: 12,
			price:    3000,
			cookie:   "JSAoPdaAsdasjdJNPdapoSAjdasakZcs",
			mockBehaviourGetUser: func(s mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: 23,
				}, nil)
			},
			mockBehaviourGetDonates: func(s mockDomain.MockDonatesUseCase, userID uint64) {
				s.EXPECT().GetDonatesByUserID(userID).Return([]models.Donate{}, errors.New("donates not found"))
			},
			expectedErrorMessage: `code=404, message=failed to find item, internal=donates not found`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			donate := mockDomain.NewMockDonatesUseCase(ctrl)
			users := mockDomain.NewMockUsersUseCase(ctrl)

			test.mockBehaviourGetUser(*users, test.cookie)
			test.mockBehaviourGetDonates(*donate, test.userID)

			handler := NewHandler(donate, users)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "https://127.0.0.1/api/v1/donates", nil)
			if test.name != "ErrNoSession-2" {
				req.Header.Add("Cookie", "session_id="+test.cookie)
			}
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/donates")

			err := handler.GetDonates(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}

			body, err := io.ReadAll(rec.Body)
			require.NoError(t, err)

			assert.Equal(t, test.expectedResponse, strings.Trim(string(body), "\n"))
		})
	}
}
