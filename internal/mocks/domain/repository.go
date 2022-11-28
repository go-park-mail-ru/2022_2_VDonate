// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/repository.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	multipart "mime/multipart"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthRepository is a mock of AuthRepository interface.
type MockAuthRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRepositoryMockRecorder
}

// MockAuthRepositoryMockRecorder is the mock recorder for MockAuthRepository.
type MockAuthRepositoryMockRecorder struct {
	mock *MockAuthRepository
}

// NewMockAuthRepository creates a new mock instance.
func NewMockAuthRepository(ctrl *gomock.Controller) *MockAuthRepository {
	mock := &MockAuthRepository{ctrl: ctrl}
	mock.recorder = &MockAuthRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthRepository) EXPECT() *MockAuthRepositoryMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockAuthRepository) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockAuthRepositoryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAuthRepository)(nil).Close))
}

// CreateSession mocks base method.
func (m *MockAuthRepository) CreateSession(cookie models.Cookie) (models.Cookie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", cookie)
	ret0, _ := ret[0].(models.Cookie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockAuthRepositoryMockRecorder) CreateSession(cookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockAuthRepository)(nil).CreateSession), cookie)
}

// DeleteBySessionID mocks base method.
func (m *MockAuthRepository) DeleteBySessionID(sessionID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBySessionID", sessionID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBySessionID indicates an expected call of DeleteBySessionID.
func (mr *MockAuthRepositoryMockRecorder) DeleteBySessionID(sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBySessionID", reflect.TypeOf((*MockAuthRepository)(nil).DeleteBySessionID), sessionID)
}

// DeleteByUserID mocks base method.
func (m *MockAuthRepository) DeleteByUserID(id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByUserID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByUserID indicates an expected call of DeleteByUserID.
func (mr *MockAuthRepositoryMockRecorder) DeleteByUserID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByUserID", reflect.TypeOf((*MockAuthRepository)(nil).DeleteByUserID), id)
}

// GetBySessionID mocks base method.
func (m *MockAuthRepository) GetBySessionID(sessionID string) (models.Cookie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySessionID", sessionID)
	ret0, _ := ret[0].(models.Cookie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySessionID indicates an expected call of GetBySessionID.
func (mr *MockAuthRepositoryMockRecorder) GetBySessionID(sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySessionID", reflect.TypeOf((*MockAuthRepository)(nil).GetBySessionID), sessionID)
}

// GetByUserID mocks base method.
func (m *MockAuthRepository) GetByUserID(id uint64) (models.Cookie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserID", id)
	ret0, _ := ret[0].(models.Cookie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserID indicates an expected call of GetByUserID.
func (mr *MockAuthRepositoryMockRecorder) GetByUserID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserID", reflect.TypeOf((*MockAuthRepository)(nil).GetByUserID), id)
}

// GetByUsername mocks base method.
func (m *MockAuthRepository) GetByUsername(username string) (models.Cookie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsername", username)
	ret0, _ := ret[0].(models.Cookie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsername indicates an expected call of GetByUsername.
func (mr *MockAuthRepositoryMockRecorder) GetByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsername", reflect.TypeOf((*MockAuthRepository)(nil).GetByUsername), username)
}

// MockPostsRepository is a mock of PostsRepository interface.
type MockPostsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPostsRepositoryMockRecorder
}

// MockPostsRepositoryMockRecorder is the mock recorder for MockPostsRepository.
type MockPostsRepositoryMockRecorder struct {
	mock *MockPostsRepository
}

// NewMockPostsRepository creates a new mock instance.
func NewMockPostsRepository(ctrl *gomock.Controller) *MockPostsRepository {
	mock := &MockPostsRepository{ctrl: ctrl}
	mock.recorder = &MockPostsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostsRepository) EXPECT() *MockPostsRepositoryMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockPostsRepository) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockPostsRepositoryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPostsRepository)(nil).Close))
}

