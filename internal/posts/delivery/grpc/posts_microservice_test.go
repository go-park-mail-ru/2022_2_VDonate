package postsMicroservice

import (
	"context"
	"errors"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/post/protobuf"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	userProto "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	mockDomain "github.com/go-park-mail-ru/2022_2_VDonate/internal/mocks/domain"
	"github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestPostsClient_GetAllByUserID(t *testing.T) {
	type mockGetAllByUserID func(r *mockDomain.MockPostsClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption)

	tests := []struct {
		name    string
		userID  uint64
		mock    mockGetAllByUserID
		want    []models.Post
		wantErr error
	}{
		{
			name:   "OK",
			userID: 1,
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption) {
				r.EXPECT().GetAllByUserID(c, in).Return(&protobuf.PostArray{Posts: []*protobuf.Post{
					{
						ID:          1,
						UserID:      1,
						DateCreated: timestamppb.New(time.Time{}),
					},
				}}, nil)
			},
			want: []models.Post{
				{
					ID:     1,
					UserID: 1,
				},
			},
			wantErr: nil,
		},
		{
			name:   "Error",
			userID: 1,
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption) {
				r.EXPECT().GetAllByUserID(c, in).Return(nil, status.Error(codes.Canceled, "canceled"))
			},
			want:    nil,
			wantErr: status.Error(codes.Canceled, "canceled"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockPostsClient(ctrl)

			test.mock(mock, context.Background(), &userProto.UserID{UserId: test.userID})

			m := &PostsMicroservice{
				client: mock,
			}
			got, err := m.GetAllByUserID(test.userID)

			require.Equal(t, got, test.want)
			if err != nil {
				require.Equal(t, err.Error(), test.wantErr.Error())
			}
		})
	}
}

func TestPostClient_GetPostByID(t *testing.T) {
	type mockGetPostByID func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.PostID, opts ...grpc.CallOption)

	tests := []struct {
		name    string
		postID  uint64
		mock    mockGetPostByID
		want    models.Post
		wantErr error
	}{
		{
			name:   "OK",
			postID: 1,
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.PostID, opts ...grpc.CallOption) {
				r.EXPECT().GetPostByID(c, in).Return(&protobuf.Post{
					ID:          1,
					UserID:      1,
					DateCreated: timestamppb.New(time.Time{}),
				}, nil)
			},
			want: models.Post{
				ID:     1,
				UserID: 1,
			},
			wantErr: nil,
		},
		{
			name:   "Error",
			postID: 1,
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.PostID, opts ...grpc.CallOption) {
				r.EXPECT().GetPostByID(c, in).Return(nil, status.Error(codes.Canceled, "canceled"))
			},
			want:    models.Post{},
			wantErr: status.Error(codes.Canceled, "canceled"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockPostsClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.PostID{PostID: test.postID})

			m := &PostsMicroservice{
				client: mock,
			}
			got, err := m.GetPostByID(test.postID)

			require.Equal(t, got, test.want)
			if err != nil {
				require.Equal(t, err.Error(), test.wantErr.Error())
			}
		})
	}
}

func TestPostClient_CreatePost(t *testing.T) {
	type mockCreatePost func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.Post, opts ...grpc.CallOption)

	tests := []struct {
		name    string
		post    models.Post
		mock    mockCreatePost
		want    uint64
		wantErr error
	}{
		{
			name: "OK",
			post: models.Post{
				ID:     1,
				UserID: 1,
			},
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.Post, opts ...grpc.CallOption) {
				r.EXPECT().Create(c, in).Return(&protobuf.PostID{
					PostID: 1,
				}, nil)
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "Error",
			post: models.Post{
				ID:     1,
				UserID: 1,
			},
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.Post, opts ...grpc.CallOption) {
				r.EXPECT().Create(c, in).Return(nil, status.Error(codes.Canceled, "canceled"))
			},
			wantErr: status.Error(codes.Canceled, "canceled"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockPostsClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.Post{
				ID:          test.post.ID,
				UserID:      test.post.UserID,
				DateCreated: timestamppb.New(time.Time{}),
				Author:      &userProto.LessUser{},
			})

			m := &PostsMicroservice{
				client: mock,
			}
			got, err := m.Create(test.post)

			require.Equal(t, got, test.want)
			if err != nil {
				require.Equal(t, err.Error(), test.wantErr.Error())
			}
		})
	}
}

