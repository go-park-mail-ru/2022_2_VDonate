// Code generated by MockGen. DO NOT EDIT.
// Source: internal/microservices/post/protobuf/posts_grpc.pb.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	reflect "reflect"

	protobuf "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/post/protobuf"
	protobuf0 "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockPostsClient is a mock of PostsClient interface.
type MockPostsClient struct {
	ctrl     *gomock.Controller
	recorder *MockPostsClientMockRecorder
}

// MockPostsClientMockRecorder is the mock recorder for MockPostsClient.
type MockPostsClientMockRecorder struct {
	mock *MockPostsClient
}

// NewMockPostsClient creates a new mock instance.
func NewMockPostsClient(ctrl *gomock.Controller) *MockPostsClient {
	mock := &MockPostsClient{ctrl: ctrl}
	mock.recorder = &MockPostsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostsClient) EXPECT() *MockPostsClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPostsClient) Create(ctx context.Context, in *protobuf.Post, opts ...grpc.CallOption) (*protobuf.Post, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*protobuf.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPostsClientMockRecorder) Create(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPostsClient)(nil).Create), varargs...)
}

// CreateComment mocks base method.
func (m *MockPostsClient) CreateComment(ctx context.Context, in *protobuf.Comment, opts ...grpc.CallOption) (*protobuf.CommentPairIdDate, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateComment", varargs...)
	ret0, _ := ret[0].(*protobuf.CommentPairIdDate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockPostsClientMockRecorder) CreateComment(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockPostsClient)(nil).CreateComment), varargs...)
}

