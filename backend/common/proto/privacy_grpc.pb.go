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
const _ = grpc.SupportPackageIsVersion7

// PrivacyClient is the client API for Privacy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrivacyClient interface {
	CreatePrivacy(ctx context.Context, in *CreatePrivacyRequest, opts ...grpc.CallOption) (*EmptyResponsePrivacy, error)
	UpdatePrivacy(ctx context.Context, in *CreatePrivacyRequest, opts ...grpc.CallOption) (*EmptyResponsePrivacy, error)
	BlockUser(ctx context.Context, in *CreateBlockRequest, opts ...grpc.CallOption) (*EmptyResponsePrivacy, error)
	UnBlockUser(ctx context.Context, in *CreateBlockRequest, opts ...grpc.CallOption) (*EmptyResponsePrivacy, error)
	CheckIfBlocked(ctx context.Context, in *CreateBlockRequest, opts ...grpc.CallOption) (*BooleanResponse, error)
	CheckUserProfilePublic(ctx context.Context, in *PrivacyRequest, opts ...grpc.CallOption) (*BooleanResponse, error)
	GetAllPublicUsers(ctx context.Context, in *RequestIdPrivacy, opts ...grpc.CallOption) (*StringArray, error)
	GetUserPrivacy(ctx context.Context, in *RequestIdPrivacy, opts ...grpc.CallOption) (*PrivacyMessage, error)
}

type privacyClient struct {
	cc grpc.ClientConnInterface
}

func NewPrivacyClient(cc grpc.ClientConnInterface) PrivacyClient {
	return &privacyClient{cc}
}

