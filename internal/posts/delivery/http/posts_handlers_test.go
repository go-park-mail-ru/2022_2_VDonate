package httpPosts

import (
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/labstack/echo/v4"
)

func TestHandler_GetPosts(t *testing.T) {
	type mockBehaviorGet func(s mockDomain.MockPostsUseCase, userID uint64)
	type mockBehaviorImage func(s mockDomain.MockImageUseCase, bucket, filename string)
	type mockBehaviourPost func(s mockDomain.MockPostsUseCase, postID uint64)
	type mockBehaviourIsLike func(s mockDomain.MockPostsUseCase, userID, postID uint64)
	type mockBehaviorCookie func(s mockDomain.MockUsersUseCase, sessionID string)

	tests := []struct {
		name                 string
		cookie               string
		userID               int
		postID               uint64
		mockBehaviorGet      mockBehaviorGet
		mockBehaviorImage    mockBehaviorImage
		mockBehaviourPost    mockBehaviourPost
		mockBehaviourIsLike  mockBehaviourIsLike
		mockBehaviorCookie   mockBehaviorCookie
		expectedRequestBody  string
		expectedErrorMessage string
	}{
		{
			name:   "OK",
			cookie: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			userID: 123,
			postID: 0,
			mockBehaviorImage: func(s mockDomain.MockImageUseCase, bucket, filename string) {
				s.EXPECT().GetImage(filename).Return("", nil)
			},
			// mockBehaviorGet: func(s mockDomain.MockPostsUseCase, userID uint64) {
			// 	s.EXPECT().GetPostsByUserID(userID).Return([]models.Post{
			// 		{
			// 			UserID: userID,
			// 			Img:    "path/to/img",
			// 			Title:  "Look at this!!!",
			// 			Text:   "Some text about my work",
			// 		},
			// 	}, nil)
			// },
			mockBehaviourPost: func(s mockDomain.MockPostsUseCase, postID uint64) {
				s.EXPECT().GetLikesNum(postID).Return(uint64(0), nil)
			},
			mockBehaviourIsLike: func(s mockDomain.MockPostsUseCase, userID, postID uint64) {
				s.EXPECT().IsPostLiked(userID, postID).Return(true)
			},
			mockBehaviorCookie: func(s mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: 123,
				}, nil)
			},
			expectedRequestBody: `[{"postID":0,"userID":123,"img":"","title":"Look at this!!!","text":"Some text about my work","likesNum":0,"isLiked":true}]`,
		},
		{
			name:              "OK-Empty",
			cookie:            "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			userID:            123,
			postID:            0,
			mockBehaviorImage: func(s mockDomain.MockImageUseCase, bucket, filename string) {},
			// mockBehaviorGet: func(s mockDomain.MockPostsUseCase, userID uint64) {
			// 	s.EXPECT().GetPostsByUserID(userID).Return([]models.Post{}, nil)
			// },
			mockBehaviourPost:   func(s mockDomain.MockPostsUseCase, postID uint64) {},
			mockBehaviourIsLike: func(s mockDomain.MockPostsUseCase, userID, postID uint64) {},
			mockBehaviorCookie:  func(s mockDomain.MockUsersUseCase, sessionID string) {},
			expectedRequestBody: `{}`,
		},
		{
			name:   "ServerError",
			userID: 123,
			// mockBehaviorGet: func(s mockDomain.MockPostsUseCase, userID uint64) {
			// 	s.EXPECT().GetPostsByUserID(userID).Return([]models.Post{}, domain.ErrNotFound)
			// },
			mockBehaviorImage:    func(s mockDomain.MockImageUseCase, bucket, filename string) {},
			mockBehaviourPost:    func(s mockDomain.MockPostsUseCase, postID uint64) {},
			mockBehaviourIsLike:  func(s mockDomain.MockPostsUseCase, userID, postID uint64) {},
			mockBehaviorCookie:   func(s mockDomain.MockUsersUseCase, sessionID string) {},
			expectedErrorMessage: "code=404, message=failed to find item, internal=failed to find item",
		},
		{
			name:                 "BadId",
			userID:               -1,
			mockBehaviorGet:      func(s mockDomain.MockPostsUseCase, userID uint64) {},
			mockBehaviorImage:    func(s mockDomain.MockImageUseCase, bucket, filename string) {},
			mockBehaviourPost:    func(s mockDomain.MockPostsUseCase, postID uint64) {},
			mockBehaviourIsLike:  func(s mockDomain.MockPostsUseCase, userID, postID uint64) {},
			mockBehaviorCookie:   func(s mockDomain.MockUsersUseCase, sessionID string) {},
			expectedErrorMessage: "code=400, message=bad request, internal=strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
		{
			name:   "ErrInternal-Likes",
			cookie: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			userID: 123,
			postID: 0,
			mockBehaviorImage: func(s mockDomain.MockImageUseCase, bucket, filename string) {
				s.EXPECT().GetImage(filename).Return("", nil)
			},
			// mockBehaviorGet: func(s mockDomain.MockPostsUseCase, userID uint64) {
			// 	s.EXPECT().GetPostsByUserID(userID).Return([]models.Post{
			// 		{
			// 			UserID: userID,
			// 			Img:    "path/to/img",
			// 			Title:  "Look at this!!!",
			// 			Text:   "Some text about my work",
			// 		},
			// 	}, nil)
			// },
			mockBehaviourPost: func(s mockDomain.MockPostsUseCase, postID uint64) {
				s.EXPECT().GetLikesNum(postID).Return(uint64(0), domain.ErrInternal)
			},
			mockBehaviourIsLike: func(s mockDomain.MockPostsUseCase, userID, postID uint64) {},
			mockBehaviorCookie: func(s mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: 123,
				}, nil)
			},
			expectedErrorMessage: "code=500, message=server error, internal=server error",
		},
		{
			name:   "ErrInternal-Likes",
			cookie: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			userID: 123,
			postID: 0,
			mockBehaviorImage: func(s mockDomain.MockImageUseCase, bucket, filename string) {
				s.EXPECT().GetImage(filename).Return("", domain.ErrInternal)
			},
			// mockBehaviorGet: func(s mockDomain.MockPostsUseCase, userID uint64) {
			// 	s.EXPECT().GetPostsByUserID(userID).Return([]models.Post{
			// 		{
			// 			UserID: userID,
			// 			Img:    "path/to/img",
			// 			Title:  "Look at this!!!",
			// 			Text:   "Some text about my work",
			// 		},
			// 	}, nil)
			// },
			mockBehaviourPost:   func(s mockDomain.MockPostsUseCase, postID uint64) {},
			mockBehaviourIsLike: func(s mockDomain.MockPostsUseCase, userID, postID uint64) {},
			mockBehaviorCookie: func(s mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID: 123,
				}, nil)
			},
			expectedErrorMessage: "code=500, message=server error, internal=server error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// ctrl := gomock.NewController(t)
			// defer ctrl.Finish()
			//
			// post := mockDomain.NewMockPostsUseCase(ctrl)
			// users := mockDomain.NewMockUsersUseCase(ctrl)
			// image := mockDomain.NewMockImageUseCase(ctrl)
			//
			// test.mockBehaviorGet(*post, uint64(test.userID))
			// test.mockBehaviorImage(*image, "image", "path/to/img")
			// test.mockBehaviourPost(*post, test.postID)
			// test.mockBehaviourIsLike(*post, uint64(test.userID), test.postID)
			// test.mockBehaviorCookie(*users, test.cookie)
			//
			// handler := NewHandler(post, users, image)
			//
			// e := echo.New()
			// req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/posts", nil)
			// req.Header.Add("Cookie", "session_id="+test.cookie)
			// v := req.URL.Query()
			// v.Add("user_id", fmt.Sprint(test.userID))
			// req.URL.RawQuery = v.Encode()
			// rec := httptest.NewRecorder()
			//
			// c := e.NewContext(req, rec)
			// c.SetPath("https://127.0.0.1/api/v1/posts")
			// c.Set("bucket", "image")
			//
			// err := handler.GetPosts(c)
			// if err != nil {
			// 	assert.Equal(t, test.expectedErrorMessage, err.Error())
			// }
			//
			// body, err := io.ReadAll(rec.Body)
			// require.NoError(t, err)
			//
			// assert.Equal(t, test.expectedRequestBody, strings.Trim(string(body), "\n"))
		})
	}
}

