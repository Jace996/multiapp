// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: user/api/permission/v1/permission_internal.proto

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
	PermissionInternalService_CheckForSubjects_FullMethodName        = "/user.api.permission.v1.PermissionInternalService/CheckForSubjects"
	PermissionInternalService_AddSubjectPermission_FullMethodName    = "/user.api.permission.v1.PermissionInternalService/AddSubjectPermission"
	PermissionInternalService_ListSubjectPermission_FullMethodName   = "/user.api.permission.v1.PermissionInternalService/ListSubjectPermission"
	PermissionInternalService_UpdateSubjectPermission_FullMethodName = "/user.api.permission.v1.PermissionInternalService/UpdateSubjectPermission"
	PermissionInternalService_RemoveSubjectPermission_FullMethodName = "/user.api.permission.v1.PermissionInternalService/RemoveSubjectPermission"
)

// PermissionInternalServiceClient is the client API for PermissionInternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionInternalServiceClient interface {
	CheckForSubjects(ctx context.Context, in *CheckSubjectsPermissionRequest, opts ...grpc.CallOption) (*CheckSubjectsPermissionReply, error)
	// management add
	AddSubjectPermission(ctx context.Context, in *AddSubjectPermissionRequest, opts ...grpc.CallOption) (*AddSubjectPermissionResponse, error)
	// management list
	ListSubjectPermission(ctx context.Context, in *ListSubjectPermissionRequest, opts ...grpc.CallOption) (*ListSubjectPermissionResponse, error)
	// management update
	UpdateSubjectPermission(ctx context.Context, in *UpdateSubjectPermissionRequest, opts ...grpc.CallOption) (*UpdateSubjectPermissionResponse, error)
	// management remove
	RemoveSubjectPermission(ctx context.Context, in *RemoveSubjectPermissionRequest, opts ...grpc.CallOption) (*RemoveSubjectPermissionReply, error)
}

type permissionInternalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionInternalServiceClient(cc grpc.ClientConnInterface) PermissionInternalServiceClient {
	return &permissionInternalServiceClient{cc}
}

