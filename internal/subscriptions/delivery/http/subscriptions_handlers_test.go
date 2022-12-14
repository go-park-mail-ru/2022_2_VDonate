package httpSubscriptions

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"

	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"

	"github.com/stretchr/testify/assert"

	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
)

func TestHandler_GetSubscriptions(t *testing.T) {
	type mockBehaviorSubscriptions func(u *mockDomain.MockSubscriptionsUseCase, userID uint64)

	tests := []struct {
		name                      string
		userID                    int
		mockBehaviorSubscriptions mockBehaviorSubscriptions
		expectedResponseBody      string
		expectedErrorMessage      string
	}{
		{
			name:   "OK",
			userID: 1,
			mockBehaviorSubscriptions: func(u *mockDomain.MockSubscriptionsUseCase, userID uint64) {
				u.EXPECT().GetSubscriptionsByUserID(userID).Return([]models.AuthorSubscription{
					{Img: "path/to/img"},
				}, nil)
			},
			expectedResponseBody: "[{\"id\":0,\"authorID\":0,\"img\":\"path/to/img\",\"tier\":0,\"title\":\"\",\"text\":\"\",\"price\":0}]",
		},
		{
			name:   "OK-Empty",
			userID: 1,
			mockBehaviorSubscriptions: func(u *mockDomain.MockSubscriptionsUseCase, userID uint64) {
				u.EXPECT().GetSubscriptionsByUserID(userID).Return([]models.AuthorSubscription{}, nil)
			},
			expectedResponseBody: "[]",
		},
		{
			name:   "ErrInternal-subscriptions",
			userID: 1,
			mockBehaviorSubscriptions: func(u *mockDomain.MockSubscriptionsUseCase, userID uint64) {
				u.EXPECT().GetSubscriptionsByUserID(userID).Return(nil, domain.ErrInternal)
			},
			expectedErrorMessage: "code=500, message=server error, internal=server error",
		},
		{
			name:   "ErrInternal-images",
			userID: 1,
			mockBehaviorSubscriptions: func(u *mockDomain.MockSubscriptionsUseCase, userID uint64) {
				u.EXPECT().GetSubscriptionsByUserID(userID).Return([]models.AuthorSubscription{}, domain.ErrInternal)
			},
			expectedErrorMessage: "code=500, message=server error, internal=server error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			user := mockDomain.NewMockUsersUseCase(ctrl)
			subscription := mockDomain.NewMockSubscriptionsUseCase(ctrl)
			image := mockDomain.NewMockImageUseCase(ctrl)

			handler := NewHandler(subscription, user, image)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/", nil)
			val := req.URL.Query()
			val.Add("user_id", fmt.Sprint(test.userID))
			req.URL.RawQuery = val.Encode()

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1")
			c.Set("bucket", "image")

			test.mockBehaviorSubscriptions(subscription, uint64(test.userID))

			err := handler.GetSubscriptions(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			} else {
				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				assert.Equal(t, test.expectedResponseBody, strings.Trim(string(body), "\n"))
			}
		})
	}
}