func TestHandler_DeletePost(t *testing.T) {
	type mockDelete func(u *mockDomain.MockPostsUseCase, postID uint64)

	tests := []struct {
		name                 string
		postID               int
		mockDelete           mockDelete
		expectedRequestBody  string
		expectedErrorMessage string
	}{
		{
			name:   "OK",
			postID: 10,
			mockDelete: func(u *mockDomain.MockPostsUseCase, postID uint64) {
				u.EXPECT().DeleteByID(postID).Return(nil)
			},
			expectedRequestBody: "{}",
		},
		{
			name:                 "ErrBadRequest",
			postID:               -1,
			mockDelete:           func(u *mockDomain.MockPostsUseCase, postID uint64) {},
			expectedErrorMessage: "code=400, message=bad request, internal=strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
		{
			name:   "ErrDelete",
			postID: 10,
			mockDelete: func(u *mockDomain.MockPostsUseCase, postID uint64) {
				u.EXPECT().DeleteByID(postID).Return(domain.ErrDelete)
			},
			expectedErrorMessage: "code=500, message=server error, internal=failed to delete item",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// ctrl := gomock.NewController(t)
			// defer ctrl.Finish()
			//
			// post := mockDomain.NewMockPostsUseCase(ctrl)
			// users := mockDomain.NewMockUsersUseCase(ctrl)
			// image := mockDomain.NewMockImageUseCase(ctrl)
			//
			// test.mockDelete(post, uint64(test.postID))
			//
			// handler := NewHandler(post, users, image)
			//
			// e := echo.New()
			// req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/", nil)
			// rec := httptest.NewRecorder()
			//
			// c := e.NewContext(req, rec)
			// c.SetPath("https://127.0.0.1/api/v1/:postID")
			// c.SetParamNames("id")
			// c.SetParamValues(strconv.FormatInt(int64(test.postID), 10))
			//
			// err := handler.DeletePost(c)
			// if err != nil {
			// 	assert.Equal(t, test.expectedErrorMessage, err.Error())
			// } else {
			// 	body, err := io.ReadAll(rec.Body)
			// 	require.NoError(t, err)
			//
			// 	assert.Equal(t, test.expectedRequestBody, strings.Trim(string(body), "\n"))
			// }
		})
	}
}

