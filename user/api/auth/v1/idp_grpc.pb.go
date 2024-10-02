// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: user/api/auth/v1/idp.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	WeChatAuthService_MiniProgramCode_FullMethodName      = "/user.api.auth.v1.WeChatAuthService/MiniProgramCode"
	WeChatAuthService_MiniProgramPhoneCode_FullMethodName = "/user.api.auth.v1.WeChatAuthService/MiniProgramPhoneCode"
)

// WeChatAuthServiceClient is the client API for WeChatAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeChatAuthServiceClient interface {
	MiniProgramCode(ctx context.Context, in *WechatMiniProgramCodeReq, opts ...grpc.CallOption) (*WeChatLoginReply, error)
	MiniProgramPhoneCode(ctx context.Context, in *WechatMiniProgramPhoneCodeReq, opts ...grpc.CallOption) (*WeChatLoginReply, error)
}

type weChatAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWeChatAuthServiceClient(cc grpc.ClientConnInterface) WeChatAuthServiceClient {
	return &weChatAuthServiceClient{cc}
}

func (c *weChatAuthServiceClient) MiniProgramCode(ctx context.Context, in *WechatMiniProgramCodeReq, opts ...grpc.CallOption) (*WeChatLoginReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WeChatLoginReply)
	err := c.cc.Invoke(ctx, WeChatAuthService_MiniProgramCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weChatAuthServiceClient) MiniProgramPhoneCode(ctx context.Context, in *WechatMiniProgramPhoneCodeReq, opts ...grpc.CallOption) (*WeChatLoginReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WeChatLoginReply)
	err := c.cc.Invoke(ctx, WeChatAuthService_MiniProgramPhoneCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeChatAuthServiceServer is the server API for WeChatAuthService service.
// All implementations should embed UnimplementedWeChatAuthServiceServer
// for forward compatibility.
type WeChatAuthServiceServer interface {
	MiniProgramCode(context.Context, *WechatMiniProgramCodeReq) (*WeChatLoginReply, error)
	MiniProgramPhoneCode(context.Context, *WechatMiniProgramPhoneCodeReq) (*WeChatLoginReply, error)
}

// UnimplementedWeChatAuthServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWeChatAuthServiceServer struct{}

func (UnimplementedWeChatAuthServiceServer) MiniProgramCode(context.Context, *WechatMiniProgramCodeReq) (*WeChatLoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MiniProgramCode not implemented")
}
func (UnimplementedWeChatAuthServiceServer) MiniProgramPhoneCode(context.Context, *WechatMiniProgramPhoneCodeReq) (*WeChatLoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MiniProgramPhoneCode not implemented")
}
func (UnimplementedWeChatAuthServiceServer) testEmbeddedByValue() {}

// UnsafeWeChatAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeChatAuthServiceServer will
// result in compilation errors.
type UnsafeWeChatAuthServiceServer interface {
	mustEmbedUnimplementedWeChatAuthServiceServer()
}

func RegisterWeChatAuthServiceServer(s grpc.ServiceRegistrar, srv WeChatAuthServiceServer) {
	// If the following call pancis, it indicates UnimplementedWeChatAuthServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WeChatAuthService_ServiceDesc, srv)
}

func _WeChatAuthService_MiniProgramCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WechatMiniProgramCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeChatAuthServiceServer).MiniProgramCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WeChatAuthService_MiniProgramCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeChatAuthServiceServer).MiniProgramCode(ctx, req.(*WechatMiniProgramCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeChatAuthService_MiniProgramPhoneCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WechatMiniProgramPhoneCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeChatAuthServiceServer).MiniProgramPhoneCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WeChatAuthService_MiniProgramPhoneCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeChatAuthServiceServer).MiniProgramPhoneCode(ctx, req.(*WechatMiniProgramPhoneCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// WeChatAuthService_ServiceDesc is the grpc.ServiceDesc for WeChatAuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WeChatAuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.api.auth.v1.WeChatAuthService",
	HandlerType: (*WeChatAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MiniProgramCode",
			Handler:    _WeChatAuthService_MiniProgramCode_Handler,
		},
		{
			MethodName: "MiniProgramPhoneCode",
			Handler:    _WeChatAuthService_MiniProgramPhoneCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/api/auth/v1/idp.proto",
}
