// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/users.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	multipart "mime/multipart"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockUsersMicroservice is a mock of UsersMicroservice interface.
type MockUsersMicroservice struct {
	ctrl     *gomock.Controller
	recorder *MockUsersMicroserviceMockRecorder
}

// MockUsersMicroserviceMockRecorder is the mock recorder for MockUsersMicroservice.
type MockUsersMicroserviceMockRecorder struct {
	mock *MockUsersMicroservice
}

// NewMockUsersMicroservice creates a new mock instance.
func NewMockUsersMicroservice(ctrl *gomock.Controller) *MockUsersMicroservice {
	mock := &MockUsersMicroservice{ctrl: ctrl}
	mock.recorder = &MockUsersMicroserviceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersMicroservice) EXPECT() *MockUsersMicroserviceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUsersMicroservice) Create(user models.User) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUsersMicroserviceMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUsersMicroservice)(nil).Create), user)
}

// DropBalance mocks base method.
func (m *MockUsersMicroservice) DropBalance(userID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DropBalance", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropBalance indicates an expected call of DropBalance.
func (mr *MockUsersMicroserviceMockRecorder) DropBalance(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropBalance", reflect.TypeOf((*MockUsersMicroservice)(nil).DropBalance), userID)
}

// GetAllAuthors mocks base method.
func (m *MockUsersMicroservice) GetAllAuthors() ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAuthors")
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAuthors indicates an expected call of GetAllAuthors.
func (mr *MockUsersMicroserviceMockRecorder) GetAllAuthors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAuthors", reflect.TypeOf((*MockUsersMicroservice)(nil).GetAllAuthors))
}

// GetAuthorByUsername mocks base method.
func (m *MockUsersMicroservice) GetAuthorByUsername(username string) ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorByUsername", username)
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorByUsername indicates an expected call of GetAuthorByUsername.
func (mr *MockUsersMicroserviceMockRecorder) GetAuthorByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorByUsername", reflect.TypeOf((*MockUsersMicroservice)(nil).GetAuthorByUsername), username)
}

// GetByEmail mocks base method.
func (m *MockUsersMicroservice) GetByEmail(email string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", email)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockUsersMicroserviceMockRecorder) GetByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockUsersMicroservice)(nil).GetByEmail), email)
}

// GetByID mocks base method.
func (m *MockUsersMicroservice) GetByID(id uint64) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockUsersMicroserviceMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUsersMicroservice)(nil).GetByID), id)
}

// GetBySessionID mocks base method.
func (m *MockUsersMicroservice) GetBySessionID(sessionID string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySessionID", sessionID)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySessionID indicates an expected call of GetBySessionID.
func (mr *MockUsersMicroserviceMockRecorder) GetBySessionID(sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySessionID", reflect.TypeOf((*MockUsersMicroservice)(nil).GetBySessionID), sessionID)
}

// GetByUsername mocks base method.
func (m *MockUsersMicroservice) GetByUsername(username string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsername", username)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsername indicates an expected call of GetByUsername.
func (mr *MockUsersMicroserviceMockRecorder) GetByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsername", reflect.TypeOf((*MockUsersMicroservice)(nil).GetByUsername), username)
}

// GetPostsNum mocks base method.
func (m *MockUsersMicroservice) GetPostsNum(UserID uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsNum", UserID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsNum indicates an expected call of GetPostsNum.
func (mr *MockUsersMicroserviceMockRecorder) GetPostsNum(UserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsNum", reflect.TypeOf((*MockUsersMicroservice)(nil).GetPostsNum), UserID)
}

// GetProfitForMounth mocks base method.
func (m *MockUsersMicroservice) GetProfitForMounth(UserID uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfitForMounth", UserID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfitForMounth indicates an expected call of GetProfitForMounth.
func (mr *MockUsersMicroserviceMockRecorder) GetProfitForMounth(UserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfitForMounth", reflect.TypeOf((*MockUsersMicroservice)(nil).GetProfitForMounth), UserID)
}

// GetSubscribersNum mocks base method.
func (m *MockUsersMicroservice) GetSubscribersNum(UserID uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscribersNum", UserID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscribersNum indicates an expected call of GetSubscribersNum.
func (mr *MockUsersMicroserviceMockRecorder) GetSubscribersNum(UserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribersNum", reflect.TypeOf((*MockUsersMicroservice)(nil).GetSubscribersNum), UserID)
}

// GetUserByPostID mocks base method.
func (m *MockUsersMicroservice) GetUserByPostID(postID uint64) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByPostID", postID)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByPostID indicates an expected call of GetUserByPostID.
func (mr *MockUsersMicroserviceMockRecorder) GetUserByPostID(postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByPostID", reflect.TypeOf((*MockUsersMicroservice)(nil).GetUserByPostID), postID)
}

// Update mocks base method.
func (m *MockUsersMicroservice) Update(user models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUsersMicroserviceMockRecorder) Update(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUsersMicroservice)(nil).Update), user)
}

// MockUsersUseCase is a mock of UsersUseCase interface.
type MockUsersUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUsersUseCaseMockRecorder
}

