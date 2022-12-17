package posts

import (
	"errors"
	"testing"

	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBlurContent(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "single image",
			input:    `<img src="https://wsrv.nl/?url=http://95.163.209.195:9000/e4da3b7fbbce2345d7772b0674a318d5/7c95af43-de59-4c34-b2f6-19cd5b134a65.png" class="post-content__image">`,
			expected: `<img src="https://wsrv.nl/?url=http://95.163.209.195:9000/e4da3b7fbbce2345d7772b0674a318d5/blur_7c95af43-de59-4c34-b2f6-19cd5b134a65.png" class="post-content__image">`,
		},
		{
			name:     "multiple images",
			input:    `<h1>3-ий пост</h1><div>фото первое:</div><img src="https://wsrv.nl/?url=http://95.163.209.195:9000/45c48cce2e2d7fbdea1afc51c7c6ad26/952eff70-5545-4d9c-a4a1-3101039cdd09.jpg" class="post-content__image"><div>фото 2-ое:</div><img src="https://wsrv.nl/?url=http://95.163.209.195:9000/45c48cce2e2d7fbdea1afc51c7c6ad26/bf66f759-60fb-4202-9c0f-0fcd06c173d9.jpg" class="post-content__image"><div>и 3-е с 4-ым:</div><img src="https://wsrv.nl/?url=http://95.163.209.195:9000/8f14e45fceea167a5a36dedd4bea2543/98141f7f-f33e-4fae-9fb0-205be040b8a7.jpg" class="post-content__image"><img src="https://wsrv.nl/?url=http://95.163.209.195:9000/e1671797c52e15f763380b45e841ec32/05792d15-51a0-4bb3-a925-3c707e0eeade.jpg" class="post-content__image">`,
			expected: `<img src="https://wsrv.nl/?url=http://95.163.209.195:9000/45c48cce2e2d7fbdea1afc51c7c6ad26/blur_952eff70-5545-4d9c-a4a1-3101039cdd09.jpg" class="post-content__image">`,
		},
		{
			name:     "no images",
			input:    `<h1>3-ий пост</h1><div>фото первое:</div><div>фото 2-ое:</div><div>и 3-е с 4-ым:</div>`,
			expected: ``,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, blurContent(test.input))
		})
	}
}

func TestSanitizeContent(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "onmouseover",
			input:    `<b onmouseover=alert('Wufff!')>click me!</b>`,
			expected: "<b>click me!</b>",
		},
		{
			name:     "onerror",
			input:    `<img src="http://url.to.file.which/not.exist" onerror=alert(document.cookie);>`,
			expected: "<img src=\"http://url.to.file.which/not.exist\">",
		},
		{
			name:  "script",
			input: `<script>alert('Wufff!')</script>`,
		},
		{
			name:  "iframe",
			input: `<iframe src="javascript:alert('Wufff!');"></iframe>`,
		},
		{
			name:     "real_img",
			input:    `<img src="https://wsrv.nl/?url=http://95.163.209.195:9000/e4da3b7fbbce2345d7772b0674a318d5/7c95af43-de59-4c34-b2f6-19cd5b134a65.png" class="post-content__image">`,
			expected: "<img src=\"https://wsrv.nl/?url=http://95.163.209.195:9000/e4da3b7fbbce2345d7772b0674a318d5/7c95af43-de59-4c34-b2f6-19cd5b134a65.png\" class=\"post-content__image\">",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, SanitizeContent(test.input, false))
		})
	}
}

