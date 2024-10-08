// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: oidc/api/client/v1/client.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ClientService_ListOAuth2Clients_FullMethodName  = "/oidc.api.client.ClientService/ListOAuth2Clients"
	ClientService_GetOAuth2Client_FullMethodName    = "/oidc.api.client.ClientService/GetOAuth2Client"
	ClientService_CreateOAuth2Client_FullMethodName = "/oidc.api.client.ClientService/CreateOAuth2Client"
	ClientService_DeleteOAuth2Client_FullMethodName = "/oidc.api.client.ClientService/DeleteOAuth2Client"
	ClientService_PatchOAuth2Client_FullMethodName  = "/oidc.api.client.ClientService/PatchOAuth2Client"
	ClientService_UpdateOAuth2Client_FullMethodName = "/oidc.api.client.ClientService/UpdateOAuth2Client"
)

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServiceClient interface {
	ListOAuth2Clients(ctx context.Context, in *ListClientRequest, opts ...grpc.CallOption) (*OAuth2ClientList, error)
	GetOAuth2Client(ctx context.Context, in *GetOAuth2ClientRequest, opts ...grpc.CallOption) (*OAuth2Client, error)
	CreateOAuth2Client(ctx context.Context, in *OAuth2Client, opts ...grpc.CallOption) (*OAuth2Client, error)
	DeleteOAuth2Client(ctx context.Context, in *DeleteOAuth2ClientRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PatchOAuth2Client(ctx context.Context, in *PatchOAuth2ClientRequest, opts ...grpc.CallOption) (*OAuth2Client, error)
	UpdateOAuth2Client(ctx context.Context, in *UpdateOAuth2ClientRequest, opts ...grpc.CallOption) (*OAuth2Client, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) ListOAuth2Clients(ctx context.Context, in *ListClientRequest, opts ...grpc.CallOption) (*OAuth2ClientList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OAuth2ClientList)
	err := c.cc.Invoke(ctx, ClientService_ListOAuth2Clients_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetOAuth2Client(ctx context.Context, in *GetOAuth2ClientRequest, opts ...grpc.CallOption) (*OAuth2Client, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OAuth2Client)
	err := c.cc.Invoke(ctx, ClientService_GetOAuth2Client_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) CreateOAuth2Client(ctx context.Context, in *OAuth2Client, opts ...grpc.CallOption) (*OAuth2Client, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OAuth2Client)
	err := c.cc.Invoke(ctx, ClientService_CreateOAuth2Client_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) DeleteOAuth2Client(ctx context.Context, in *DeleteOAuth2ClientRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ClientService_DeleteOAuth2Client_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) PatchOAuth2Client(ctx context.Context, in *PatchOAuth2ClientRequest, opts ...grpc.CallOption) (*OAuth2Client, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OAuth2Client)
	err := c.cc.Invoke(ctx, ClientService_PatchOAuth2Client_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) UpdateOAuth2Client(ctx context.Context, in *UpdateOAuth2ClientRequest, opts ...grpc.CallOption) (*OAuth2Client, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OAuth2Client)
	err := c.cc.Invoke(ctx, ClientService_UpdateOAuth2Client_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
// All implementations should embed UnimplementedClientServiceServer
// for forward compatibility.
type ClientServiceServer interface {
	ListOAuth2Clients(context.Context, *ListClientRequest) (*OAuth2ClientList, error)
	GetOAuth2Client(context.Context, *GetOAuth2ClientRequest) (*OAuth2Client, error)
	CreateOAuth2Client(context.Context, *OAuth2Client) (*OAuth2Client, error)
	DeleteOAuth2Client(context.Context, *DeleteOAuth2ClientRequest) (*emptypb.Empty, error)
	PatchOAuth2Client(context.Context, *PatchOAuth2ClientRequest) (*OAuth2Client, error)
	UpdateOAuth2Client(context.Context, *UpdateOAuth2ClientRequest) (*OAuth2Client, error)
}

// UnimplementedClientServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedClientServiceServer struct{}

func (UnimplementedClientServiceServer) ListOAuth2Clients(context.Context, *ListClientRequest) (*OAuth2ClientList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOAuth2Clients not implemented")
}
func (UnimplementedClientServiceServer) GetOAuth2Client(context.Context, *GetOAuth2ClientRequest) (*OAuth2Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOAuth2Client not implemented")
}
func (UnimplementedClientServiceServer) CreateOAuth2Client(context.Context, *OAuth2Client) (*OAuth2Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOAuth2Client not implemented")
}
func (UnimplementedClientServiceServer) DeleteOAuth2Client(context.Context, *DeleteOAuth2ClientRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOAuth2Client not implemented")
}
func (UnimplementedClientServiceServer) PatchOAuth2Client(context.Context, *PatchOAuth2ClientRequest) (*OAuth2Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchOAuth2Client not implemented")
}
func (UnimplementedClientServiceServer) UpdateOAuth2Client(context.Context, *UpdateOAuth2ClientRequest) (*OAuth2Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOAuth2Client not implemented")
}
func (UnimplementedClientServiceServer) testEmbeddedByValue() {}

// UnsafeClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServiceServer will
// result in compilation errors.
type UnsafeClientServiceServer interface {
	mustEmbedUnimplementedClientServiceServer()
}

func RegisterClientServiceServer(s grpc.ServiceRegistrar, srv ClientServiceServer) {
	// If the following call pancis, it indicates UnimplementedClientServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ClientService_ServiceDesc, srv)
}

func _ClientService_ListOAuth2Clients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).ListOAuth2Clients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientService_ListOAuth2Clients_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).ListOAuth2Clients(ctx, req.(*ListClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetOAuth2Client_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOAuth2ClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetOAuth2Client(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientService_GetOAuth2Client_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetOAuth2Client(ctx, req.(*GetOAuth2ClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_CreateOAuth2Client_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OAuth2Client)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).CreateOAuth2Client(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientService_CreateOAuth2Client_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).CreateOAuth2Client(ctx, req.(*OAuth2Client))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_DeleteOAuth2Client_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOAuth2ClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).DeleteOAuth2Client(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientService_DeleteOAuth2Client_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).DeleteOAuth2Client(ctx, req.(*DeleteOAuth2ClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_PatchOAuth2Client_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchOAuth2ClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).PatchOAuth2Client(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientService_PatchOAuth2Client_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).PatchOAuth2Client(ctx, req.(*PatchOAuth2ClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_UpdateOAuth2Client_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOAuth2ClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).UpdateOAuth2Client(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientService_UpdateOAuth2Client_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).UpdateOAuth2Client(ctx, req.(*UpdateOAuth2ClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientService_ServiceDesc is the grpc.ServiceDesc for ClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "oidc.api.client.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListOAuth2Clients",
			Handler:    _ClientService_ListOAuth2Clients_Handler,
		},
		{
			MethodName: "GetOAuth2Client",
			Handler:    _ClientService_GetOAuth2Client_Handler,
		},
		{
			MethodName: "CreateOAuth2Client",
			Handler:    _ClientService_CreateOAuth2Client_Handler,
		},
		{
			MethodName: "DeleteOAuth2Client",
			Handler:    _ClientService_DeleteOAuth2Client_Handler,
		},
		{
			MethodName: "PatchOAuth2Client",
			Handler:    _ClientService_PatchOAuth2Client_Handler,
		},
		{
			MethodName: "UpdateOAuth2Client",
			Handler:    _ClientService_UpdateOAuth2Client_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oidc/api/client/v1/client.proto",
}