func TestHangler_GetPost(t *testing.T) {
	type mockBehaviorGet func(s mockDomain.MockPostsUseCase, userID uint64)
	type mockBehaviorImage func(s mockDomain.MockImageUseCase, bucket, filename string)
	type mockBehaviourPost func(s mockDomain.MockPostsUseCase, postID uint64)
	type mockBehaviourIsLike func(s mockDomain.MockPostsUseCase, userID, postID uint64)
	type mockBehaviorCookie func(s mockDomain.MockUsersUseCase, sessionID string)

	tests := []struct {
		name                 string
		method               string
		postID               int
		userID               uint64
		cookie               string
		mockBehaviorGet      mockBehaviorGet
		mockBehaviorImage    mockBehaviorImage
		mockBehaviourPost    mockBehaviourPost
		mockBehaviourIsLike  mockBehaviourIsLike
		mockBehaviorCookie   mockBehaviorCookie
		expectedRequestBody  string
		expectedErrorMessage string
	}{
		{
			name:   "OK",
			userID: 0,
			postID: 123,

			mockBehaviorImage: func(s mockDomain.MockImageUseCase, bucket, filename string) {
				s.EXPECT().GetImage(filename).Return("", nil)
			},
			mockBehaviourPost: func(s mockDomain.MockPostsUseCase, postID uint64) {
				s.EXPECT().GetLikesNum(postID).Return(uint64(0), nil)
			},
			mockBehaviourIsLike: func(s mockDomain.MockPostsUseCase, userID, postID uint64) {
				s.EXPECT().IsPostLiked(userID, postID).Return(false)
			},
			mockBehaviorCookie: func(s mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID:       0,
					Username: "username",
				}, nil)
			},
			expectedRequestBody: `{"postID":0,"userID":0,"img":"","title":"Look at this!!!","text":"Some text about my work","likesNum":0,"isLiked":false}`,
		},
		{
			name:   "NotFound",
			postID: 123,

			mockBehaviorImage:    func(s mockDomain.MockImageUseCase, bucket, filename string) {},
			mockBehaviourPost:    func(s mockDomain.MockPostsUseCase, postID uint64) {},
			mockBehaviourIsLike:  func(s mockDomain.MockPostsUseCase, userID, postID uint64) {},
			mockBehaviorCookie:   func(s mockDomain.MockUsersUseCase, sessionID string) {},
			expectedErrorMessage: "code=404, message=failed to find item, internal=failed to find item",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// ctrl := gomock.NewController(t)
			// defer ctrl.Finish()
			//
			// post := mockDomain.NewMockPostsUseCase(ctrl)
			// users := mockDomain.NewMockUsersUseCase(ctrl)
			// image := mockDomain.NewMockImageUseCase(ctrl)
			//
			// test.mockBehaviorGet(*post, uint64(test.postID))
			// test.mockBehaviorImage(*image, "image", "path/to/img")
			// test.mockBehaviourPost(*post, uint64(test.postID))
			// test.mockBehaviourIsLike(*post, test.userID, uint64(test.postID))
			// test.mockBehaviorCookie(*users, test.cookie)
			//
			// handler := NewHandler(post, users, image)
			//
			// e := echo.New()
			// req := httptest.NewRequest(http.MethodGet, "https://127.0.0.1/api/v1/", nil)
			// req.Header.Add("Cookie", "session_id="+test.cookie)
			// rec := httptest.NewRecorder()
			//
			// c := e.NewContext(req, rec)
			// c.SetPath("https://127.0.0.1/api/v1/:postID")
			// c.SetParamNames("id")
			// c.SetParamValues(strconv.FormatInt(int64(test.postID), 10))
			// c.Set("bucket", "image")
			//
			// err := handler.GetPost(c)
			// if err != nil {
			// 	assert.Equal(t, test.expectedErrorMessage, err.Error())
			// }
			//
			// body, err := io.ReadAll(rec.Body)
			// require.NoError(t, err)
			//
			// assert.Equal(t, test.expectedRequestBody, strings.Trim(string(body), "\n"))
		})
	}
}