func TestPostClient_UpdatePost(t *testing.T) {
	type mockUpdatePost func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.Post, opts ...grpc.CallOption)

	tests := []struct {
		name    string
		post    models.Post
		mock    mockUpdatePost
		wantErr error
	}{
		{
			name: "OK",
			post: models.Post{
				ID:     1,
				UserID: 1,
			},
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.Post, opts ...grpc.CallOption) {
				r.EXPECT().Update(c, in).Return(&emptypb.Empty{}, nil)
			},
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockPostsClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.Post{
				ID:          test.post.ID,
				UserID:      test.post.UserID,
				DateCreated: timestamppb.New(time.Time{}),
				Author:      &userProto.LessUser{},
			})

			m := &PostsMicroservice{
				client: mock,
			}
			err := m.Update(test.post)
			if err != nil {
				require.Equal(t, err.Error(), test.wantErr.Error())
			}
		})
	}
}

func TestPostClient_DeleteByID(t *testing.T) {
	type mockDeletePost func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.PostID, opts ...grpc.CallOption)

	tests := []struct {
		name    string
		postID  uint64
		mock    mockDeletePost
		wantErr error
	}{
		{
			name:   "OK",
			postID: 1,
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.PostID, opts ...grpc.CallOption) {
				r.EXPECT().DeleteByID(c, in).Return(&emptypb.Empty{}, nil)
			},
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockPostsClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.PostID{PostID: test.postID})

			m := &PostsMicroservice{
				client: mock,
			}
			err := m.DeleteByID(test.postID)
			if err != nil {
				require.Equal(t, err.Error(), test.wantErr.Error())
			}
		})
	}
}

func TestPostsClient_GetPostsBySubscriptions(t *testing.T) {
	type mockGetPostsBySubscriptions func(r *mockDomain.MockPostsClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption)

	tests := []struct {
		name    string
		id      uint64
		mock    mockGetPostsBySubscriptions
		want    []models.Post
		wantErr error
	}{
		{
			name: "OK",
			id:   1,
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption) {
				r.EXPECT().GetPostsBySubscriptions(c, in).Return(&protobuf.PostArray{
					Posts: []*protobuf.Post{
						{
							ID:          1,
							UserID:      1,
							DateCreated: timestamppb.New(time.Time{}),
						},
					},
				}, nil)
			},
			want: []models.Post{
				{
					ID:     1,
					UserID: 1,
				},
			},
			wantErr: nil,
		},
		{
			name: "Error",
			id:   1,
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *userProto.UserID, opts ...grpc.CallOption) {
				r.EXPECT().GetPostsBySubscriptions(c, in).Return(nil, errors.New("error"))
			},
			want:    nil,
			wantErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockPostsClient(ctrl)

			test.mock(mock, context.Background(), &userProto.UserID{UserId: test.id})

			m := &PostsMicroservice{
				client: mock,
			}
			got, err := m.GetPostsBySubscriptions(test.id)

			require.Equal(t, got, test.want)
			if err != nil {
				require.Equal(t, err.Error(), test.wantErr.Error())
			}
		})
	}
}

