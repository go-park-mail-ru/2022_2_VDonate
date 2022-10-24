package httpUsers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHadler_GetUser(t *testing.T) {
	type mockBehavior func(r *mockDomain.MockUsersUseCase, id uint64)

	tests := []struct {
		name                 string
		redirectID           int
		mockBehavior         mockBehavior
		expectedResponseBody string
		expectedErrorMessage string
	}{
		{
			name:       "OK",
			redirectID: 24,
			mockBehavior: func(r *mockDomain.MockUsersUseCase, id uint64) {
				r.EXPECT().GetByID(id).Return(&models.User{
					ID:       id,
					Username: "themilchenko",
					Email:    "example@ex.com",
					IsAuthor: false,
				}, nil)
			},
			expectedResponseBody: `{"id":24,"username":"themilchenko","email":"example@ex.com","is_author":false}`,
		},
		{
			name:                 "BadID",
			redirectID:           -1,
			mockBehavior:         func(r *mockDomain.MockUsersUseCase, id uint64) {},
			expectedErrorMessage: "code=400, message=bad request, internal=strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
		{
			name:       "NotFound",
			redirectID: 24,
			mockBehavior: func(r *mockDomain.MockUsersUseCase, id uint64) {
				r.EXPECT().GetByID(id).Return(nil, domain.ErrNotFound)
			},
			expectedErrorMessage: "code=404, message=failed to find item, internal=failed to find item",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			user := mockDomain.NewMockUsersUseCase(ctrl)
			auth := mockDomain.NewMockAuthUseCase(ctrl)
			test.mockBehavior(user, uint64(test.redirectID))

			handler := NewHandler(user, auth)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/users", nil)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(strconv.FormatInt(int64(test.redirectID), 10))

			err := handler.GetUser(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}

			body, _ := io.ReadAll(rec.Body)

			assert.Equal(t, test.expectedResponseBody, strings.Trim(string(body), "\n"))
		})
	}
}

func TestHandler_PutUser(t *testing.T) {
	type mockBehavior func(r *mockDomain.MockUsersUseCase, id uint64, user models.User)

	tests := []struct {
		name                 string
		userID               int
		requestBody          string
		userModel            models.User
		mockBehavior         mockBehavior
		expectedErrorMessage string
		expectedResponseBody string
	}{
		{
			name:        "OK",
			requestBody: `{"id":345,"username":"superuser","first_name":"Vasya","last_name":"Pupkin","about":"I love sport"}`,
			userID:      345,
			userModel: models.User{
				ID:        345,
				Username:  "superuser",
				FirstName: "Vasya",
				LastName:  "Pupkin",
				About:     "I love sport",
			},
			mockBehavior: func(r *mockDomain.MockUsersUseCase, id uint64, user models.User) {
				r.EXPECT().Update(user).Return(&models.User{
					ID:        id,
					Username:  "superuser",
					FirstName: "Vasya",
					LastName:  "Pupkin",
					IsAuthor:  true,
					About:     "I love sport",
				}, nil)
			},
			expectedResponseBody: `{"id":345,"username":"superuser","first_name":"Vasya","last_name":"Pupkin","email":"","is_author":true,"about":"I love sport"}`,
		},
		{
			name:        "BadRequest-ID",
			requestBody: `{"id":345,"username":"superuser","first_name":"Vasya","last_name":"Pupkin","about":"I love sport"}`,
			userID:      -1,
			userModel: models.User{
				ID:        345,
				Username:  "superuser",
				FirstName: "Vasya",
				LastName:  "Pupkin",
				About:     "I love sport",
			},
			mockBehavior:         func(r *mockDomain.MockUsersUseCase, id uint64, user models.User) {},
			expectedErrorMessage: "code=400, message=bad request, internal=strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
		{
			name:         "BadRequest-Bind",
			requestBody:  `NotJSON`,
			userID:       345,
			userModel:    models.User{},
			mockBehavior: func(r *mockDomain.MockUsersUseCase, id uint64, user models.User) {},
			expectedErrorMessage: "code=400, " +
				"message=bad request, " +
				"internal=code=400, " +
				"message=Syntax error: offset=1, " +
				"error=invalid character 'N' looking for beginning of value, " +
				"internal=invalid character 'N' looking for beginning of value",
		},
		{
			name:        "ErrUpdate",
			requestBody: `{"id":345,"username":"superuser","first_name":"Vasya","last_name":"Pupkin","about":"I love sport"}`,
			userID:      345,
			userModel: models.User{
				ID:        345,
				Username:  "superuser",
				FirstName: "Vasya",
				LastName:  "Pupkin",
				About:     "I love sport",
			},
			mockBehavior: func(r *mockDomain.MockUsersUseCase, id uint64, user models.User) {
				r.EXPECT().Update(user).Return(nil, domain.ErrUpdate)
			},
			expectedErrorMessage: "code=500, message=failed to update item",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			user := mockDomain.NewMockUsersUseCase(ctrl)
			auth := mockDomain.NewMockAuthUseCase(ctrl)
			test.mockBehavior(user, test.userModel.ID, test.userModel)

			handler := NewHandler(user, auth)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "https://127.0.0.1/api/v1/users/", strings.NewReader(test.requestBody))
			rec := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(strconv.FormatInt(int64(test.userID), 10))

			err := handler.PutUser(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}

			body, err := io.ReadAll(rec.Body)
			require.NoError(t, err)

			assert.Equal(t, test.expectedResponseBody, strings.Trim(string(body), "\n"))
		})
	}
}