func TestHandler_CreatePosts(t *testing.T) {
	type mockBehaviorPut func(s *mockDomain.MockPostsUseCase, post models.Post, postID uint64)

	type mockBehaviorImage func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context)

	type mockBehaviorGetImage func(s *mockDomain.MockImageUseCase, name, bucket string)

	tests := []struct {
		name                 string
		postID               int
		inputPost            models.Post
		mockBehaviorImage    mockBehaviorImage
		mockBehaviorPut      mockBehaviorPut
		mockBehaviorGetImage mockBehaviorGetImage
		expectedErrorMessage string
	}{
		{
			name:              "OK",
			postID:            10,
			inputPost:         models.Post{},
			mockBehaviorImage: func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {},
			mockBehaviorPut: func(s *mockDomain.MockPostsUseCase, post models.Post, postID uint64) {
				s.EXPECT().Update(post, postID).Return(nil)
			},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, name, bucket string) {
				s.EXPECT().GetImage(name).Return("path/to/img", nil)
			},
		},
		{
			name:                 "BadRequest-ID",
			postID:               -1,
			inputPost:            models.Post{},
			mockBehaviorImage:    func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {},
			mockBehaviorPut:      func(s *mockDomain.MockPostsUseCase, post models.Post, postID uint64) {},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, name, bucket string) {},
			expectedErrorMessage: "code=400, message=bad request, internal=strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
		{
			name:                 "BadRequest-Bind",
			postID:               10,
			inputPost:            models.Post{},
			mockBehaviorImage:    func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {},
			mockBehaviorPut:      func(s *mockDomain.MockPostsUseCase, post models.Post, postID uint64) {},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, name, bucket string) {},
			expectedErrorMessage: "code=400, message=bad request, internal=code=400, message=strconv.ParseUint: parsing \"�\": invalid syntax, internal=strconv.ParseUint: parsing \"�\": invalid syntax",
		},
		{
			name:                 "ErrCreate",
			postID:               10,
			inputPost:            models.Post{},
			mockBehaviorImage:    func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {},
			mockBehaviorPut:      func(s *mockDomain.MockPostsUseCase, post models.Post, postID uint64) {},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, name, bucket string) {},
			expectedErrorMessage: "code=500, message=failed to create item, internal=failed to create item",
		},
		{
			name:              "Update",
			postID:            10,
			inputPost:         models.Post{},
			mockBehaviorImage: func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {},
			mockBehaviorPut: func(s *mockDomain.MockPostsUseCase, post models.Post, postID uint64) {
				s.EXPECT().Update(post, postID).Return(domain.ErrUpdate)
			},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, name, bucket string) {},
			expectedErrorMessage: "code=500, message=failed to update item, internal=failed to update item",
		},
		{
			name:              "ErrInternal",
			postID:            10,
			inputPost:         models.Post{},
			mockBehaviorImage: func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {},
			mockBehaviorPut: func(s *mockDomain.MockPostsUseCase, post models.Post, postID uint64) {
				s.EXPECT().Update(post, postID).Return(nil)
			},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, name, bucket string) {
				s.EXPECT().GetImage(name).Return("", domain.ErrInternal)
			},
			expectedErrorMessage: "code=500, message=server error, internal=server error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// ctrl := gomock.NewController(t)
			// defer ctrl.Finish()
			//
			// post := mockDomain.NewMockPostsUseCase(ctrl)
			// user := mockDomain.NewMockUsersUseCase(ctrl)
			// image := mockDomain.NewMockImageUseCase(ctrl)
			//
			// test.mockBehaviorPut(post, test.inputPost, uint64(test.postID))
			// test.mockBehaviorGetImage(image, "../../../test/test.txt", "image")
			//
			// handler := NewHandler(post, user, image)
			//
			// e := echo.New()
			// body := new(bytes.Buffer)
			// writer := multipart.NewWriter(body)
			//
			// if test.name == "BadRequest-Bind" {
			// 	err := writer.WriteField("userID", string(rune(-1)))
			// 	assert.NoError(t, err)
			// } else {
			// 	err := writer.WriteField("userID", strconv.FormatUint(test.inputPost.UserID, 10))
			// 	assert.NoError(t, err)
			// }
			//
			// var formFile io.Writer
			//
			// formFile, err := writer.CreateFormFile("file", "../../../test/test.txt")
			// assert.NoError(t, err)
			//
			// sample, err := os.Open("../../../test/test.txt")
			// assert.NoError(t, err)
			//
			// _, err = io.Copy(formFile, sample)
			// assert.NoError(t, err)
			//
			// err = writer.Close()
			// assert.NoError(t, err)
			//
			// req := httptest.NewRequest(http.MethodPost, "https://127.0.0.1/api/v1/posts/", body)
			// rec := httptest.NewRecorder()
			// req.Header.Set("Content-Type", writer.FormDataContentType())
			//
			// c := e.NewContext(req, rec)
			// c.SetPath("https://127.0.0.1/api/v1/posts/:post_id")
			// c.SetParamNames("id")
			// c.SetParamValues(strconv.FormatInt(int64(test.postID), 10))
			// c.Set("bucket", "image")
			//
			// test.mockBehaviorImage(image, "image", c)
			//
			// if err = handler.PutPost(c); err != nil {
			// 	assert.Equal(t, test.expectedErrorMessage, err.Error())
			// }
		})
	}
}

