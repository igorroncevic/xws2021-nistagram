// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// FollowersClient is the client API for Followers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FollowersClient interface {
	CreateUserConnection(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*EmptyResponseFollowers, error)
	CreateUser(ctx context.Context, in *CreateUserRequestFollowers, opts ...grpc.CallOption) (*EmptyResponseFollowers, error)
	DeleteDirectedConnection(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*EmptyResponseFollowers, error)
	DeleteBiDirectedConnection(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*EmptyResponseFollowers, error)
	GetAllFollowers(ctx context.Context, in *CreateUserRequestFollowers, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetAllFollowing(ctx context.Context, in *CreateUserRequestFollowers, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetAllFollowingsForHomepage(ctx context.Context, in *CreateUserRequestFollowers, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetCloseFriends(ctx context.Context, in *RequestIdFollowers, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetCloseFriendsReversed(ctx context.Context, in *RequestIdFollowers, opts ...grpc.CallOption) (*CreateUserResponse, error)
	UpdateUserConnection(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*CreateFollowerResponse, error)
	AcceptFollowRequest(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*CreateFollowerResponse, error)
	GetFollowersConnection(ctx context.Context, in *Follower, opts ...grpc.CallOption) (*Follower, error)
	GetUsersForNotificationEnabled(ctx context.Context, in *RequestForNotification, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type followersClient struct {
	cc grpc.ClientConnInterface
}

func NewFollowersClient(cc grpc.ClientConnInterface) FollowersClient {
	return &followersClient{cc}
}

func (c *followersClient) CreateUserConnection(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*EmptyResponseFollowers, error) {
	out := new(EmptyResponseFollowers)
	err := c.cc.Invoke(ctx, "/proto.Followers/CreateUserConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersClient) CreateUser(ctx context.Context, in *CreateUserRequestFollowers, opts ...grpc.CallOption) (*EmptyResponseFollowers, error) {
	out := new(EmptyResponseFollowers)
	err := c.cc.Invoke(ctx, "/proto.Followers/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersClient) DeleteDirectedConnection(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*EmptyResponseFollowers, error) {
	out := new(EmptyResponseFollowers)
	err := c.cc.Invoke(ctx, "/proto.Followers/DeleteDirectedConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersClient) DeleteBiDirectedConnection(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*EmptyResponseFollowers, error) {
	out := new(EmptyResponseFollowers)
	err := c.cc.Invoke(ctx, "/proto.Followers/DeleteBiDirectedConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersClient) GetAllFollowers(ctx context.Context, in *CreateUserRequestFollowers, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/proto.Followers/GetAllFollowers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersClient) GetAllFollowing(ctx context.Context, in *CreateUserRequestFollowers, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/proto.Followers/GetAllFollowing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersClient) GetAllFollowingsForHomepage(ctx context.Context, in *CreateUserRequestFollowers, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/proto.Followers/GetAllFollowingsForHomepage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersClient) GetCloseFriends(ctx context.Context, in *RequestIdFollowers, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/proto.Followers/GetCloseFriends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersClient) GetCloseFriendsReversed(ctx context.Context, in *RequestIdFollowers, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/proto.Followers/GetCloseFriendsReversed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersClient) UpdateUserConnection(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*CreateFollowerResponse, error) {
	out := new(CreateFollowerResponse)
	err := c.cc.Invoke(ctx, "/proto.Followers/UpdateUserConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersClient) AcceptFollowRequest(ctx context.Context, in *CreateFollowerRequest, opts ...grpc.CallOption) (*CreateFollowerResponse, error) {
	out := new(CreateFollowerResponse)
	err := c.cc.Invoke(ctx, "/proto.Followers/AcceptFollowRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersClient) GetFollowersConnection(ctx context.Context, in *Follower, opts ...grpc.CallOption) (*Follower, error) {
	out := new(Follower)
	err := c.cc.Invoke(ctx, "/proto.Followers/GetFollowersConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersClient) GetUsersForNotificationEnabled(ctx context.Context, in *RequestForNotification, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/proto.Followers/GetUsersForNotificationEnabled", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FollowersServer is the server API for Followers service.
// All implementations must embed UnimplementedFollowersServer
// for forward compatibility
type FollowersServer interface {
	CreateUserConnection(context.Context, *CreateFollowerRequest) (*EmptyResponseFollowers, error)
	CreateUser(context.Context, *CreateUserRequestFollowers) (*EmptyResponseFollowers, error)
	DeleteDirectedConnection(context.Context, *CreateFollowerRequest) (*EmptyResponseFollowers, error)
	DeleteBiDirectedConnection(context.Context, *CreateFollowerRequest) (*EmptyResponseFollowers, error)
	GetAllFollowers(context.Context, *CreateUserRequestFollowers) (*CreateUserResponse, error)
	GetAllFollowing(context.Context, *CreateUserRequestFollowers) (*CreateUserResponse, error)
	GetAllFollowingsForHomepage(context.Context, *CreateUserRequestFollowers) (*CreateUserResponse, error)
	GetCloseFriends(context.Context, *RequestIdFollowers) (*CreateUserResponse, error)
	GetCloseFriendsReversed(context.Context, *RequestIdFollowers) (*CreateUserResponse, error)
	UpdateUserConnection(context.Context, *CreateFollowerRequest) (*CreateFollowerResponse, error)
	AcceptFollowRequest(context.Context, *CreateFollowerRequest) (*CreateFollowerResponse, error)
	GetFollowersConnection(context.Context, *Follower) (*Follower, error)
	GetUsersForNotificationEnabled(context.Context, *RequestForNotification) (*CreateUserResponse, error)
	mustEmbedUnimplementedFollowersServer()
}

// UnimplementedFollowersServer must be embedded to have forward compatible implementations.
type UnimplementedFollowersServer struct {
}

func (UnimplementedFollowersServer) CreateUserConnection(context.Context, *CreateFollowerRequest) (*EmptyResponseFollowers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserConnection not implemented")
}
func (UnimplementedFollowersServer) CreateUser(context.Context, *CreateUserRequestFollowers) (*EmptyResponseFollowers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedFollowersServer) DeleteDirectedConnection(context.Context, *CreateFollowerRequest) (*EmptyResponseFollowers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDirectedConnection not implemented")
}
func (UnimplementedFollowersServer) DeleteBiDirectedConnection(context.Context, *CreateFollowerRequest) (*EmptyResponseFollowers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBiDirectedConnection not implemented")
}
func (UnimplementedFollowersServer) GetAllFollowers(context.Context, *CreateUserRequestFollowers) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFollowers not implemented")
}
func (UnimplementedFollowersServer) GetAllFollowing(context.Context, *CreateUserRequestFollowers) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFollowing not implemented")
}
func (UnimplementedFollowersServer) GetAllFollowingsForHomepage(context.Context, *CreateUserRequestFollowers) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFollowingsForHomepage not implemented")
}
func (UnimplementedFollowersServer) GetCloseFriends(context.Context, *RequestIdFollowers) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCloseFriends not implemented")
}
func (UnimplementedFollowersServer) GetCloseFriendsReversed(context.Context, *RequestIdFollowers) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCloseFriendsReversed not implemented")
}
func (UnimplementedFollowersServer) UpdateUserConnection(context.Context, *CreateFollowerRequest) (*CreateFollowerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserConnection not implemented")
}
func (UnimplementedFollowersServer) AcceptFollowRequest(context.Context, *CreateFollowerRequest) (*CreateFollowerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptFollowRequest not implemented")
}
func (UnimplementedFollowersServer) GetFollowersConnection(context.Context, *Follower) (*Follower, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowersConnection not implemented")
}
func (UnimplementedFollowersServer) GetUsersForNotificationEnabled(context.Context, *RequestForNotification) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersForNotificationEnabled not implemented")
}
func (UnimplementedFollowersServer) mustEmbedUnimplementedFollowersServer() {}

// UnsafeFollowersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FollowersServer will
// result in compilation errors.
type UnsafeFollowersServer interface {
	mustEmbedUnimplementedFollowersServer()
}

func RegisterFollowersServer(s grpc.ServiceRegistrar, srv FollowersServer) {
	s.RegisterService(&Followers_ServiceDesc, srv)
}

func _Followers_CreateUserConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFollowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).CreateUserConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Followers/CreateUserConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).CreateUserConnection(ctx, req.(*CreateFollowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followers_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequestFollowers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Followers/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).CreateUser(ctx, req.(*CreateUserRequestFollowers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followers_DeleteDirectedConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFollowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).DeleteDirectedConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Followers/DeleteDirectedConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).DeleteDirectedConnection(ctx, req.(*CreateFollowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followers_DeleteBiDirectedConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFollowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).DeleteBiDirectedConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Followers/DeleteBiDirectedConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).DeleteBiDirectedConnection(ctx, req.(*CreateFollowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followers_GetAllFollowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequestFollowers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).GetAllFollowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Followers/GetAllFollowers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).GetAllFollowers(ctx, req.(*CreateUserRequestFollowers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followers_GetAllFollowing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequestFollowers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).GetAllFollowing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Followers/GetAllFollowing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).GetAllFollowing(ctx, req.(*CreateUserRequestFollowers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followers_GetAllFollowingsForHomepage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequestFollowers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).GetAllFollowingsForHomepage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Followers/GetAllFollowingsForHomepage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).GetAllFollowingsForHomepage(ctx, req.(*CreateUserRequestFollowers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followers_GetCloseFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestIdFollowers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).GetCloseFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Followers/GetCloseFriends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).GetCloseFriends(ctx, req.(*RequestIdFollowers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followers_GetCloseFriendsReversed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestIdFollowers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).GetCloseFriendsReversed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Followers/GetCloseFriendsReversed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).GetCloseFriendsReversed(ctx, req.(*RequestIdFollowers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followers_UpdateUserConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFollowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).UpdateUserConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Followers/UpdateUserConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).UpdateUserConnection(ctx, req.(*CreateFollowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followers_AcceptFollowRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFollowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).AcceptFollowRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Followers/AcceptFollowRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).AcceptFollowRequest(ctx, req.(*CreateFollowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followers_GetFollowersConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Follower)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).GetFollowersConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Followers/GetFollowersConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).GetFollowersConnection(ctx, req.(*Follower))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followers_GetUsersForNotificationEnabled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestForNotification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServer).GetUsersForNotificationEnabled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Followers/GetUsersForNotificationEnabled",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServer).GetUsersForNotificationEnabled(ctx, req.(*RequestForNotification))
	}
	return interceptor(ctx, in, info, handler)
}

// Followers_ServiceDesc is the grpc.ServiceDesc for Followers service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Followers_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Followers",
	HandlerType: (*FollowersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserConnection",
			Handler:    _Followers_CreateUserConnection_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Followers_CreateUser_Handler,
		},
		{
			MethodName: "DeleteDirectedConnection",
			Handler:    _Followers_DeleteDirectedConnection_Handler,
		},
		{
			MethodName: "DeleteBiDirectedConnection",
			Handler:    _Followers_DeleteBiDirectedConnection_Handler,
		},
		{
			MethodName: "GetAllFollowers",
			Handler:    _Followers_GetAllFollowers_Handler,
		},
		{
			MethodName: "GetAllFollowing",
			Handler:    _Followers_GetAllFollowing_Handler,
		},
		{
			MethodName: "GetAllFollowingsForHomepage",
			Handler:    _Followers_GetAllFollowingsForHomepage_Handler,
		},
		{
			MethodName: "GetCloseFriends",
			Handler:    _Followers_GetCloseFriends_Handler,
		},
		{
			MethodName: "GetCloseFriendsReversed",
			Handler:    _Followers_GetCloseFriendsReversed_Handler,
		},
		{
			MethodName: "UpdateUserConnection",
			Handler:    _Followers_UpdateUserConnection_Handler,
		},
		{
			MethodName: "AcceptFollowRequest",
			Handler:    _Followers_AcceptFollowRequest_Handler,
		},
		{
			MethodName: "GetFollowersConnection",
			Handler:    _Followers_GetFollowersConnection_Handler,
		},
		{
			MethodName: "GetUsersForNotificationEnabled",
			Handler:    _Followers_GetUsersForNotificationEnabled_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "followers.proto",
}
