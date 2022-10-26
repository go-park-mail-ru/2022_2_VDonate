package httpUsers

import (
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	images "github.com/go-park-mail-ru/2022_2_VDonate/internal/images/usecase"
	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestHadler_GetUser(t *testing.T) {
	type mockUserBehavior func(r *mockDomain.MockUsersUseCase, id uint64)

	type mockImageBehavior func(r *mockDomain.MockImageUseCase, bucket, filename string)

	tests := []struct {
		name                 string
		redirectID           int
		mockUserBehavior     mockUserBehavior
		mockImageBehavior    mockImageBehavior
		expectedResponseBody string
		expectedErrorMessage string
	}{
		{
			name:       "OK",
			redirectID: 24,
			mockUserBehavior: func(r *mockDomain.MockUsersUseCase, id uint64) {
				r.EXPECT().GetByID(id).Return(&models.User{
					ID:       id,
					Username: "themilchenko",
					Avatar:   "filename",
					Email:    "example@ex.com",
					IsAuthor: false,
				}, nil)
			},
			mockImageBehavior: func(r *mockDomain.MockImageUseCase, bucket, filename string) {
				r.EXPECT().GetImage(bucket, filename).Return(&url.URL{}, nil)
			},
			expectedResponseBody: `{"id":24,"username":"themilchenko","email":"example@ex.com","is_author":false}`,
		},
		{
			name:                 "BadID",
			redirectID:           -1,
			mockUserBehavior:     func(r *mockDomain.MockUsersUseCase, id uint64) {},
			mockImageBehavior:    func(r *mockDomain.MockImageUseCase, bucket, filename string) {},
			expectedErrorMessage: "code=400, message=bad request, internal=strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
		{
			name:       "NotFound",
			redirectID: 24,
			mockUserBehavior: func(r *mockDomain.MockUsersUseCase, id uint64) {
				r.EXPECT().GetByID(id).Return(nil, domain.ErrNotFound)
			},
			mockImageBehavior:    func(r *mockDomain.MockImageUseCase, bucket, filename string) {},
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

			test.mockUserBehavior(user, uint64(test.redirectID))
			test.mockImageBehavior(image, "avatar", "filename")

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
	type mockImagesBehavior func(r *mockDomain.MockImageUseCase, bucket string, c echo.Context)

	tests := []struct {
		name                 string
		userID               int
		requestBody          multipart.Form
		inputUser            models.User
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
			inputUser: models.User{
				ID:       345,
				Username: "superuser",
				About:    "I love sport",
			},
			mockUserBehavior: func(r *mockDomain.MockUsersUseCase, id uint64, user models.User) {
				r.EXPECT().Update(user).Return(nil)
			},
			mockImagesBehavior: func(r *mockDomain.MockImageUseCase, bucket string, c echo.Context) {
				file, err := images.GetFileFromContext(c)
				assert.NoError(t, err)
				r.EXPECT().CreateImage(file, bucket)
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
			inputUser: models.User{
				ID:       345,
				Username: "superuser",
				About:    "I love sport",
			},
			mockUserBehavior:   func(r *mockDomain.MockUsersUseCase, id uint64, user models.User) {},
			mockImagesBehavior: func(r *mockDomain.MockImageUseCase, bucket string, c echo.Context) {},
			expectedErrorMessage: "" +
				"code=400, " +
				"message=bad request, " +
				"internal=strconv.ParseUint: parsing \"-1\": invalid syntax",
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
			inputUser:          models.User{},
			mockUserBehavior:   func(r *mockDomain.MockUsersUseCase, id uint64, user models.User) {},
			mockImagesBehavior: func(r *mockDomain.MockImageUseCase, bucket string, c echo.Context) {},
			expectedErrorMessage: "" +
				"code=400, " +
				"message=bad request, " +
				"internal=code=400, " +
				"message=strconv.ParseUint: parsing \"�\": invalid syntax, internal=strconv.ParseUint: parsing \"�\": invalid syntax",
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
			inputUser: models.User{
				ID:       345,
				Username: "superuser",
				About:    "I love sport",
			},
			mockUserBehavior: func(r *mockDomain.MockUsersUseCase, id uint64, user models.User) {
				r.EXPECT().Update(user).Return(domain.ErrUpdate)
			},
			mockImagesBehavior: func(r *mockDomain.MockImageUseCase, bucket string, c echo.Context) {
				file, err := images.GetFileFromContext(c)
				assert.NoError(t, err)
				r.EXPECT().CreateImage(file, bucket)
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

			test.mockUserBehavior(user, test.inputUser.ID, test.inputUser)

			handler := NewHandler(user, auth, image, "avatar")

			e := echo.New()
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)

			if test.name == "BadRequest-Bind" {
				err := writer.WriteField("id", string(rune(-1)))
				assert.NoError(t, err)
			} else {
				err := writer.WriteField("id", strconv.FormatUint(test.inputUser.ID, 10))
				assert.NoError(t, err)
			}

			formFile, err := writer.CreateFormFile("file", "../../../test/test.txt")
			assert.NoError(t, err)

			sample, err := os.Open("../../../test/test.txt")
			assert.NoError(t, err)

			_, err = io.Copy(formFile, sample)
			assert.NoError(t, err)

			err = writer.WriteField("username", test.inputUser.Username)
			assert.NoError(t, err)

			err = writer.WriteField("about", test.inputUser.About)
			assert.NoError(t, err)

			err = writer.Close()
			assert.NoError(t, err)

			req := httptest.NewRequest(http.MethodPut, "https://127.0.0.1/api/v1/users/", body)
			rec := httptest.NewRecorder()
			req.Header.Set("Content-Type", writer.FormDataContentType())

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(strconv.FormatInt(int64(test.userID), 10))

			test.mockImagesBehavior(image, "avatar", c)
			err = handler.PutUser(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}
		})
	}
}
