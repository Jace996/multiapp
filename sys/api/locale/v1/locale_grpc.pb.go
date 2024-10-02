// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: sys/api/locale/v1/locale.proto

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
	LocaleService_ListMessages_FullMethodName = "/sys.api.locale.v1.LocaleService/ListMessages"
)

// LocaleServiceClient is the client API for LocaleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocaleServiceClient interface {
	ListMessages(ctx context.Context, in *ListMessageRequest, opts ...grpc.CallOption) (*ListMessageReply, error)
}

type localeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocaleServiceClient(cc grpc.ClientConnInterface) LocaleServiceClient {
	return &localeServiceClient{cc}
}

func (c *localeServiceClient) ListMessages(ctx context.Context, in *ListMessageRequest, opts ...grpc.CallOption) (*ListMessageReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListMessageReply)
	err := c.cc.Invoke(ctx, LocaleService_ListMessages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocaleServiceServer is the server API for LocaleService service.
// All implementations should embed UnimplementedLocaleServiceServer
// for forward compatibility.
type LocaleServiceServer interface {
	ListMessages(context.Context, *ListMessageRequest) (*ListMessageReply, error)
}

// UnimplementedLocaleServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLocaleServiceServer struct{}

func (UnimplementedLocaleServiceServer) ListMessages(context.Context, *ListMessageRequest) (*ListMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMessages not implemented")
}
func (UnimplementedLocaleServiceServer) testEmbeddedByValue() {}

// UnsafeLocaleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocaleServiceServer will
// result in compilation errors.
type UnsafeLocaleServiceServer interface {
	mustEmbedUnimplementedLocaleServiceServer()
}

func RegisterLocaleServiceServer(s grpc.ServiceRegistrar, srv LocaleServiceServer) {
	// If the following call pancis, it indicates UnimplementedLocaleServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LocaleService_ServiceDesc, srv)
}

func _LocaleService_ListMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocaleServiceServer).ListMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocaleService_ListMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocaleServiceServer).ListMessages(ctx, req.(*ListMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LocaleService_ServiceDesc is the grpc.ServiceDesc for LocaleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocaleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sys.api.locale.v1.LocaleService",
	HandlerType: (*LocaleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMessages",
			Handler:    _LocaleService_ListMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sys/api/locale/v1/locale.proto",
}
