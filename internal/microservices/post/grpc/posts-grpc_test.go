package grpcPosts

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/post/protobuf"
	userProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	mock_domain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestConvertToModel(t *testing.T) {
	input := &protobuf.Post{
		ID:          1,
		UserID:      1,
		Content:     "test",
		Tier:        1,
		IsAllowed:   true,
		DateCreated: timestamppb.New(time.Time{}),
		Tags:        []string{"test"},
	}

	expected := models.Post{
		ID:          1,
		UserID:      1,
		Content:     "test",
		Tier:        1,
		IsAllowed:   true,
		DateCreated: time.Time{},
		Tags:        []string{"test"},
	}

	actual := ConvertToModel(input)

	assert.Equal(t, expected, actual)
}

func TestConvertToProto(t *testing.T) {
	input := models.Post{
		ID:          1,
		UserID:      1,
		Content:     "test",
		Tier:        1,
		IsAllowed:   true,
		DateCreated: time.Time{},
		Tags:        []string{"test"},
	}

	expected := &protobuf.Post{
		ID:          1,
		UserID:      1,
		Content:     "test",
		Tier:        1,
		IsAllowed:   true,
		DateCreated: timestamppb.New(time.Time{}),
		Tags:        []string{"test"},
		Author:      &userProto.LessUser{},
	}

	actual := ConvertToProto(input)

	assert.Equal(t, expected, actual)
}

func TestPostsService_GetAllByUserID(t *testing.T) {
	type mockBehaviorGetAllByUserID func(r *mock_domain.MockPostsRepository, authorID uint64)

	tests := []struct {
		name          string
		authorID      uint64
		mockBehavior  mockBehaviorGetAllByUserID
		expectedPosts *protobuf.PostArray
		expectedError string
	}{
		{
			name:     "OK",
			authorID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, authorID uint64) {
				r.EXPECT().GetAllByUserID(authorID).Return([]models.Post{
					{
						ID:          1,
						UserID:      1,
						Content:     "test",
						Tier:        1,
						IsAllowed:   true,
						DateCreated: time.Unix(0, 0),
						Tags:        []string{"test"},
					},
				}, nil)
			},
			expectedPosts: &protobuf.PostArray{
				Posts: []*protobuf.Post{
					{
						ID:        1,
						UserID:    1,
						Content:   "test",
						Tier:      1,
						IsAllowed: true,
						DateCreated: &timestamppb.Timestamp{
							Seconds: 0,
							Nanos:   0,
						},
						Tags:   []string{"test"},
						Author: &userProto.LessUser{},
					},
				},
			},
		},
		{
			name:     "Error",
			authorID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, authorID uint64) {
				r.EXPECT().GetAllByUserID(authorID).Return(nil, domain.ErrInternal)
			},
			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockPostsRepository(ctrl)
			test.mockBehavior(repo, test.authorID)

			s := New(repo)

			posts, err := s.GetAllByUserID(context.Background(), &userProto.UserID{UserId: test.authorID})

			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			} else {
				assert.Equal(t, test.expectedPosts, posts)
			}
		})
	}
}

func TestPostsService_GetPostByID(t *testing.T) {
	type mockBehaviorGetPostByID func(r *mock_domain.MockPostsRepository, postID uint64)

	tests := []struct {
		name          string
		postID        uint64
		mockBehavior  mockBehaviorGetPostByID
		expectedPosts *protobuf.Post
		expectedError string
	}{
		{
			name:   "OK",
			postID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, postID uint64) {
				r.EXPECT().GetPostByID(postID).Return(models.Post{
					ID:          1,
					UserID:      1,
					Content:     "test",
					Tier:        1,
					IsAllowed:   true,
					DateCreated: time.Unix(0, 0),
					Tags:        []string{"test"},
				}, nil)
			},
			expectedPosts: &protobuf.Post{
				ID:        1,
				UserID:    1,
				Content:   "test",
				Tier:      1,
				IsAllowed: true,
				DateCreated: &timestamppb.Timestamp{
					Seconds: 0,
					Nanos:   0,
				},
				Tags:   []string{"test"},
				Author: &userProto.LessUser{},
			},
		},
		{
			name:   "Error",
			postID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, postID uint64) {
				r.EXPECT().GetPostByID(postID).Return(models.Post{}, domain.ErrInternal)
			},
			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockPostsRepository(ctrl)
			test.mockBehavior(repo, test.postID)

			s := New(repo)

			post, err := s.GetPostByID(context.Background(), &protobuf.PostID{PostID: test.postID})

			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			} else {
				assert.Equal(t, test.expectedPosts, post)
			}
		})
	}
}

