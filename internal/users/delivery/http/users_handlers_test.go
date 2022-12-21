package httpUsers

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
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
	type mockUserBehavior func(r *mockDomain.MockUsersUseCase, id uint64)
	type mockImageBehavior func(r *mockDomain.MockImageUseCase, bucket, filename string)
	type mockSubscribersBehavior func(r *mockDomain.MockSubscribersUseCase, userID uint64)
	type mockSubscriptionsBehavior func(r *mockDomain.MockSubscriptionsUseCase, userID uint64)
	type mockStatisticsBehavior func(r *mockDomain.MockUsersUseCase, userID uint64)
	type mockGetBySession func(r *mockDomain.MockUsersUseCase, sessionID string)

	tests := []struct {
		name                        string
		redirectID                  int
		mockUserBehavior            mockUserBehavior
		mockImageBehavior           mockImageBehavior
		mockSubscriptionsBehavior   mockSubscriptionsBehavior
		mockSubscribersBehavior     mockSubscribersBehavior
		mockGetPostStat             mockStatisticsBehavior
		mockGetSubscribersStat      mockStatisticsBehavior
		mockGetProfitStat           mockStatisticsBehavior
		mockGetSubscribersMountStat mockStatisticsBehavior
		mockGetBySession            mockGetBySession
		expectedResponseBody        string
		expectedErrorMessage        string
	}{
		{
			name:       "OK",
			redirectID: 24,
			mockUserBehavior: func(r *mockDomain.MockUsersUseCase, id uint64) {
				r.EXPECT().GetByID(id).Return(models.User{
					ID:       id,
					Username: "themilchenko",
					Avatar:   "avatar",
					Email:    "example@ex.com",
					IsAuthor: true,
				}, nil)
			},
			mockImageBehavior: func(r *mockDomain.MockImageUseCase, bucket, filename string) {
				r.EXPECT().GetImage(bucket).Return("", nil)
			},
			mockSubscriptionsBehavior: func(r *mockDomain.MockSubscriptionsUseCase, userID uint64) {
				r.EXPECT().GetSubscriptionsByUserID(userID).Return([]models.AuthorSubscription{
					{AuthorID: 25},
				}, nil)
			},
			mockSubscribersBehavior: func(r *mockDomain.MockSubscribersUseCase, userID uint64) {
				r.EXPECT().GetSubscribers(userID).Return([]models.User{
					{ID: 24},
				}, nil)
			},
			mockGetPostStat: func(r *mockDomain.MockUsersUseCase, userID uint64) {
				r.EXPECT().GetPostsNum(userID).Return(uint64(1), nil)
			},
			mockGetSubscribersStat: func(r *mockDomain.MockUsersUseCase, userID uint64) {
				r.EXPECT().GetSubscribersNum(userID).Return(uint64(1), nil)
			},
			mockGetProfitStat: func(r *mockDomain.MockUsersUseCase, userID uint64) {
				r.EXPECT().GetProfitForMounth(userID).Return(uint64(1), nil)
			},
			mockGetBySession: func(r *mockDomain.MockUsersUseCase, sessionID string) {
				r.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: 24,
				}, nil)
			},
			expectedResponseBody: `{"id":24,"username":"themilchenko","email":"example@ex.com","avatar":"","isAuthor":true,"about":"","countSubscriptions":1,"countSubscribers":1,"countPosts":1,"countSubscribersMounth":1,"countProfitMounth":1}`,
		},
		{
			name:                      "BadID",
			redirectID:                -1,
			mockUserBehavior:          func(r *mockDomain.MockUsersUseCase, id uint64) {},
			mockImageBehavior:         func(r *mockDomain.MockImageUseCase, bucket, filename string) {},
			mockSubscriptionsBehavior: func(r *mockDomain.MockSubscriptionsUseCase, userID uint64) {},
			mockSubscribersBehavior:   func(r *mockDomain.MockSubscribersUseCase, userID uint64) {},
			mockGetPostStat:           func(r *mockDomain.MockUsersUseCase, userID uint64) {},
			mockGetSubscribersStat:    func(r *mockDomain.MockUsersUseCase, userID uint64) {},
			mockGetProfitStat:         func(r *mockDomain.MockUsersUseCase, userID uint64) {},
			mockGetBySession:          func(r *mockDomain.MockUsersUseCase, sessionID string) {},
			expectedErrorMessage:      "code=400, message=bad request, internal=strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
		{
			name:       "NotFound",
			redirectID: 24,
			mockUserBehavior: func(r *mockDomain.MockUsersUseCase, id uint64) {
				r.EXPECT().GetByID(id).Return(models.User{}, domain.ErrNotFound)
			},
			mockImageBehavior:         func(r *mockDomain.MockImageUseCase, bucket, filename string) {},
			mockSubscriptionsBehavior: func(r *mockDomain.MockSubscriptionsUseCase, userID uint64) {},
			mockSubscribersBehavior:   func(r *mockDomain.MockSubscribersUseCase, userID uint64) {},
			mockGetPostStat:           func(r *mockDomain.MockUsersUseCase, userID uint64) {},
			mockGetSubscribersStat:    func(r *mockDomain.MockUsersUseCase, userID uint64) {},
			mockGetProfitStat:         func(r *mockDomain.MockUsersUseCase, userID uint64) {},
			mockGetBySession: func(r *mockDomain.MockUsersUseCase, sessionID string) {
				r.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: 24,
				}, nil)
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
			subscriber := mockDomain.NewMockSubscribersUseCase(ctrl)
			subscription := mockDomain.NewMockSubscriptionsUseCase(ctrl)

			test.mockUserBehavior(user, uint64(test.redirectID))
			test.mockImageBehavior(image, "avatar", "filename")
			test.mockSubscribersBehavior(subscriber, uint64(test.redirectID))
			test.mockSubscriptionsBehavior(subscription, uint64(test.redirectID))
			test.mockGetPostStat(user, uint64(test.redirectID))
			test.mockGetSubscribersStat(user, uint64(test.redirectID))
			test.mockGetProfitStat(user, uint64(test.redirectID))
			test.mockGetBySession(user, "session_id")

			handler := NewHandler(user, auth, image, subscription, subscriber)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/users", nil)
			req.AddCookie(&http.Cookie{Name: "session_id", Value: "session_id"})
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("https://127.0.0.1/api/v1/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(strconv.FormatInt(int64(test.redirectID), 10))
			c.Set("bucket", "avatar")
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
	type mockUserBehavior func(r *mockDomain.MockUsersUseCase, u models.User, f *multipart.FileHeader, id uint64)
	type mockGetImageBehaviout func(r *mockDomain.MockImageUseCase, avatart string)

	tests := []struct {
		name                  string
		userID                int
		requestBody           multipart.Form
		inputUser             models.User
		mockUserBehavior      mockUserBehavior
		mockGetImageBehaviout mockGetImageBehaviout
		expectedErrorMessage  string
		expectedResponseBody  string
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
			mockUserBehavior: func(r *mockDomain.MockUsersUseCase, u models.User, f *multipart.FileHeader, id uint64) {
				r.EXPECT().Update(u, f, id).Return(models.User{}, nil)
			},
			mockGetImageBehaviout: func(r *mockDomain.MockImageUseCase, avatar string) {
				r.EXPECT().GetImage(avatar).Return("", nil)
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
			mockUserBehavior:      func(r *mockDomain.MockUsersUseCase, u models.User, f *multipart.FileHeader, id uint64) {},
			mockGetImageBehaviout: func(r *mockDomain.MockImageUseCase, avatar string) {},
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
			userID:                345,
			inputUser:             models.User{},
			mockUserBehavior:      func(r *mockDomain.MockUsersUseCase, u models.User, f *multipart.FileHeader, id uint64) {},
			mockGetImageBehaviout: func(r *mockDomain.MockImageUseCase, avatar string) {},
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
			mockUserBehavior: func(r *mockDomain.MockUsersUseCase, u models.User, f *multipart.FileHeader, id uint64) {
				r.EXPECT().Update(u, f, id).Return(models.User{}, domain.ErrUpdate)
			},
			mockGetImageBehaviout: func(r *mockDomain.MockImageUseCase, avatar string) {},
			expectedErrorMessage:  "code=500, message=failed to update item, internal=failed to update item",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			user := mockDomain.NewMockUsersUseCase(ctrl)
			auth := mockDomain.NewMockAuthUseCase(ctrl)
			image := mockDomain.NewMockImageUseCase(ctrl)
			subscriber := mockDomain.NewMockSubscribersUseCase(ctrl)
			subscription := mockDomain.NewMockSubscriptionsUseCase(ctrl)

			handler := NewHandler(user, auth, image, subscription, subscriber)

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

			formFile, err := writer.CreateFormFile("file", "../../../../test/test.txt")
			assert.NoError(t, err)

			sample, err := os.Open("../../../../test/test.txt")
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
			c.Set("bucket", "avatar")

			f, err := c.FormFile("file")
			assert.NoError(t, err)

			test.mockUserBehavior(user, test.inputUser, f, uint64(test.userID))
			test.mockGetImageBehaviout(image, "")

			err = handler.PutUser(c)
			if err != nil {
				assert.Equal(t, test.expectedErrorMessage, err.Error())
			}
		})
	}
}