func TestHandler_GetAuthorSubscriptions(t *testing.T) {
	type mockBehaviorSubscriptions func(u *mockDomain.MockSubscriptionsUseCase, authorID uint64)

	tests := []struct {
		name                      string
		authorID                  int
		mockBehaviorSubscriptions mockBehaviorSubscriptions
		expectedResponseBody      string
		expectedErrorMessage      string
	}{
		{
			name:     "OK",
			authorID: 1,
			mockBehaviorSubscriptions: func(u *mockDomain.MockSubscriptionsUseCase, authorID uint64) {
				u.EXPECT().GetAuthorSubscriptionsByAuthorID(authorID).Return([]models.AuthorSubscription{
					{Img: "path/to/img"},
				}, nil)
			},
			expectedResponseBody: "[{\"id\":0,\"authorID\":0,\"img\":\"path/to/img\",\"tier\":0,\"title\":\"\",\"text\":\"\",\"price\":0}]",
		},
		{
			name:                      "ErrBadRequest",
			authorID:                  -1,
			mockBehaviorSubscriptions: func(u *mockDomain.MockSubscriptionsUseCase, subID uint64) {},
			expectedErrorMessage:      "code=400, message=bad request, internal=strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
		{
			name:     "ErrNotFound",
			authorID: 1,
			mockBehaviorSubscriptions: func(u *mockDomain.MockSubscriptionsUseCase, authorID uint64) {
				u.EXPECT().GetAuthorSubscriptionsByAuthorID(authorID).Return([]models.AuthorSubscription{}, domain.ErrNotFound)
			},
			expectedErrorMessage: "code=404, message=failed to find item, internal=failed to find item",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			user := mockDomain.NewMockUsersUseCase(ctrl)
			subscription := mockDomain.NewMockSubscriptionsUseCase(ctrl)
			image := mockDomain.NewMockImageUseCase(ctrl)

			handler := NewHandler(subscription, user, image)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/", nil)
			val := req.URL.Query()
			val.Add("author_id", fmt.Sprint(test.authorID))
			req.URL.RawQuery = val.Encode()

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1")
			c.Set("bucket", "image")

			test.mockBehaviorSubscriptions(subscription, uint64(test.authorID))

			err := handler.GetAuthorSubscriptions(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			} else {
				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				assert.Equal(t, test.expectedResponseBody, strings.Trim(string(body), "\n"))
			}
		})
	}
}

func TestHandler_GetAuthorSubscription(t *testing.T) {
	type mockBehaviorSubscriptions func(u *mockDomain.MockSubscriptionsUseCase, subID uint64)

	tests := []struct {
		name                      string
		subID                     int
		mockBehaviorSubscriptions mockBehaviorSubscriptions
		expectedResponseBody      string
		expectedErrorMessage      string
	}{
		{
			name:  "OK",
			subID: 1,
			mockBehaviorSubscriptions: func(u *mockDomain.MockSubscriptionsUseCase, subID uint64) {
				u.EXPECT().GetAuthorSubscriptionByID(subID).Return(models.AuthorSubscription{
					Img: "path/to/img",
				}, nil)
			},
			expectedResponseBody: "{\"id\":0,\"authorID\":0,\"img\":\"path/to/img\",\"tier\":0,\"title\":\"\",\"text\":\"\",\"price\":0}",
		},
		{
			name:                      "ErrBadRequest",
			subID:                     -1,
			mockBehaviorSubscriptions: func(u *mockDomain.MockSubscriptionsUseCase, subID uint64) {},
			expectedErrorMessage:      "code=400, message=bad request, internal=strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
		{
			name:  "ErrNotFound",
			subID: 1,
			mockBehaviorSubscriptions: func(u *mockDomain.MockSubscriptionsUseCase, subID uint64) {
				u.EXPECT().GetAuthorSubscriptionByID(subID).Return(models.AuthorSubscription{}, domain.ErrNotFound)
			},
			expectedErrorMessage: "code=404, message=failed to find item, internal=failed to find item",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			user := mockDomain.NewMockUsersUseCase(ctrl)
			subscription := mockDomain.NewMockSubscriptionsUseCase(ctrl)
			image := mockDomain.NewMockImageUseCase(ctrl)

			handler := NewHandler(subscription, user, image)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/", nil)

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1")
			c.Set("bucket", "image")
			c.SetParamNames("id")
			c.SetParamValues(fmt.Sprint(test.subID))

			test.mockBehaviorSubscriptions(subscription, uint64(test.subID))

			err := handler.GetAuthorSubscription(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			} else {
				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				assert.Equal(t, test.expectedResponseBody, strings.Trim(string(body), "\n"))
			}
		})
	}
}