func (c *privacyClient) CreatePrivacy(ctx context.Context, in *CreatePrivacyRequest, opts ...grpc.CallOption) (*EmptyResponsePrivacy, error) {
	out := new(EmptyResponsePrivacy)
	err := c.cc.Invoke(ctx, "/proto.Privacy/CreatePrivacy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyClient) UpdatePrivacy(ctx context.Context, in *CreatePrivacyRequest, opts ...grpc.CallOption) (*EmptyResponsePrivacy, error) {
	out := new(EmptyResponsePrivacy)
	err := c.cc.Invoke(ctx, "/proto.Privacy/UpdatePrivacy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyClient) BlockUser(ctx context.Context, in *CreateBlockRequest, opts ...grpc.CallOption) (*EmptyResponsePrivacy, error) {
	out := new(EmptyResponsePrivacy)
	err := c.cc.Invoke(ctx, "/proto.Privacy/BlockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyClient) UnBlockUser(ctx context.Context, in *CreateBlockRequest, opts ...grpc.CallOption) (*EmptyResponsePrivacy, error) {
	out := new(EmptyResponsePrivacy)
	err := c.cc.Invoke(ctx, "/proto.Privacy/UnBlockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyClient) CheckIfBlocked(ctx context.Context, in *CreateBlockRequest, opts ...grpc.CallOption) (*BooleanResponse, error) {
	out := new(BooleanResponse)
	err := c.cc.Invoke(ctx, "/proto.Privacy/CheckIfBlocked", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyClient) CheckUserProfilePublic(ctx context.Context, in *PrivacyRequest, opts ...grpc.CallOption) (*BooleanResponse, error) {
	out := new(BooleanResponse)
	err := c.cc.Invoke(ctx, "/proto.Privacy/CheckUserProfilePublic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyClient) GetAllPublicUsers(ctx context.Context, in *RequestIdPrivacy, opts ...grpc.CallOption) (*StringArray, error) {
	out := new(StringArray)
	err := c.cc.Invoke(ctx, "/proto.Privacy/GetAllPublicUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privacyClient) GetUserPrivacy(ctx context.Context, in *RequestIdPrivacy, opts ...grpc.CallOption) (*PrivacyMessage, error) {
	out := new(PrivacyMessage)
	err := c.cc.Invoke(ctx, "/proto.Privacy/GetUserPrivacy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrivacyServer is the server API for Privacy service.
// All implementations must embed UnimplementedPrivacyServer
// for forward compatibility
type PrivacyServer interface {
	CreatePrivacy(context.Context, *CreatePrivacyRequest) (*EmptyResponsePrivacy, error)
	UpdatePrivacy(context.Context, *CreatePrivacyRequest) (*EmptyResponsePrivacy, error)
	BlockUser(context.Context, *CreateBlockRequest) (*EmptyResponsePrivacy, error)
	UnBlockUser(context.Context, *CreateBlockRequest) (*EmptyResponsePrivacy, error)
	CheckIfBlocked(context.Context, *CreateBlockRequest) (*BooleanResponse, error)
	CheckUserProfilePublic(context.Context, *PrivacyRequest) (*BooleanResponse, error)
	GetAllPublicUsers(context.Context, *RequestIdPrivacy) (*StringArray, error)
	GetUserPrivacy(context.Context, *RequestIdPrivacy) (*PrivacyMessage, error)
	mustEmbedUnimplementedPrivacyServer()
}

// UnimplementedPrivacyServer must be embedded to have forward compatible implementations.
type UnimplementedPrivacyServer struct {
}

func (UnimplementedPrivacyServer) CreatePrivacy(context.Context, *CreatePrivacyRequest) (*EmptyResponsePrivacy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePrivacy not implemented")
}
func (UnimplementedPrivacyServer) UpdatePrivacy(context.Context, *CreatePrivacyRequest) (*EmptyResponsePrivacy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePrivacy not implemented")
}
func (UnimplementedPrivacyServer) BlockUser(context.Context, *CreateBlockRequest) (*EmptyResponsePrivacy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockUser not implemented")
}
func (UnimplementedPrivacyServer) UnBlockUser(context.Context, *CreateBlockRequest) (*EmptyResponsePrivacy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnBlockUser not implemented")
}
func (UnimplementedPrivacyServer) CheckIfBlocked(context.Context, *CreateBlockRequest) (*BooleanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIfBlocked not implemented")
}
func (UnimplementedPrivacyServer) CheckUserProfilePublic(context.Context, *PrivacyRequest) (*BooleanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserProfilePublic not implemented")
}
func (UnimplementedPrivacyServer) GetAllPublicUsers(context.Context, *RequestIdPrivacy) (*StringArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPublicUsers not implemented")
}
func (UnimplementedPrivacyServer) GetUserPrivacy(context.Context, *RequestIdPrivacy) (*PrivacyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPrivacy not implemented")
}
func (UnimplementedPrivacyServer) mustEmbedUnimplementedPrivacyServer() {}

// UnsafePrivacyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrivacyServer will
// result in compilation errors.
type UnsafePrivacyServer interface {
	mustEmbedUnimplementedPrivacyServer()
}

func RegisterPrivacyServer(s *grpc.Server, srv PrivacyServer) {
	s.RegisterService(&_Privacy_serviceDesc, srv)
}

func _Privacy_CreatePrivacy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePrivacyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyServer).CreatePrivacy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Privacy/CreatePrivacy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyServer).CreatePrivacy(ctx, req.(*CreatePrivacyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Privacy_UpdatePrivacy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePrivacyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyServer).UpdatePrivacy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Privacy/UpdatePrivacy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyServer).UpdatePrivacy(ctx, req.(*CreatePrivacyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Privacy_BlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyServer).BlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Privacy/BlockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyServer).BlockUser(ctx, req.(*CreateBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Privacy_UnBlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyServer).UnBlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Privacy/UnBlockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyServer).UnBlockUser(ctx, req.(*CreateBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Privacy_CheckIfBlocked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyServer).CheckIfBlocked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Privacy/CheckIfBlocked",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyServer).CheckIfBlocked(ctx, req.(*CreateBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Privacy_CheckUserProfilePublic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivacyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyServer).CheckUserProfilePublic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Privacy/CheckUserProfilePublic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyServer).CheckUserProfilePublic(ctx, req.(*PrivacyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Privacy_GetAllPublicUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestIdPrivacy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyServer).GetAllPublicUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Privacy/GetAllPublicUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyServer).GetAllPublicUsers(ctx, req.(*RequestIdPrivacy))
	}
	return interceptor(ctx, in, info, handler)
}

func _Privacy_GetUserPrivacy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestIdPrivacy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivacyServer).GetUserPrivacy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Privacy/GetUserPrivacy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivacyServer).GetUserPrivacy(ctx, req.(*RequestIdPrivacy))
	}
	return interceptor(ctx, in, info, handler)
}

var _Privacy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Privacy",
	HandlerType: (*PrivacyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePrivacy",
			Handler:    _Privacy_CreatePrivacy_Handler,
		},
		{
			MethodName: "UpdatePrivacy",
			Handler:    _Privacy_UpdatePrivacy_Handler,
		},
		{
			MethodName: "BlockUser",
			Handler:    _Privacy_BlockUser_Handler,
		},
		{
			MethodName: "UnBlockUser",
			Handler:    _Privacy_UnBlockUser_Handler,
		},
		{
			MethodName: "CheckIfBlocked",
			Handler:    _Privacy_CheckIfBlocked_Handler,
		},
		{
			MethodName: "CheckUserProfilePublic",
			Handler:    _Privacy_CheckUserProfilePublic_Handler,
		},
		{
			MethodName: "GetAllPublicUsers",
			Handler:    _Privacy_GetAllPublicUsers_Handler,
		},
		{
			MethodName: "GetUserPrivacy",
			Handler:    _Privacy_GetUserPrivacy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "privacy.proto",
}