func TestPostsService_Create(t *testing.T) {
	type mockBehaviorCreate func(r *mock_domain.MockPostsRepository, post models.Post)

	tests := []struct {
		name          string
		post          models.Post
		mockBehavior  mockBehaviorCreate
		expectedError string
	}{
		{
			name: "OK",
			post: models.Post{
				ID:          1,
				UserID:      1,
				Content:     "test",
				Tier:        1,
				IsAllowed:   true,
				DateCreated: time.Time{},
				Tags:        []string{"test"},
			},
			mockBehavior: func(r *mock_domain.MockPostsRepository, post models.Post) {
				r.EXPECT().Create(post).Return(post, nil)
			},
		},
		{
			name: "Error",
			post: models.Post{
				ID:          1,
				UserID:      1,
				Content:     "test",
				Tier:        1,
				IsAllowed:   true,
				DateCreated: time.Time{},
				Tags:        []string{"test"},
			},
			mockBehavior: func(r *mock_domain.MockPostsRepository, post models.Post) {
				r.EXPECT().Create(post).Return(post, domain.ErrInternal)
			},
			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockPostsRepository(ctrl)
			test.mockBehavior(repo, test.post)

			s := New(repo)

			_, err := s.Create(context.Background(), &protobuf.Post{
				ID:        test.post.ID,
				UserID:    test.post.UserID,
				Content:   test.post.Content,
				Tier:      test.post.Tier,
				IsAllowed: test.post.IsAllowed,
				DateCreated: &timestamppb.Timestamp{
					Seconds: test.post.DateCreated.Unix(),
					Nanos:   int32(test.post.DateCreated.Nanosecond()),
				},
				Tags: test.post.Tags,
			})
			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			}
		})
	}
}

func TestPostsService_Update(t *testing.T) {
	type mockBehaviorUpdate func(r *mock_domain.MockPostsRepository, post models.Post)

	tests := []struct {
		name          string
		post          models.Post
		mockBehavior  mockBehaviorUpdate
		expectedError string
	}{
		{
			name: "OK",
			post: models.Post{
				ID:          1,
				UserID:      1,
				Content:     "test",
				Tier:        1,
				IsAllowed:   true,
				DateCreated: time.Time{},
				Tags:        []string{"test"},
			},
			mockBehavior: func(r *mock_domain.MockPostsRepository, post models.Post) {
				r.EXPECT().Update(post).Return(nil)
			},
		},
		{
			name: "Error",
			post: models.Post{
				ID:          1,
				UserID:      1,
				Content:     "test",
				Tier:        1,
				IsAllowed:   true,
				DateCreated: time.Time{},
				Tags:        []string{"test"},
			},
			mockBehavior: func(r *mock_domain.MockPostsRepository, post models.Post) {
				r.EXPECT().Update(post).Return(domain.ErrInternal)
			},
			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockPostsRepository(ctrl)
			test.mockBehavior(repo, test.post)

			s := New(repo)

			_, err := s.Update(context.Background(), &protobuf.Post{
				ID:        test.post.ID,
				UserID:    test.post.UserID,
				Content:   test.post.Content,
				Tier:      test.post.Tier,
				IsAllowed: test.post.IsAllowed,
				DateCreated: &timestamppb.Timestamp{
					Seconds: test.post.DateCreated.Unix(),
					Nanos:   int32(test.post.DateCreated.Nanosecond()),
				},
				Tags: test.post.Tags,
			})
			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			}
		})
	}
}