// func TestUsecase_GetPostsByUserID(t *testing.T) {
// 	type mockPost func(s *mockDomain.MockPostsMicroservice, id uint64)
// 	type mockUser func(s *mockDomain.MockUsersMicroservice, id uint64)
// 	type mockSubscription func(s *mockDomain.MockSubscriptionMicroservice, userID, authorID uint64)
// 	type mockTagDeps func(s *mockDomain.MockPostsMicroservice, postID uint64)
// 	type mockTags func(s *mockDomain.MockPostsMicroservice, tagID uint64)
// 	type mockImg func(s *mockDomain.MockImageUseCase, img string)
// 	type mockLike func(s *mockDomain.MockPostsMicroservice, postID uint64)
// 	type mockIsLike func(s *mockDomain.MockPostsMicroservice, userID, postID uint64)
//
// 	tests := []struct {
// 		name                 string
// 		userID               uint64
// 		postID               uint64
// 		mockPost             mockPost
// 		mockUser             mockUser
// 		mockSubscription     mockSubscription
// 		mockTagDeps          mockTagDeps
// 		mockTags             mockTags
// 		mockImg              mockImg
// 		mockLike             mockLike
// 		mockIsLike           mockIsLike
// 		response             models.Post
// 		responseErrorMessage string
// 	}{
// 		{
// 			name:   "OKBlur",
// 			userID: 200,
// 			postID: 1,
// 			mockPost: func(s *mockDomain.MockPostsMicroservice, id uint64) {
// 				s.EXPECT().GetPostByID(id).Return(models.Post{
// 					ID:     1,
// 					UserID: 200,
// 					Tier:   1,
// 				}, nil)
// 			},
// 			mockUser: func(s *mockDomain.MockUsersMicroservice, id uint64) {
// 				s.EXPECT().GetByID(id).Return(models.User{
// 					ID:     200,
// 					Avatar: "img",
// 				}, nil)
// 			},
// 			mockSubscription: func(s *mockDomain.MockSubscriptionMicroservice, userID, authorID uint64) {
// 				s.EXPECT().GetSubscriptionByUserAndAuthorID(userID, authorID).Return(models.AuthorSubscription{
// 					ID:       1,
// 					AuthorID: 200,
// 				}, nil)
// 			},
// 			mockTagDeps: func(s *mockDomain.MockPostsMicroservice, postID uint64) {
// 				s.EXPECT().GetTagDepsByPostId(postID).Return([]models.TagDep{
// 					{
// 						TagID:  1,
// 						PostID: 1,
// 					},
// 				}, nil)
// 			},
// 			mockTags: func(s *mockDomain.MockPostsMicroservice, tagID uint64) {
// 				s.EXPECT().GetTagById(tagID).Return(models.Tag{
// 					ID:      1,
// 					TagName: "tag",
// 				}, nil)
// 			},
// 			mockImg: func(s *mockDomain.MockImageUseCase, img string) {
// 				s.EXPECT().GetImage(img).Return("img", nil)
// 			},
// 			mockLike: func(s *mockDomain.MockPostsMicroservice, postID uint64) {
// 				s.EXPECT().GetAllLikesByPostID(postID).Return([]models.Like{
// 					{
// 						UserID: 200,
// 						PostID: 1,
// 					},
// 				}, nil)
// 			},
// 			mockIsLike: func(s *mockDomain.MockPostsMicroservice, userID, postID uint64) {
// 				s.EXPECT().GetLikeByUserAndPostID(userID, postID).Return(models.Like{
// 					UserID: 1,
// 					PostID: 1,
// 				}, nil)
// 			},
// 			response: models.Post{
// 				ID:     1,
// 				UserID: 200,
// 				Author: models.ResponseImageUsers{
// 					UserID:  200,
// 					ImgPath: "img",
// 				},
// 				LikesNum: 1,
// 				IsLiked:  true,
// 				Tags: []string{
// 					"tag",
// 				},
// 				Content: `<img src="https://wsrv.nl/?url=http://95.163.209.195:9000/eccbc87e4b5ce2fe28308fd9f2a7baf3/blur_f59a3f54-5fe5-44c8-abfc-72e92590d083.jpg" class="post-content__image">`,
// 				Tier:    1,
// 			},
// 		},
// 		{
// 			name:   "OKNoBlur",
// 			userID: 200,
// 			postID: 1,
// 			mockPost: func(s *mockDomain.MockPostsMicroservice, id uint64) {
// 				s.EXPECT().GetPostByID(id).Return(models.Post{
// 					ID:     1,
// 					UserID: 200,
// 				}, nil)
// 			},
// 			mockUser: func(s *mockDomain.MockUsersMicroservice, id uint64) {
// 				s.EXPECT().GetByID(id).Return(models.User{
// 					ID:     200,
// 					Avatar: "img",
// 				}, nil)
// 			},
// 			mockSubscription: func(s *mockDomain.MockSubscriptionMicroservice, userID, authorID uint64) {
// 				s.EXPECT().GetSubscriptionByUserAndAuthorID(userID, authorID).Return(models.AuthorSubscription{
// 					ID:       1,
// 					AuthorID: 200,
// 				}, nil)
// 			},
// 			mockTagDeps: func(s *mockDomain.MockPostsMicroservice, postID uint64) {
// 				s.EXPECT().GetTagDepsByPostId(postID).Return([]models.TagDep{
// 					{
// 						TagID:  1,
// 						PostID: 1,
// 					},
// 				}, nil)
// 			},
// 			mockTags: func(s *mockDomain.MockPostsMicroservice, tagID uint64) {
// 				s.EXPECT().GetTagById(tagID).Return(models.Tag{
// 					ID:      1,
// 					TagName: "tag",
// 				}, nil)
// 			},
// 			mockImg: func(s *mockDomain.MockImageUseCase, img string) {
// 				s.EXPECT().GetImage(img).Return("img", nil)
// 			},
// 			mockLike: func(s *mockDomain.MockPostsMicroservice, postID uint64) {
// 				s.EXPECT().GetAllLikesByPostID(postID).Return([]models.Like{
// 					{
// 						UserID: 200,
// 						PostID: 1,
// 					},
// 				}, nil)
// 			},
// 			mockIsLike: func(s *mockDomain.MockPostsMicroservice, userID, postID uint64) {
// 				s.EXPECT().GetLikeByUserAndPostID(userID, postID).Return(models.Like{
// 					UserID: 1,
// 					PostID: 1,
// 				}, nil)
// 			},
// 			response: models.Post{
// 				ID:     1,
// 				UserID: 200,
// 				Author: models.ResponseImageUsers{
// 					UserID:  200,
// 					ImgPath: "img",
// 				},
// 				LikesNum: 1,
// 				IsLiked:  true,
// 				Tags: []string{
// 					"tag",
// 				},
// 				IsAllowed: true,
// 				Content:   `<img src="https://wsrv.nl/?url=http://95.163.209.195:9000/eccbc87e4b5ce2fe28308fd9f2a7baf3/f59a3f54-5fe5-44c8-abfc-72e92590d083.jpg" class="post-content__image">`,
// 			},
// 		},
// 		{
// 			name:   "ErrGetByPostID",
// 			userID: 200,
// 			postID: 1,
// 			mockPost: func(s *mockDomain.MockPostsMicroservice, id uint64) {
// 				s.EXPECT().GetPostByID(id).Return(models.Post{}, errors.New("err"))
// 			},
// 			mockUser:             func(s *mockDomain.MockUsersMicroservice, id uint64) {},
// 			mockSubscription:     func(s *mockDomain.MockSubscriptionMicroservice, userID, authorID uint64) {},
// 			mockTagDeps:          func(s *mockDomain.MockPostsMicroservice, postID uint64) {},
// 			mockTags:             func(s *mockDomain.MockPostsMicroservice, tagID uint64) {},
// 			mockImg:              func(s *mockDomain.MockImageUseCase, img string) {},
// 			mockLike:             func(s *mockDomain.MockPostsMicroservice, postID uint64) {},
// 			mockIsLike:           func(s *mockDomain.MockPostsMicroservice, userID, postID uint64) {},
// 			responseErrorMessage: "err",
// 		},
// 		{
// 			name: "ErrGetByID",
// 			mockPost: func(s *mockDomain.MockPostsMicroservice, id uint64) {
// 				s.EXPECT().GetPostByID(id).Return(models.Post{}, nil)
// 			},
// 			mockUser: func(s *mockDomain.MockUsersMicroservice, id uint64) {
// 				s.EXPECT().GetByID(id).Return(models.User{}, errors.New("err"))
// 			},
// 			mockSubscription:     func(s *mockDomain.MockSubscriptionMicroservice, userID, authorID uint64) {},
// 			mockTagDeps:          func(s *mockDomain.MockPostsMicroservice, postID uint64) {},
// 			mockTags:             func(s *mockDomain.MockPostsMicroservice, tagID uint64) {},
// 			mockImg:              func(s *mockDomain.MockImageUseCase, img string) {},
// 			mockLike:             func(s *mockDomain.MockPostsMicroservice, postID uint64) {},
// 			mockIsLike:           func(s *mockDomain.MockPostsMicroservice, userID, postID uint64) {},
// 			responseErrorMessage: "err",
// 		},
// 		{
// 			name: "ErrGetSubscriptionByUserAndAuthorID",
// 			mockPost: func(s *mockDomain.MockPostsMicroservice, id uint64) {
// 				s.EXPECT().GetPostByID(id).Return(models.Post{}, nil)
// 			},
// 			mockUser: func(s *mockDomain.MockUsersMicroservice, id uint64) {
// 				s.EXPECT().GetByID(id).Return(models.User{}, nil)
// 			},
// 			mockSubscription: func(s *mockDomain.MockSubscriptionMicroservice, userID, authorID uint64) {
// 				s.EXPECT().GetSubscriptionByUserAndAuthorID(userID, authorID).Return(models.AuthorSubscription{}, errors.New("err"))
// 			},
// 			mockTagDeps:          func(s *mockDomain.MockPostsMicroservice, postID uint64) {},
// 			mockTags:             func(s *mockDomain.MockPostsMicroservice, tagID uint64) {},
// 			mockImg:              func(s *mockDomain.MockImageUseCase, img string) {},
// 			mockLike:             func(s *mockDomain.MockPostsMicroservice, postID uint64) {},
// 			mockIsLike:           func(s *mockDomain.MockPostsMicroservice, userID, postID uint64) {},
// 			responseErrorMessage: "err",
// 		},
// 	}
//
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()
//
// 			postMock := mockDomain.NewMockPostsMicroservice(ctrl)
// 			userMock := mockDomain.NewMockUsersMicroservice(ctrl)
// 			imgMock := mockDomain.NewMockImageUseCase(ctrl)
// 			subscriptionMock := mockDomain.NewMockSubscriptionMicroservice(ctrl)
//
// 			test.mockPost(postMock, test.postID)
// 			test.mockUser(userMock, test.userID)
// 			test.mockSubscription(subscriptionMock, test.userID, test.userID)
// 			test.mockTagDeps(postMock, test.postID)
// 			test.mockTags(postMock, uint64(1))
// 			test.mockImg(imgMock, "img")
// 			test.mockLike(postMock, test.postID)
// 			test.mockIsLike(postMock, test.userID, test.postID)
//
// 			usecase := New(postMock, userMock, imgMock, subscriptionMock)
//
// 			post, err := usecase.GetPostByID(test.postID, test.userID)
// 			if err != nil {
// 				require.Equal(t, test.responseErrorMessage, err.Error())
// 			}
// 			require.Equal(t, test.response, post)
// 		})
// 	}
// }

