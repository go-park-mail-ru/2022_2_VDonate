// Code generated by MockGen. DO NOT EDIT.
// Source: internal/microservices/subscribers/protobuf/subscribers_grpc.pb.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	reflect "reflect"

	protobuf "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/subscribers/protobuf"
	protobuf0 "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockSubscribersClient is a mock of SubscribersClient interface.
type MockSubscribersClient struct {
	ctrl     *gomock.Controller
	recorder *MockSubscribersClientMockRecorder
}

// MockSubscribersClientMockRecorder is the mock recorder for MockSubscribersClient.
type MockSubscribersClientMockRecorder struct {
	mock *MockSubscribersClient
}

// NewMockSubscribersClient creates a new mock instance.
func NewMockSubscribersClient(ctrl *gomock.Controller) *MockSubscribersClient {
	mock := &MockSubscribersClient{ctrl: ctrl}
	mock.recorder = &MockSubscribersClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscribersClient) EXPECT() *MockSubscribersClientMockRecorder {
	return m.recorder
}

// ChangePaymentStatus mocks base method.
func (m *MockSubscribersClient) ChangePaymentStatus(ctx context.Context, in *protobuf.StatusAndID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangePaymentStatus", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangePaymentStatus indicates an expected call of ChangePaymentStatus.
func (mr *MockSubscribersClientMockRecorder) ChangePaymentStatus(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePaymentStatus", reflect.TypeOf((*MockSubscribersClient)(nil).ChangePaymentStatus), varargs...)
}

// GetSubscribers mocks base method.
func (m *MockSubscribersClient) GetSubscribers(ctx context.Context, in *protobuf0.UserID, opts ...grpc.CallOption) (*protobuf0.UserIDs, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSubscribers", varargs...)
	ret0, _ := ret[0].(*protobuf0.UserIDs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscribers indicates an expected call of GetSubscribers.
func (mr *MockSubscribersClientMockRecorder) GetSubscribers(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribers", reflect.TypeOf((*MockSubscribersClient)(nil).GetSubscribers), varargs...)
}

// Subscribe mocks base method.
func (m *MockSubscribersClient) Subscribe(ctx context.Context, in *protobuf.Payment, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockSubscribersClientMockRecorder) Subscribe(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockSubscribersClient)(nil).Subscribe), varargs...)
}

// Unsubscribe mocks base method.
func (m *MockSubscribersClient) Unsubscribe(ctx context.Context, in *protobuf0.UserAuthorPair, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Unsubscribe", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockSubscribersClientMockRecorder) Unsubscribe(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockSubscribersClient)(nil).Unsubscribe), varargs...)
}

// MockSubscribersServer is a mock of SubscribersServer interface.
type MockSubscribersServer struct {
	ctrl     *gomock.Controller
	recorder *MockSubscribersServerMockRecorder
}

// MockSubscribersServerMockRecorder is the mock recorder for MockSubscribersServer.
type MockSubscribersServerMockRecorder struct {
	mock *MockSubscribersServer
}

// NewMockSubscribersServer creates a new mock instance.
func NewMockSubscribersServer(ctrl *gomock.Controller) *MockSubscribersServer {
	mock := &MockSubscribersServer{ctrl: ctrl}
	mock.recorder = &MockSubscribersServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscribersServer) EXPECT() *MockSubscribersServerMockRecorder {
	return m.recorder
}

// ChangePaymentStatus mocks base method.
func (m *MockSubscribersServer) ChangePaymentStatus(arg0 context.Context, arg1 *protobuf.StatusAndID) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePaymentStatus", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangePaymentStatus indicates an expected call of ChangePaymentStatus.
func (mr *MockSubscribersServerMockRecorder) ChangePaymentStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePaymentStatus", reflect.TypeOf((*MockSubscribersServer)(nil).ChangePaymentStatus), arg0, arg1)
}

// GetSubscribers mocks base method.
func (m *MockSubscribersServer) GetSubscribers(arg0 context.Context, arg1 *protobuf0.UserID) (*protobuf0.UserIDs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscribers", arg0, arg1)
	ret0, _ := ret[0].(*protobuf0.UserIDs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscribers indicates an expected call of GetSubscribers.
func (mr *MockSubscribersServerMockRecorder) GetSubscribers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribers", reflect.TypeOf((*MockSubscribersServer)(nil).GetSubscribers), arg0, arg1)
}

// Subscribe mocks base method.
func (m *MockSubscribersServer) Subscribe(arg0 context.Context, arg1 *protobuf.Payment) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockSubscribersServerMockRecorder) Subscribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockSubscribersServer)(nil).Subscribe), arg0, arg1)
}

// Unsubscribe mocks base method.
func (m *MockSubscribersServer) Unsubscribe(arg0 context.Context, arg1 *protobuf0.UserAuthorPair) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockSubscribersServerMockRecorder) Unsubscribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockSubscribersServer)(nil).Unsubscribe), arg0, arg1)
}

// mustEmbedUnimplementedSubscribersServer mocks base method.
func (m *MockSubscribersServer) mustEmbedUnimplementedSubscribersServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedSubscribersServer")
}

// mustEmbedUnimplementedSubscribersServer indicates an expected call of mustEmbedUnimplementedSubscribersServer.
func (mr *MockSubscribersServerMockRecorder) mustEmbedUnimplementedSubscribersServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedSubscribersServer", reflect.TypeOf((*MockSubscribersServer)(nil).mustEmbedUnimplementedSubscribersServer))
}

// MockUnsafeSubscribersServer is a mock of UnsafeSubscribersServer interface.
type MockUnsafeSubscribersServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeSubscribersServerMockRecorder
}

// MockUnsafeSubscribersServerMockRecorder is the mock recorder for MockUnsafeSubscribersServer.
type MockUnsafeSubscribersServerMockRecorder struct {
	mock *MockUnsafeSubscribersServer
}

// NewMockUnsafeSubscribersServer creates a new mock instance.
func NewMockUnsafeSubscribersServer(ctrl *gomock.Controller) *MockUnsafeSubscribersServer {
	mock := &MockUnsafeSubscribersServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeSubscribersServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeSubscribersServer) EXPECT() *MockUnsafeSubscribersServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedSubscribersServer mocks base method.
func (m *MockUnsafeSubscribersServer) mustEmbedUnimplementedSubscribersServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedSubscribersServer")
}

// mustEmbedUnimplementedSubscribersServer indicates an expected call of mustEmbedUnimplementedSubscribersServer.
func (mr *MockUnsafeSubscribersServerMockRecorder) mustEmbedUnimplementedSubscribersServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedSubscribersServer", reflect.TypeOf((*MockUnsafeSubscribersServer)(nil).mustEmbedUnimplementedSubscribersServer))
}