func TestPostsService_DeleteByID(t *testing.T) {
	type mockBehaviorDeleteByID func(r *mock_domain.MockPostsRepository, id uint64)

	tests := []struct {
		name          string
		id            uint64
		mockBehavior  mockBehaviorDeleteByID
		expectedError string
	}{
		{
			name: "OK",
			id:   1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, id uint64) {
				r.EXPECT().DeleteByID(id).Return(nil)
			},
		},
		{
			name: "Error",
			id:   1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, id uint64) {
				r.EXPECT().DeleteByID(id).Return(domain.ErrInternal)
			},
			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockPostsRepository(ctrl)
			test.mockBehavior(repo, test.id)

			s := New(repo)

			_, err := s.DeleteByID(context.Background(), &protobuf.PostID{PostID: test.id})
			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			}
		})
	}
}

func TestPostsService_GetPostsBySubscriptions(t *testing.T) {
	type mockBehaviorGetPostsBySubscriptions func(r *mock_domain.MockPostsRepository, userID uint64)

	tests := []struct {
		name          string
		userID        uint64
		mockBehavior  mockBehaviorGetPostsBySubscriptions
		expectedError string
	}{
		{
			name:   "OK",
			userID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, userID uint64) {
				r.EXPECT().GetPostsBySubscriptions(userID).Return([]models.Post{
					{
						ID:          1,
						UserID:      1,
						Content:     "test",
						Tier:        1,
						IsAllowed:   true,
						DateCreated: time.Time{},
						Tags:        []string{"test"},
					},
				}, nil)
			},
		},
		{
			name:   "Error",
			userID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, userID uint64) {
				r.EXPECT().GetPostsBySubscriptions(userID).Return(nil, domain.ErrInternal)
			},
			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockPostsRepository(ctrl)
			test.mockBehavior(repo, test.userID)

			s := New(repo)

			_, err := s.GetPostsBySubscriptions(context.Background(), &userProto.UserID{UserId: test.userID})
			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			}
		})
	}
}

func TestPostsService_GetLikeByUserAndPostID(t *testing.T) {
	type mockBehaviorGetLikeByUserAndPostID func(r *mock_domain.MockPostsRepository, userID, postID uint64)

	tests := []struct {
		name          string
		userID        uint64
		postID        uint64
		mockBehavior  mockBehaviorGetLikeByUserAndPostID
		expectedError string
	}{
		{
			name:   "OK",
			userID: 1,
			postID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, userID, postID uint64) {
				r.EXPECT().GetLikeByUserAndPostID(userID, postID).Return(models.Like{
					UserID: 1,
					PostID: 1,
				}, nil)
			},
		},
		{
			name:   "Error",
			userID: 1,
			postID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, userID, postID uint64) {
				r.EXPECT().GetLikeByUserAndPostID(userID, postID).Return(models.Like{}, domain.ErrInternal)
			},
			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockPostsRepository(ctrl)
			test.mockBehavior(repo, test.userID, test.postID)

			s := New(repo)

			_, err := s.GetLikeByUserAndPostID(context.Background(), &protobuf.PostUserIDs{
				UserID: test.userID,
				PostID: test.postID,
			})
			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			}
		})
	}
}

func TestPostsService_GetAllLikesByPostID(t *testing.T) {
	type mockBehaviorGetAllLikesByPostID func(r *mock_domain.MockPostsRepository, postID uint64)

	tests := []struct {
		name          string
		postID        uint64
		mockBehavior  mockBehaviorGetAllLikesByPostID
		expectedError string
	}{
		{
			name:   "OK",
			postID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, postID uint64) {
				r.EXPECT().GetAllLikesByPostID(postID).Return([]models.Like{
					{
						UserID: 1,
						PostID: 1,
					},
				}, nil)
			},
		},
		{
			name:   "Error",
			postID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, postID uint64) {
				r.EXPECT().GetAllLikesByPostID(postID).Return(nil, domain.ErrInternal)
			},
			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockPostsRepository(ctrl)
			test.mockBehavior(repo, test.postID)

			s := New(repo)

			_, err := s.GetAllLikesByPostID(context.Background(), &protobuf.PostID{PostID: test.postID})
			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			}
		})
	}
}