func TestHandler_CreateAuthorSubscription(t *testing.T) {
	type mockGetBySessionID func(u *mockDomain.MockUsersUseCase, sessionID string)
	type mockCreateImage func(u *mockDomain.MockImageUseCase, image *multipart.FileHeader, oldName string)
	type mockAddAuthorSubscription func(u *mockDomain.MockSubscriptionsUseCase, s models.AuthorSubscription, authorID uint64)
	type mockGetAvatar func(u *mockDomain.MockImageUseCase, filename string)

	tests := []struct {
		name                      string
		authorID                  int
		sessionID                 string
		subscription              models.AuthorSubscription
		authorImg                 string
		mockGetBySessionID        mockGetBySessionID
		mockCreateImage           mockCreateImage
		mockAddAuthorSubscription mockAddAuthorSubscription
		mockGetAvatar             mockGetAvatar
		expectedErrorMessage      string
	}{
		{
			name:      "OK",
			authorID:  10,
			sessionID: "some cookie",
			subscription: models.AuthorSubscription{
				Img: "path/to/img",
			},
			authorImg: "path/to/img",
			mockGetBySessionID: func(u *mockDomain.MockUsersUseCase, sessionID string) {
				u.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID:     10,
					Avatar: "avatar",
				}, nil)
			},
			mockCreateImage: func(u *mockDomain.MockImageUseCase, image *multipart.FileHeader, oldName string) {
				u.EXPECT().CreateOrUpdateImage(image, oldName).Return("path/to/img", nil)
			},
			mockAddAuthorSubscription: func(u *mockDomain.MockSubscriptionsUseCase, s models.AuthorSubscription, authorID uint64) {
				u.EXPECT().AddAuthorSubscription(s, authorID).Return(authorID, nil)
			},
			mockGetAvatar: func(u *mockDomain.MockImageUseCase, filename string) {
				u.EXPECT().GetImage(filename).Return("../../../test/test.txt", nil)
			},
		},
		{
			name:      "NoSession-Cookie",
			authorID:  10,
			sessionID: "",
			subscription: models.AuthorSubscription{
				Img: "../../../test/test.txt",
			},
			authorImg:                 "avatar",
			mockGetBySessionID:        func(u *mockDomain.MockUsersUseCase, sessionID string) {},
			mockCreateImage:           func(u *mockDomain.MockImageUseCase, image *multipart.FileHeader, oldName string) {},
			mockAddAuthorSubscription: func(u *mockDomain.MockSubscriptionsUseCase, s models.AuthorSubscription, authorID uint64) {},
			mockGetAvatar:             func(u *mockDomain.MockImageUseCase, filename string) {},
			expectedErrorMessage:      "code=401, message=no existing session, internal=http: named cookie not present",
		},
		{
			name:      "NoSession-User",
			authorID:  10,
			sessionID: "some cookie",
			subscription: models.AuthorSubscription{
				Img: "../../../test/test.txt",
			},
			authorImg: "avatar",
			mockGetBySessionID: func(u *mockDomain.MockUsersUseCase, sessionID string) {
				u.EXPECT().GetBySessionID(sessionID).Return(models.User{}, domain.ErrNoSession)
			},
			mockCreateImage:           func(u *mockDomain.MockImageUseCase, image *multipart.FileHeader, oldName string) {},
			mockAddAuthorSubscription: func(u *mockDomain.MockSubscriptionsUseCase, s models.AuthorSubscription, authorID uint64) {},
			mockGetAvatar:             func(u *mockDomain.MockImageUseCase, filename string) {},
			expectedErrorMessage:      "code=401, message=no existing session, internal=no existing session",
		},
		{
			name:      "BadRequest-Bind",
			authorID:  10,
			sessionID: "some cookie",
			subscription: models.AuthorSubscription{
				Img: "../../../test/test.txt",
			},
			authorImg: "avatar",
			mockGetBySessionID: func(u *mockDomain.MockUsersUseCase, sessionID string) {
				u.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: 10,
				}, nil)
			},
			mockCreateImage:           func(u *mockDomain.MockImageUseCase, image *multipart.FileHeader, oldName string) {},
			mockAddAuthorSubscription: func(u *mockDomain.MockSubscriptionsUseCase, s models.AuthorSubscription, authorID uint64) {},
			mockGetAvatar:             func(u *mockDomain.MockImageUseCase, filename string) {},
			expectedErrorMessage: "code=400, " +
				"message=bad request, " +
				"internal=code=400, " +
				"message=strconv.ParseUint: parsing \"???\": invalid syntax, internal=strconv.ParseUint: parsing \"???\": invalid syntax",
		},
		{
			name:      "ErrCreate-Image",
			authorID:  10,
			sessionID: "some cookie",
			subscription: models.AuthorSubscription{
				Img: "../../../test/test.txt",
			},
			authorImg: "avatar",
			mockGetBySessionID: func(u *mockDomain.MockUsersUseCase, sessionID string) {
				u.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: 10,
				}, nil)
			},
			mockCreateImage: func(u *mockDomain.MockImageUseCase, image *multipart.FileHeader, oldName string) {
				u.EXPECT().CreateOrUpdateImage(image, oldName).Return("", domain.ErrCreate)
			},
			mockAddAuthorSubscription: func(u *mockDomain.MockSubscriptionsUseCase, s models.AuthorSubscription, authorID uint64) {},
			mockGetAvatar:             func(u *mockDomain.MockImageUseCase, filename string) {},
			expectedErrorMessage:      "code=500, message=failed to create item, internal=failed to create item",
		},
		{
			name:      "ErrCreate-Subscription",
			authorID:  10,
			sessionID: "some cookie",
			subscription: models.AuthorSubscription{
				Img: "../../../test/test.txt",
			},
			authorImg: "avatar",
			mockGetBySessionID: func(u *mockDomain.MockUsersUseCase, sessionID string) {
				u.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: 10,
				}, nil)
			},
			mockAddAuthorSubscription: func(u *mockDomain.MockSubscriptionsUseCase, s models.AuthorSubscription, authorID uint64) {
				u.EXPECT().AddAuthorSubscription(s, authorID).Return(uint64(0), domain.ErrCreate)
			},
			mockCreateImage: func(u *mockDomain.MockImageUseCase, image *multipart.FileHeader, oldName string) {
				u.EXPECT().CreateOrUpdateImage(image, oldName).Return("../../../test/test.txt", nil)
			},
			mockGetAvatar:        func(u *mockDomain.MockImageUseCase, filename string) {},
			expectedErrorMessage: "code=500, message=failed to create item, internal=failed to create item",
		},
		{
			name:      "ErrInternal-GetImage",
			authorID:  10,
			sessionID: "some cookie",
			subscription: models.AuthorSubscription{
				Img: "avatar",
			},
			authorImg: "avatar",
			mockGetBySessionID: func(u *mockDomain.MockUsersUseCase, sessionID string) {
				u.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID:     10,
					Avatar: "avatar",
				}, nil)
			},
			mockCreateImage: func(u *mockDomain.MockImageUseCase, image *multipart.FileHeader, oldName string) {
				u.EXPECT().CreateOrUpdateImage(image, oldName).Return("avatar", nil)
			},
			mockAddAuthorSubscription: func(u *mockDomain.MockSubscriptionsUseCase, s models.AuthorSubscription, authorID uint64) {
				u.EXPECT().AddAuthorSubscription(s, authorID).Return(authorID, nil)
			},
			mockGetAvatar: func(u *mockDomain.MockImageUseCase, filename string) {
				u.EXPECT().GetImage(filename).Return("", domain.ErrInternal)
			},
			expectedErrorMessage: "code=500, message=server error, internal=server error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			subscription := mockDomain.NewMockSubscriptionsUseCase(ctrl)
			user := mockDomain.NewMockUsersUseCase(ctrl)
			image := mockDomain.NewMockImageUseCase(ctrl)

			test.mockGetBySessionID(user, test.sessionID)
			test.mockAddAuthorSubscription(subscription, test.subscription, uint64(test.authorID))
			test.mockGetAvatar(image, test.authorImg)

			handler := NewHandler(subscription, user, image)

			e := echo.New()
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)

			var formFile io.Writer

			var err error

			formFile, err = writer.CreateFormFile("file", "../../../test/test.txt")
			assert.NoError(t, err)

			sample, err := os.Open("../../../../test/test.txt")
			assert.NoError(t, err)

			_, err = io.Copy(formFile, sample)
			assert.NoError(t, err)

			err = writer.WriteField("img", test.subscription.Img)
			assert.NoError(t, err)

			err = writer.WriteField("text", test.subscription.Text)
			assert.NoError(t, err)

			err = writer.WriteField("title", test.subscription.Title)
			assert.NoError(t, err)

			err = writer.WriteField("tier", strconv.FormatUint(test.subscription.Tier, 10))
			assert.NoError(t, err)

			if test.name == "BadRequest-Bind" {
				err = writer.WriteField("price", string(rune(-1)))
				assert.NoError(t, err)
			} else {
				err = writer.WriteField("price", strconv.FormatUint(test.subscription.Price, 10))
				assert.NoError(t, err)
			}

			err = writer.Close()
			assert.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "https://127.0.0.1/api/v1/posts/", body)
			if test.name != "NoSession-Cookie" {
				req.AddCookie(&http.Cookie{
					Name:  "session_id",
					Value: test.sessionID,
				})
			}

			rec := httptest.NewRecorder()
			req.Header.Set("Content-Type", writer.FormDataContentType())

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/posts/:post_id")
			c.Set("bucket", "image")

			f, errFile := c.FormFile("file")
			assert.NoError(t, errFile)

			test.mockCreateImage(image, f, "")

			if err = handler.CreateAuthorSubscription(c); err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}
		})
	}
}