// MockUsersUseCaseMockRecorder is the mock recorder for MockUsersUseCase.
type MockUsersUseCaseMockRecorder struct {
	mock *MockUsersUseCase
}

// NewMockUsersUseCase creates a new mock instance.
func NewMockUsersUseCase(ctrl *gomock.Controller) *MockUsersUseCase {
	mock := &MockUsersUseCase{ctrl: ctrl}
	mock.recorder = &MockUsersUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersUseCase) EXPECT() *MockUsersUseCaseMockRecorder {
	return m.recorder
}

// CheckIDAndPassword mocks base method.
func (m *MockUsersUseCase) CheckIDAndPassword(id uint64, password string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIDAndPassword", id, password)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckIDAndPassword indicates an expected call of CheckIDAndPassword.
func (mr *MockUsersUseCaseMockRecorder) CheckIDAndPassword(id, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIDAndPassword", reflect.TypeOf((*MockUsersUseCase)(nil).CheckIDAndPassword), id, password)
}

// Create mocks base method.
func (m *MockUsersUseCase) Create(user models.User) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUsersUseCaseMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUsersUseCase)(nil).Create), user)
}

// FindAuthors mocks base method.
func (m *MockUsersUseCase) FindAuthors(keyword string) ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAuthors", keyword)
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAuthors indicates an expected call of FindAuthors.
func (mr *MockUsersUseCaseMockRecorder) FindAuthors(keyword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAuthors", reflect.TypeOf((*MockUsersUseCase)(nil).FindAuthors), keyword)
}

// GetByEmail mocks base method.
func (m *MockUsersUseCase) GetByEmail(email string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", email)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockUsersUseCaseMockRecorder) GetByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockUsersUseCase)(nil).GetByEmail), email)
}

// GetByID mocks base method.
func (m *MockUsersUseCase) GetByID(id uint64) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockUsersUseCaseMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUsersUseCase)(nil).GetByID), id)
}

// GetBySessionID mocks base method.
func (m *MockUsersUseCase) GetBySessionID(sessionID string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySessionID", sessionID)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySessionID indicates an expected call of GetBySessionID.
func (mr *MockUsersUseCaseMockRecorder) GetBySessionID(sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySessionID", reflect.TypeOf((*MockUsersUseCase)(nil).GetBySessionID), sessionID)
}

// GetByUsername mocks base method.
func (m *MockUsersUseCase) GetByUsername(username string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsername", username)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsername indicates an expected call of GetByUsername.
func (mr *MockUsersUseCaseMockRecorder) GetByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsername", reflect.TypeOf((*MockUsersUseCase)(nil).GetByUsername), username)
}

// GetPostsNum mocks base method.
func (m *MockUsersUseCase) GetPostsNum(userID uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsNum", userID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsNum indicates an expected call of GetPostsNum.
func (mr *MockUsersUseCaseMockRecorder) GetPostsNum(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsNum", reflect.TypeOf((*MockUsersUseCase)(nil).GetPostsNum), userID)
}

// GetProfitForMounth mocks base method.
func (m *MockUsersUseCase) GetProfitForMounth(userID uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfitForMounth", userID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfitForMounth indicates an expected call of GetProfitForMounth.
func (mr *MockUsersUseCaseMockRecorder) GetProfitForMounth(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfitForMounth", reflect.TypeOf((*MockUsersUseCase)(nil).GetProfitForMounth), userID)
}

// GetSubscribersNum mocks base method.
func (m *MockUsersUseCase) GetSubscribersNum(userID uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscribersNum", userID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscribersNum indicates an expected call of GetSubscribersNum.
func (mr *MockUsersUseCaseMockRecorder) GetSubscribersNum(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribersNum", reflect.TypeOf((*MockUsersUseCase)(nil).GetSubscribersNum), userID)
}

// GetUserByPostID mocks base method.
func (m *MockUsersUseCase) GetUserByPostID(postID uint64) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByPostID", postID)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByPostID indicates an expected call of GetUserByPostID.
func (mr *MockUsersUseCaseMockRecorder) GetUserByPostID(postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByPostID", reflect.TypeOf((*MockUsersUseCase)(nil).GetUserByPostID), postID)
}

// IsExistUsernameAndEmail mocks base method.
func (m *MockUsersUseCase) IsExistUsernameAndEmail(username, email string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExistUsernameAndEmail", username, email)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsExistUsernameAndEmail indicates an expected call of IsExistUsernameAndEmail.
func (mr *MockUsersUseCaseMockRecorder) IsExistUsernameAndEmail(username, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExistUsernameAndEmail", reflect.TypeOf((*MockUsersUseCase)(nil).IsExistUsernameAndEmail), username, email)
}

// Update mocks base method.
func (m *MockUsersUseCase) Update(user models.User, file *multipart.FileHeader, id uint64) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", user, file, id)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUsersUseCaseMockRecorder) Update(user, file, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUsersUseCase)(nil).Update), user, file, id)
}