func TestPostsService_CreateLike(t *testing.T) {
	type mockBehaviorCreateLike func(r *mock_domain.MockPostsRepository, userID, postID uint64)

	tests := []struct {
		name          string
		userID        uint64
		postID        uint64
		mockBehavior  mockBehaviorCreateLike
		expectedError string
	}{
		{
			name:   "OK",
			userID: 1,
			postID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, userID, postID uint64) {
				r.EXPECT().CreateLike(userID, postID).Return(nil)
			},
		},
		{
			name:   "Error",
			userID: 1,
			postID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, userID, postID uint64) {
				r.EXPECT().CreateLike(userID, postID).Return(domain.ErrInternal)
			},
			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockPostsRepository(ctrl)
			test.mockBehavior(repo, test.userID, test.postID)

			s := New(repo)

			_, err := s.CreateLike(context.Background(), &protobuf.PostUserIDs{
				UserID: test.userID,
				PostID: test.postID,
			})
			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			}
		})
	}
}

func TestPostsService_DeleteLikeByID(t *testing.T) {
	type mockBehaviorDeleteLikeByID func(r *mock_domain.MockPostsRepository, userID, postID uint64)

	tests := []struct {
		name          string
		userID        uint64
		postID        uint64
		mockBehavior  mockBehaviorDeleteLikeByID
		expectedError string
	}{
		{
			name:   "OK",
			userID: 1,
			postID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, userID, postID uint64) {
				r.EXPECT().DeleteLikeByID(userID, postID).Return(nil)
			},
		},
		{
			name:   "Error",
			userID: 1,
			postID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, userID, postID uint64) {
				r.EXPECT().DeleteLikeByID(userID, postID).Return(domain.ErrInternal)
			},
			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockPostsRepository(ctrl)
			test.mockBehavior(repo, test.userID, test.postID)

			s := New(repo)

			_, err := s.DeleteLikeByID(context.Background(), &protobuf.PostUserIDs{
				UserID: test.userID,
				PostID: test.postID,
			})
			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			}
		})
	}
}

func TestPostsService_CreateTag(t *testing.T) {
	type mockBehaviorCreateTag func(r *mock_domain.MockPostsRepository, tag string, tagID uint64)

	tests := []struct {
		name          string
		tag           string
		tagID         uint64
		mockBehavior  mockBehaviorCreateTag
		expectedError string
	}{
		{
			name:  "OK",
			tag:   "tag",
			tagID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, tag string, tagID uint64) {
				r.EXPECT().CreateTag(tag).Return(tagID, nil)
			},
		},
		{
			name:  "Error",
			tag:   "tag",
			tagID: 0,
			mockBehavior: func(r *mock_domain.MockPostsRepository, tag string, tagID uint64) {
				r.EXPECT().CreateTag(tag).Return(tagID, domain.ErrInternal)
			},
			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockPostsRepository(ctrl)
			test.mockBehavior(repo, test.tag, test.tagID)

			s := New(repo)

			_, err := s.CreateTag(context.Background(), &protobuf.TagName{TagName: test.tag})
			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			}
		})
	}
}

func TestPostsService_GetTagById(t *testing.T) {
	type mockBehaviorGetTagById func(r *mock_domain.MockPostsRepository, tagID uint64)

	tests := []struct {
		name          string
		tagID         uint64
		mockBehavior  mockBehaviorGetTagById
		expectedError string
	}{
		{
			name:  "OK",
			tagID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, tagID uint64) {
				r.EXPECT().GetTagById(tagID).Return(models.Tag{
					ID: tagID,
				}, nil)
			},
		},
		{
			name:  "Error",
			tagID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, tagID uint64) {
				r.EXPECT().GetTagById(tagID).Return(models.Tag{}, domain.ErrInternal)
			},
			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockPostsRepository(ctrl)
			test.mockBehavior(repo, test.tagID)

			s := New(repo)

			_, err := s.GetTagById(context.Background(), &protobuf.TagID{TagID: test.tagID})
			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			}
		})
	}
}