func TestUsecase_GetPostsByFilter(t *testing.T) {
	type mockPostSub func(s *mockDomain.MockPostsMicroservice, userID uint64)
	type mockUser func(s *mockDomain.MockUsersMicroservice, authorID uint64)
	type mockSub func(s *mockDomain.MockSubscriptionMicroservice, userID, authorID uint64)
	type mockGetUser func(s *mockDomain.MockUsersMicroservice, userID uint64)
	type mockTagDep func(s *mockDomain.MockPostsMicroservice, postID uint64)
	type mockTag func(s *mockDomain.MockPostsMicroservice, tagID uint64)
	type mockImg func(s *mockDomain.MockImageUseCase, img string)
	type mockLike func(s *mockDomain.MockPostsMicroservice, postID uint64)
	type mockIsLike func(s *mockDomain.MockPostsMicroservice, userID, postID uint64)
	type mockComment func(s *mockDomain.MockPostsMicroservice, postID uint64)
	type mockGetPost func(s *mockDomain.MockPostsMicroservice, postID uint64)

	tests := []struct {
		name                 string
		userID               uint64
		authorID             uint64
		postID               uint64
		mockPost             mockPostSub
		mockUser             mockUser
		mockSubscription     mockSub
		mockGetUser          mockGetUser
		mockTagDep           mockTagDep
		mockTag              mockTag
		mockImg              mockImg
		mockLike             mockLike
		mockIsLike           mockIsLike
		mockComment          mockComment
		mockGetPost          mockGetPost
		response             []models.Post
		responseErrorMessage string
	}{
		{
			name:     "OK",
			userID:   200,
			authorID: 100,
			postID:   1,
			mockPost: func(s *mockDomain.MockPostsMicroservice, userID uint64) {
				s.EXPECT().GetAllByUserID(userID).Return([]models.Post{
					{
						ID:     1,
						UserID: 200,
					},
				}, nil)
			},
			mockUser: func(s *mockDomain.MockUsersMicroservice, authorID uint64) {},
			mockSubscription: func(s *mockDomain.MockSubscriptionMicroservice, userID, authorID uint64) {
				s.EXPECT().GetSubscriptionByUserAndAuthorID(userID, authorID).Return(models.AuthorSubscription{
					AuthorID: 100,
				}, nil)
			},
			mockGetUser: func(s *mockDomain.MockUsersMicroservice, userID uint64) {
				s.EXPECT().GetByID(userID).Return(models.User{
					ID:     100,
					Avatar: "img",
				}, nil)
			},
			mockTagDep: func(s *mockDomain.MockPostsMicroservice, postID uint64) {
				s.EXPECT().GetTagDepsByPostId(postID).Return([]models.TagDep{
					{
						PostID: 1,
						TagID:  1,
					},
				}, nil)
			},
			mockTag: func(s *mockDomain.MockPostsMicroservice, tagID uint64) {
				s.EXPECT().GetTagById(tagID).Return(models.Tag{
					ID: 1,
				}, nil)
			},
			mockImg: func(s *mockDomain.MockImageUseCase, img string) {
				s.EXPECT().GetImage(img).Return("img", nil)
			},
			mockLike: func(s *mockDomain.MockPostsMicroservice, postID uint64) {
				s.EXPECT().GetAllLikesByPostID(postID).Return([]models.Like{
					{
						UserID: 100,
						PostID: 1,
					},
				}, nil)
			},
			mockIsLike: func(s *mockDomain.MockPostsMicroservice, userID, postID uint64) {
				s.EXPECT().GetLikeByUserAndPostID(userID, postID).Return(models.Like{
					UserID: 100,
					PostID: 1,
				}, nil)
			},
			mockComment: func(s *mockDomain.MockPostsMicroservice, postID uint64) {
				s.EXPECT().GetCommentsByPostID(postID).Return([]models.Comment{
					{
						UserID: 100,
						PostID: 1,
					},
				}, nil)
			},
			mockGetPost: func(s *mockDomain.MockPostsMicroservice, postID uint64) {
				s.EXPECT().GetPostByID(postID).Return(models.Post{
					ID:     1,
					UserID: 200,
				}, nil)
			},
			response: []models.Post{
				{
					ID:        1,
					UserID:    200,
					IsAllowed: true,
					Author: models.ResponseImageUsers{
						UserID:  100,
						ImgPath: "img",
					},
					Tags:        []string{""},
					LikesNum:    1,
					IsLiked:     true,
					CommentsNum: 1,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			postMock := mockDomain.NewMockPostsMicroservice(ctrl)
			userMock := mockDomain.NewMockUsersMicroservice(ctrl)
			imgMock := mockDomain.NewMockImageUseCase(ctrl)
			subscriptionMock := mockDomain.NewMockSubscriptionMicroservice(ctrl)

			test.mockPost(postMock, test.authorID)
			test.mockUser(userMock, test.authorID)
			test.mockSubscription(subscriptionMock, test.userID, test.authorID)
			test.mockGetUser(userMock, test.userID)
			test.mockTagDep(postMock, test.postID)
			test.mockTag(postMock, uint64(1))
			test.mockImg(imgMock, "img")
			test.mockLike(postMock, test.postID)
			test.mockIsLike(postMock, test.userID, test.postID)
			test.mockComment(postMock, test.postID)
			test.mockGetPost(postMock, test.postID)

			usecase := New(postMock, userMock, imgMock, subscriptionMock)

			post, err := usecase.GetPostsByFilter(test.userID, test.authorID)
			if err != nil {
				require.Equal(t, test.responseErrorMessage, err.Error())
			}
			require.Equal(t, test.response, post)
		})
	}
}

func TestUsecase_UpdateTags(t *testing.T) {
	type mockTagDep func(s *mockDomain.MockPostsMicroservice, postID uint64)
	type mockTag func(s *mockDomain.MockPostsMicroservice, tagID uint64)
	type mockTagName func(s *mockDomain.MockPostsMicroservice, tagName string)
	type mockCreateTag func(s *mockDomain.MockPostsMicroservice, tag string)
	type mockCreateTagDep func(s *mockDomain.MockPostsMicroservice, postID, tagID uint64)
	type mockDeleteTagDep func(s *mockDomain.MockPostsMicroservice, dep models.TagDep)

	tests := []struct {
		name                 string
		postID               uint64
		tagID                uint64
		tags                 []string
		mockTagDep           mockTagDep
		mockTag              mockTag
		mockTagName          mockTagName
		mockCreateTag        mockCreateTag
		mockCreateTagDep     mockCreateTagDep
		mockDeleteTagDep     mockDeleteTagDep
		responseErrorMessage string
	}{
		{
			name:   "OK",
			postID: 1,
			tagID:  1,
			tags:   []string{"tag"},
			mockTagDep: func(s *mockDomain.MockPostsMicroservice, postID uint64) {
				s.EXPECT().GetTagDepsByPostId(postID).Return([]models.TagDep{}, nil)
			},
			mockTag: func(s *mockDomain.MockPostsMicroservice, tagID uint64) {},
			mockTagName: func(s *mockDomain.MockPostsMicroservice, tagName string) {
				s.EXPECT().GetTagByName(tagName).Return(models.Tag{}, errors.New("error"))
			},
			mockCreateTag: func(s *mockDomain.MockPostsMicroservice, tag string) {
				s.EXPECT().CreateTag(tag).Return(uint64(1), nil)
			},
			mockCreateTagDep: func(s *mockDomain.MockPostsMicroservice, postID, tagID uint64) {
				s.EXPECT().CreateDepTag(postID, tagID).Return(nil)
			},
			mockDeleteTagDep: func(s *mockDomain.MockPostsMicroservice, dep models.TagDep) {},
		},
		{
			name:   "OKtag",
			postID: 1,
			tagID:  0,
			tags:   []string{"tag"},
			mockTagDep: func(s *mockDomain.MockPostsMicroservice, tagID uint64) {
				s.EXPECT().GetTagDepsByPostId(tagID).Return([]models.TagDep{
					{
						TagID:  0,
						PostID: 1,
					},
				}, nil)
			},
			mockTag: func(s *mockDomain.MockPostsMicroservice, tagID uint64) {
				s.EXPECT().GetTagById(tagID).Return(models.Tag{}, nil)
			},
			mockTagName: func(s *mockDomain.MockPostsMicroservice, tagName string) {
				s.EXPECT().GetTagByName(tagName).Return(models.Tag{}, nil)
			},
			mockCreateTag: func(s *mockDomain.MockPostsMicroservice, tag string) {},
			mockCreateTagDep: func(s *mockDomain.MockPostsMicroservice, postID, tagID uint64) {
				s.EXPECT().CreateDepTag(postID, tagID).Return(nil)
			},
			mockDeleteTagDep: func(s *mockDomain.MockPostsMicroservice, dep models.TagDep) {
				s.EXPECT().DeleteDepTag(dep).Return(nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			postMock := mockDomain.NewMockPostsMicroservice(ctrl)

			test.mockTagDep(postMock, test.postID)
			test.mockTag(postMock, test.tagID)
			test.mockTagName(postMock, test.tags[0])
			test.mockCreateTag(postMock, test.tags[0])
			test.mockCreateTagDep(postMock, test.postID, test.tagID)
			test.mockDeleteTagDep(postMock, models.TagDep{
				TagID:  test.tagID,
				PostID: test.postID,
			})

			usecase := New(postMock, nil, nil, nil)

			err := usecase.UpdateTags(test.tags, test.postID)
			if err != nil {
				require.Equal(t, test.responseErrorMessage, err.Error())
			}
		})
	}
}
