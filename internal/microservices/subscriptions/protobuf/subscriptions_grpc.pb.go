// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: internal/microservices/subscriptions/protobuf/subscriptions.proto

package protobuf

import (
	context "context"
	protobuf "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SubscriptionsClient is the client API for Subscriptions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscriptionsClient interface {
	GetSubscriptionsByUserID(ctx context.Context, in *protobuf.UserID, opts ...grpc.CallOption) (*SubArray, error)
	GetSubscriptionsByAuthorID(ctx context.Context, in *protobuf.UserID, opts ...grpc.CallOption) (*SubArray, error)
	GetSubscriptionByID(ctx context.Context, in *AuthorSubscriptionID, opts ...grpc.CallOption) (*AuthorSubscription, error)
	GetSubscriptionByUserAndAuthorID(ctx context.Context, in *protobuf.UserAuthorPair, opts ...grpc.CallOption) (*AuthorSubscription, error)
	AddSubscription(ctx context.Context, in *AuthorSubscription, opts ...grpc.CallOption) (*AuthorSubscriptionID, error)
	UpdateSubscription(ctx context.Context, in *AuthorSubscription, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteSubscription(ctx context.Context, in *AuthorSubscriptionID, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type subscriptionsClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriptionsClient(cc grpc.ClientConnInterface) SubscriptionsClient {
	return &subscriptionsClient{cc}
}

func (c *subscriptionsClient) GetSubscriptionsByUserID(ctx context.Context, in *protobuf.UserID, opts ...grpc.CallOption) (*SubArray, error) {
	out := new(SubArray)
	err := c.cc.Invoke(ctx, "/subscriptions.Subscriptions/GetSubscriptionsByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionsClient) GetSubscriptionsByAuthorID(ctx context.Context, in *protobuf.UserID, opts ...grpc.CallOption) (*SubArray, error) {
	out := new(SubArray)
	err := c.cc.Invoke(ctx, "/subscriptions.Subscriptions/GetSubscriptionsByAuthorID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionsClient) GetSubscriptionByID(ctx context.Context, in *AuthorSubscriptionID, opts ...grpc.CallOption) (*AuthorSubscription, error) {
	out := new(AuthorSubscription)
	err := c.cc.Invoke(ctx, "/subscriptions.Subscriptions/GetSubscriptionByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionsClient) GetSubscriptionByUserAndAuthorID(ctx context.Context, in *protobuf.UserAuthorPair, opts ...grpc.CallOption) (*AuthorSubscription, error) {
	out := new(AuthorSubscription)
	err := c.cc.Invoke(ctx, "/subscriptions.Subscriptions/GetSubscriptionByUserAndAuthorID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionsClient) AddSubscription(ctx context.Context, in *AuthorSubscription, opts ...grpc.CallOption) (*AuthorSubscriptionID, error) {
	out := new(AuthorSubscriptionID)
	err := c.cc.Invoke(ctx, "/subscriptions.Subscriptions/AddSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionsClient) UpdateSubscription(ctx context.Context, in *AuthorSubscription, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/subscriptions.Subscriptions/UpdateSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionsClient) DeleteSubscription(ctx context.Context, in *AuthorSubscriptionID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/subscriptions.Subscriptions/DeleteSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionsServer is the server API for Subscriptions service.
// All implementations must embed UnimplementedSubscriptionsServer
// for forward compatibility
type SubscriptionsServer interface {
	GetSubscriptionsByUserID(context.Context, *protobuf.UserID) (*SubArray, error)
	GetSubscriptionsByAuthorID(context.Context, *protobuf.UserID) (*SubArray, error)
	GetSubscriptionByID(context.Context, *AuthorSubscriptionID) (*AuthorSubscription, error)
	GetSubscriptionByUserAndAuthorID(context.Context, *protobuf.UserAuthorPair) (*AuthorSubscription, error)
	AddSubscription(context.Context, *AuthorSubscription) (*AuthorSubscriptionID, error)
	UpdateSubscription(context.Context, *AuthorSubscription) (*emptypb.Empty, error)
	DeleteSubscription(context.Context, *AuthorSubscriptionID) (*emptypb.Empty, error)
	mustEmbedUnimplementedSubscriptionsServer()
}

// UnimplementedSubscriptionsServer must be embedded to have forward compatible implementations.
type UnimplementedSubscriptionsServer struct {
}

func (UnimplementedSubscriptionsServer) GetSubscriptionsByUserID(context.Context, *protobuf.UserID) (*SubArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriptionsByUserID not implemented")
}
func (UnimplementedSubscriptionsServer) GetSubscriptionsByAuthorID(context.Context, *protobuf.UserID) (*SubArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriptionsByAuthorID not implemented")
}
func (UnimplementedSubscriptionsServer) GetSubscriptionByID(context.Context, *AuthorSubscriptionID) (*AuthorSubscription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriptionByID not implemented")
}
func (UnimplementedSubscriptionsServer) GetSubscriptionByUserAndAuthorID(context.Context, *protobuf.UserAuthorPair) (*AuthorSubscription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriptionByUserAndAuthorID not implemented")
}
func (UnimplementedSubscriptionsServer) AddSubscription(context.Context, *AuthorSubscription) (*AuthorSubscriptionID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSubscription not implemented")
}
func (UnimplementedSubscriptionsServer) UpdateSubscription(context.Context, *AuthorSubscription) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubscription not implemented")
}
func (UnimplementedSubscriptionsServer) DeleteSubscription(context.Context, *AuthorSubscriptionID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubscription not implemented")
}
func (UnimplementedSubscriptionsServer) mustEmbedUnimplementedSubscriptionsServer() {}

// UnsafeSubscriptionsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriptionsServer will
// result in compilation errors.
type UnsafeSubscriptionsServer interface {
	mustEmbedUnimplementedSubscriptionsServer()
}

func RegisterSubscriptionsServer(s grpc.ServiceRegistrar, srv SubscriptionsServer) {
	s.RegisterService(&Subscriptions_ServiceDesc, srv)
}

func _Subscriptions_GetSubscriptionsByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionsServer).GetSubscriptionsByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscriptions.Subscriptions/GetSubscriptionsByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionsServer).GetSubscriptionsByUserID(ctx, req.(*protobuf.UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscriptions_GetSubscriptionsByAuthorID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionsServer).GetSubscriptionsByAuthorID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscriptions.Subscriptions/GetSubscriptionsByAuthorID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionsServer).GetSubscriptionsByAuthorID(ctx, req.(*protobuf.UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscriptions_GetSubscriptionByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorSubscriptionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionsServer).GetSubscriptionByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscriptions.Subscriptions/GetSubscriptionByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionsServer).GetSubscriptionByID(ctx, req.(*AuthorSubscriptionID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscriptions_GetSubscriptionByUserAndAuthorID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.UserAuthorPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionsServer).GetSubscriptionByUserAndAuthorID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscriptions.Subscriptions/GetSubscriptionByUserAndAuthorID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionsServer).GetSubscriptionByUserAndAuthorID(ctx, req.(*protobuf.UserAuthorPair))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscriptions_AddSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorSubscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionsServer).AddSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscriptions.Subscriptions/AddSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionsServer).AddSubscription(ctx, req.(*AuthorSubscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscriptions_UpdateSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorSubscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionsServer).UpdateSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscriptions.Subscriptions/UpdateSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionsServer).UpdateSubscription(ctx, req.(*AuthorSubscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscriptions_DeleteSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorSubscriptionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionsServer).DeleteSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscriptions.Subscriptions/DeleteSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionsServer).DeleteSubscription(ctx, req.(*AuthorSubscriptionID))
	}
	return interceptor(ctx, in, info, handler)
}

// Subscriptions_ServiceDesc is the grpc.ServiceDesc for Subscriptions service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Subscriptions_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "subscriptions.Subscriptions",
	HandlerType: (*SubscriptionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSubscriptionsByUserID",
			Handler:    _Subscriptions_GetSubscriptionsByUserID_Handler,
		},
		{
			MethodName: "GetSubscriptionsByAuthorID",
			Handler:    _Subscriptions_GetSubscriptionsByAuthorID_Handler,
		},
		{
			MethodName: "GetSubscriptionByID",
			Handler:    _Subscriptions_GetSubscriptionByID_Handler,
		},
		{
			MethodName: "GetSubscriptionByUserAndAuthorID",
			Handler:    _Subscriptions_GetSubscriptionByUserAndAuthorID_Handler,
		},
		{
			MethodName: "AddSubscription",
			Handler:    _Subscriptions_AddSubscription_Handler,
		},
		{
			MethodName: "UpdateSubscription",
			Handler:    _Subscriptions_UpdateSubscription_Handler,
		},
		{
			MethodName: "DeleteSubscription",
			Handler:    _Subscriptions_DeleteSubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/microservices/subscriptions/protobuf/subscriptions.proto",
}
