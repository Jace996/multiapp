// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: payment/api/checkout/v1/checkout.proto

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
	CheckoutService_CheckoutNow_FullMethodName   = "/payment.api.checkout.v1.CheckoutService/CheckoutNow"
	CheckoutService_CheckoutOrder_FullMethodName = "/payment.api.checkout.v1.CheckoutService/CheckoutOrder"
)

// CheckoutServiceClient is the client API for CheckoutService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckoutServiceClient interface {
	CheckoutNow(ctx context.Context, in *CheckoutNowRequest, opts ...grpc.CallOption) (*CheckoutNowReply, error)
	CheckoutOrder(ctx context.Context, in *CheckOutOrderRequest, opts ...grpc.CallOption) (*CheckoutOrderReply, error)
}

type checkoutServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckoutServiceClient(cc grpc.ClientConnInterface) CheckoutServiceClient {
	return &checkoutServiceClient{cc}
}

func (c *checkoutServiceClient) CheckoutNow(ctx context.Context, in *CheckoutNowRequest, opts ...grpc.CallOption) (*CheckoutNowReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckoutNowReply)
	err := c.cc.Invoke(ctx, CheckoutService_CheckoutNow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkoutServiceClient) CheckoutOrder(ctx context.Context, in *CheckOutOrderRequest, opts ...grpc.CallOption) (*CheckoutOrderReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckoutOrderReply)
	err := c.cc.Invoke(ctx, CheckoutService_CheckoutOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckoutServiceServer is the server API for CheckoutService service.
// All implementations should embed UnimplementedCheckoutServiceServer
// for forward compatibility.
type CheckoutServiceServer interface {
	CheckoutNow(context.Context, *CheckoutNowRequest) (*CheckoutNowReply, error)
	CheckoutOrder(context.Context, *CheckOutOrderRequest) (*CheckoutOrderReply, error)
}

// UnimplementedCheckoutServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCheckoutServiceServer struct{}

func (UnimplementedCheckoutServiceServer) CheckoutNow(context.Context, *CheckoutNowRequest) (*CheckoutNowReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckoutNow not implemented")
}
func (UnimplementedCheckoutServiceServer) CheckoutOrder(context.Context, *CheckOutOrderRequest) (*CheckoutOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckoutOrder not implemented")
}
func (UnimplementedCheckoutServiceServer) testEmbeddedByValue() {}

// UnsafeCheckoutServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckoutServiceServer will
// result in compilation errors.
type UnsafeCheckoutServiceServer interface {
	mustEmbedUnimplementedCheckoutServiceServer()
}

func RegisterCheckoutServiceServer(s grpc.ServiceRegistrar, srv CheckoutServiceServer) {
	// If the following call pancis, it indicates UnimplementedCheckoutServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CheckoutService_ServiceDesc, srv)
}

func _CheckoutService_CheckoutNow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckoutNowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckoutServiceServer).CheckoutNow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CheckoutService_CheckoutNow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckoutServiceServer).CheckoutNow(ctx, req.(*CheckoutNowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckoutService_CheckoutOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckOutOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckoutServiceServer).CheckoutOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CheckoutService_CheckoutOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckoutServiceServer).CheckoutOrder(ctx, req.(*CheckOutOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CheckoutService_ServiceDesc is the grpc.ServiceDesc for CheckoutService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CheckoutService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.api.checkout.v1.CheckoutService",
	HandlerType: (*CheckoutServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckoutNow",
			Handler:    _CheckoutService_CheckoutNow_Handler,
		},
		{
			MethodName: "CheckoutOrder",
			Handler:    _CheckoutService_CheckoutOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment/api/checkout/v1/checkout.proto",
}