func TestHandler_PutPost(t *testing.T) {
	type mockBehaviorUsers func(s *mockDomain.MockUsersUseCase, cookie string)

	type mockBehaviorCreate func(s *mockDomain.MockPostsUseCase, postID models.Post)

	type mockBehaviorImage func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context)

	type mockBehaviorGetImage func(s *mockDomain.MockImageUseCase, bucket string)

	tests := []struct {
		name                 string
		userID               int
		cookie               string
		inputPost            models.Post
		mockBehaviorUsers    mockBehaviorUsers
		mockBehaviorCreate   mockBehaviorCreate
		mockBehaviorImage    mockBehaviorImage
		mockBehaviorGetImage mockBehaviorGetImage
		expectedStatusCode   int
		expectedErrorMessage string
	}{
		{
			name:      "OK",
			cookie:    "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			inputPost: models.Post{},
			mockBehaviorUsers: func(s *mockDomain.MockUsersUseCase, cookie string) {
				s.EXPECT().GetBySessionID(cookie).Return(models.User{
					ID: 100,
				}, nil)
			},
			mockBehaviorImage: func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {},
			mockBehaviorCreate: func(s *mockDomain.MockPostsUseCase, post models.Post) {
				s.EXPECT().Create(post, post.UserID).Return(uint64(1), nil)
			},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, bucket string) {
				s.EXPECT().GetImage("../../../test/test.txt").Return("", nil)
			},
		},
		{
			name:                 "NoSession",
			userID:               -1,
			cookie:               "",
			inputPost:            models.Post{},
			mockBehaviorUsers:    func(s *mockDomain.MockUsersUseCase, cookie string) {},
			mockBehaviorCreate:   func(s *mockDomain.MockPostsUseCase, post models.Post) {},
			mockBehaviorImage:    func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, bucket string) {},
			expectedErrorMessage: "code=401, message=no existing session, internal=http: named cookie not present",
		},
		{
			name:      "NoSessionForUser",
			userID:    100,
			cookie:    "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			inputPost: models.Post{},
			mockBehaviorUsers: func(s *mockDomain.MockUsersUseCase, cookie string) {
				s.EXPECT().GetBySessionID(cookie).Return(models.User{}, domain.ErrNoSession)
			},
			mockBehaviorCreate:   func(s *mockDomain.MockPostsUseCase, post models.Post) {},
			mockBehaviorImage:    func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, bucket string) {},
			expectedErrorMessage: "code=401, message=no existing session, internal=no existing session",
		},
		{
			name:      "ErrBind",
			userID:    100,
			cookie:    "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			inputPost: models.Post{},
			mockBehaviorUsers: func(s *mockDomain.MockUsersUseCase, cookie string) {
				s.EXPECT().GetBySessionID(cookie).Return(models.User{
					ID: 100,
				}, nil)
			},
			mockBehaviorImage:    func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {},
			mockBehaviorCreate:   func(s *mockDomain.MockPostsUseCase, post models.Post) {},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, bucket string) {},
			expectedErrorMessage: "code=400," +
				" message=bad request," +
				" internal=code=400," +
				" message=strconv.ParseUint: parsing \"�\": invalid syntax, internal=strconv.ParseUint: parsing \"�\": invalid syntax",
		},
		{
			name:      "ErrCreate",
			userID:    100,
			cookie:    "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			inputPost: models.Post{},
			mockBehaviorUsers: func(s *mockDomain.MockUsersUseCase, cookie string) {
				s.EXPECT().GetBySessionID(cookie).Return(models.User{
					ID: 100,
				}, nil)
			},
			mockBehaviorImage: func(s *mockDomain.MockImageUseCase, bucket string, c echo.Context) {},
			mockBehaviorCreate: func(s *mockDomain.MockPostsUseCase, post models.Post) {
				s.EXPECT().Create(post, post.UserID).Return(uint64(0), domain.ErrCreate)
			},
			mockBehaviorGetImage: func(s *mockDomain.MockImageUseCase, bucket string) {},
			expectedErrorMessage: "code=500, message=failed to create item, internal=failed to create item",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// ctrl := gomock.NewController(t)
			// defer ctrl.Finish()
			//
			// post := mockDomain.NewMockPostsUseCase(ctrl)
			// user := mockDomain.NewMockUsersUseCase(ctrl)
			// image := mockDomain.NewMockImageUseCase(ctrl)
			//
			// test.mockBehaviorUsers(user, test.cookie)
			// test.mockBehaviorCreate(post, test.inputPost)
			//
			// handler := NewHandler(post, user, image)
			//
			// e := echo.New()
			// body := new(bytes.Buffer)
			// writer := multipart.NewWriter(body)
			//
			// if test.name == "ErrBind" {
			// 	err := writer.WriteField("userID", string(rune(-1)))
			// 	assert.NoError(t, err)
			// } else {
			// 	err := writer.WriteField("userID", strconv.FormatUint(test.inputPost.UserID, 10))
			// 	assert.NoError(t, err)
			// }
			//
			// formFile, err := writer.CreateFormFile("file", "../../../test/test.txt")
			// assert.NoError(t, err)
			//
			// sample, err := os.Open("../../../test/test.txt")
			// assert.NoError(t, err)
			//
			// _, err = io.Copy(formFile, sample)
			// assert.NoError(t, err)
			//
			// err = writer.Close()
			// assert.NoError(t, err)
			//
			// req := httptest.NewRequest(http.MethodPost, "https://127.0.0.1/api/v1/posts/", body)
			// rec := httptest.NewRecorder()
			// req.Header.Set("Content-Type", writer.FormDataContentType())
			// if len(test.cookie) != 0 {
			// 	req.Header.Add("Cookie", "session_id="+test.cookie)
			// }
			//
			// c := e.NewContext(req, rec)
			// c.SetPath("https://127.0.0.1/api/v1/posts/:postID")
			// c.SetParamNames("postID")
			// c.SetParamValues(strconv.FormatInt(int64(test.userID), 10))
			// c.Set("bucket", "image")
			//
			// test.mockBehaviorImage(image, "image", c)
			// test.mockBehaviorGetImage(image, "image")
			//
			// if err = handler.CreatePost(c); err != nil {
			// 	assert.Equal(t, test.expectedErrorMessage, err.Error())
			// }
		})
	}
}

