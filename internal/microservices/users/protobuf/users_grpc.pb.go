// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: internal/microservices/users/protobuf/users.proto

package protobuf

import (
	context "context"
	protobuf "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/auth/protobuf"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserID, error)
	GetAuthorByUsername(ctx context.Context, in *Keyword, opts ...grpc.CallOption) (*UsersArray, error)
	GetByID(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*User, error)
	GetBySessionID(ctx context.Context, in *protobuf.SessionID, opts ...grpc.CallOption) (*User, error)
	GetByEmail(ctx context.Context, in *Email, opts ...grpc.CallOption) (*User, error)
	GetByUsername(ctx context.Context, in *Username, opts ...grpc.CallOption) (*User, error)
	GetUserByPostID(ctx context.Context, in *PostID, opts ...grpc.CallOption) (*User, error)
	GetAllAuthors(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UsersArray, error)
	GetPostsNum(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*PostsNum, error)
	GetSubscribersNumForMounth(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*SubscribersNum, error)
	GetProfitForMounth(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Profit, error)
	DropBalance(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.Users/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserID, error) {
	out := new(UserID)
	err := c.cc.Invoke(ctx, "/user.Users/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetAuthorByUsername(ctx context.Context, in *Keyword, opts ...grpc.CallOption) (*UsersArray, error) {
	out := new(UsersArray)
	err := c.cc.Invoke(ctx, "/user.Users/GetAuthorByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetByID(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.Users/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetBySessionID(ctx context.Context, in *protobuf.SessionID, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.Users/GetBySessionID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetByEmail(ctx context.Context, in *Email, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.Users/GetByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetByUsername(ctx context.Context, in *Username, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.Users/GetByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUserByPostID(ctx context.Context, in *PostID, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.Users/GetUserByPostID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetAllAuthors(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UsersArray, error) {
	out := new(UsersArray)
	err := c.cc.Invoke(ctx, "/user.Users/GetAllAuthors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetPostsNum(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*PostsNum, error) {
	out := new(PostsNum)
	err := c.cc.Invoke(ctx, "/user.Users/GetPostsNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetSubscribersNumForMounth(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*SubscribersNum, error) {
	out := new(SubscribersNum)
	err := c.cc.Invoke(ctx, "/user.Users/GetSubscribersNumForMounth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetProfitForMounth(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Profit, error) {
	out := new(Profit)
	err := c.cc.Invoke(ctx, "/user.Users/GetProfitForMounth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) DropBalance(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.Users/DropBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
// All implementations must embed UnimplementedUsersServer
// for forward compatibility
type UsersServer interface {
	Update(context.Context, *User) (*emptypb.Empty, error)
	Create(context.Context, *User) (*UserID, error)
	GetAuthorByUsername(context.Context, *Keyword) (*UsersArray, error)
	GetByID(context.Context, *UserID) (*User, error)
	GetBySessionID(context.Context, *protobuf.SessionID) (*User, error)
	GetByEmail(context.Context, *Email) (*User, error)
	GetByUsername(context.Context, *Username) (*User, error)
	GetUserByPostID(context.Context, *PostID) (*User, error)
	GetAllAuthors(context.Context, *emptypb.Empty) (*UsersArray, error)
	GetPostsNum(context.Context, *UserID) (*PostsNum, error)
	GetSubscribersNumForMounth(context.Context, *UserID) (*SubscribersNum, error)
	GetProfitForMounth(context.Context, *UserID) (*Profit, error)
	DropBalance(context.Context, *UserID) (*emptypb.Empty, error)
	mustEmbedUnimplementedUsersServer()
}

// UnimplementedUsersServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (UnimplementedUsersServer) Update(context.Context, *User) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUsersServer) Create(context.Context, *User) (*UserID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUsersServer) GetAuthorByUsername(context.Context, *Keyword) (*UsersArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorByUsername not implemented")
}
func (UnimplementedUsersServer) GetByID(context.Context, *UserID) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedUsersServer) GetBySessionID(context.Context, *protobuf.SessionID) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBySessionID not implemented")
}
func (UnimplementedUsersServer) GetByEmail(context.Context, *Email) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByEmail not implemented")
}
func (UnimplementedUsersServer) GetByUsername(context.Context, *Username) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUsername not implemented")
}
func (UnimplementedUsersServer) GetUserByPostID(context.Context, *PostID) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByPostID not implemented")
}
func (UnimplementedUsersServer) GetAllAuthors(context.Context, *emptypb.Empty) (*UsersArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAuthors not implemented")
}
func (UnimplementedUsersServer) GetPostsNum(context.Context, *UserID) (*PostsNum, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostsNum not implemented")
}
func (UnimplementedUsersServer) GetSubscribersNumForMounth(context.Context, *UserID) (*SubscribersNum, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscribersNumForMounth not implemented")
}
func (UnimplementedUsersServer) GetProfitForMounth(context.Context, *UserID) (*Profit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfitForMounth not implemented")
}
func (UnimplementedUsersServer) DropBalance(context.Context, *UserID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropBalance not implemented")
}
func (UnimplementedUsersServer) mustEmbedUnimplementedUsersServer() {}

// UnsafeUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServer will
// result in compilation errors.
type UnsafeUsersServer interface {
	mustEmbedUnimplementedUsersServer()
}

func RegisterUsersServer(s grpc.ServiceRegistrar, srv UsersServer) {
	s.RegisterService(&Users_ServiceDesc, srv)
}

func _Users_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Update(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Create(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetAuthorByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Keyword)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetAuthorByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/GetAuthorByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetAuthorByUsername(ctx, req.(*Keyword))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetByID(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetBySessionID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.SessionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetBySessionID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/GetBySessionID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetBySessionID(ctx, req.(*protobuf.SessionID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/GetByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetByEmail(ctx, req.(*Email))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Username)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/GetByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetByUsername(ctx, req.(*Username))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUserByPostID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserByPostID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/GetUserByPostID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserByPostID(ctx, req.(*PostID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetAllAuthors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetAllAuthors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/GetAllAuthors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetAllAuthors(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetPostsNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetPostsNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/GetPostsNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetPostsNum(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetSubscribersNumForMounth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetSubscribersNumForMounth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/GetSubscribersNumForMounth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetSubscribersNumForMounth(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetProfitForMounth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetProfitForMounth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/GetProfitForMounth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetProfitForMounth(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_DropBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).DropBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/DropBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).DropBalance(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

// Users_ServiceDesc is the grpc.ServiceDesc for Users service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Users_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _Users_Update_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Users_Create_Handler,
		},
		{
			MethodName: "GetAuthorByUsername",
			Handler:    _Users_GetAuthorByUsername_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _Users_GetByID_Handler,
		},
		{
			MethodName: "GetBySessionID",
			Handler:    _Users_GetBySessionID_Handler,
		},
		{
			MethodName: "GetByEmail",
			Handler:    _Users_GetByEmail_Handler,
		},
		{
			MethodName: "GetByUsername",
			Handler:    _Users_GetByUsername_Handler,
		},
		{
			MethodName: "GetUserByPostID",
			Handler:    _Users_GetUserByPostID_Handler,
		},
		{
			MethodName: "GetAllAuthors",
			Handler:    _Users_GetAllAuthors_Handler,
		},
		{
			MethodName: "GetPostsNum",
			Handler:    _Users_GetPostsNum_Handler,
		},
		{
			MethodName: "GetSubscribersNumForMounth",
			Handler:    _Users_GetSubscribersNumForMounth_Handler,
		},
		{
			MethodName: "GetProfitForMounth",
			Handler:    _Users_GetProfitForMounth_Handler,
		},
		{
			MethodName: "DropBalance",
			Handler:    _Users_DropBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/microservices/users/protobuf/users.proto",
}