// CreateDepTag mocks base method.
func (m *MockPostsClient) CreateDepTag(ctx context.Context, in *protobuf.TagDep, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateDepTag", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDepTag indicates an expected call of CreateDepTag.
func (mr *MockPostsClientMockRecorder) CreateDepTag(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDepTag", reflect.TypeOf((*MockPostsClient)(nil).CreateDepTag), varargs...)
}

// CreateLike mocks base method.
func (m *MockPostsClient) CreateLike(ctx context.Context, in *protobuf.PostUserIDs, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateLike", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLike indicates an expected call of CreateLike.
func (mr *MockPostsClientMockRecorder) CreateLike(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLike", reflect.TypeOf((*MockPostsClient)(nil).CreateLike), varargs...)
}

// CreateTag mocks base method.
func (m *MockPostsClient) CreateTag(ctx context.Context, in *protobuf.TagName, opts ...grpc.CallOption) (*protobuf.TagID, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTag", varargs...)
	ret0, _ := ret[0].(*protobuf.TagID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTag indicates an expected call of CreateTag.
func (mr *MockPostsClientMockRecorder) CreateTag(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTag", reflect.TypeOf((*MockPostsClient)(nil).CreateTag), varargs...)
}

// DeleteByID mocks base method.
func (m *MockPostsClient) DeleteByID(ctx context.Context, in *protobuf.PostID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteByID", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockPostsClientMockRecorder) DeleteByID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockPostsClient)(nil).DeleteByID), varargs...)
}

// DeleteCommentByID mocks base method.
func (m *MockPostsClient) DeleteCommentByID(ctx context.Context, in *protobuf.CommentID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteCommentByID", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCommentByID indicates an expected call of DeleteCommentByID.
func (mr *MockPostsClientMockRecorder) DeleteCommentByID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCommentByID", reflect.TypeOf((*MockPostsClient)(nil).DeleteCommentByID), varargs...)
}

// DeleteDepTag mocks base method.
func (m *MockPostsClient) DeleteDepTag(ctx context.Context, in *protobuf.TagDep, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteDepTag", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDepTag indicates an expected call of DeleteDepTag.
func (mr *MockPostsClientMockRecorder) DeleteDepTag(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDepTag", reflect.TypeOf((*MockPostsClient)(nil).DeleteDepTag), varargs...)
}

// DeleteLikeByID mocks base method.
func (m *MockPostsClient) DeleteLikeByID(ctx context.Context, in *protobuf.PostUserIDs, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteLikeByID", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLikeByID indicates an expected call of DeleteLikeByID.
func (mr *MockPostsClientMockRecorder) DeleteLikeByID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLikeByID", reflect.TypeOf((*MockPostsClient)(nil).DeleteLikeByID), varargs...)
}

// GetAllByUserID mocks base method.
func (m *MockPostsClient) GetAllByUserID(ctx context.Context, in *protobuf0.UserID, opts ...grpc.CallOption) (*protobuf.PostArray, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllByUserID", varargs...)
	ret0, _ := ret[0].(*protobuf.PostArray)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByUserID indicates an expected call of GetAllByUserID.
func (mr *MockPostsClientMockRecorder) GetAllByUserID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByUserID", reflect.TypeOf((*MockPostsClient)(nil).GetAllByUserID), varargs...)
}

// GetAllLikesByPostID mocks base method.
func (m *MockPostsClient) GetAllLikesByPostID(ctx context.Context, in *protobuf.PostID, opts ...grpc.CallOption) (*protobuf.Likes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllLikesByPostID", varargs...)
	ret0, _ := ret[0].(*protobuf.Likes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllLikesByPostID indicates an expected call of GetAllLikesByPostID.
func (mr *MockPostsClientMockRecorder) GetAllLikesByPostID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllLikesByPostID", reflect.TypeOf((*MockPostsClient)(nil).GetAllLikesByPostID), varargs...)
}

// GetCommentByID mocks base method.
func (m *MockPostsClient) GetCommentByID(ctx context.Context, in *protobuf.CommentID, opts ...grpc.CallOption) (*protobuf.Comment, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCommentByID", varargs...)
	ret0, _ := ret[0].(*protobuf.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentByID indicates an expected call of GetCommentByID.
func (mr *MockPostsClientMockRecorder) GetCommentByID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentByID", reflect.TypeOf((*MockPostsClient)(nil).GetCommentByID), varargs...)
}

// GetCommentsByPostID mocks base method.
func (m *MockPostsClient) GetCommentsByPostID(ctx context.Context, in *protobuf.PostID, opts ...grpc.CallOption) (*protobuf.CommentArray, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCommentsByPostID", varargs...)
	ret0, _ := ret[0].(*protobuf.CommentArray)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentsByPostID indicates an expected call of GetCommentsByPostID.
func (mr *MockPostsClientMockRecorder) GetCommentsByPostID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentsByPostID", reflect.TypeOf((*MockPostsClient)(nil).GetCommentsByPostID), varargs...)
}

// GetLikeByUserAndPostID mocks base method.
func (m *MockPostsClient) GetLikeByUserAndPostID(ctx context.Context, in *protobuf.PostUserIDs, opts ...grpc.CallOption) (*protobuf.Like, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLikeByUserAndPostID", varargs...)
	ret0, _ := ret[0].(*protobuf.Like)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLikeByUserAndPostID indicates an expected call of GetLikeByUserAndPostID.
func (mr *MockPostsClientMockRecorder) GetLikeByUserAndPostID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLikeByUserAndPostID", reflect.TypeOf((*MockPostsClient)(nil).GetLikeByUserAndPostID), varargs...)
}

// GetPostByID mocks base method.
func (m *MockPostsClient) GetPostByID(ctx context.Context, in *protobuf.PostID, opts ...grpc.CallOption) (*protobuf.Post, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPostByID", varargs...)
	ret0, _ := ret[0].(*protobuf.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostByID indicates an expected call of GetPostByID.
func (mr *MockPostsClientMockRecorder) GetPostByID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostByID", reflect.TypeOf((*MockPostsClient)(nil).GetPostByID), varargs...)
}

// GetPostsBySubscriptions mocks base method.
func (m *MockPostsClient) GetPostsBySubscriptions(ctx context.Context, in *protobuf0.UserID, opts ...grpc.CallOption) (*protobuf.PostArray, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPostsBySubscriptions", varargs...)
	ret0, _ := ret[0].(*protobuf.PostArray)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsBySubscriptions indicates an expected call of GetPostsBySubscriptions.
func (mr *MockPostsClientMockRecorder) GetPostsBySubscriptions(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsBySubscriptions", reflect.TypeOf((*MockPostsClient)(nil).GetPostsBySubscriptions), varargs...)
}

// GetTagById mocks base method.
func (m *MockPostsClient) GetTagById(ctx context.Context, in *protobuf.TagID, opts ...grpc.CallOption) (*protobuf.Tag, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTagById", varargs...)
	ret0, _ := ret[0].(*protobuf.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagById indicates an expected call of GetTagById.
func (mr *MockPostsClientMockRecorder) GetTagById(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagById", reflect.TypeOf((*MockPostsClient)(nil).GetTagById), varargs...)
}

// GetTagByName mocks base method.
func (m *MockPostsClient) GetTagByName(ctx context.Context, in *protobuf.TagName, opts ...grpc.CallOption) (*protobuf.Tag, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTagByName", varargs...)
	ret0, _ := ret[0].(*protobuf.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagByName indicates an expected call of GetTagByName.
func (mr *MockPostsClientMockRecorder) GetTagByName(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagByName", reflect.TypeOf((*MockPostsClient)(nil).GetTagByName), varargs...)
}

// GetTagDepsByPostId mocks base method.
func (m *MockPostsClient) GetTagDepsByPostId(ctx context.Context, in *protobuf.PostID, opts ...grpc.CallOption) (*protobuf.TagDeps, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTagDepsByPostId", varargs...)
	ret0, _ := ret[0].(*protobuf.TagDeps)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagDepsByPostId indicates an expected call of GetTagDepsByPostId.
func (mr *MockPostsClientMockRecorder) GetTagDepsByPostId(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagDepsByPostId", reflect.TypeOf((*MockPostsClient)(nil).GetTagDepsByPostId), varargs...)
}

// Update mocks base method.
func (m *MockPostsClient) Update(ctx context.Context, in *protobuf.Post, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPostsClientMockRecorder) Update(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPostsClient)(nil).Update), varargs...)
}

// UpdateComment mocks base method.
func (m *MockPostsClient) UpdateComment(ctx context.Context, in *protobuf.Comment, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateComment", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateComment indicates an expected call of UpdateComment.
func (mr *MockPostsClientMockRecorder) UpdateComment(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComment", reflect.TypeOf((*MockPostsClient)(nil).UpdateComment), varargs...)
}

// MockPostsServer is a mock of PostsServer interface.
type MockPostsServer struct {
	ctrl     *gomock.Controller
	recorder *MockPostsServerMockRecorder
}

// MockPostsServerMockRecorder is the mock recorder for MockPostsServer.
type MockPostsServerMockRecorder struct {
	mock *MockPostsServer
}

// NewMockPostsServer creates a new mock instance.
func NewMockPostsServer(ctrl *gomock.Controller) *MockPostsServer {
	mock := &MockPostsServer{ctrl: ctrl}
	mock.recorder = &MockPostsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostsServer) EXPECT() *MockPostsServerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPostsServer) Create(arg0 context.Context, arg1 *protobuf.Post) (*protobuf.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*protobuf.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPostsServerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPostsServer)(nil).Create), arg0, arg1)
}

// CreateComment mocks base method.
func (m *MockPostsServer) CreateComment(arg0 context.Context, arg1 *protobuf.Comment) (*protobuf.CommentPairIdDate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", arg0, arg1)
	ret0, _ := ret[0].(*protobuf.CommentPairIdDate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockPostsServerMockRecorder) CreateComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockPostsServer)(nil).CreateComment), arg0, arg1)
}

// CreateDepTag mocks base method.
func (m *MockPostsServer) CreateDepTag(arg0 context.Context, arg1 *protobuf.TagDep) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDepTag", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDepTag indicates an expected call of CreateDepTag.
func (mr *MockPostsServerMockRecorder) CreateDepTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDepTag", reflect.TypeOf((*MockPostsServer)(nil).CreateDepTag), arg0, arg1)
}

// CreateLike mocks base method.
func (m *MockPostsServer) CreateLike(arg0 context.Context, arg1 *protobuf.PostUserIDs) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLike", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLike indicates an expected call of CreateLike.
func (mr *MockPostsServerMockRecorder) CreateLike(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLike", reflect.TypeOf((*MockPostsServer)(nil).CreateLike), arg0, arg1)
}

// CreateTag mocks base method.
func (m *MockPostsServer) CreateTag(arg0 context.Context, arg1 *protobuf.TagName) (*protobuf.TagID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTag", arg0, arg1)
	ret0, _ := ret[0].(*protobuf.TagID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTag indicates an expected call of CreateTag.
func (mr *MockPostsServerMockRecorder) CreateTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTag", reflect.TypeOf((*MockPostsServer)(nil).CreateTag), arg0, arg1)
}

// DeleteByID mocks base method.
func (m *MockPostsServer) DeleteByID(arg0 context.Context, arg1 *protobuf.PostID) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockPostsServerMockRecorder) DeleteByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockPostsServer)(nil).DeleteByID), arg0, arg1)
}

