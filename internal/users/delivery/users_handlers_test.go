package httpUsers

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/auth/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/users/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHadler_GetUser(t *testing.T) {
	type mockBehavior func(r *mock_users.MockUseCase, id uint64)

	tests := []struct {
		name                 string
		redirectId           int
		mockBehavior         mockBehavior
		expectedResponseBody string
	}{
		{
			name:       "OK-ById",
			redirectId: 24,
			mockBehavior: func(r *mock_users.MockUseCase, id uint64) {
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
			name:                 "Bad-BadId",
			redirectId:           -1,
			mockBehavior:         func(r *mock_users.MockUseCase, id uint64) {},
			expectedResponseBody: `{"message":"unable to convert id"}`,
		},
		{
			name:       "OK-ById",
			redirectId: 24,
			mockBehavior: func(r *mock_users.MockUseCase, id uint64) {
				r.EXPECT().GetByID(id).Return(&models.User{}, errors.New("user not found"))
			},
			expectedResponseBody: `{"message":"user not found"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			user := mock_users.NewMockUseCase(ctrl)
			auth := mock_auth.NewMockUseCase(ctrl)
			test.mockBehavior(user, uint64(test.redirectId))

			handler := NewHandler(user, auth)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/", nil)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/:id")
			c.SetParamNames("id")
			c.SetParamValues(strconv.FormatInt(int64(test.redirectId), 10))

			err := handler.GetUser(c)
			require.NoError(t, err)

			body, _ := ioutil.ReadAll(rec.Body)

			assert.Equal(t, test.expectedResponseBody, strings.Trim(string(body), "\n"))
		})
	}
}

func TestHandler_PutUser(t *testing.T) {
	type mockBehavior func(r *mock_users.MockUseCase, id uint64, user models.User)

	tests := []struct {
		name                 string
		userId               int
		requestBody          string
		userModel            models.User
		mockBehavior         mockBehavior
		expectedResponseBody string
	}{
		{
			name:        "OK",
			requestBody: `{"id":345,"username":"superuser","first_name":"Vasya","last_name":"Pupkin","about":"I love sport"}`,
			userId:      345,
			userModel: models.User{
				ID:        345,
				Username:  "superuser",
				FirstName: "Vasya",
				LastName:  "Pupkin",
				About:     "I love sport",
			},
			mockBehavior: func(r *mock_users.MockUseCase, id uint64, user models.User) {
				r.EXPECT().Update(id, &user).Return(&models.User{
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
			name:        "UserNotFound",
			requestBody: `{"id":345,"username":"superuser","first_name":"Vasya","last_name":"Pupkin","about":"I love sport"}`,
			userId:      345,
			userModel: models.User{
				ID:        345,
				Username:  "superuser",
				FirstName: "Vasya",
				LastName:  "Pupkin",
				About:     "I love sport",
			},
			mockBehavior: func(r *mock_users.MockUseCase, id uint64, user models.User) {
				r.EXPECT().Update(id, &user).Return(&models.User{}, errors.New("failed to update"))
			},
			expectedResponseBody: `{"message":"failed to update"}`,
		},
		{
			name:                 "BadId",
			requestBody:          `{"id":-100,"username":"superuser","first_name":"Vasya","last_name":"Pupkin","about":"I love sport"}`,
			userId:               -100,
			userModel:            models.User{},
			mockBehavior:         func(r *mock_users.MockUseCase, id uint64, user models.User) {},
			expectedResponseBody: `{"message":"unable to convert id"}`,
		},
		{
			name:                 "BindError",
			requestBody:          `mopdapodsooasmp2312nlk14`,
			userId:               100,
			userModel:            models.User{},
			mockBehavior:         func(r *mock_users.MockUseCase, id uint64, user models.User) {},
			expectedResponseBody: `{"message":"bad request"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			user := mock_users.NewMockUseCase(ctrl)
			auth := mock_auth.NewMockUseCase(ctrl)
			test.mockBehavior(user, test.userModel.ID, test.userModel)

			handler := NewHandler(user, auth)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "https://127.0.0.1/api/v1/users/", strings.NewReader(test.requestBody))
			rec := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(strconv.FormatInt(int64(test.userId), 10))

			err := handler.PutUser(c)
			require.NoError(t, err)

			body, err := ioutil.ReadAll(rec.Body)
			require.NoError(t, err)

			assert.Equal(t, test.expectedResponseBody, strings.Trim(string(body), "\n"))
		})
	}
}