func TestHandler_GetAuthors(t *testing.T) {
	type MockAuthors func(r *mockDomain.MockUsersUseCase, keyword string)
	type MockSubscribers func(r *mockDomain.MockSubscribersUseCase, authorID uint64)

	tests := []struct {
		name            string
		keyword         string
		mockAuthors     MockAuthors
		mockSubscribers MockSubscribers
		responseMessage string
		responseError   string
	}{
		{
			name:    "OK",
			keyword: "superuser",
			mockAuthors: func(r *mockDomain.MockUsersUseCase, keyword string) {
				r.EXPECT().FindAuthors(keyword).Return([]models.User{
					{
						ID:       345,
						Username: "superuser",
					},
				}, nil)
			},
			mockSubscribers: func(r *mockDomain.MockSubscribersUseCase, authorID uint64) {
				r.EXPECT().GetSubscribers(authorID).Return([]models.User{}, nil)
			},
			responseMessage: `[{"id":345,"username":"superuser","email":"","avatar":"","isAuthor":true,"about":"","countSubscriptions":0,"countSubscribers":0,"countPosts":0,"countSubscribersMounth":0,"countProfitMounth":0}]`,
		},
		{
			name:    "ErrFindAuthors",
			keyword: "superuser",
			mockAuthors: func(r *mockDomain.MockUsersUseCase, keyword string) {
				r.EXPECT().FindAuthors(keyword).Return([]models.User{}, domain.ErrInternal)
			},
			mockSubscribers: func(r *mockDomain.MockSubscribersUseCase, authorID uint64) {},
			responseError:   "code=500, message=server error, internal=server error",
		},
		{
			name:    "ErrGetSubscribers",
			keyword: "superuser",
			mockAuthors: func(r *mockDomain.MockUsersUseCase, keyword string) {
				r.EXPECT().FindAuthors(keyword).Return([]models.User{
					{
						ID:       345,
						Username: "superuser",
					},
				}, nil)
			},
			mockSubscribers: func(r *mockDomain.MockSubscribersUseCase, authorID uint64) {
				r.EXPECT().GetSubscribers(authorID).Return([]models.User{}, domain.ErrInternal)
			},
			responseError: "code=500, message=server error, internal=server error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			user := mockDomain.NewMockUsersUseCase(ctrl)
			auth := mockDomain.NewMockAuthUseCase(ctrl)
			image := mockDomain.NewMockImageUseCase(ctrl)
			subscriber := mockDomain.NewMockSubscribersUseCase(ctrl)
			subscription := mockDomain.NewMockSubscriptionsUseCase(ctrl)

			handler := NewHandler(user, auth, image, subscription, subscriber)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/users/authors?keyword="+test.keyword, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("keyword")
			c.SetParamValues(test.keyword)

			test.mockAuthors(user, test.keyword)
			test.mockSubscribers(subscriber, 345)

			err := handler.GetAuthors(c)
			if err != nil {
				assert.Equal(t, test.responseError, err.Error())
			}

			body, _ := io.ReadAll(rec.Body)

			assert.Equal(t, test.responseMessage, strings.Trim(string(body), "\n"))
		})
	}
}