func TestPostsClient_GetLikeByUserAndPostID(t *testing.T) {
	type mockGetLikeByUserAndPostID func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.PostUserIDs, opts ...grpc.CallOption)

	tests := []struct {
		name    string
		like    models.Like
		mock    mockGetLikeByUserAndPostID
		want    models.Like
		wantErr error
	}{
		{
			name: "OK",
			like: models.Like{
				UserID: 1,
				PostID: 1,
			},
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.PostUserIDs, opts ...grpc.CallOption) {
				r.EXPECT().GetLikeByUserAndPostID(c, in).Return(&protobuf.Like{
					UserID: 1,
					PostID: 1,
				}, nil)
			},
			want: models.Like{
				UserID: 1,
				PostID: 1,
			},
			wantErr: nil,
		},
		{
			name: "Error",
			like: models.Like{
				UserID: 1,
				PostID: 1,
			},
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.PostUserIDs, opts ...grpc.CallOption) {
				r.EXPECT().GetLikeByUserAndPostID(c, in).Return(nil, errors.New("error"))
			},
			want:    models.Like{},
			wantErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockPostsClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.PostUserIDs{
				UserID: test.like.UserID,
				PostID: test.like.PostID,
			})

			m := &PostsMicroservice{
				client: mock,
			}
			got, err := m.GetLikeByUserAndPostID(test.like.PostID, test.like.UserID)

			require.Equal(t, got, test.want)
			if err != nil {
				require.Equal(t, err.Error(), test.wantErr.Error())
			}
		})
	}
}

func TestPostsClient_GetAllLikesByPostID(t *testing.T) {
	type mockGetLikesByPostID func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.PostID, opts ...grpc.CallOption)

	tests := []struct {
		name    string
		postID  uint64
		mock    mockGetLikesByPostID
		want    []models.Like
		wantErr error
	}{
		{
			name:   "OK",
			postID: 1,
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.PostID, opts ...grpc.CallOption) {
				r.EXPECT().GetAllLikesByPostID(c, in).Return(&protobuf.Likes{
					Likes: []*protobuf.Like{
						{
							UserID: 1,
							PostID: 1,
						},
					},
				}, nil)
			},
			want: []models.Like{
				{
					UserID: 1,
					PostID: 1,
				},
			},
			wantErr: nil,
		},
		{
			name:   "Error",
			postID: 1,
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.PostID, opts ...grpc.CallOption) {
				r.EXPECT().GetAllLikesByPostID(c, in).Return(nil, errors.New("error"))
			},
			want:    nil,
			wantErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockPostsClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.PostID{PostID: test.postID})

			m := &PostsMicroservice{
				client: mock,
			}
			got, err := m.GetAllLikesByPostID(test.postID)

			require.Equal(t, got, test.want)
			if err != nil {
				require.Equal(t, err.Error(), test.wantErr.Error())
			}
		})
	}
}

func TestPostsClient_CreateLike(t *testing.T) {
	type mockCreateLike func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.PostUserIDs, opts ...grpc.CallOption)

	tests := []struct {
		name    string
		like    models.Like
		mock    mockCreateLike
		want    models.Like
		wantErr error
	}{
		{
			name: "OK",
			like: models.Like{
				UserID: 1,
				PostID: 1,
			},
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.PostUserIDs, opts ...grpc.CallOption) {
				r.EXPECT().CreateLike(c, in).Return(&emptypb.Empty{}, nil)
			},
			want: models.Like{
				UserID: 1,
				PostID: 1,
			},
			wantErr: nil,
		},
		{
			name: "Error",
			like: models.Like{
				UserID: 1,
				PostID: 1,
			},
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.PostUserIDs, opts ...grpc.CallOption) {
				r.EXPECT().CreateLike(c, in).Return(nil, errors.New("error"))
			},
			want:    models.Like{},
			wantErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockPostsClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.PostUserIDs{
				PostID: test.like.PostID,
				UserID: test.like.UserID,
			})

			m := &PostsMicroservice{
				client: mock,
			}
			err := m.CreateLike(test.like.UserID, test.like.PostID)
			if err != nil {
				require.Equal(t, err.Error(), test.wantErr.Error())
			}
		})
	}
}

func TestPostsClient_DeleteLikeByID(t *testing.T) {
	type mockDeleteLike func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.PostUserIDs, opts ...grpc.CallOption)

	tests := []struct {
		name    string
		like    models.Like
		mock    mockDeleteLike
		wantErr error
	}{
		{
			name: "OK",
			like: models.Like{
				UserID: 1,
				PostID: 1,
			},
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.PostUserIDs, opts ...grpc.CallOption) {
				r.EXPECT().DeleteLikeByID(c, in).Return(&emptypb.Empty{}, nil)
			},
			wantErr: nil,
		},
		{
			name: "Error",
			like: models.Like{
				UserID: 1,
				PostID: 1,
			},
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.PostUserIDs, opts ...grpc.CallOption) {
				r.EXPECT().DeleteLikeByID(c, in).Return(nil, errors.New("error"))
			},
			wantErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockPostsClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.PostUserIDs{
				PostID: test.like.PostID,
				UserID: test.like.UserID,
			})

			m := &PostsMicroservice{
				client: mock,
			}
			err := m.DeleteLikeByID(test.like.UserID, test.like.PostID)
			if err != nil {
				require.Equal(t, err.Error(), test.wantErr.Error())
			}
		})
	}
}