func TestHandler_GetLikes(t *testing.T) {
	type mockBehaviourLikes func(s *mockDomain.MockPostsUseCase, postID uint64)

	tests := []struct {
		name                 string
		postID               int64
		mockBehaviourLikes   mockBehaviourLikes
		expectedResponse     string
		expectedErrorMessage string
	}{
		{
			name:   "OK-1",
			postID: 23,
			mockBehaviourLikes: func(s *mockDomain.MockPostsUseCase, postID uint64) {
				s.EXPECT().GetLikesByPostID(postID).Return([]models.Like{
					{
						UserID: 12,
						PostID: 23,
					},
				}, nil)
			},
			expectedResponse: `[{"userID":12,"postID":23}]`,
		},
		{
			name:   "OK-2",
			postID: 12,
			mockBehaviourLikes: func(s *mockDomain.MockPostsUseCase, postID uint64) {
				s.EXPECT().GetLikesByPostID(postID).Return([]models.Like{}, nil)
			},
			expectedResponse: "[]",
		},
		{
			name:   "ErrNotFound",
			postID: 100,
			mockBehaviourLikes: func(s *mockDomain.MockPostsUseCase, postID uint64) {
				s.EXPECT().GetLikesByPostID(postID).Return([]models.Like{}, errors.New(""))
			},
			expectedErrorMessage: `code=404, message=failed to find item, internal=`,
		},
		{
			name:                 "BadRequest",
			postID:               -1,
			mockBehaviourLikes:   func(s *mockDomain.MockPostsUseCase, postID uint64) {},
			expectedErrorMessage: `code=400, message=bad request, internal=strconv.ParseUint: parsing "-1": invalid syntax`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// ctrl := gomock.NewController(t)
			// defer ctrl.Finish()
			//
			// post := mockDomain.NewMockPostsUseCase(ctrl)
			// users := mockDomain.NewMockUsersUseCase(ctrl)
			// image := mockDomain.NewMockImageUseCase(ctrl)
			//
			// test.mockBehaviourLikes(post, uint64(test.postID))
			//
			// handler := NewHandler(post, users, image)
			//
			// e := echo.New()
			// req := httptest.NewRequest(http.MethodPut, "https://127.0.0.1/api/v1/posts/:id/likes", nil)
			// rec := httptest.NewRecorder()
			//
			// c := e.NewContext(req, rec)
			// c.SetPath("https://127.0.0.1/api/v1/posts/:id/likes")
			// c.SetParamNames("id")
			// c.SetParamValues(strconv.FormatInt(test.postID, 10))
			//
			// err := handler.GetLikes(c)
			// if err != nil {
			// 	assert.Equal(t, test.expectedErrorMessage, err.Error())
			// }
			//
			// body, err := io.ReadAll(rec.Body)
			// require.NoError(t, err)
			//
			// assert.Equal(t, test.expectedResponse, strings.Trim(string(body), "\n"))
		})
	}
}

