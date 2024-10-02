// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: user/api/user/v1/user_admin.proto

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
	UserAdminService_ListUsersAdmin_FullMethodName  = "/user.api.user.v1.UserAdminService/ListUsersAdmin"
	UserAdminService_GetUserAdmin_FullMethodName    = "/user.api.user.v1.UserAdminService/GetUserAdmin"
	UserAdminService_CreateUserAdmin_FullMethodName = "/user.api.user.v1.UserAdminService/CreateUserAdmin"
	UserAdminService_UpdateUserAdmin_FullMethodName = "/user.api.user.v1.UserAdminService/UpdateUserAdmin"
	UserAdminService_DeleteUserAdmin_FullMethodName = "/user.api.user.v1.UserAdminService/DeleteUserAdmin"
)

// UserAdminServiceClient is the client API for UserAdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAdminServiceClient interface {
	// ListUsers
	// authz: user.admin.user,*,list
	ListUsersAdmin(ctx context.Context, in *AdminListUsersRequest, opts ...grpc.CallOption) (*AdminListUsersResponse, error)
	// GetUser
	// authz: user.admin.user,id,get
	GetUserAdmin(ctx context.Context, in *AdminGetUserRequest, opts ...grpc.CallOption) (*User, error)
	// CreateUser
	// authz: user.admin.user,*,create
	CreateUserAdmin(ctx context.Context, in *AdminCreateUserRequest, opts ...grpc.CallOption) (*User, error)
	// UpdateUser
	// authz: user.admin.user,id,update
	UpdateUserAdmin(ctx context.Context, in *AdminUpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	// DeleteUser
	// authz: user.admin.user,id,delete
	DeleteUserAdmin(ctx context.Context, in *AdminDeleteUserRequest, opts ...grpc.CallOption) (*AdminDeleteUserResponse, error)
}

type userAdminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAdminServiceClient(cc grpc.ClientConnInterface) UserAdminServiceClient {
	return &userAdminServiceClient{cc}
}

func (c *userAdminServiceClient) ListUsersAdmin(ctx context.Context, in *AdminListUsersRequest, opts ...grpc.CallOption) (*AdminListUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminListUsersResponse)
	err := c.cc.Invoke(ctx, UserAdminService_ListUsersAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAdminServiceClient) GetUserAdmin(ctx context.Context, in *AdminGetUserRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, UserAdminService_GetUserAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAdminServiceClient) CreateUserAdmin(ctx context.Context, in *AdminCreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, UserAdminService_CreateUserAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAdminServiceClient) UpdateUserAdmin(ctx context.Context, in *AdminUpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, UserAdminService_UpdateUserAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAdminServiceClient) DeleteUserAdmin(ctx context.Context, in *AdminDeleteUserRequest, opts ...grpc.CallOption) (*AdminDeleteUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminDeleteUserResponse)
	err := c.cc.Invoke(ctx, UserAdminService_DeleteUserAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAdminServiceServer is the server API for UserAdminService service.
// All implementations should embed UnimplementedUserAdminServiceServer
// for forward compatibility.
type UserAdminServiceServer interface {
	// ListUsers
	// authz: user.admin.user,*,list
	ListUsersAdmin(context.Context, *AdminListUsersRequest) (*AdminListUsersResponse, error)
	// GetUser
	// authz: user.admin.user,id,get
	GetUserAdmin(context.Context, *AdminGetUserRequest) (*User, error)
	// CreateUser
	// authz: user.admin.user,*,create
	CreateUserAdmin(context.Context, *AdminCreateUserRequest) (*User, error)
	// UpdateUser
	// authz: user.admin.user,id,update
	UpdateUserAdmin(context.Context, *AdminUpdateUserRequest) (*User, error)
	// DeleteUser
	// authz: user.admin.user,id,delete
	DeleteUserAdmin(context.Context, *AdminDeleteUserRequest) (*AdminDeleteUserResponse, error)
}

// UnimplementedUserAdminServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserAdminServiceServer struct{}

func (UnimplementedUserAdminServiceServer) ListUsersAdmin(context.Context, *AdminListUsersRequest) (*AdminListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsersAdmin not implemented")
}
func (UnimplementedUserAdminServiceServer) GetUserAdmin(context.Context, *AdminGetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAdmin not implemented")
}
func (UnimplementedUserAdminServiceServer) CreateUserAdmin(context.Context, *AdminCreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserAdmin not implemented")
}
func (UnimplementedUserAdminServiceServer) UpdateUserAdmin(context.Context, *AdminUpdateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserAdmin not implemented")
}
func (UnimplementedUserAdminServiceServer) DeleteUserAdmin(context.Context, *AdminDeleteUserRequest) (*AdminDeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserAdmin not implemented")
}
func (UnimplementedUserAdminServiceServer) testEmbeddedByValue() {}

// UnsafeUserAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAdminServiceServer will
// result in compilation errors.
type UnsafeUserAdminServiceServer interface {
	mustEmbedUnimplementedUserAdminServiceServer()
}

func RegisterUserAdminServiceServer(s grpc.ServiceRegistrar, srv UserAdminServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserAdminServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserAdminService_ServiceDesc, srv)
}

func _UserAdminService_ListUsersAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAdminServiceServer).ListUsersAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAdminService_ListUsersAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAdminServiceServer).ListUsersAdmin(ctx, req.(*AdminListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAdminService_GetUserAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAdminServiceServer).GetUserAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAdminService_GetUserAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAdminServiceServer).GetUserAdmin(ctx, req.(*AdminGetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAdminService_CreateUserAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAdminServiceServer).CreateUserAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAdminService_CreateUserAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAdminServiceServer).CreateUserAdmin(ctx, req.(*AdminCreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAdminService_UpdateUserAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAdminServiceServer).UpdateUserAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAdminService_UpdateUserAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAdminServiceServer).UpdateUserAdmin(ctx, req.(*AdminUpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAdminService_DeleteUserAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminDeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAdminServiceServer).DeleteUserAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAdminService_DeleteUserAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAdminServiceServer).DeleteUserAdmin(ctx, req.(*AdminDeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserAdminService_ServiceDesc is the grpc.ServiceDesc for UserAdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserAdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.api.user.v1.UserAdminService",
	HandlerType: (*UserAdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUsersAdmin",
			Handler:    _UserAdminService_ListUsersAdmin_Handler,
		},
		{
			MethodName: "GetUserAdmin",
			Handler:    _UserAdminService_GetUserAdmin_Handler,
		},
		{
			MethodName: "CreateUserAdmin",
			Handler:    _UserAdminService_CreateUserAdmin_Handler,
		},
		{
			MethodName: "UpdateUserAdmin",
			Handler:    _UserAdminService_UpdateUserAdmin_Handler,
		},
		{
			MethodName: "DeleteUserAdmin",
			Handler:    _UserAdminService_DeleteUserAdmin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/api/user/v1/user_admin.proto",
}
