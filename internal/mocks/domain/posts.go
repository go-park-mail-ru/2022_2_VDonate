// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/posts.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	reflect "reflect"

	models "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockPostsUseCase is a mock of PostsUseCase interface.
type MockPostsUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockPostsUseCaseMockRecorder
}

// MockPostsUseCaseMockRecorder is the mock recorder for MockPostsUseCase.
type MockPostsUseCaseMockRecorder struct {
	mock *MockPostsUseCase
}

// NewMockPostsUseCase creates a new mock instance.
func NewMockPostsUseCase(ctrl *gomock.Controller) *MockPostsUseCase {
	mock := &MockPostsUseCase{ctrl: ctrl}
	mock.recorder = &MockPostsUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostsUseCase) EXPECT() *MockPostsUseCaseMockRecorder {
	return m.recorder
}

// ConvertTagsToStrSlice mocks base method.
func (m *MockPostsUseCase) ConvertTagsToStrSlice(tags []models.Tag) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertTagsToStrSlice", tags)
	ret0, _ := ret[0].([]string)
	return ret0
}

// ConvertTagsToStrSlice indicates an expected call of ConvertTagsToStrSlice.
func (mr *MockPostsUseCaseMockRecorder) ConvertTagsToStrSlice(tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertTagsToStrSlice", reflect.TypeOf((*MockPostsUseCase)(nil).ConvertTagsToStrSlice), tags)
}

// Create mocks base method.
func (m *MockPostsUseCase) Create(post models.Post, userID uint64) (models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", post, userID)
	ret0, _ := ret[0].(models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPostsUseCaseMockRecorder) Create(post, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPostsUseCase)(nil).Create), post, userID)
}

// CreateTags mocks base method.
func (m *MockPostsUseCase) CreateTags(tagNames []string, postID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTags", tagNames, postID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTags indicates an expected call of CreateTags.
func (mr *MockPostsUseCaseMockRecorder) CreateTags(tagNames, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTags", reflect.TypeOf((*MockPostsUseCase)(nil).CreateTags), tagNames, postID)
}

// DeleteByID mocks base method.
func (m *MockPostsUseCase) DeleteByID(postID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", postID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockPostsUseCaseMockRecorder) DeleteByID(postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockPostsUseCase)(nil).DeleteByID), postID)
}

// DeleteTagDeps mocks base method.
func (m *MockPostsUseCase) DeleteTagDeps(postID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTagDeps", postID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTagDeps indicates an expected call of DeleteTagDeps.
func (mr *MockPostsUseCaseMockRecorder) DeleteTagDeps(postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTagDeps", reflect.TypeOf((*MockPostsUseCase)(nil).DeleteTagDeps), postID)
}

// GetLikeByUserAndPostID mocks base method.
func (m *MockPostsUseCase) GetLikeByUserAndPostID(userID, postID uint64) (models.Like, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLikeByUserAndPostID", userID, postID)
	ret0, _ := ret[0].(models.Like)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLikeByUserAndPostID indicates an expected call of GetLikeByUserAndPostID.
func (mr *MockPostsUseCaseMockRecorder) GetLikeByUserAndPostID(userID, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLikeByUserAndPostID", reflect.TypeOf((*MockPostsUseCase)(nil).GetLikeByUserAndPostID), userID, postID)
}

// GetLikesByPostID mocks base method.
func (m *MockPostsUseCase) GetLikesByPostID(postID uint64) ([]models.Like, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLikesByPostID", postID)
	ret0, _ := ret[0].([]models.Like)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLikesByPostID indicates an expected call of GetLikesByPostID.
func (mr *MockPostsUseCaseMockRecorder) GetLikesByPostID(postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLikesByPostID", reflect.TypeOf((*MockPostsUseCase)(nil).GetLikesByPostID), postID)
}

// GetLikesNum mocks base method.
func (m *MockPostsUseCase) GetLikesNum(postID uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLikesNum", postID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLikesNum indicates an expected call of GetLikesNum.
func (mr *MockPostsUseCaseMockRecorder) GetLikesNum(postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLikesNum", reflect.TypeOf((*MockPostsUseCase)(nil).GetLikesNum), postID)
}

// GetPostByID mocks base method.
func (m *MockPostsUseCase) GetPostByID(postID, userID uint64) (models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostByID", postID, userID)
	ret0, _ := ret[0].(models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostByID indicates an expected call of GetPostByID.
func (mr *MockPostsUseCaseMockRecorder) GetPostByID(postID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostByID", reflect.TypeOf((*MockPostsUseCase)(nil).GetPostByID), postID, userID)
}

// GetPostsByFilter mocks base method.
func (m *MockPostsUseCase) GetPostsByFilter(userID, authorID uint64) ([]models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsByFilter", userID, authorID)
	ret0, _ := ret[0].([]models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsByFilter indicates an expected call of GetPostsByFilter.
func (mr *MockPostsUseCaseMockRecorder) GetPostsByFilter(userID, authorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsByFilter", reflect.TypeOf((*MockPostsUseCase)(nil).GetPostsByFilter), userID, authorID)
}

// GetTagsByPostID mocks base method.
func (m *MockPostsUseCase) GetTagsByPostID(postID uint64) ([]models.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagsByPostID", postID)
	ret0, _ := ret[0].([]models.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagsByPostID indicates an expected call of GetTagsByPostID.
func (mr *MockPostsUseCaseMockRecorder) GetTagsByPostID(postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagsByPostID", reflect.TypeOf((*MockPostsUseCase)(nil).GetTagsByPostID), postID)
}

// IsPostLiked mocks base method.
func (m *MockPostsUseCase) IsPostLiked(userID, postID uint64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPostLiked", userID, postID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPostLiked indicates an expected call of IsPostLiked.
func (mr *MockPostsUseCaseMockRecorder) IsPostLiked(userID, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPostLiked", reflect.TypeOf((*MockPostsUseCase)(nil).IsPostLiked), userID, postID)
}

// LikePost mocks base method.
func (m *MockPostsUseCase) LikePost(userID, postID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LikePost", userID, postID)
	ret0, _ := ret[0].(error)
	return ret0
}

// LikePost indicates an expected call of LikePost.
func (mr *MockPostsUseCaseMockRecorder) LikePost(userID, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LikePost", reflect.TypeOf((*MockPostsUseCase)(nil).LikePost), userID, postID)
}

// UnlikePost mocks base method.
func (m *MockPostsUseCase) UnlikePost(userID, postID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnlikePost", userID, postID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnlikePost indicates an expected call of UnlikePost.
func (mr *MockPostsUseCaseMockRecorder) UnlikePost(userID, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlikePost", reflect.TypeOf((*MockPostsUseCase)(nil).UnlikePost), userID, postID)
}

// Update mocks base method.
func (m *MockPostsUseCase) Update(post models.Post, postID uint64) (models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", post, postID)
	ret0, _ := ret[0].(models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPostsUseCaseMockRecorder) Update(post, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPostsUseCase)(nil).Update), post, postID)
}

// UpdateTags mocks base method.
func (m *MockPostsUseCase) UpdateTags(tagNames []string, postID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTags", tagNames, postID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTags indicates an expected call of UpdateTags.
func (mr *MockPostsUseCaseMockRecorder) UpdateTags(tagNames, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTags", reflect.TypeOf((*MockPostsUseCase)(nil).UpdateTags), tagNames, postID)
}
