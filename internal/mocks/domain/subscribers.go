// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/subscribers.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	reflect "reflect"

	models "github.com/go-park-mail-ru/2022_2_VDonate/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockSubscribersUseCase is a mock of SubscribersUseCase interface.
type MockSubscribersUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSubscribersUseCaseMockRecorder
}

// MockSubscribersUseCaseMockRecorder is the mock recorder for MockSubscribersUseCase.
type MockSubscribersUseCaseMockRecorder struct {
	mock *MockSubscribersUseCase
}

// NewMockSubscribersUseCase creates a new mock instance.
func NewMockSubscribersUseCase(ctrl *gomock.Controller) *MockSubscribersUseCase {
	mock := &MockSubscribersUseCase{ctrl: ctrl}
	mock.recorder = &MockSubscribersUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscribersUseCase) EXPECT() *MockSubscribersUseCaseMockRecorder {
	return m.recorder
}

// CardValidation mocks base method.
func (m *MockSubscribersUseCase) CardValidation(card string) (models.WithdrawValidation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CardValidation", card)
	ret0, _ := ret[0].(models.WithdrawValidation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CardValidation indicates an expected call of CardValidation.
func (mr *MockSubscribersUseCaseMockRecorder) CardValidation(card interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CardValidation", reflect.TypeOf((*MockSubscribersUseCase)(nil).CardValidation), card)
}

// GetSubscribers mocks base method.
func (m *MockSubscribersUseCase) GetSubscribers(authorID uint64) ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscribers", authorID)
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscribers indicates an expected call of GetSubscribers.
func (mr *MockSubscribersUseCaseMockRecorder) GetSubscribers(authorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribers", reflect.TypeOf((*MockSubscribersUseCase)(nil).GetSubscribers), authorID)
}

// IsSubscriber mocks base method.
func (m *MockSubscribersUseCase) IsSubscriber(userID, authorID uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSubscriber", userID, authorID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSubscriber indicates an expected call of IsSubscriber.
func (mr *MockSubscribersUseCaseMockRecorder) IsSubscriber(userID, authorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSubscriber", reflect.TypeOf((*MockSubscribersUseCase)(nil).IsSubscriber), userID, authorID)
}

// Subscribe mocks base method.
func (m *MockSubscribersUseCase) Subscribe(subscription models.Subscription, userID uint64, as models.AuthorSubscription) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", subscription, userID, as)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockSubscribersUseCaseMockRecorder) Subscribe(subscription, userID, as interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockSubscribersUseCase)(nil).Subscribe), subscription, userID, as)
}

// Unsubscribe mocks base method.
func (m *MockSubscribersUseCase) Unsubscribe(userID, authorID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", userID, authorID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockSubscribersUseCaseMockRecorder) Unsubscribe(userID, authorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockSubscribersUseCase)(nil).Unsubscribe), userID, authorID)
}

// Withdraw mocks base method.
func (m *MockSubscribersUseCase) Withdraw(userID uint64, phone, card string) (models.WithdrawInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Withdraw", userID, phone, card)
	ret0, _ := ret[0].(models.WithdrawInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Withdraw indicates an expected call of Withdraw.
func (mr *MockSubscribersUseCaseMockRecorder) Withdraw(userID, phone, card interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Withdraw", reflect.TypeOf((*MockSubscribersUseCase)(nil).Withdraw), userID, phone, card)
}

// WithdrawCard mocks base method.
func (m *MockSubscribersUseCase) WithdrawCard(userID uint64, card, provider string) (models.WithdrawInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithdrawCard", userID, card, provider)
	ret0, _ := ret[0].(models.WithdrawInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithdrawCard indicates an expected call of WithdrawCard.
func (mr *MockSubscribersUseCaseMockRecorder) WithdrawCard(userID, card, provider interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawCard", reflect.TypeOf((*MockSubscribersUseCase)(nil).WithdrawCard), userID, card, provider)
}

// WithdrawQiwi mocks base method.
func (m *MockSubscribersUseCase) WithdrawQiwi(userID uint64, phone string) (models.WithdrawInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithdrawQiwi", userID, phone)
	ret0, _ := ret[0].(models.WithdrawInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithdrawQiwi indicates an expected call of WithdrawQiwi.
func (mr *MockSubscribersUseCaseMockRecorder) WithdrawQiwi(userID, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawQiwi", reflect.TypeOf((*MockSubscribersUseCase)(nil).WithdrawQiwi), userID, phone)
}

// MockSubscribersMicroservice is a mock of SubscribersMicroservice interface.
type MockSubscribersMicroservice struct {
	ctrl     *gomock.Controller
	recorder *MockSubscribersMicroserviceMockRecorder
}

// MockSubscribersMicroserviceMockRecorder is the mock recorder for MockSubscribersMicroservice.
type MockSubscribersMicroserviceMockRecorder struct {
	mock *MockSubscribersMicroservice
}

// NewMockSubscribersMicroservice creates a new mock instance.
func NewMockSubscribersMicroservice(ctrl *gomock.Controller) *MockSubscribersMicroservice {
	mock := &MockSubscribersMicroservice{ctrl: ctrl}
	mock.recorder = &MockSubscribersMicroserviceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscribersMicroservice) EXPECT() *MockSubscribersMicroserviceMockRecorder {
	return m.recorder
}

// GetSubscribers mocks base method.
func (m *MockSubscribersMicroservice) GetSubscribers(authorID uint64) ([]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscribers", authorID)
	ret0, _ := ret[0].([]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscribers indicates an expected call of GetSubscribers.
func (mr *MockSubscribersMicroserviceMockRecorder) GetSubscribers(authorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribers", reflect.TypeOf((*MockSubscribersMicroservice)(nil).GetSubscribers), authorID)
}

// Subscribe mocks base method.
func (m *MockSubscribersMicroservice) Subscribe(payment models.Payment) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Subscribe", payment)
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockSubscribersMicroserviceMockRecorder) Subscribe(payment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockSubscribersMicroservice)(nil).Subscribe), payment)
}

// Unsubscribe mocks base method.
func (m *MockSubscribersMicroservice) Unsubscribe(subscription models.Subscription) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", subscription)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockSubscribersMicroserviceMockRecorder) Unsubscribe(subscription interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockSubscribersMicroservice)(nil).Unsubscribe), subscription)
}
