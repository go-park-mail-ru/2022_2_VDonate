// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/auth.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	reflect "reflect"

	protobuf "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/auth/protobuf"
	models "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthUseCase is a mock of AuthUseCase interface.
type MockAuthUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockAuthUseCaseMockRecorder
}

// MockAuthUseCaseMockRecorder is the mock recorder for MockAuthUseCase.
type MockAuthUseCaseMockRecorder struct {
	mock *MockAuthUseCase
}

// NewMockAuthUseCase creates a new mock instance.
func NewMockAuthUseCase(ctrl *gomock.Controller) *MockAuthUseCase {
	mock := &MockAuthUseCase{ctrl: ctrl}
	mock.recorder = &MockAuthUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthUseCase) EXPECT() *MockAuthUseCaseMockRecorder {
	return m.recorder
}

// Auth mocks base method.
func (m *MockAuthUseCase) Auth(sessionID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auth", sessionID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Auth indicates an expected call of Auth.
func (mr *MockAuthUseCaseMockRecorder) Auth(sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auth", reflect.TypeOf((*MockAuthUseCase)(nil).Auth), sessionID)
}

// IsSameSession mocks base method.
func (m *MockAuthUseCase) IsSameSession(sessionID string, userID uint64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSameSession", sessionID, userID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSameSession indicates an expected call of IsSameSession.
func (mr *MockAuthUseCaseMockRecorder) IsSameSession(sessionID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSameSession", reflect.TypeOf((*MockAuthUseCase)(nil).IsSameSession), sessionID, userID)
}

// Login mocks base method.
func (m *MockAuthUseCase) Login(login, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", login, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthUseCaseMockRecorder) Login(login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthUseCase)(nil).Login), login, password)
}

// Logout mocks base method.
func (m *MockAuthUseCase) Logout(sessionID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", sessionID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Logout indicates an expected call of Logout.
func (mr *MockAuthUseCaseMockRecorder) Logout(sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAuthUseCase)(nil).Logout), sessionID)
}

// SignUp mocks base method.
func (m *MockAuthUseCase) SignUp(user models.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockAuthUseCaseMockRecorder) SignUp(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockAuthUseCase)(nil).SignUp), user)
}

// MockAuthMicroservice is a mock of AuthMicroservice interface.
type MockAuthMicroservice struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMicroserviceMockRecorder
}

// MockAuthMicroserviceMockRecorder is the mock recorder for MockAuthMicroservice.
type MockAuthMicroserviceMockRecorder struct {
	mock *MockAuthMicroservice
}

// NewMockAuthMicroservice creates a new mock instance.
func NewMockAuthMicroservice(ctrl *gomock.Controller) *MockAuthMicroservice {
	mock := &MockAuthMicroservice{ctrl: ctrl}
	mock.recorder = &MockAuthMicroserviceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthMicroservice) EXPECT() *MockAuthMicroserviceMockRecorder {
	return m.recorder
}

// CreateSession mocks base method.
func (m *MockAuthMicroservice) CreateSession(userID uint64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockAuthMicroserviceMockRecorder) CreateSession(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockAuthMicroservice)(nil).CreateSession), userID)
}

// DeleteBySessionID mocks base method.
func (m *MockAuthMicroservice) DeleteBySessionID(sessionID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBySessionID", sessionID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBySessionID indicates an expected call of DeleteBySessionID.
func (mr *MockAuthMicroserviceMockRecorder) DeleteBySessionID(sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBySessionID", reflect.TypeOf((*MockAuthMicroservice)(nil).DeleteBySessionID), sessionID)
}

// GetBySessionID mocks base method.
func (m *MockAuthMicroservice) GetBySessionID(sessionID string) (*protobuf.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySessionID", sessionID)
	ret0, _ := ret[0].(*protobuf.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySessionID indicates an expected call of GetBySessionID.
func (mr *MockAuthMicroserviceMockRecorder) GetBySessionID(sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySessionID", reflect.TypeOf((*MockAuthMicroservice)(nil).GetBySessionID), sessionID)
}