func (c *permissionInternalServiceClient) CheckForSubjects(ctx context.Context, in *CheckSubjectsPermissionRequest, opts ...grpc.CallOption) (*CheckSubjectsPermissionReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckSubjectsPermissionReply)
	err := c.cc.Invoke(ctx, PermissionInternalService_CheckForSubjects_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionInternalServiceClient) AddSubjectPermission(ctx context.Context, in *AddSubjectPermissionRequest, opts ...grpc.CallOption) (*AddSubjectPermissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddSubjectPermissionResponse)
	err := c.cc.Invoke(ctx, PermissionInternalService_AddSubjectPermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionInternalServiceClient) ListSubjectPermission(ctx context.Context, in *ListSubjectPermissionRequest, opts ...grpc.CallOption) (*ListSubjectPermissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSubjectPermissionResponse)
	err := c.cc.Invoke(ctx, PermissionInternalService_ListSubjectPermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionInternalServiceClient) UpdateSubjectPermission(ctx context.Context, in *UpdateSubjectPermissionRequest, opts ...grpc.CallOption) (*UpdateSubjectPermissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateSubjectPermissionResponse)
	err := c.cc.Invoke(ctx, PermissionInternalService_UpdateSubjectPermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionInternalServiceClient) RemoveSubjectPermission(ctx context.Context, in *RemoveSubjectPermissionRequest, opts ...grpc.CallOption) (*RemoveSubjectPermissionReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveSubjectPermissionReply)
	err := c.cc.Invoke(ctx, PermissionInternalService_RemoveSubjectPermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionInternalServiceServer is the server API for PermissionInternalService service.
// All implementations should embed UnimplementedPermissionInternalServiceServer
// for forward compatibility.
type PermissionInternalServiceServer interface {
	CheckForSubjects(context.Context, *CheckSubjectsPermissionRequest) (*CheckSubjectsPermissionReply, error)
	// management add
	AddSubjectPermission(context.Context, *AddSubjectPermissionRequest) (*AddSubjectPermissionResponse, error)
	// management list
	ListSubjectPermission(context.Context, *ListSubjectPermissionRequest) (*ListSubjectPermissionResponse, error)
	// management update
	UpdateSubjectPermission(context.Context, *UpdateSubjectPermissionRequest) (*UpdateSubjectPermissionResponse, error)
	// management remove
	RemoveSubjectPermission(context.Context, *RemoveSubjectPermissionRequest) (*RemoveSubjectPermissionReply, error)
}

// UnimplementedPermissionInternalServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPermissionInternalServiceServer struct{}

func (UnimplementedPermissionInternalServiceServer) CheckForSubjects(context.Context, *CheckSubjectsPermissionRequest) (*CheckSubjectsPermissionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckForSubjects not implemented")
}
func (UnimplementedPermissionInternalServiceServer) AddSubjectPermission(context.Context, *AddSubjectPermissionRequest) (*AddSubjectPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSubjectPermission not implemented")
}
func (UnimplementedPermissionInternalServiceServer) ListSubjectPermission(context.Context, *ListSubjectPermissionRequest) (*ListSubjectPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubjectPermission not implemented")
}
func (UnimplementedPermissionInternalServiceServer) UpdateSubjectPermission(context.Context, *UpdateSubjectPermissionRequest) (*UpdateSubjectPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubjectPermission not implemented")
}
func (UnimplementedPermissionInternalServiceServer) RemoveSubjectPermission(context.Context, *RemoveSubjectPermissionRequest) (*RemoveSubjectPermissionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSubjectPermission not implemented")
}
func (UnimplementedPermissionInternalServiceServer) testEmbeddedByValue() {}

// UnsafePermissionInternalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionInternalServiceServer will
// result in compilation errors.
type UnsafePermissionInternalServiceServer interface {
	mustEmbedUnimplementedPermissionInternalServiceServer()
}

func RegisterPermissionInternalServiceServer(s grpc.ServiceRegistrar, srv PermissionInternalServiceServer) {
	// If the following call pancis, it indicates UnimplementedPermissionInternalServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PermissionInternalService_ServiceDesc, srv)
}

func _PermissionInternalService_CheckForSubjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSubjectsPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionInternalServiceServer).CheckForSubjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionInternalService_CheckForSubjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionInternalServiceServer).CheckForSubjects(ctx, req.(*CheckSubjectsPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionInternalService_AddSubjectPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSubjectPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionInternalServiceServer).AddSubjectPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionInternalService_AddSubjectPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionInternalServiceServer).AddSubjectPermission(ctx, req.(*AddSubjectPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionInternalService_ListSubjectPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubjectPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionInternalServiceServer).ListSubjectPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionInternalService_ListSubjectPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionInternalServiceServer).ListSubjectPermission(ctx, req.(*ListSubjectPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionInternalService_UpdateSubjectPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSubjectPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionInternalServiceServer).UpdateSubjectPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionInternalService_UpdateSubjectPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionInternalServiceServer).UpdateSubjectPermission(ctx, req.(*UpdateSubjectPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionInternalService_RemoveSubjectPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSubjectPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionInternalServiceServer).RemoveSubjectPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionInternalService_RemoveSubjectPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionInternalServiceServer).RemoveSubjectPermission(ctx, req.(*RemoveSubjectPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PermissionInternalService_ServiceDesc is the grpc.ServiceDesc for PermissionInternalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PermissionInternalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.api.permission.v1.PermissionInternalService",
	HandlerType: (*PermissionInternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckForSubjects",
			Handler:    _PermissionInternalService_CheckForSubjects_Handler,
		},
		{
			MethodName: "AddSubjectPermission",
			Handler:    _PermissionInternalService_AddSubjectPermission_Handler,
		},
		{
			MethodName: "ListSubjectPermission",
			Handler:    _PermissionInternalService_ListSubjectPermission_Handler,
		},
		{
			MethodName: "UpdateSubjectPermission",
			Handler:    _PermissionInternalService_UpdateSubjectPermission_Handler,
		},
		{
			MethodName: "RemoveSubjectPermission",
			Handler:    _PermissionInternalService_RemoveSubjectPermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/api/permission/v1/permission_internal.proto",
}