// DeleteCommentByID mocks base method.
func (m *MockPostsServer) DeleteCommentByID(arg0 context.Context, arg1 *protobuf.CommentID) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCommentByID", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCommentByID indicates an expected call of DeleteCommentByID.
func (mr *MockPostsServerMockRecorder) DeleteCommentByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCommentByID", reflect.TypeOf((*MockPostsServer)(nil).DeleteCommentByID), arg0, arg1)
}

// DeleteDepTag mocks base method.
func (m *MockPostsServer) DeleteDepTag(arg0 context.Context, arg1 *protobuf.TagDep) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDepTag", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDepTag indicates an expected call of DeleteDepTag.
func (mr *MockPostsServerMockRecorder) DeleteDepTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDepTag", reflect.TypeOf((*MockPostsServer)(nil).DeleteDepTag), arg0, arg1)
}

// DeleteLikeByID mocks base method.
func (m *MockPostsServer) DeleteLikeByID(arg0 context.Context, arg1 *protobuf.PostUserIDs) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLikeByID", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLikeByID indicates an expected call of DeleteLikeByID.
func (mr *MockPostsServerMockRecorder) DeleteLikeByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLikeByID", reflect.TypeOf((*MockPostsServer)(nil).DeleteLikeByID), arg0, arg1)
}

