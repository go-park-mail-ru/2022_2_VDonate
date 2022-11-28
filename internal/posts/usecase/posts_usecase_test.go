package posts

// func TestUsecase_GetPostsByUserID(t *testing.T) {
// 	type mockBehaviour func(s *mockDomain.MockPostsRepository, userID uint64)
//
// 	tests := []struct {
// 		name                 string
// 		userID               uint64
// 		mockBehaviour        mockBehaviour
// 		response             []models.Post
// 		responseErrorMessage string
// 	}{
// 		{
// 			name:   "OK",
// 			userID: 200,
// 			mockBehaviour: func(s *mockDomain.MockPostsRepository, userID uint64) {
// 				s.EXPECT().GetAllByUserID(userID).Return([]models.Post{
// 					{
// 						ID:     1,
// 						UserID: 200,
// 						Title:  "title",
// 						Text:   "text",
// 					},
// 				}, nil)
// 			},
// 			response: []models.Post{
// 				{
// 					ID:     1,
// 					UserID: 200,
// 					Title:  "title",
// 					Text:   "text",
// 				},
// 			},
// 		},
// 		{
// 			name:   "NoPosts",
// 			userID: 200,
// 			mockBehaviour: func(s *mockDomain.MockPostsRepository, userID uint64) {
// 				s.EXPECT().GetAllByUserID(userID).Return([]models.Post{}, nil)
// 			},
// 			response:             []models.Post{},
// 			responseErrorMessage: "no posts were found",
// 		},
// 	}
//
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()
//
// 			postMock := mockDomain.NewMockPostsRepository(ctrl)
// 			userMock := mockDomain.NewMockUsersRepository(ctrl)
// 			imgMock := mockDomain.NewMockImageUseCase(ctrl)
//
// 			test.mockBehaviour(postMock, test.userID)
//
// 			usecase := New(postMock, userMock, imgMock)
// 			if err != nil {
// 				assert.Equal(t, test.responseErrorMessage, err.Error())
// 			}
//
// 			require.Equal(t, test.response, post)
// 		})
// 	}
// }