func TestHandler_CreateLike(t *testing.T) {
	type mockBehaviourLikes func(s *mockDomain.MockPostsUseCase, userID, postID uint64)
	type mockBehaviourUsers func(s *mockDomain.MockUsersUseCase, sessionID string)

	tests := []struct {
		name                 string
		cookie               string
		postID               int64
		userID               uint64
		mockBehaviourUsers   mockBehaviourUsers
		mockBehaviourLikes   mockBehaviourLikes
		expectedResponse     string
		expectedErrorMessage string
	}{
		{
			name:   "OK",
			cookie: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
			postID: 21,
			userID: 2,
			mockBehaviourUsers: func(s *mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID:       2,
					Username: "a",
					Email:    "a@a.ru",
				}, nil)
			},
			mockBehaviourLikes: func(s *mockDomain.MockPostsUseCase, userID, postID uint64) {
				s.EXPECT().LikePost(userID, postID).Return(nil)
			},
			expectedResponse: `{}`,
		},
		{
			name:   "ErrNoSession-1",
			cookie: "JSAoPdaAsdasjdJNPdapoSAjdasakZcs",
			userID: 1,
			mockBehaviourUsers: func(s *mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{}, errors.New("no session"))
			},
			mockBehaviourLikes:   func(s *mockDomain.MockPostsUseCase, userID, postID uint64) {},
			expectedErrorMessage: "code=401, message=no existing session, internal=no session",
		},
		{
			name:                 "ErrNoSession-2",
			cookie:               "JSAoPdaAsdasjdJNPdapoSAjdasakZcs",
			userID:               1,
			mockBehaviourUsers:   func(s *mockDomain.MockUsersUseCase, sessionID string) {},
			mockBehaviourLikes:   func(s *mockDomain.MockPostsUseCase, userID, postID uint64) {},
			expectedErrorMessage: "code=401, message=no existing session, internal=http: named cookie not present",
		},
		{
			name:   "ErrLikeExist",
			cookie: "JSAoPdaAsdasjdJNPdapoSAjdasakZcs",
			userID: 2,
			mockBehaviourUsers: func(s *mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID:       2,
					Username: "a",
					Email:    "a@a.ru",
				}, nil)
			},
			mockBehaviourLikes: func(s *mockDomain.MockPostsUseCase, userID, postID uint64) {
				s.EXPECT().LikePost(userID, postID).Return(errors.New(""))
			},
			expectedErrorMessage: "code=500, message=like alredy exist, internal=",
		},
		{
			name:                 "ErrBadId",
			postID:               -1,
			mockBehaviourUsers:   func(s *mockDomain.MockUsersUseCase, sessionID string) {},
			mockBehaviourLikes:   func(s *mockDomain.MockPostsUseCase, userID, postID uint64) {},
			expectedErrorMessage: `code=400, message=bad request, internal=strconv.ParseUint: parsing "-1": invalid syntax`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// ctrl := gomock.NewController(t)
			// defer ctrl.Finish()
			//
			// post := mockDomain.NewMockPostsUseCase(ctrl)
			// users := mockDomain.NewMockUsersUseCase(ctrl)
			// image := mockDomain.NewMockImageUseCase(ctrl)
			//
			// test.mockBehaviourUsers(users, test.cookie)
			// test.mockBehaviourLikes(post, test.userID, uint64(test.postID))
			//
			// handler := NewHandler(post, users, image)
			//
			// e := echo.New()
			// req := httptest.NewRequest(http.MethodPut, "https://127.0.0.1/api/v1/posts/:id/likes", nil)
			// if test.name != "ErrNoSession-2" {
			// 	req.Header.Add("Cookie", "session_id="+test.cookie)
			// }
			// rec := httptest.NewRecorder()
			//
			// c := e.NewContext(req, rec)
			// c.SetPath("https://127.0.0.1/api/v1/posts/:id/likes")
			// c.SetParamNames("id")
			// c.SetParamValues(strconv.FormatInt(test.postID, 10))
			//
			// err := handler.CreateLike(c)
			// if err != nil {
			// 	assert.Equal(t, test.expectedErrorMessage, err.Error())
			// }
			//
			// body, err := io.ReadAll(rec.Body)
			// require.NoError(t, err)
			//
			// assert.Equal(t, test.expectedResponse, strings.Trim(string(body), "\n"))
		})
	}
}

