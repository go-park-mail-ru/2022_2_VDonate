// Code generated by MockGen. DO NOT EDIT.
// Source: internal/microservices/donates/protobuf/donates_grpc.pb.go

// Package mock_protobuf is a generated GoMock package.
package mock_domain

import (
	context "context"
	reflect "reflect"

	protobuf "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/donates/protobuf"
	protobuf0 "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockDonatesClient is a mock of DonatesClient interface.
type MockDonatesClient struct {
	ctrl     *gomock.Controller
	recorder *MockDonatesClientMockRecorder
}

// MockDonatesClientMockRecorder is the mock recorder for MockDonatesClient.
type MockDonatesClientMockRecorder struct {
	mock *MockDonatesClient
}

// NewMockDonatesClient creates a new mock instance.
func NewMockDonatesClient(ctrl *gomock.Controller) *MockDonatesClient {
	mock := &MockDonatesClient{ctrl: ctrl}
	mock.recorder = &MockDonatesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDonatesClient) EXPECT() *MockDonatesClientMockRecorder {
	return m.recorder
}

// GetDonateByID mocks base method.
func (m *MockDonatesClient) GetDonateByID(ctx context.Context, in *protobuf.DonateID, opts ...grpc.CallOption) (*protobuf.Donate, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDonateByID", varargs...)
	ret0, _ := ret[0].(*protobuf.Donate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDonateByID indicates an expected call of GetDonateByID.
func (mr *MockDonatesClientMockRecorder) GetDonateByID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDonateByID", reflect.TypeOf((*MockDonatesClient)(nil).GetDonateByID), varargs...)
}

// GetDonatesByUserID mocks base method.
func (m *MockDonatesClient) GetDonatesByUserID(ctx context.Context, in *protobuf0.UserID, opts ...grpc.CallOption) (*protobuf.DonateArray, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDonatesByUserID", varargs...)
	ret0, _ := ret[0].(*protobuf.DonateArray)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDonatesByUserID indicates an expected call of GetDonatesByUserID.
func (mr *MockDonatesClientMockRecorder) GetDonatesByUserID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDonatesByUserID", reflect.TypeOf((*MockDonatesClient)(nil).GetDonatesByUserID), varargs...)
}

// SendDonate mocks base method.
func (m *MockDonatesClient) SendDonate(ctx context.Context, in *protobuf.Donate, opts ...grpc.CallOption) (*protobuf.Donate, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendDonate", varargs...)
	ret0, _ := ret[0].(*protobuf.Donate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendDonate indicates an expected call of SendDonate.
func (mr *MockDonatesClientMockRecorder) SendDonate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendDonate", reflect.TypeOf((*MockDonatesClient)(nil).SendDonate), varargs...)
}

// MockDonatesServer is a mock of DonatesServer interface.
type MockDonatesServer struct {
	ctrl     *gomock.Controller
	recorder *MockDonatesServerMockRecorder
}

// MockDonatesServerMockRecorder is the mock recorder for MockDonatesServer.
type MockDonatesServerMockRecorder struct {
	mock *MockDonatesServer
}

// NewMockDonatesServer creates a new mock instance.
func NewMockDonatesServer(ctrl *gomock.Controller) *MockDonatesServer {
	mock := &MockDonatesServer{ctrl: ctrl}
	mock.recorder = &MockDonatesServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDonatesServer) EXPECT() *MockDonatesServerMockRecorder {
	return m.recorder
}

// GetDonateByID mocks base method.
func (m *MockDonatesServer) GetDonateByID(arg0 context.Context, arg1 *protobuf.DonateID) (*protobuf.Donate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDonateByID", arg0, arg1)
	ret0, _ := ret[0].(*protobuf.Donate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDonateByID indicates an expected call of GetDonateByID.
func (mr *MockDonatesServerMockRecorder) GetDonateByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDonateByID", reflect.TypeOf((*MockDonatesServer)(nil).GetDonateByID), arg0, arg1)
}

// GetDonatesByUserID mocks base method.
func (m *MockDonatesServer) GetDonatesByUserID(arg0 context.Context, arg1 *protobuf0.UserID) (*protobuf.DonateArray, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDonatesByUserID", arg0, arg1)
	ret0, _ := ret[0].(*protobuf.DonateArray)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDonatesByUserID indicates an expected call of GetDonatesByUserID.
func (mr *MockDonatesServerMockRecorder) GetDonatesByUserID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDonatesByUserID", reflect.TypeOf((*MockDonatesServer)(nil).GetDonatesByUserID), arg0, arg1)
}

// SendDonate mocks base method.
func (m *MockDonatesServer) SendDonate(arg0 context.Context, arg1 *protobuf.Donate) (*protobuf.Donate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendDonate", arg0, arg1)
	ret0, _ := ret[0].(*protobuf.Donate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendDonate indicates an expected call of SendDonate.
func (mr *MockDonatesServerMockRecorder) SendDonate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendDonate", reflect.TypeOf((*MockDonatesServer)(nil).SendDonate), arg0, arg1)
}

// mustEmbedUnimplementedDonatesServer mocks base method.
func (m *MockDonatesServer) mustEmbedUnimplementedDonatesServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedDonatesServer")
}

// mustEmbedUnimplementedDonatesServer indicates an expected call of mustEmbedUnimplementedDonatesServer.
func (mr *MockDonatesServerMockRecorder) mustEmbedUnimplementedDonatesServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedDonatesServer", reflect.TypeOf((*MockDonatesServer)(nil).mustEmbedUnimplementedDonatesServer))
}

// MockUnsafeDonatesServer is a mock of UnsafeDonatesServer interface.
type MockUnsafeDonatesServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeDonatesServerMockRecorder
}

// MockUnsafeDonatesServerMockRecorder is the mock recorder for MockUnsafeDonatesServer.
type MockUnsafeDonatesServerMockRecorder struct {
	mock *MockUnsafeDonatesServer
}

// NewMockUnsafeDonatesServer creates a new mock instance.
func NewMockUnsafeDonatesServer(ctrl *gomock.Controller) *MockUnsafeDonatesServer {
	mock := &MockUnsafeDonatesServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeDonatesServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeDonatesServer) EXPECT() *MockUnsafeDonatesServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedDonatesServer mocks base method.
func (m *MockUnsafeDonatesServer) mustEmbedUnimplementedDonatesServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedDonatesServer")
}

// mustEmbedUnimplementedDonatesServer indicates an expected call of mustEmbedUnimplementedDonatesServer.
func (mr *MockUnsafeDonatesServerMockRecorder) mustEmbedUnimplementedDonatesServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedDonatesServer", reflect.TypeOf((*MockUnsafeDonatesServer)(nil).mustEmbedUnimplementedDonatesServer))
}
