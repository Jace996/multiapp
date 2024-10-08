// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: user/api/user/v1/user.proto

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
	UserService_ListUsers_FullMethodName            = "/user.api.user.v1.UserService/ListUsers"
	UserService_GetUser_FullMethodName              = "/user.api.user.v1.UserService/GetUser"
	UserService_CreateUser_FullMethodName           = "/user.api.user.v1.UserService/CreateUser"
	UserService_UpdateUser_FullMethodName           = "/user.api.user.v1.UserService/UpdateUser"
	UserService_DeleteUser_FullMethodName           = "/user.api.user.v1.UserService/DeleteUser"
	UserService_GetUserRoles_FullMethodName         = "/user.api.user.v1.UserService/GetUserRoles"
	UserService_GetUserPermission_FullMethodName    = "/user.api.user.v1.UserService/GetUserPermission"
	UserService_UpdateUserPermission_FullMethodName = "/user.api.user.v1.UserService/UpdateUserPermission"
	UserService_InviteUser_FullMethodName           = "/user.api.user.v1.UserService/InviteUser"
	UserService_PublicSearchUser_FullMethodName     = "/user.api.user.v1.UserService/PublicSearchUser"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// ListUsers
	// authz: user.user,*,list
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	// GetUser
	// authz: user.user,id,get
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	// CreateUser
	// authz: user.user,*,create
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	// UpdateUser
	// authz: user.user,id,update
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	// DeleteUser
	// authz: user.user,id,delete
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	// GetUserRoles
	// authz: user.user,id,get
	GetUserRoles(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*GetUserRoleReply, error)
	GetUserPermission(ctx context.Context, in *GetUserPermissionRequest, opts ...grpc.CallOption) (*GetUserPermissionReply, error)
	UpdateUserPermission(ctx context.Context, in *UpdateUserPermissionRequest, opts ...grpc.CallOption) (*UpdateUserPermissionReply, error)
	// InviteUser
	// authz: user.user,*,create
	InviteUser(ctx context.Context, in *InviteUserRequest, opts ...grpc.CallOption) (*InviteUserReply, error)
	PublicSearchUser(ctx context.Context, in *SearchUserRequest, opts ...grpc.CallOption) (*SearchUserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, UserService_ListUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, UserService_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, UserService_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, UserService_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, UserService_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserRoles(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*GetUserRoleReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserRoleReply)
	err := c.cc.Invoke(ctx, UserService_GetUserRoles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserPermission(ctx context.Context, in *GetUserPermissionRequest, opts ...grpc.CallOption) (*GetUserPermissionReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserPermissionReply)
	err := c.cc.Invoke(ctx, UserService_GetUserPermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserPermission(ctx context.Context, in *UpdateUserPermissionRequest, opts ...grpc.CallOption) (*UpdateUserPermissionReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserPermissionReply)
	err := c.cc.Invoke(ctx, UserService_UpdateUserPermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) InviteUser(ctx context.Context, in *InviteUserRequest, opts ...grpc.CallOption) (*InviteUserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InviteUserReply)
	err := c.cc.Invoke(ctx, UserService_InviteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) PublicSearchUser(ctx context.Context, in *SearchUserRequest, opts ...grpc.CallOption) (*SearchUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchUserResponse)
	err := c.cc.Invoke(ctx, UserService_PublicSearchUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	// ListUsers
	// authz: user.user,*,list
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	// GetUser
	// authz: user.user,id,get
	GetUser(context.Context, *GetUserRequest) (*User, error)
	// CreateUser
	// authz: user.user,*,create
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// UpdateUser
	// authz: user.user,id,update
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	// DeleteUser
	// authz: user.user,id,delete
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	// GetUserRoles
	// authz: user.user,id,get
	GetUserRoles(context.Context, *GetUserRoleRequest) (*GetUserRoleReply, error)
	GetUserPermission(context.Context, *GetUserPermissionRequest) (*GetUserPermissionReply, error)
	UpdateUserPermission(context.Context, *UpdateUserPermissionRequest) (*UpdateUserPermissionReply, error)
	// InviteUser
	// authz: user.user,*,create
	InviteUser(context.Context, *InviteUserRequest) (*InviteUserReply, error)
	PublicSearchUser(context.Context, *SearchUserRequest) (*SearchUserResponse, error)
}

// UnimplementedUserServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserRoles(context.Context, *GetUserRoleRequest) (*GetUserRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRoles not implemented")
}
func (UnimplementedUserServiceServer) GetUserPermission(context.Context, *GetUserPermissionRequest) (*GetUserPermissionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPermission not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserPermission(context.Context, *UpdateUserPermissionRequest) (*UpdateUserPermissionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserPermission not implemented")
}
func (UnimplementedUserServiceServer) InviteUser(context.Context, *InviteUserRequest) (*InviteUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteUser not implemented")
}
func (UnimplementedUserServiceServer) PublicSearchUser(context.Context, *SearchUserRequest) (*SearchUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublicSearchUser not implemented")
}
func (UnimplementedUserServiceServer) testEmbeddedByValue() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ListUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserRoles(ctx, req.(*GetUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserPermission(ctx, req.(*GetUserPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUserPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserPermission(ctx, req.(*UpdateUserPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_InviteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).InviteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_InviteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).InviteUser(ctx, req.(*InviteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_PublicSearchUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).PublicSearchUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_PublicSearchUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).PublicSearchUser(ctx, req.(*SearchUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.api.user.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUsers",
			Handler:    _UserService_ListUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "GetUserRoles",
			Handler:    _UserService_GetUserRoles_Handler,
		},
		{
			MethodName: "GetUserPermission",
			Handler:    _UserService_GetUserPermission_Handler,
		},
		{
			MethodName: "UpdateUserPermission",
			Handler:    _UserService_UpdateUserPermission_Handler,
		},
		{
			MethodName: "InviteUser",
			Handler:    _UserService_InviteUser_Handler,
		},
		{
			MethodName: "PublicSearchUser",
			Handler:    _UserService_PublicSearchUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/api/user/v1/user.proto",
}

const (
	UserInternalService_CreateTenant_FullMethodName               = "/user.api.user.v1.UserInternalService/CreateTenant"
	UserInternalService_CheckUserTenant_FullMethodName            = "/user.api.user.v1.UserInternalService/CheckUserTenant"
	UserInternalService_FindOrCreateStripeCustomer_FullMethodName = "/user.api.user.v1.UserInternalService/FindOrCreateStripeCustomer"
)

// UserInternalServiceClient is the client API for UserInternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserInternalServiceClient interface {
	CreateTenant(ctx context.Context, in *UserInternalCreateTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// CheckUserTenant internal api for checking whether user is allowed in this tenant
	CheckUserTenant(ctx context.Context, in *CheckUserTenantRequest, opts ...grpc.CallOption) (*CheckUserTenantReply, error)
	FindOrCreateStripeCustomer(ctx context.Context, in *FindOrCreateStripeCustomerRequest, opts ...grpc.CallOption) (*FindOrCreateStripeCustomerReply, error)
}

type userInternalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserInternalServiceClient(cc grpc.ClientConnInterface) UserInternalServiceClient {
	return &userInternalServiceClient{cc}
}

func (c *userInternalServiceClient) CreateTenant(ctx context.Context, in *UserInternalCreateTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserInternalService_CreateTenant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInternalServiceClient) CheckUserTenant(ctx context.Context, in *CheckUserTenantRequest, opts ...grpc.CallOption) (*CheckUserTenantReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckUserTenantReply)
	err := c.cc.Invoke(ctx, UserInternalService_CheckUserTenant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInternalServiceClient) FindOrCreateStripeCustomer(ctx context.Context, in *FindOrCreateStripeCustomerRequest, opts ...grpc.CallOption) (*FindOrCreateStripeCustomerReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindOrCreateStripeCustomerReply)
	err := c.cc.Invoke(ctx, UserInternalService_FindOrCreateStripeCustomer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInternalServiceServer is the server API for UserInternalService service.
// All implementations should embed UnimplementedUserInternalServiceServer
// for forward compatibility.
type UserInternalServiceServer interface {
	CreateTenant(context.Context, *UserInternalCreateTenantRequest) (*emptypb.Empty, error)
	// CheckUserTenant internal api for checking whether user is allowed in this tenant
	CheckUserTenant(context.Context, *CheckUserTenantRequest) (*CheckUserTenantReply, error)
	FindOrCreateStripeCustomer(context.Context, *FindOrCreateStripeCustomerRequest) (*FindOrCreateStripeCustomerReply, error)
}

// UnimplementedUserInternalServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserInternalServiceServer struct{}

func (UnimplementedUserInternalServiceServer) CreateTenant(context.Context, *UserInternalCreateTenantRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTenant not implemented")
}
func (UnimplementedUserInternalServiceServer) CheckUserTenant(context.Context, *CheckUserTenantRequest) (*CheckUserTenantReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserTenant not implemented")
}
func (UnimplementedUserInternalServiceServer) FindOrCreateStripeCustomer(context.Context, *FindOrCreateStripeCustomerRequest) (*FindOrCreateStripeCustomerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOrCreateStripeCustomer not implemented")
}
func (UnimplementedUserInternalServiceServer) testEmbeddedByValue() {}

// UnsafeUserInternalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserInternalServiceServer will
// result in compilation errors.
type UnsafeUserInternalServiceServer interface {
	mustEmbedUnimplementedUserInternalServiceServer()
}

func RegisterUserInternalServiceServer(s grpc.ServiceRegistrar, srv UserInternalServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserInternalServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserInternalService_ServiceDesc, srv)
}

func _UserInternalService_CreateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInternalCreateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInternalServiceServer).CreateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInternalService_CreateTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInternalServiceServer).CreateTenant(ctx, req.(*UserInternalCreateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInternalService_CheckUserTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInternalServiceServer).CheckUserTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInternalService_CheckUserTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInternalServiceServer).CheckUserTenant(ctx, req.(*CheckUserTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInternalService_FindOrCreateStripeCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOrCreateStripeCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInternalServiceServer).FindOrCreateStripeCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInternalService_FindOrCreateStripeCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInternalServiceServer).FindOrCreateStripeCustomer(ctx, req.(*FindOrCreateStripeCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserInternalService_ServiceDesc is the grpc.ServiceDesc for UserInternalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserInternalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.api.user.v1.UserInternalService",
	HandlerType: (*UserInternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTenant",
			Handler:    _UserInternalService_CreateTenant_Handler,
		},
		{
			MethodName: "CheckUserTenant",
			Handler:    _UserInternalService_CheckUserTenant_Handler,
		},
		{
			MethodName: "FindOrCreateStripeCustomer",
			Handler:    _UserInternalService_FindOrCreateStripeCustomer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/api/user/v1/user.proto",
}
