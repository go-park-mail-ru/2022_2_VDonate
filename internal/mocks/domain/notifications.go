// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/notifications.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	reflect "reflect"

	models "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockNotificationsUseCase is a mock of NotificationsUseCase interface.
type MockNotificationsUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationsUseCaseMockRecorder
}

// MockNotificationsUseCaseMockRecorder is the mock recorder for MockNotificationsUseCase.
type MockNotificationsUseCaseMockRecorder struct {
	mock *MockNotificationsUseCase
}

// NewMockNotificationsUseCase creates a new mock instance.
func NewMockNotificationsUseCase(ctrl *gomock.Controller) *MockNotificationsUseCase {
	mock := &MockNotificationsUseCase{ctrl: ctrl}
	mock.recorder = &MockNotificationsUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationsUseCase) EXPECT() *MockNotificationsUseCaseMockRecorder {
	return m.recorder
}

// DeleteNotifications mocks base method.
func (m *MockNotificationsUseCase) DeleteNotifications(userID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNotifications", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNotifications indicates an expected call of DeleteNotifications.
func (mr *MockNotificationsUseCaseMockRecorder) DeleteNotifications(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNotifications", reflect.TypeOf((*MockNotificationsUseCase)(nil).DeleteNotifications), userID)
}

// GetNotifications mocks base method.
func (m *MockNotificationsUseCase) GetNotifications(userID uint64) ([]models.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotifications", userID)
	ret0, _ := ret[0].([]models.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotifications indicates an expected call of GetNotifications.
func (mr *MockNotificationsUseCaseMockRecorder) GetNotifications(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotifications", reflect.TypeOf((*MockNotificationsUseCase)(nil).GetNotifications), userID)
}