func TestPostsService_GetTagByName(t *testing.T) {
	type mockBehaviorGetTagByName func(r *mock_domain.MockPostsRepository, tagName string)

	tests := []struct {
		name          string
		tagName       string
		mockBehavior  mockBehaviorGetTagByName
		expectedError string
	}{
		{
			name:    "OK",
			tagName: "tag",
			mockBehavior: func(r *mock_domain.MockPostsRepository, tagName string) {
				r.EXPECT().GetTagByName(tagName).Return(models.Tag{
					TagName: tagName,
				}, nil)
			},
		},
		{
			name:    "Error",
			tagName: "tag",
			mockBehavior: func(r *mock_domain.MockPostsRepository, tagName string) {
				r.EXPECT().GetTagByName(tagName).Return(models.Tag{}, domain.ErrInternal)
			},
			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockPostsRepository(ctrl)
			test.mockBehavior(repo, test.tagName)

			s := New(repo)

			_, err := s.GetTagByName(context.Background(), &protobuf.TagName{TagName: test.tagName})
			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			}
		})
	}
}

func TestPostsService_CreateDepTag(t *testing.T) {
	type mockBehaviorCreateDepTag func(r *mock_domain.MockPostsRepository, tagID uint64, postID uint64)

	tests := []struct {
		name          string
		tagID         uint64
		postID        uint64
		mockBehavior  mockBehaviorCreateDepTag
		expectedError string
	}{
		{
			name:   "OK",
			tagID:  1,
			postID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, tagID uint64, postID uint64) {
				r.EXPECT().CreateDepTag(tagID, postID).Return(nil)
			},
		},
		{
			name:   "Error",
			tagID:  1,
			postID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, tagID uint64, postID uint64) {
				r.EXPECT().CreateDepTag(tagID, postID).Return(domain.ErrInternal)
			},
			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockPostsRepository(ctrl)
			test.mockBehavior(repo, test.tagID, test.postID)

			s := New(repo)

			_, err := s.CreateDepTag(context.Background(), &protobuf.TagDep{TagID: test.tagID, PostID: test.postID})
			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			}
		})
	}
}

func TestPostsService_GetTagDepsByPostId(t *testing.T) {
	type mockBehaviorGetTagDepsByPostId func(r *mock_domain.MockPostsRepository, postID uint64)

	tests := []struct {
		name          string
		postID        uint64
		mockBehavior  mockBehaviorGetTagDepsByPostId
		expectedError string
	}{
		{
			name:   "OK",
			postID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, postID uint64) {
				r.EXPECT().GetTagDepsByPostId(postID).Return([]models.TagDep{
					{
						TagID:  1,
						PostID: 1,
					},
				}, nil)
			},
		},
		{
			name:   "Error",
			postID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, postID uint64) {
				r.EXPECT().GetTagDepsByPostId(postID).Return([]models.TagDep{}, domain.ErrInternal)
			},
			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockPostsRepository(ctrl)
			test.mockBehavior(repo, test.postID)

			s := New(repo)

			_, err := s.GetTagDepsByPostId(context.Background(), &protobuf.PostID{PostID: test.postID})
			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			}
		})
	}
}

func TestPostsService_DeleteDepTag(t *testing.T) {
	type mockBehaviorDeleteDepTag func(r *mock_domain.MockPostsRepository, tagID uint64, postID uint64)

	tests := []struct {
		name          string
		tagID         uint64
		postID        uint64
		mockBehavior  mockBehaviorDeleteDepTag
		expectedError string
	}{
		{
			name:   "OK",
			tagID:  1,
			postID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, tagID uint64, postID uint64) {
				r.EXPECT().DeleteDepTag(models.TagDep{
					PostID: postID,
					TagID:  tagID,
				}).Return(nil)
			},
		},
		{
			name:   "Error",
			tagID:  1,
			postID: 1,
			mockBehavior: func(r *mock_domain.MockPostsRepository, tagID uint64, postID uint64) {
				r.EXPECT().DeleteDepTag(models.TagDep{
					PostID: postID,
					TagID:  tagID,
				}).Return(domain.ErrInternal)
			},
			expectedError: status.Error(codes.Internal, domain.ErrInternal.Error()).Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_domain.NewMockPostsRepository(ctrl)
			test.mockBehavior(repo, test.tagID, test.postID)

			s := New(repo)

			_, err := s.DeleteDepTag(context.Background(), &protobuf.TagDep{TagID: test.tagID, PostID: test.postID})
			if err != nil {
				assert.Equal(t, test.expectedError, err.Error())
			}
		})
	}
}
