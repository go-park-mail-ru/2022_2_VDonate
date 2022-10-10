package httpPosts

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"github.com/labstack/echo/v4"

// 	"github.com/golang/mock/gomock"
// 	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
// 	"github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/posts/usecase"
// )

// func TestGetPosts(t *testing.T) {
// 	type mockBehavior func(s *mock_posts.MockUseCase, post models.PostDB)

// 	tests := []struct {
// 		name string
// 		inputBody string
// 		inputPost models.PostDB
// 		mockBehavior mockBehavior
// 		expectedStatusCode int
// 		expectedRequestBody  string
// 	}{
// 		{
// 			name: "OK",
// 			inputBody: "",
// 			inputPost: models.PostDB{
// 				ID: 1,
// 				UserID: 2,
// 				Title: "Test",
// 				Text: "Test text.",
// 			},
// 			mockBehavior: func(s *mock_posts.MockUseCase, post models.PostDB) {
// 				s.EXPECT().Create(post).Return(1, nil)
// 			},
// 			expectedStatusCode: http.StatusOK,
// 			expectedRequestBody: `{"id": 1}`,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			// Init dependencies
// 			ctrl := gomock.NewController(t)
// 		 	defer ctrl.Finish()

// 			// mock_posts.NewMockUseCae()
// 			usecase := mock_posts.NewMockUseCase(ctrl)
// 			test.mockBehavior(usecase, test.inputPost)

// 			// service := posts.UseCase{usecase}
// 			// handler := NewHandler(usecase)

// 			// Test server
// 			// e := echo.New()
// 			// req := httptest.NewRequest(http.MethodGet, models.UrlTestPosts, nil)
// 			// rec := httptest.NewRecorder()
// 			// e.GET("http://127.0.0.1/api/v1/posts/users/1", handler)
// 		})
// 	}
// }



// // tests := []models.Test{
// 	// 	models.Test{
// 	// 		Name: "get-post-1",
// 	// 		Request: []byte(`{"ID":1,"UserID":2, "Title":"Test", "Text":"Test text."}`),
// 	// 		Response: &models.PostDB{
// 	// 			ID: 1,
// 	// 			UserID: 2,
// 	// 			Title:  "Test",
// 	// 			Text:   "Test text.",
// 	// 		},
// 	// 		ExpectedError: nil,
// 	// 	},
// 	// }


// // 	ctrl := gomock.NewController(t)
// 		// 	defer ctrl.Finish()

// 		// 	var parsedReqest models.PostDB
// 		// 	err := json.Unmarshal(test.Request, &parsedReqest)
// 		// 	if err != nil {
// 		// 		t.Error(err.Error())
// 		// 	}
			
// 		// 	usecase := mocks.NewMockUseCase(ctrl)
// 		// 	usecase.EXPECT().
// 		// 		GetPostByID(gomock.Eq(parsedReqest.ID)).
// 		// 		Return(&models.PostDB{
// 		// 			ID: test.Response.ID,
// 		// 			UserID: test.Response.UserID,
// 		// 			Title: test.Response.Title,
// 		// 			Text: test.Response.Text,
// 		// 		}, nil)
// 		// 	handler := Handler{postsUseCase: usecase}

// 		// 	e := echo.New()

// 		// 	req := httptest.NewRequest(http.MethodGet, models.UrlTestPosts, nil)
// 		// 	rec := httptest.NewRecorder()

// 		// 	c := e.NewContext(req, rec)
// 		// 	c.SetPath("http://127.0.0.1/api/v1/posts/users/" + strconv.FormatUint(parsedReqest.ID, 10))
// 		// 	c.SetParamNames("id", "UserId", "Title", "Text")
// 		// 	c.SetParamValues(strconv.FormatUint(parsedReqest.UserID, 10), 
// 		// 					strconv.FormatUint(parsedReqest.ID, 10), 
// 		// 					parsedReqest.Title, 
// 		// 					parsedReqest.Text)

// 		// 	err = handler.GetPosts(c)
// 		// 	if err != nil {
// 		// 		t.Errorf("Error is not nil: %s", err)
// 		// 	}

// 		// 	body, err := ioutil.ReadAll(rec.Body)
// 		// 	if err != nil {
// 		// 		t.Errorf("Something went wrong: %s", err)
// 		// 	}

// 		// 	expectedVal, err := json.Marshal(test.Response)
// 		// 	if err != nil {
// 		// 		t.Errorf("Something went wrong: %s", err)
// 		// 	}

// 		// 	if strings.Trim(string(body), "\n") != string(expectedVal) {
// 		// 		t.Errorf("Expected: %s, got: %s", expectedVal, string(body))
// 		// 	}