func TestPostsClient_CreateTag(t *testing.T) {
	type mockCreateTag func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.TagName, opts ...grpc.CallOption)

	tests := []struct {
		name    string
		tag     models.Tag
		mock    mockCreateTag
		want    uint64
		wantErr error
	}{
		{
			name: "OK",
			tag: models.Tag{
				TagName: "tag",
			},
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.TagName, opts ...grpc.CallOption) {
				r.EXPECT().CreateTag(c, in).Return(&protobuf.TagID{TagID: 1}, nil)
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "Error",
			tag: models.Tag{
				TagName: "tag",
			},
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.TagName, opts ...grpc.CallOption) {
				r.EXPECT().CreateTag(c, in).Return(nil, errors.New("error"))
			},
			wantErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockPostsClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.TagName{TagName: test.tag.TagName})

			m := &PostsMicroservice{
				client: mock,
			}
			got, err := m.CreateTag(test.tag.TagName)
			require.Equal(t, got, test.want)
			if err != nil {
				require.Equal(t, err.Error(), test.wantErr.Error())
			}
		})
	}
}

func TestPostsClient_GetTagByID(t *testing.T) {
	type mockGetTagByID func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.TagID, opts ...grpc.CallOption)

	tests := []struct {
		name    string
		tagID   uint64
		mock    mockGetTagByID
		want    models.Tag
		wantErr error
	}{
		{
			name:  "OK",
			tagID: 1,
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.TagID, opts ...grpc.CallOption) {
				r.EXPECT().GetTagById(c, in).Return(&protobuf.Tag{TagName: "tag", Id: 1}, nil)
			},
			want: models.Tag{
				ID:      1,
				TagName: "tag",
			},
			wantErr: nil,
		},
		{
			name:  "Error",
			tagID: 1,
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.TagID, opts ...grpc.CallOption) {
				r.EXPECT().GetTagById(c, in).Return(nil, errors.New("error"))
			},
			want:    models.Tag{},
			wantErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockPostsClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.TagID{TagID: test.tagID})

			m := &PostsMicroservice{
				client: mock,
			}
			got, err := m.GetTagById(test.tagID)
			require.Equal(t, got, test.want)
			if err != nil {
				require.Equal(t, err.Error(), test.wantErr.Error())
			}
		})
	}
}

func TestPostsClient_GetTagByName(t *testing.T) {
	type mockGetTagByName func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.TagName, opts ...grpc.CallOption)

	tests := []struct {
		name    string
		tagName string
		mock    mockGetTagByName
		want    models.Tag
		wantErr error
	}{
		{
			name:    "OK",
			tagName: "tag",
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.TagName, opts ...grpc.CallOption) {
				r.EXPECT().GetTagByName(c, in).Return(&protobuf.Tag{TagName: "tag", Id: 1}, nil)
			},
			want: models.Tag{
				ID:      1,
				TagName: "tag",
			},
			wantErr: nil,
		},
		{
			name:    "Error",
			tagName: "tag",
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.TagName, opts ...grpc.CallOption) {
				r.EXPECT().GetTagByName(c, in).Return(nil, errors.New("error"))
			},
			want:    models.Tag{},
			wantErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockPostsClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.TagName{TagName: test.tagName})

			m := &PostsMicroservice{
				client: mock,
			}
			got, err := m.GetTagByName(test.tagName)
			require.Equal(t, got, test.want)
			if err != nil {
				require.Equal(t, err.Error(), test.wantErr.Error())
			}
		})
	}
}