// GetAllByUserID mocks base method.
func (m *MockPostsServer) GetAllByUserID(arg0 context.Context, arg1 *protobuf0.UserID) (*protobuf.PostArray, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByUserID", arg0, arg1)
	ret0, _ := ret[0].(*protobuf.PostArray)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByUserID indicates an expected call of GetAllByUserID.
func (mr *MockPostsServerMockRecorder) GetAllByUserID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByUserID", reflect.TypeOf((*MockPostsServer)(nil).GetAllByUserID), arg0, arg1)
}

// GetAllLikesByPostID mocks base method.
func (m *MockPostsServer) GetAllLikesByPostID(arg0 context.Context, arg1 *protobuf.PostID) (*protobuf.Likes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllLikesByPostID", arg0, arg1)
	ret0, _ := ret[0].(*protobuf.Likes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllLikesByPostID indicates an expected call of GetAllLikesByPostID.
func (mr *MockPostsServerMockRecorder) GetAllLikesByPostID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllLikesByPostID", reflect.TypeOf((*MockPostsServer)(nil).GetAllLikesByPostID), arg0, arg1)
}

// GetCommentByID mocks base method.
func (m *MockPostsServer) GetCommentByID(arg0 context.Context, arg1 *protobuf.CommentID) (*protobuf.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentByID", arg0, arg1)
	ret0, _ := ret[0].(*protobuf.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentByID indicates an expected call of GetCommentByID.
func (mr *MockPostsServerMockRecorder) GetCommentByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentByID", reflect.TypeOf((*MockPostsServer)(nil).GetCommentByID), arg0, arg1)
}

// GetCommentsByPostID mocks base method.
func (m *MockPostsServer) GetCommentsByPostID(arg0 context.Context, arg1 *protobuf.PostID) (*protobuf.CommentArray, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentsByPostID", arg0, arg1)
	ret0, _ := ret[0].(*protobuf.CommentArray)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentsByPostID indicates an expected call of GetCommentsByPostID.
func (mr *MockPostsServerMockRecorder) GetCommentsByPostID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentsByPostID", reflect.TypeOf((*MockPostsServer)(nil).GetCommentsByPostID), arg0, arg1)
}

// GetLikeByUserAndPostID mocks base method.
func (m *MockPostsServer) GetLikeByUserAndPostID(arg0 context.Context, arg1 *protobuf.PostUserIDs) (*protobuf.Like, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLikeByUserAndPostID", arg0, arg1)
	ret0, _ := ret[0].(*protobuf.Like)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLikeByUserAndPostID indicates an expected call of GetLikeByUserAndPostID.
func (mr *MockPostsServerMockRecorder) GetLikeByUserAndPostID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLikeByUserAndPostID", reflect.TypeOf((*MockPostsServer)(nil).GetLikeByUserAndPostID), arg0, arg1)
}

// GetPostByID mocks base method.
func (m *MockPostsServer) GetPostByID(arg0 context.Context, arg1 *protobuf.PostID) (*protobuf.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostByID", arg0, arg1)
	ret0, _ := ret[0].(*protobuf.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostByID indicates an expected call of GetPostByID.
func (mr *MockPostsServerMockRecorder) GetPostByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostByID", reflect.TypeOf((*MockPostsServer)(nil).GetPostByID), arg0, arg1)
}

// GetPostsBySubscriptions mocks base method.
func (m *MockPostsServer) GetPostsBySubscriptions(arg0 context.Context, arg1 *protobuf0.UserID) (*protobuf.PostArray, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsBySubscriptions", arg0, arg1)
	ret0, _ := ret[0].(*protobuf.PostArray)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsBySubscriptions indicates an expected call of GetPostsBySubscriptions.
func (mr *MockPostsServerMockRecorder) GetPostsBySubscriptions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsBySubscriptions", reflect.TypeOf((*MockPostsServer)(nil).GetPostsBySubscriptions), arg0, arg1)
}

// GetTagById mocks base method.
func (m *MockPostsServer) GetTagById(arg0 context.Context, arg1 *protobuf.TagID) (*protobuf.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagById", arg0, arg1)
	ret0, _ := ret[0].(*protobuf.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagById indicates an expected call of GetTagById.
func (mr *MockPostsServerMockRecorder) GetTagById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagById", reflect.TypeOf((*MockPostsServer)(nil).GetTagById), arg0, arg1)
}

// GetTagByName mocks base method.
func (m *MockPostsServer) GetTagByName(arg0 context.Context, arg1 *protobuf.TagName) (*protobuf.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagByName", arg0, arg1)
	ret0, _ := ret[0].(*protobuf.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagByName indicates an expected call of GetTagByName.
func (mr *MockPostsServerMockRecorder) GetTagByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagByName", reflect.TypeOf((*MockPostsServer)(nil).GetTagByName), arg0, arg1)
}

// GetTagDepsByPostId mocks base method.
func (m *MockPostsServer) GetTagDepsByPostId(arg0 context.Context, arg1 *protobuf.PostID) (*protobuf.TagDeps, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagDepsByPostId", arg0, arg1)
	ret0, _ := ret[0].(*protobuf.TagDeps)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagDepsByPostId indicates an expected call of GetTagDepsByPostId.
func (mr *MockPostsServerMockRecorder) GetTagDepsByPostId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagDepsByPostId", reflect.TypeOf((*MockPostsServer)(nil).GetTagDepsByPostId), arg0, arg1)
}

// Update mocks base method.
func (m *MockPostsServer) Update(arg0 context.Context, arg1 *protobuf.Post) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPostsServerMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPostsServer)(nil).Update), arg0, arg1)
}

// UpdateComment mocks base method.
func (m *MockPostsServer) UpdateComment(arg0 context.Context, arg1 *protobuf.Comment) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateComment", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateComment indicates an expected call of UpdateComment.
func (mr *MockPostsServerMockRecorder) UpdateComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComment", reflect.TypeOf((*MockPostsServer)(nil).UpdateComment), arg0, arg1)
}

