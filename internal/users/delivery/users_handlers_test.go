package httpUsers

import (
	"bytes"
	"io"
	"mime/multipart"
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
			image := mockDomain.NewMockImageUseCase(ctrl)

			test.mockBehavior(user, uint64(test.redirectID))

			handler := NewHandler(user, auth, image, "avatar")

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
	type mockUserBehavior func(r *mockDomain.MockUsersUseCase, id uint64, user models.User)
	type mockImagesBehavior func(r *mockDomain.MockImageUseCase, id uint64, user models.User)

	tests := []struct {
		name                 string
		userID               int
		requestBody          multipart.Form
		userModel            models.User
		mockUserBehavior     mockUserBehavior
		mockImagesBehavior   mockImagesBehavior
		expectedErrorMessage string
		expectedResponseBody string
	}{
		{
			name: "OK",
			requestBody: multipart.Form{
				Value: map[string][]string{
					"id":       {"345"},
					"username": {"superuser"},
					"about":    {"I love sport"},
				},
			},
			userID: 345,
			userModel: models.User{
				ID:       345,
				Username: "superuser",
				About:    "I love sport",
			},
			mockImagesBehavior: func(r *mockDomain.MockImageUseCase, id uint64, user models.User) {},
			mockUserBehavior: func(r *mockDomain.MockUsersUseCase, id uint64, user models.User) {
				r.EXPECT().Update(user).Return(nil)
			},
			expectedResponseBody: `{"id":345,"username":"superuser","email":"","is_author":true,"about":"I love sport"}`,
		},
		{
			name: "BadRequest-ID",
			requestBody: multipart.Form{
				Value: map[string][]string{
					"id":       {"345"},
					"username": {"superuser"},
					"about":    {"I love sport"},
				},
			},
			userID: -1,
			userModel: models.User{
				ID:       345,
				Username: "superuser",
				About:    "I love sport",
			},
			mockImagesBehavior:   func(r *mockDomain.MockImageUseCase, id uint64, user models.User) {},
			mockUserBehavior:     func(r *mockDomain.MockUsersUseCase, id uint64, user models.User) {},
			expectedErrorMessage: "code=400, message=bad request, internal=strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
		{
			name: "BadRequest-Bind",
			requestBody: multipart.Form{
				Value: map[string][]string{
					"id":       {"345", "123124"},
					"username": {"superuser"},
					"about":    {"I love sport"},
				},
			},
			userID:             345,
			userModel:          models.User{},
			mockImagesBehavior: func(r *mockDomain.MockImageUseCase, id uint64, user models.User) {},
			mockUserBehavior:   func(r *mockDomain.MockUsersUseCase, id uint64, user models.User) {},
			expectedErrorMessage: "code=400, " +
				"message=bad request, " +
				"internal=code=400, " +
				"message=Syntax error: offset=1, " +
				"error=invalid character 'N' looking for beginning of value, " +
				"internal=invalid character 'N' looking for beginning of value",
		},
		{
			name: "ErrUpdate",
			requestBody: multipart.Form{
				Value: map[string][]string{
					"id":       {"345"},
					"username": {"superuser"},
					"about":    {"I love sport"},
				},
			},
			userID: 345,
			userModel: models.User{
				ID:       345,
				Username: "superuser",
				About:    "I love sport",
			},
			mockImagesBehavior: func(r *mockDomain.MockImageUseCase, id uint64, user models.User) {},
			mockUserBehavior: func(r *mockDomain.MockUsersUseCase, id uint64, user models.User) {
				r.EXPECT().Update(user).Return(domain.ErrUpdate)
			},
			expectedErrorMessage: "code=500, message=failed to update item, internal=failed to update item",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			user := mockDomain.NewMockUsersUseCase(ctrl)
			auth := mockDomain.NewMockAuthUseCase(ctrl)
			image := mockDomain.NewMockImageUseCase(ctrl)

			test.mockUserBehavior(user, test.userModel.ID, test.userModel)

			handler := NewHandler(user, auth, image, "avatar")

			e := echo.New()

			body := new(bytes.Buffer)
			w := multipart.NewWriter(body)
			err := w.WriteField("id", test.requestBody.Value["id"][0])
			assert.NoError(t, err)

			req := httptest.NewRequest(http.MethodPut, "https://127.0.0.1/api/v1/users/", body)
			rec := httptest.NewRecorder()
			req.Header.Set("Content-Type", "multipart/form-data")

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(strconv.FormatInt(int64(test.userID), 10))

			err = handler.PutUser(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}
		})
	}
}