func TestPostsClient_CreateDepTag(t *testing.T) {
	type mockCreateDepTag func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.TagDep, opts ...grpc.CallOption)

	tests := []struct {
		name    string
		tagID   uint64
		postID  uint64
		mock    mockCreateDepTag
		want    models.TagDep
		wantErr error
	}{
		{
			name:   "OK",
			tagID:  1,
			postID: 1,
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.TagDep, opts ...grpc.CallOption) {
				r.EXPECT().CreateDepTag(c, in).Return(&emptypb.Empty{}, nil)
			},
			want: models.TagDep{
				TagID:  1,
				PostID: 1,
			},
			wantErr: nil,
		},
		{
			name:   "Error",
			tagID:  1,
			postID: 1,
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.TagDep, opts ...grpc.CallOption) {
				r.EXPECT().CreateDepTag(c, in).Return(nil, errors.New("error"))
			},
			want:    models.TagDep{},
			wantErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockPostsClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.TagDep{TagID: test.tagID, PostID: test.postID})

			m := &PostsMicroservice{
				client: mock,
			}
			err := m.CreateDepTag(test.postID, test.tagID)
			if err != nil {
				require.Equal(t, err.Error(), test.wantErr.Error())
			}
		})
	}
}

func TestPostsClient_GetTagDepsByPostId(t *testing.T) {
	type mockGetTagDepsByPostId func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.PostID, opts ...grpc.CallOption)

	tests := []struct {
		name    string
		postID  uint64
		mock    mockGetTagDepsByPostId
		want    []models.TagDep
		wantErr error
	}{
		{
			name:   "OK",
			postID: 1,
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.PostID, opts ...grpc.CallOption) {
				r.EXPECT().GetTagDepsByPostId(c, in).Return(&protobuf.TagDeps{TagDeps: []*protobuf.TagDep{{TagID: 1, PostID: 1}}}, nil)
			},
			want: []models.TagDep{
				{
					TagID:  1,
					PostID: 1,
				},
			},
			wantErr: nil,
		},
		{
			name:   "Error",
			postID: 1,
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.PostID, opts ...grpc.CallOption) {
				r.EXPECT().GetTagDepsByPostId(c, in).Return(nil, errors.New("error"))
			},
			want:    nil,
			wantErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockPostsClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.PostID{PostID: test.postID})

			m := &PostsMicroservice{
				client: mock,
			}
			got, err := m.GetTagDepsByPostId(test.postID)
			require.Equal(t, got, test.want)
			if err != nil {
				require.Equal(t, err.Error(), test.wantErr.Error())
			}
		})
	}
}

func TestPostsClient_DeleteDepTag(t *testing.T) {
	type mockDeleteDepTag func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.TagDep, opts ...grpc.CallOption)

	tests := []struct {
		name    string
		tagID   uint64
		postID  uint64
		mock    mockDeleteDepTag
		want    models.TagDep
		wantErr error
	}{
		{
			name:   "OK",
			tagID:  1,
			postID: 1,
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.TagDep, opts ...grpc.CallOption) {
				r.EXPECT().DeleteDepTag(c, in).Return(&emptypb.Empty{}, nil)
			},
			want: models.TagDep{
				TagID:  1,
				PostID: 1,
			},
			wantErr: nil,
		},
		{
			name:   "Error",
			tagID:  1,
			postID: 1,
			mock: func(r *mockDomain.MockPostsClient, c context.Context, in *protobuf.TagDep, opts ...grpc.CallOption) {
				r.EXPECT().DeleteDepTag(c, in).Return(nil, errors.New("error"))
			},
			want:    models.TagDep{},
			wantErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mockDomain.NewMockPostsClient(ctrl)

			test.mock(mock, context.Background(), &protobuf.TagDep{TagID: test.tagID, PostID: test.postID})

			m := &PostsMicroservice{
				client: mock,
			}
			err := m.DeleteDepTag(models.TagDep{
				TagID:  test.tagID,
				PostID: test.postID,
			})
			if err != nil {
				require.Equal(t, err.Error(), test.wantErr.Error())
			}
		})
	}
}
