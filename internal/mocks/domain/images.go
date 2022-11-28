// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/images.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	multipart "mime/multipart"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockImageUseCase is a mock of ImageUseCase interface.
type MockImageUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockImageUseCaseMockRecorder
}

// MockImageUseCaseMockRecorder is the mock recorder for MockImageUseCase.
type MockImageUseCaseMockRecorder struct {
	mock *MockImageUseCase
}

// NewMockImageUseCase creates a new mock instance.
func NewMockImageUseCase(ctrl *gomock.Controller) *MockImageUseCase {
	mock := &MockImageUseCase{ctrl: ctrl}
	mock.recorder = &MockImageUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageUseCase) EXPECT() *MockImageUseCaseMockRecorder {
	return m.recorder
}

// CreateOrUpdateImage mocks base method.
func (m *MockImageUseCase) CreateOrUpdateImage(image *multipart.FileHeader, oldFilename string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateImage", image, oldFilename)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdateImage indicates an expected call of CreateOrUpdateImage.
func (mr *MockImageUseCaseMockRecorder) CreateOrUpdateImage(image, oldFilename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateImage", reflect.TypeOf((*MockImageUseCase)(nil).CreateOrUpdateImage), image, oldFilename)
}

// GetImage mocks base method.
func (m *MockImageUseCase) GetImage(filename string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImage", filename)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImage indicates an expected call of GetImage.
func (mr *MockImageUseCaseMockRecorder) GetImage(filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImage", reflect.TypeOf((*MockImageUseCase)(nil).GetImage), filename)
}
