// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/users.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	reflect "reflect"

	models "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	gomock "github.com/golang/mock/gomock"
)

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

// DeleteByEmail mocks base method.
func (m *MockUsersUseCase) DeleteByEmail(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByEmail", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByEmail indicates an expected call of DeleteByEmail.
func (mr *MockUsersUseCaseMockRecorder) DeleteByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByEmail", reflect.TypeOf((*MockUsersUseCase)(nil).DeleteByEmail), email)
}

// DeleteByID mocks base method.
func (m *MockUsersUseCase) DeleteByID(id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockUsersUseCaseMockRecorder) DeleteByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockUsersUseCase)(nil).DeleteByID), id)
}

// DeleteByUsername mocks base method.
func (m *MockUsersUseCase) DeleteByUsername(username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByUsername", username)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByUsername indicates an expected call of DeleteByUsername.
func (mr *MockUsersUseCaseMockRecorder) DeleteByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByUsername", reflect.TypeOf((*MockUsersUseCase)(nil).DeleteByUsername), username)
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
func (m *MockUsersUseCase) Update(user models.User, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", user, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUsersUseCaseMockRecorder) Update(user, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUsersUseCase)(nil).Update), user, id)
}
