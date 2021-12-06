// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserSeviceClient is the client API for UserSevice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserSeviceClient interface {
	PostNewUser(ctx context.Context, in *PostUser, opts ...grpc.CallOption) (*User, error)
	GetUserBasicInfo(ctx context.Context, in *GetBasicInfo, opts ...grpc.CallOption) (*User, error)
	GetUsersList(ctx context.Context, in *ListOfPostedUsersReq, opts ...grpc.CallOption) (*ListOfPostedUsersRes, error)
}

type userSeviceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserSeviceClient(cc grpc.ClientConnInterface) UserSeviceClient {
	return &userSeviceClient{cc}
}

func (c *userSeviceClient) PostNewUser(ctx context.Context, in *PostUser, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.UserSevice/PostNewUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSeviceClient) GetUserBasicInfo(ctx context.Context, in *GetBasicInfo, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.UserSevice/GetUserBasicInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSeviceClient) GetUsersList(ctx context.Context, in *ListOfPostedUsersReq, opts ...grpc.CallOption) (*ListOfPostedUsersRes, error) {
	out := new(ListOfPostedUsersRes)
	err := c.cc.Invoke(ctx, "/user.UserSevice/GetUsersList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSeviceServer is the server API for UserSevice service.
// All implementations must embed UnimplementedUserSeviceServer
// for forward compatibility
type UserSeviceServer interface {
	PostNewUser(context.Context, *PostUser) (*User, error)
	GetUserBasicInfo(context.Context, *GetBasicInfo) (*User, error)
	GetUsersList(context.Context, *ListOfPostedUsersReq) (*ListOfPostedUsersRes, error)
	mustEmbedUnimplementedUserSeviceServer()
}

// UnimplementedUserSeviceServer must be embedded to have forward compatible implementations.
type UnimplementedUserSeviceServer struct {
}

func (UnimplementedUserSeviceServer) PostNewUser(context.Context, *PostUser) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostNewUser not implemented")
}
func (UnimplementedUserSeviceServer) GetUserBasicInfo(context.Context, *GetBasicInfo) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBasicInfo not implemented")
}
func (UnimplementedUserSeviceServer) GetUsersList(context.Context, *ListOfPostedUsersReq) (*ListOfPostedUsersRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersList not implemented")
}
func (UnimplementedUserSeviceServer) mustEmbedUnimplementedUserSeviceServer() {}

// UnsafeUserSeviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserSeviceServer will
// result in compilation errors.
type UnsafeUserSeviceServer interface {
	mustEmbedUnimplementedUserSeviceServer()
}

func RegisterUserSeviceServer(s grpc.ServiceRegistrar, srv UserSeviceServer) {
	s.RegisterService(&UserSevice_ServiceDesc, srv)
}

func _UserSevice_PostNewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSeviceServer).PostNewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserSevice/PostNewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSeviceServer).PostNewUser(ctx, req.(*PostUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSevice_GetUserBasicInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBasicInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSeviceServer).GetUserBasicInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserSevice/GetUserBasicInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSeviceServer).GetUserBasicInfo(ctx, req.(*GetBasicInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSevice_GetUsersList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOfPostedUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSeviceServer).GetUsersList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserSevice/GetUsersList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSeviceServer).GetUsersList(ctx, req.(*ListOfPostedUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserSevice_ServiceDesc is the grpc.ServiceDesc for UserSevice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserSevice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserSevice",
	HandlerType: (*UserSeviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostNewUser",
			Handler:    _UserSevice_PostNewUser_Handler,
		},
		{
			MethodName: "GetUserBasicInfo",
			Handler:    _UserSevice_GetUserBasicInfo_Handler,
		},
		{
			MethodName: "GetUsersList",
			Handler:    _UserSevice_GetUsersList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto_files/user/user.proto",
}