func TestHandler_DeleteLike(t *testing.T) {
	type mockBehaviourUsers func(s *mockDomain.MockUsersUseCase, sessionID string)
	type mockBehaviourLikes func(s *mockDomain.MockPostsUseCase, userID, postID uint64)

	tests := []struct {
		name                 string
		postID               int64
		userID               uint64
		cookie               string
		mockBehaviourLikes   mockBehaviourLikes
		mockBehaviourUsers   mockBehaviourUsers
		expectedResponse     string
		expectedErrorMessage string
	}{
		{
			name:   "OK",
			postID: 3,
			userID: 21,
			cookie: "JSAoPdaAsdasjdJNPdapoSAjdasakZcs",
			mockBehaviourUsers: func(s *mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID:       21,
					Username: "user",
					Email:    "a@a.ru",
				}, nil)
			},
			mockBehaviourLikes: func(s *mockDomain.MockPostsUseCase, userID, postID uint64) {
				s.EXPECT().UnlikePost(userID, postID).Return(nil)
			},
			expectedResponse: `{}`,
		},
		{
			name:   "OK",
			postID: 3,
			userID: 21,
			cookie: "JSAoPdaAsdasjdJNPdapoSAjdasakZcs",
			mockBehaviourUsers: func(s *mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{
					ID:       21,
					Username: "user",
					Email:    "a@a.ru",
				}, nil)
			},
			mockBehaviourLikes: func(s *mockDomain.MockPostsUseCase, userID, postID uint64) {
				s.EXPECT().UnlikePost(userID, postID).Return(errors.New("like not found"))
			},
			expectedErrorMessage: "code=404, message=failed to find item, internal=like not found",
		},
		{
			name:   "ErrNoSession-1",
			postID: 3,
			userID: 21,
			cookie: "JSAoPdaAsdasjdJNPdapoSAjdasakZcs",
			mockBehaviourUsers: func(s *mockDomain.MockUsersUseCase, sessionID string) {
				s.EXPECT().GetBySessionID(sessionID).Return(models.User{}, errors.New("no session"))
			},
			mockBehaviourLikes:   func(s *mockDomain.MockPostsUseCase, userID, postID uint64) {},
			expectedErrorMessage: "code=401, message=no existing session, internal=no session",
		},
		{
			name:                 "ErrNoSession-2",
			cookie:               "JSAoPdaAsdasjdJNPdapoSAjdasakZcs",
			mockBehaviourUsers:   func(s *mockDomain.MockUsersUseCase, sessionID string) {},
			mockBehaviourLikes:   func(s *mockDomain.MockPostsUseCase, userID, postID uint64) {},
			expectedErrorMessage: "code=401, message=no existing session, internal=http: named cookie not present",
		},
		{
			name:                 "BadId",
			postID:               -1,
			mockBehaviourUsers:   func(s *mockDomain.MockUsersUseCase, sessionID string) {},
			mockBehaviourLikes:   func(s *mockDomain.MockPostsUseCase, userID, postID uint64) {},
			expectedErrorMessage: "code=400, message=bad request, internal=strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// ctrl := gomock.NewController(t)
			// defer ctrl.Finish()
			//
			// users := mockDomain.NewMockUsersUseCase(ctrl)
			// post := mockDomain.NewMockPostsUseCase(ctrl)
			// image := mockDomain.NewMockImageUseCase(ctrl)
			//
			// test.mockBehaviourUsers(users, test.cookie)
			// test.mockBehaviourLikes(post, test.userID, uint64(test.postID))
			//
			// handler := NewHandler(post, users, image)
			//
			// e := echo.New()
			// req := httptest.NewRequest(http.MethodPut, "https://127.0.0.1/api/v1/posts/:id/likes", nil)
			// if test.name != "ErrNoSession-2" {
			// 	req.Header.Add("Cookie", "session_id="+test.cookie)
			// }
			// rec := httptest.NewRecorder()
			//
			// c := e.NewContext(req, rec)
			// c.SetPath("https://127.0.0.1/api/v1/posts/:id/likes")
			// c.SetParamNames("id")
			// c.SetParamValues(strconv.FormatInt(test.postID, 10))
			//
			// err := handler.DeleteLike(c)
			// if err != nil {
			// 	assert.Equal(t, test.expectedErrorMessage, err.Error())
			// }
			//
			// body, err := io.ReadAll(rec.Body)
			// require.NoError(t, err)
			//
			// assert.Equal(t, test.expectedResponse, strings.Trim(string(body), "\n"))
		})
	}
}
