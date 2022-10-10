package httpUsers

import (
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

func TestHadler_getuser(t *testing.T) {
	type mockBehavior func(r *mock_users.MockUseCase, id uint64)

	tests := []struct {
		name                 string
		redirectId 			 uint64
		mockBehavior         mockBehavior
		expectedResponseBody string
	} {
		{
			name:      "OK-ById-1",
			redirectId: 24,
			mockBehavior: func(r *mock_users.MockUseCase, id uint64) {
				r.EXPECT().GetByID(id).Return(&models.User{
					ID: id,
					Username: "themilchenko",
					Email: "example@ex.com",
					IsAuthor: false,
				}, nil)
			},
			expectedResponseBody: `{"id":24,"username":"themilchenko","email":"example@ex.com","is_author":false}`,
		},
		{
			name:      "Bad-ById-2",
			redirectId: 15,
			mockBehavior: func(r *mock_users.MockUseCase, id uint64) {
				r.EXPECT().GetByID(id).Return(&models.User{
					ID: id,
					Username: "themilchenko",
					Email: "example@ex.com",
					IsAuthor: false,
				}, nil)
			},
			expectedResponseBody: `{"id":15,"username":"themilchenko","email":"example@ex.com","is_author":false}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			if test.redirectId == 15 {

			}

			user := mock_users.NewMockUseCase(ctrl)
			auth := mock_auth.NewMockUseCase(ctrl)
			test.mockBehavior(user, test.redirectId)

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