// Create mocks base method.
func (m *MockPostsRepository) Create(post models.Post) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", post)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPostsRepositoryMockRecorder) Create(post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPostsRepository)(nil).Create), post)
}

// CreateDepTag mocks base method.
func (m *MockPostsRepository) CreateDepTag(postID, tagID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDepTag", postID, tagID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDepTag indicates an expected call of CreateDepTag.
func (mr *MockPostsRepositoryMockRecorder) CreateDepTag(postID, tagID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDepTag", reflect.TypeOf((*MockPostsRepository)(nil).CreateDepTag), postID, tagID)
}

// CreateLike mocks base method.
func (m *MockPostsRepository) CreateLike(userID, postID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLike", userID, postID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLike indicates an expected call of CreateLike.
func (mr *MockPostsRepositoryMockRecorder) CreateLike(userID, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLike", reflect.TypeOf((*MockPostsRepository)(nil).CreateLike), userID, postID)
}

// CreateTag mocks base method.
func (m *MockPostsRepository) CreateTag(tagName string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTag", tagName)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTag indicates an expected call of CreateTag.
func (mr *MockPostsRepositoryMockRecorder) CreateTag(tagName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTag", reflect.TypeOf((*MockPostsRepository)(nil).CreateTag), tagName)
}

// DeleteByID mocks base method.
func (m *MockPostsRepository) DeleteByID(postID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", postID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockPostsRepositoryMockRecorder) DeleteByID(postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockPostsRepository)(nil).DeleteByID), postID)
}

// DeleteDepTag mocks base method.
func (m *MockPostsRepository) DeleteDepTag(tagDep models.TagDep) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDepTag", tagDep)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDepTag indicates an expected call of DeleteDepTag.
func (mr *MockPostsRepositoryMockRecorder) DeleteDepTag(tagDep interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDepTag", reflect.TypeOf((*MockPostsRepository)(nil).DeleteDepTag), tagDep)
}

// DeleteLikeByID mocks base method.
func (m *MockPostsRepository) DeleteLikeByID(userID, postID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLikeByID", userID, postID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLikeByID indicates an expected call of DeleteLikeByID.
func (mr *MockPostsRepositoryMockRecorder) DeleteLikeByID(userID, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLikeByID", reflect.TypeOf((*MockPostsRepository)(nil).DeleteLikeByID), userID, postID)
}

// GetAllByUserID mocks base method.
func (m *MockPostsRepository) GetAllByUserID(authorID uint64) ([]models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByUserID", authorID)
	ret0, _ := ret[0].([]models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByUserID indicates an expected call of GetAllByUserID.
func (mr *MockPostsRepositoryMockRecorder) GetAllByUserID(authorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByUserID", reflect.TypeOf((*MockPostsRepository)(nil).GetAllByUserID), authorID)
}

// GetAllLikesByPostID mocks base method.
func (m *MockPostsRepository) GetAllLikesByPostID(postID uint64) ([]models.Like, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllLikesByPostID", postID)
	ret0, _ := ret[0].([]models.Like)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllLikesByPostID indicates an expected call of GetAllLikesByPostID.
func (mr *MockPostsRepositoryMockRecorder) GetAllLikesByPostID(postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllLikesByPostID", reflect.TypeOf((*MockPostsRepository)(nil).GetAllLikesByPostID), postID)
}

// GetLikeByUserAndPostID mocks base method.
func (m *MockPostsRepository) GetLikeByUserAndPostID(userID, postID uint64) (models.Like, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLikeByUserAndPostID", userID, postID)
	ret0, _ := ret[0].(models.Like)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLikeByUserAndPostID indicates an expected call of GetLikeByUserAndPostID.
func (mr *MockPostsRepositoryMockRecorder) GetLikeByUserAndPostID(userID, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLikeByUserAndPostID", reflect.TypeOf((*MockPostsRepository)(nil).GetLikeByUserAndPostID), userID, postID)
}

// GetPostByID mocks base method.
func (m *MockPostsRepository) GetPostByID(postID uint64) (models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostByID", postID)
	ret0, _ := ret[0].(models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostByID indicates an expected call of GetPostByID.
func (mr *MockPostsRepositoryMockRecorder) GetPostByID(postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostByID", reflect.TypeOf((*MockPostsRepository)(nil).GetPostByID), postID)
}

// GetPostsBySubscriptions mocks base method.
func (m *MockPostsRepository) GetPostsBySubscriptions(userID uint64) ([]models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsBySubscriptions", userID)
	ret0, _ := ret[0].([]models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsBySubscriptions indicates an expected call of GetPostsBySubscriptions.
func (mr *MockPostsRepositoryMockRecorder) GetPostsBySubscriptions(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsBySubscriptions", reflect.TypeOf((*MockPostsRepository)(nil).GetPostsBySubscriptions), userID)
}

// GetTagById mocks base method.
func (m *MockPostsRepository) GetTagById(tagID uint64) (models.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagById", tagID)
	ret0, _ := ret[0].(models.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagById indicates an expected call of GetTagById.
func (mr *MockPostsRepositoryMockRecorder) GetTagById(tagID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagById", reflect.TypeOf((*MockPostsRepository)(nil).GetTagById), tagID)
}

// GetTagByName mocks base method.
func (m *MockPostsRepository) GetTagByName(tagName string) (models.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagByName", tagName)
	ret0, _ := ret[0].(models.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagByName indicates an expected call of GetTagByName.
func (mr *MockPostsRepositoryMockRecorder) GetTagByName(tagName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagByName", reflect.TypeOf((*MockPostsRepository)(nil).GetTagByName), tagName)
}

// GetTagDepsByPostId mocks base method.
func (m *MockPostsRepository) GetTagDepsByPostId(postID uint64) ([]models.TagDep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagDepsByPostId", postID)
	ret0, _ := ret[0].([]models.TagDep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagDepsByPostId indicates an expected call of GetTagDepsByPostId.
func (mr *MockPostsRepositoryMockRecorder) GetTagDepsByPostId(postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagDepsByPostId", reflect.TypeOf((*MockPostsRepository)(nil).GetTagDepsByPostId), postID)
}

// Update mocks base method.
func (m *MockPostsRepository) Update(post models.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", post)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPostsRepositoryMockRecorder) Update(post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPostsRepository)(nil).Update), post)
}

// MockSubscribersRepository is a mock of SubscribersRepository interface.
type MockSubscribersRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSubscribersRepositoryMockRecorder
}

// MockSubscribersRepositoryMockRecorder is the mock recorder for MockSubscribersRepository.
type MockSubscribersRepositoryMockRecorder struct {
	mock *MockSubscribersRepository
}

// NewMockSubscribersRepository creates a new mock instance.
func NewMockSubscribersRepository(ctrl *gomock.Controller) *MockSubscribersRepository {
	mock := &MockSubscribersRepository{ctrl: ctrl}
	mock.recorder = &MockSubscribersRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscribersRepository) EXPECT() *MockSubscribersRepositoryMockRecorder {
	return m.recorder
}

// GetSubscribers mocks base method.
func (m *MockSubscribersRepository) GetSubscribers(authorID uint64) ([]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscribers", authorID)
	ret0, _ := ret[0].([]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscribers indicates an expected call of GetSubscribers.
func (mr *MockSubscribersRepositoryMockRecorder) GetSubscribers(authorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribers", reflect.TypeOf((*MockSubscribersRepository)(nil).GetSubscribers), authorID)
}

// Subscribe mocks base method.
func (m *MockSubscribersRepository) Subscribe(subscription models.Subscription) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", subscription)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockSubscribersRepositoryMockRecorder) Subscribe(subscription interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockSubscribersRepository)(nil).Subscribe), subscription)
}

// Unsubscribe mocks base method.
func (m *MockSubscribersRepository) Unsubscribe(userID, authorID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", userID, authorID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockSubscribersRepositoryMockRecorder) Unsubscribe(userID, authorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockSubscribersRepository)(nil).Unsubscribe), userID, authorID)
}

// MockSubscriptionsRepository is a mock of SubscriptionsRepository interface.
type MockSubscriptionsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionsRepositoryMockRecorder
}

// MockSubscriptionsRepositoryMockRecorder is the mock recorder for MockSubscriptionsRepository.
type MockSubscriptionsRepositoryMockRecorder struct {
	mock *MockSubscriptionsRepository
}

// NewMockSubscriptionsRepository creates a new mock instance.
func NewMockSubscriptionsRepository(ctrl *gomock.Controller) *MockSubscriptionsRepository {
	mock := &MockSubscriptionsRepository{ctrl: ctrl}
	mock.recorder = &MockSubscriptionsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscriptionsRepository) EXPECT() *MockSubscriptionsRepositoryMockRecorder {
	return m.recorder
}

// AddSubscription mocks base method.
func (m *MockSubscriptionsRepository) AddSubscription(sub models.AuthorSubscription) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSubscription", sub)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSubscription indicates an expected call of AddSubscription.
func (mr *MockSubscriptionsRepositoryMockRecorder) AddSubscription(sub interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSubscription", reflect.TypeOf((*MockSubscriptionsRepository)(nil).AddSubscription), sub)
}

// DeleteSubscription mocks base method.
func (m *MockSubscriptionsRepository) DeleteSubscription(subID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubscription", subID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSubscription indicates an expected call of DeleteSubscription.
func (mr *MockSubscriptionsRepositoryMockRecorder) DeleteSubscription(subID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubscription", reflect.TypeOf((*MockSubscriptionsRepository)(nil).DeleteSubscription), subID)
}

// GetSubscriptionByID mocks base method.
func (m *MockSubscriptionsRepository) GetSubscriptionByID(ID uint64) (models.AuthorSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptionByID", ID)
	ret0, _ := ret[0].(models.AuthorSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptionByID indicates an expected call of GetSubscriptionByID.
func (mr *MockSubscriptionsRepositoryMockRecorder) GetSubscriptionByID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptionByID", reflect.TypeOf((*MockSubscriptionsRepository)(nil).GetSubscriptionByID), ID)
}

// GetSubscriptionByUserAndAuthorID mocks base method.
func (m *MockSubscriptionsRepository) GetSubscriptionByUserAndAuthorID(userID, authorID uint64) (models.AuthorSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptionByUserAndAuthorID", userID, authorID)
	ret0, _ := ret[0].(models.AuthorSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptionByUserAndAuthorID indicates an expected call of GetSubscriptionByUserAndAuthorID.
func (mr *MockSubscriptionsRepositoryMockRecorder) GetSubscriptionByUserAndAuthorID(userID, authorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptionByUserAndAuthorID", reflect.TypeOf((*MockSubscriptionsRepository)(nil).GetSubscriptionByUserAndAuthorID), userID, authorID)
}

// GetSubscriptionsByAuthorID mocks base method.
func (m *MockSubscriptionsRepository) GetSubscriptionsByAuthorID(authorID uint64) ([]models.AuthorSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptionsByAuthorID", authorID)
	ret0, _ := ret[0].([]models.AuthorSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptionsByAuthorID indicates an expected call of GetSubscriptionsByAuthorID.
func (mr *MockSubscriptionsRepositoryMockRecorder) GetSubscriptionsByAuthorID(authorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptionsByAuthorID", reflect.TypeOf((*MockSubscriptionsRepository)(nil).GetSubscriptionsByAuthorID), authorID)
}

// GetSubscriptionsByUserID mocks base method.
func (m *MockSubscriptionsRepository) GetSubscriptionsByUserID(userID uint64) ([]models.AuthorSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptionsByUserID", userID)
	ret0, _ := ret[0].([]models.AuthorSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptionsByUserID indicates an expected call of GetSubscriptionsByUserID.
func (mr *MockSubscriptionsRepositoryMockRecorder) GetSubscriptionsByUserID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptionsByUserID", reflect.TypeOf((*MockSubscriptionsRepository)(nil).GetSubscriptionsByUserID), userID)
}

// UpdateSubscription mocks base method.
func (m *MockSubscriptionsRepository) UpdateSubscription(sub models.AuthorSubscription) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubscription", sub)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSubscription indicates an expected call of UpdateSubscription.
func (mr *MockSubscriptionsRepositoryMockRecorder) UpdateSubscription(sub interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubscription", reflect.TypeOf((*MockSubscriptionsRepository)(nil).UpdateSubscription), sub)
}

// MockUsersRepository is a mock of UsersRepository interface.
type MockUsersRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUsersRepositoryMockRecorder
}

// MockUsersRepositoryMockRecorder is the mock recorder for MockUsersRepository.
type MockUsersRepositoryMockRecorder struct {
	mock *MockUsersRepository
}

// NewMockUsersRepository creates a new mock instance.
func NewMockUsersRepository(ctrl *gomock.Controller) *MockUsersRepository {
	mock := &MockUsersRepository{ctrl: ctrl}
	mock.recorder = &MockUsersRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersRepository) EXPECT() *MockUsersRepositoryMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockUsersRepository) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockUsersRepositoryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockUsersRepository)(nil).Close))
}

// Create mocks base method.
func (m *MockUsersRepository) Create(user models.User) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUsersRepositoryMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUsersRepository)(nil).Create), user)
}

// DeleteByID mocks base method.
func (m *MockUsersRepository) DeleteByID(id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockUsersRepositoryMockRecorder) DeleteByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockUsersRepository)(nil).DeleteByID), id)
}

// GetAuthorByUsername mocks base method.
func (m *MockUsersRepository) GetAuthorByUsername(username string) ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorByUsername", username)
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorByUsername indicates an expected call of GetAuthorByUsername.
func (mr *MockUsersRepositoryMockRecorder) GetAuthorByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorByUsername", reflect.TypeOf((*MockUsersRepository)(nil).GetAuthorByUsername), username)
}

// GetByEmail mocks base method.
func (m *MockUsersRepository) GetByEmail(email string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", email)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockUsersRepositoryMockRecorder) GetByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockUsersRepository)(nil).GetByEmail), email)
}

// GetByID mocks base method.
func (m *MockUsersRepository) GetByID(id uint64) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockUsersRepositoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUsersRepository)(nil).GetByID), id)
}

// GetBySessionID mocks base method.
func (m *MockUsersRepository) GetBySessionID(sessionID string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySessionID", sessionID)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySessionID indicates an expected call of GetBySessionID.
func (mr *MockUsersRepositoryMockRecorder) GetBySessionID(sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySessionID", reflect.TypeOf((*MockUsersRepository)(nil).GetBySessionID), sessionID)
}

// GetByUsername mocks base method.
func (m *MockUsersRepository) GetByUsername(username string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsername", username)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsername indicates an expected call of GetByUsername.
func (mr *MockUsersRepositoryMockRecorder) GetByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsername", reflect.TypeOf((*MockUsersRepository)(nil).GetByUsername), username)
}

// GetUserByPostID mocks base method.
func (m *MockUsersRepository) GetUserByPostID(postID uint64) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByPostID", postID)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByPostID indicates an expected call of GetUserByPostID.
func (mr *MockUsersRepositoryMockRecorder) GetUserByPostID(postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByPostID", reflect.TypeOf((*MockUsersRepository)(nil).GetUserByPostID), postID)
}

// Update mocks base method.
func (m *MockUsersRepository) Update(user models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUsersRepositoryMockRecorder) Update(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUsersRepository)(nil).Update), user)
}

// MockDonatesRepository is a mock of DonatesRepository interface.
type MockDonatesRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDonatesRepositoryMockRecorder
}

// MockDonatesRepositoryMockRecorder is the mock recorder for MockDonatesRepository.
type MockDonatesRepositoryMockRecorder struct {
	mock *MockDonatesRepository
}

// NewMockDonatesRepository creates a new mock instance.
func NewMockDonatesRepository(ctrl *gomock.Controller) *MockDonatesRepository {
	mock := &MockDonatesRepository{ctrl: ctrl}
	mock.recorder = &MockDonatesRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDonatesRepository) EXPECT() *MockDonatesRepositoryMockRecorder {
	return m.recorder
}

// GetDonateByID mocks base method.
func (m *MockDonatesRepository) GetDonateByID(donateID uint64) (models.Donate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDonateByID", donateID)
	ret0, _ := ret[0].(models.Donate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDonateByID indicates an expected call of GetDonateByID.
func (mr *MockDonatesRepositoryMockRecorder) GetDonateByID(donateID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDonateByID", reflect.TypeOf((*MockDonatesRepository)(nil).GetDonateByID), donateID)
}

// GetDonatesByUserID mocks base method.
func (m *MockDonatesRepository) GetDonatesByUserID(userID uint64) ([]models.Donate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDonatesByUserID", userID)
	ret0, _ := ret[0].([]models.Donate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDonatesByUserID indicates an expected call of GetDonatesByUserID.
func (mr *MockDonatesRepositoryMockRecorder) GetDonatesByUserID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDonatesByUserID", reflect.TypeOf((*MockDonatesRepository)(nil).GetDonatesByUserID), userID)
}

// SendDonate mocks base method.
func (m *MockDonatesRepository) SendDonate(donate models.Donate) (models.Donate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendDonate", donate)
	ret0, _ := ret[0].(models.Donate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendDonate indicates an expected call of SendDonate.
func (mr *MockDonatesRepositoryMockRecorder) SendDonate(donate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendDonate", reflect.TypeOf((*MockDonatesRepository)(nil).SendDonate), donate)
}

// MockImagesRepository is a mock of ImagesRepository interface.
type MockImagesRepository struct {
	ctrl     *gomock.Controller
	recorder *MockImagesRepositoryMockRecorder
}

// MockImagesRepositoryMockRecorder is the mock recorder for MockImagesRepository.
type MockImagesRepositoryMockRecorder struct {
	mock *MockImagesRepository
}

// NewMockImagesRepository creates a new mock instance.
func NewMockImagesRepository(ctrl *gomock.Controller) *MockImagesRepository {
	mock := &MockImagesRepository{ctrl: ctrl}
	mock.recorder = &MockImagesRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImagesRepository) EXPECT() *MockImagesRepositoryMockRecorder {
	return m.recorder
}

// CreateOrUpdateImage mocks base method.
func (m *MockImagesRepository) CreateOrUpdateImage(image *multipart.FileHeader, oldFilename string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateImage", image, oldFilename)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdateImage indicates an expected call of CreateOrUpdateImage.
func (mr *MockImagesRepositoryMockRecorder) CreateOrUpdateImage(image, oldFilename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateImage", reflect.TypeOf((*MockImagesRepository)(nil).CreateOrUpdateImage), image, oldFilename)
}

// GetPermanentImage mocks base method.
func (m *MockImagesRepository) GetPermanentImage(filename string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPermanentImage", filename)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPermanentImage indicates an expected call of GetPermanentImage.
func (mr *MockImagesRepositoryMockRecorder) GetPermanentImage(filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPermanentImage", reflect.TypeOf((*MockImagesRepository)(nil).GetPermanentImage), filename)
}