// mustEmbedUnimplementedPostsServer mocks base method.
func (m *MockPostsServer) mustEmbedUnimplementedPostsServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedPostsServer")
}

// mustEmbedUnimplementedPostsServer indicates an expected call of mustEmbedUnimplementedPostsServer.
func (mr *MockPostsServerMockRecorder) mustEmbedUnimplementedPostsServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedPostsServer", reflect.TypeOf((*MockPostsServer)(nil).mustEmbedUnimplementedPostsServer))
}

// MockUnsafePostsServer is a mock of UnsafePostsServer interface.
type MockUnsafePostsServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafePostsServerMockRecorder
}

// MockUnsafePostsServerMockRecorder is the mock recorder for MockUnsafePostsServer.
type MockUnsafePostsServerMockRecorder struct {
	mock *MockUnsafePostsServer
}

// NewMockUnsafePostsServer creates a new mock instance.
func NewMockUnsafePostsServer(ctrl *gomock.Controller) *MockUnsafePostsServer {
	mock := &MockUnsafePostsServer{ctrl: ctrl}
	mock.recorder = &MockUnsafePostsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafePostsServer) EXPECT() *MockUnsafePostsServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedPostsServer mocks base method.
func (m *MockUnsafePostsServer) mustEmbedUnimplementedPostsServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedPostsServer")
}

// mustEmbedUnimplementedPostsServer indicates an expected call of mustEmbedUnimplementedPostsServer.
func (mr *MockUnsafePostsServerMockRecorder) mustEmbedUnimplementedPostsServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedPostsServer", reflect.TypeOf((*MockUnsafePostsServer)(nil).mustEmbedUnimplementedPostsServer))
}
