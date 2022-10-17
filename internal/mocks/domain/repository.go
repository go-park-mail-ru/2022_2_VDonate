// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/repository.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
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
func (m *MockAuthRepository) CreateSession(cookie models.Cookie) (*models.Cookie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", cookie)
	ret0, _ := ret[0].(*models.Cookie)
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
func (m *MockAuthRepository) GetBySessionID(sessionID string) (*models.Cookie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySessionID", sessionID)
	ret0, _ := ret[0].(*models.Cookie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySessionID indicates an expected call of GetBySessionID.
func (mr *MockAuthRepositoryMockRecorder) GetBySessionID(sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySessionID", reflect.TypeOf((*MockAuthRepository)(nil).GetBySessionID), sessionID)
}

// GetByUserID mocks base method.
func (m *MockAuthRepository) GetByUserID(id uint64) (*models.Cookie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserID", id)
	ret0, _ := ret[0].(*models.Cookie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserID indicates an expected call of GetByUserID.
func (mr *MockAuthRepositoryMockRecorder) GetByUserID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserID", reflect.TypeOf((*MockAuthRepository)(nil).GetByUserID), id)
}

// GetByUsername mocks base method.
func (m *MockAuthRepository) GetByUsername(username string) (*models.Cookie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsername", username)
	ret0, _ := ret[0].(*models.Cookie)
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
func (m *MockPostsRepository) Create(post models.Post) (*models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", post)
	ret0, _ := ret[0].(*models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPostsRepositoryMockRecorder) Create(post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPostsRepository)(nil).Create), post)
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

// GetAllByUserID mocks base method.
func (m *MockPostsRepository) GetAllByUserID(userID uint64) ([]*models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByUserID", userID)
	ret0, _ := ret[0].([]*models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByUserID indicates an expected call of GetAllByUserID.
func (mr *MockPostsRepositoryMockRecorder) GetAllByUserID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByUserID", reflect.TypeOf((*MockPostsRepository)(nil).GetAllByUserID), userID)
}

// GetPostByID mocks base method.
func (m *MockPostsRepository) GetPostByID(postID uint64) (*models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostByID", postID)
	ret0, _ := ret[0].(*models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostByID indicates an expected call of GetPostByID.
func (mr *MockPostsRepositoryMockRecorder) GetPostByID(postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostByID", reflect.TypeOf((*MockPostsRepository)(nil).GetPostByID), postID)
}

// Update mocks base method.
func (m *MockPostsRepository) Update(post models.Post) (*models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", post)
	ret0, _ := ret[0].(*models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
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
func (m *MockSubscriptionsRepository) AddSubscription(sub models.AuthorSubscription) (*models.AuthorSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSubscription", sub)
	ret0, _ := ret[0].(*models.AuthorSubscription)
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

// GetSubscriptionsByAuthorID mocks base method.
func (m *MockSubscriptionsRepository) GetSubscriptionsByAuthorID(authorID uint64) ([]*models.AuthorSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptionsByAuthorID", authorID)
	ret0, _ := ret[0].([]*models.AuthorSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptionsByAuthorID indicates an expected call of GetSubscriptionsByAuthorID.
func (mr *MockSubscriptionsRepositoryMockRecorder) GetSubscriptionsByAuthorID(authorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptionsByAuthorID", reflect.TypeOf((*MockSubscriptionsRepository)(nil).GetSubscriptionsByAuthorID), authorID)
}

// GetSubscriptionsByID mocks base method.
func (m *MockSubscriptionsRepository) GetSubscriptionsByID(ID uint64) (*models.AuthorSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptionsByID", ID)
	ret0, _ := ret[0].(*models.AuthorSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptionsByID indicates an expected call of GetSubscriptionsByID.
func (mr *MockSubscriptionsRepositoryMockRecorder) GetSubscriptionsByID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptionsByID", reflect.TypeOf((*MockSubscriptionsRepository)(nil).GetSubscriptionsByID), ID)
}

// UpdateSubscription mocks base method.
func (m *MockSubscriptionsRepository) UpdateSubscription(sub *models.AuthorSubscription) (*models.AuthorSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubscription", sub)
	ret0, _ := ret[0].(*models.AuthorSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
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
func (m *MockUsersRepository) Create(user *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(*models.User)
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

// GetByEmail mocks base method.
func (m *MockUsersRepository) GetByEmail(email string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", email)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockUsersRepositoryMockRecorder) GetByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockUsersRepository)(nil).GetByEmail), email)
}

// GetByID mocks base method.
func (m *MockUsersRepository) GetByID(id uint64) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockUsersRepositoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUsersRepository)(nil).GetByID), id)
}

// GetBySessionID mocks base method.
func (m *MockUsersRepository) GetBySessionID(sessionID string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySessionID", sessionID)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySessionID indicates an expected call of GetBySessionID.
func (mr *MockUsersRepositoryMockRecorder) GetBySessionID(sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySessionID", reflect.TypeOf((*MockUsersRepository)(nil).GetBySessionID), sessionID)
}

// GetByUsername mocks base method.
func (m *MockUsersRepository) GetByUsername(username string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsername", username)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsername indicates an expected call of GetByUsername.
func (mr *MockUsersRepositoryMockRecorder) GetByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsername", reflect.TypeOf((*MockUsersRepository)(nil).GetByUsername), username)
}

// GetUserByPostID mocks base method.
func (m *MockUsersRepository) GetUserByPostID(postID uint64) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByPostID", postID)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByPostID indicates an expected call of GetUserByPostID.
func (mr *MockUsersRepositoryMockRecorder) GetUserByPostID(postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByPostID", reflect.TypeOf((*MockUsersRepository)(nil).GetUserByPostID), postID)
}

// Update mocks base method.
func (m *MockUsersRepository) Update(user *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", user)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUsersRepositoryMockRecorder) Update(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUsersRepository)(nil).Update), user)
}
