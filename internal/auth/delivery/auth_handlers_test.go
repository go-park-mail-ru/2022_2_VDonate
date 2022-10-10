package httpAuth

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/auth/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/users/usecase"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandler_signup(t *testing.T) {
	type mockBehavior func(r *mock_auth.MockUseCase, user models.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            models.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	} {
		{
			name:      "Ok",
			inputBody: `{"username": "username", "email": "ex@example.com", "password": "qwerty", "IsAuthor": "false"}`,
			inputUser: models.User{
				Username: "username",
				Password: "qwerty",
				Email: "ex@example.com",
				IsAuthor: false,
			},
			mockBehavior: func(r *mock_auth.MockUseCase, user models.User) {
				r.EXPECT().SignUp(user).Return("1", nil)
			},
			expectedStatusCode:   404,
			expectedResponseBody: `{"message":"Not Found"}
`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_auth.NewMockUseCase(c)
			user := mock_users.NewMockUseCase(c)
			test.mockBehavior(repo, test.inputUser)

			handler := NewHandler(repo, user)

			// Init Endpoint
			r := echo.New()
			r.POST("http://127.0.0.1/api/v1/users", handler.SignUp)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "http://127.0.0.1/api/v1/users", bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}