func TestHandler_UpdateAuthorSubscription(t *testing.T) {
	type mockCreateImage func(u *mockDomain.MockImageUseCase, f *multipart.FileHeader, img string)
	type mockUpdateAuthor func(u *mockDomain.MockSubscriptionsUseCase, s models.AuthorSubscription, subID uint64)
	type mockGetImage func(u *mockDomain.MockImageUseCase, filename string)

	tests := []struct {
		name                 string
		subID                int
		subscription         models.AuthorSubscription
		mockCreateImage      mockCreateImage
		mockUpdateAuthor     mockUpdateAuthor
		mockGetImage         mockGetImage
		expectedErrorMessage string
	}{
		{
			name:  "OK",
			subID: 10,
			subscription: models.AuthorSubscription{
				Img: "image",
			},
			mockCreateImage: func(u *mockDomain.MockImageUseCase, f *multipart.FileHeader, img string) {
				u.EXPECT().CreateOrUpdateImage(f, img).Return("image", nil)
			},
			mockUpdateAuthor: func(u *mockDomain.MockSubscriptionsUseCase, s models.AuthorSubscription, subID uint64) {
				u.EXPECT().UpdateAuthorSubscription(s, subID).Return(nil)
			},
			mockGetImage: func(u *mockDomain.MockImageUseCase, filename string) {
				u.EXPECT().GetImage(filename).Return("", nil)
			},
		},
		{
			name:  "BadRequest-ID",
			subID: -1,
			subscription: models.AuthorSubscription{
				Img: "../../../../test/test.txt",
			},
			mockCreateImage:      func(u *mockDomain.MockImageUseCase, f *multipart.FileHeader, img string) {},
			mockUpdateAuthor:     func(u *mockDomain.MockSubscriptionsUseCase, s models.AuthorSubscription, subID uint64) {},
			mockGetImage:         func(u *mockDomain.MockImageUseCase, filename string) {},
			expectedErrorMessage: "code=400, message=bad request, internal=strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
		{
			name:  "BadRequest-Bind",
			subID: 10,
			subscription: models.AuthorSubscription{
				Img: "../../../test/test.txt",
			},
			mockCreateImage:  func(u *mockDomain.MockImageUseCase, f *multipart.FileHeader, img string) {},
			mockUpdateAuthor: func(u *mockDomain.MockSubscriptionsUseCase, s models.AuthorSubscription, subID uint64) {},
			mockGetImage:     func(u *mockDomain.MockImageUseCase, filename string) {},
			expectedErrorMessage: "code=400, " +
				"message=bad request, " +
				"internal=code=400, " +
				"message=strconv.ParseUint: parsing \"???\": invalid syntax, internal=strconv.ParseUint: parsing \"???\": invalid syntax",
		},
		{
			name:  "ErrCreate-Image",
			subID: 10,
			subscription: models.AuthorSubscription{
				Img: "../../../test/test.txt",
			},
			mockCreateImage: func(u *mockDomain.MockImageUseCase, f *multipart.FileHeader, img string) {
				u.EXPECT().CreateOrUpdateImage(f, img).Return("", domain.ErrCreate)
			},
			mockUpdateAuthor:     func(u *mockDomain.MockSubscriptionsUseCase, s models.AuthorSubscription, subID uint64) {},
			mockGetImage:         func(u *mockDomain.MockImageUseCase, filename string) {},
			expectedErrorMessage: "code=500, message=failed to create item, internal=failed to create item",
		},
		{
			name:  "ErrUpdate-Subscription",
			subID: 10,
			subscription: models.AuthorSubscription{
				Img: "image",
			},
			mockCreateImage: func(u *mockDomain.MockImageUseCase, f *multipart.FileHeader, img string) {
				u.EXPECT().CreateOrUpdateImage(f, img).Return("image", nil)
			},
			mockUpdateAuthor: func(u *mockDomain.MockSubscriptionsUseCase, s models.AuthorSubscription, subID uint64) {
				u.EXPECT().UpdateAuthorSubscription(s, subID).Return(domain.ErrUpdate)
			},
			mockGetImage:         func(u *mockDomain.MockImageUseCase, filename string) {},
			expectedErrorMessage: "code=500, message=failed to update item, internal=failed to update item",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			subscription := mockDomain.NewMockSubscriptionsUseCase(ctrl)
			user := mockDomain.NewMockUsersUseCase(ctrl)
			image := mockDomain.NewMockImageUseCase(ctrl)

			test.mockUpdateAuthor(subscription, test.subscription, uint64(test.subID))

			handler := NewHandler(subscription, user, image)

			e := echo.New()
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)

			var formFile io.Writer

			var err error

			formFile, err = writer.CreateFormFile("file", "../../../../test/test.txt")
			assert.NoError(t, err)

			sample, err := os.Open("../../../../test/test.txt")
			assert.NoError(t, err)

			_, err = io.Copy(formFile, sample)
			assert.NoError(t, err)

			err = writer.WriteField("img", test.subscription.Img)
			assert.NoError(t, err)

			err = writer.WriteField("text", test.subscription.Text)
			assert.NoError(t, err)

			err = writer.WriteField("title", test.subscription.Title)
			assert.NoError(t, err)

			err = writer.WriteField("tier", strconv.FormatUint(test.subscription.Tier, 10))
			assert.NoError(t, err)

			if test.name == "BadRequest-Bind" {
				err = writer.WriteField("price", string(rune(-1)))
				assert.NoError(t, err)
			} else {
				err = writer.WriteField("price", strconv.FormatUint(test.subscription.Price, 10))
				assert.NoError(t, err)
			}

			err = writer.Close()
			assert.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "https://127.0.0.1/api/v1/posts/", body)

			rec := httptest.NewRecorder()
			req.Header.Set("Content-Type", writer.FormDataContentType())

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/posts/:post_id")
			c.Set("bucket", "image")
			c.SetParamNames("id")
			c.SetParamValues(fmt.Sprint(test.subID))

			f, err := c.FormFile("file")
			assert.NoError(t, err)

			test.mockCreateImage(image, f, "")
			test.mockGetImage(image, "image")

			if err = handler.UpdateAuthorSubscription(c); err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}
		})
	}
}
