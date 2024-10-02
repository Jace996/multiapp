// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: user/api/auth/v1/auth.proto

package v1

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

const (
	Auth_Register_FullMethodName                = "/user.api.auth.v1.Auth/Register"
	Auth_Login_FullMethodName                   = "/user.api.auth.v1.Auth/Login"
	Auth_GetLogin_FullMethodName                = "/user.api.auth.v1.Auth/GetLogin"
	Auth_Token_FullMethodName                   = "/user.api.auth.v1.Auth/Token"
	Auth_Refresh_FullMethodName                 = "/user.api.auth.v1.Auth/Refresh"
	Auth_SendPasswordlessToken_FullMethodName   = "/user.api.auth.v1.Auth/SendPasswordlessToken"
	Auth_LoginPasswordless_FullMethodName       = "/user.api.auth.v1.Auth/LoginPasswordless"
	Auth_SendForgetPasswordToken_FullMethodName = "/user.api.auth.v1.Auth/SendForgetPasswordToken"
	Auth_ForgetPassword_FullMethodName          = "/user.api.auth.v1.Auth/ForgetPassword"
	Auth_ChangePasswordByForget_FullMethodName  = "/user.api.auth.v1.Auth/ChangePasswordByForget"
	Auth_ValidatePassword_FullMethodName        = "/user.api.auth.v1.Auth/ValidatePassword"
	Auth_ChangePasswordByPre_FullMethodName     = "/user.api.auth.v1.Auth/ChangePasswordByPre"
	Auth_GetCsrfToken_FullMethodName            = "/user.api.auth.v1.Auth/GetCsrfToken"
	Auth_RefreshRememberToken_FullMethodName    = "/user.api.auth.v1.Auth/RefreshRememberToken"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	Register(ctx context.Context, in *RegisterAuthRequest, opts ...grpc.CallOption) (*RegisterAuthReply, error)
	Login(ctx context.Context, in *LoginAuthRequest, opts ...grpc.CallOption) (*LoginAuthReply, error)
	GetLogin(ctx context.Context, in *GetLoginRequest, opts ...grpc.CallOption) (*GetLoginResponse, error)
	Token(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenReply, error)
	Refresh(ctx context.Context, in *RefreshTokenAuthRequest, opts ...grpc.CallOption) (*RefreshTokenAuthReply, error)
	SendPasswordlessToken(ctx context.Context, in *PasswordlessTokenAuthRequest, opts ...grpc.CallOption) (*PasswordlessTokenAuthReply, error)
	LoginPasswordless(ctx context.Context, in *LoginPasswordlessRequest, opts ...grpc.CallOption) (*LoginPasswordlessReply, error)
	SendForgetPasswordToken(ctx context.Context, in *ForgetPasswordTokenRequest, opts ...grpc.CallOption) (*ForgetPasswordTokenReply, error)
	//ForgetPassword
	//verify forget password token
	ForgetPassword(ctx context.Context, in *ForgetPasswordRequest, opts ...grpc.CallOption) (*ForgetPasswordReply, error)
	ChangePasswordByForget(ctx context.Context, in *ChangePasswordByForgetRequest, opts ...grpc.CallOption) (*ChangePasswordByForgetReply, error)
	//ValidatePassword
	// server validation for password strength
	ValidatePassword(ctx context.Context, in *ValidatePasswordRequest, opts ...grpc.CallOption) (*ValidatePasswordReply, error)
	ChangePasswordByPre(ctx context.Context, in *ChangePasswordByPreRequest, opts ...grpc.CallOption) (*ChangePasswordByPreReply, error)
	GetCsrfToken(ctx context.Context, in *GetCsrfTokenRequest, opts ...grpc.CallOption) (*GetCsrfTokenResponse, error)
	//RefreshRememberToken internal api for refresh remember token
	RefreshRememberToken(ctx context.Context, in *RefreshRememberTokenRequest, opts ...grpc.CallOption) (*RefreshRememberTokenReply, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Register(ctx context.Context, in *RegisterAuthRequest, opts ...grpc.CallOption) (*RegisterAuthReply, error) {
	out := new(RegisterAuthReply)
	err := c.cc.Invoke(ctx, Auth_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Login(ctx context.Context, in *LoginAuthRequest, opts ...grpc.CallOption) (*LoginAuthReply, error) {
	out := new(LoginAuthReply)
	err := c.cc.Invoke(ctx, Auth_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetLogin(ctx context.Context, in *GetLoginRequest, opts ...grpc.CallOption) (*GetLoginResponse, error) {
	out := new(GetLoginResponse)
	err := c.cc.Invoke(ctx, Auth_GetLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Token(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenReply, error) {
	out := new(TokenReply)
	err := c.cc.Invoke(ctx, Auth_Token_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Refresh(ctx context.Context, in *RefreshTokenAuthRequest, opts ...grpc.CallOption) (*RefreshTokenAuthReply, error) {
	out := new(RefreshTokenAuthReply)
	err := c.cc.Invoke(ctx, Auth_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SendPasswordlessToken(ctx context.Context, in *PasswordlessTokenAuthRequest, opts ...grpc.CallOption) (*PasswordlessTokenAuthReply, error) {
	out := new(PasswordlessTokenAuthReply)
	err := c.cc.Invoke(ctx, Auth_SendPasswordlessToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) LoginPasswordless(ctx context.Context, in *LoginPasswordlessRequest, opts ...grpc.CallOption) (*LoginPasswordlessReply, error) {
	out := new(LoginPasswordlessReply)
	err := c.cc.Invoke(ctx, Auth_LoginPasswordless_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SendForgetPasswordToken(ctx context.Context, in *ForgetPasswordTokenRequest, opts ...grpc.CallOption) (*ForgetPasswordTokenReply, error) {
	out := new(ForgetPasswordTokenReply)
	err := c.cc.Invoke(ctx, Auth_SendForgetPasswordToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ForgetPassword(ctx context.Context, in *ForgetPasswordRequest, opts ...grpc.CallOption) (*ForgetPasswordReply, error) {
	out := new(ForgetPasswordReply)
	err := c.cc.Invoke(ctx, Auth_ForgetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ChangePasswordByForget(ctx context.Context, in *ChangePasswordByForgetRequest, opts ...grpc.CallOption) (*ChangePasswordByForgetReply, error) {
	out := new(ChangePasswordByForgetReply)
	err := c.cc.Invoke(ctx, Auth_ChangePasswordByForget_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ValidatePassword(ctx context.Context, in *ValidatePasswordRequest, opts ...grpc.CallOption) (*ValidatePasswordReply, error) {
	out := new(ValidatePasswordReply)
	err := c.cc.Invoke(ctx, Auth_ValidatePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ChangePasswordByPre(ctx context.Context, in *ChangePasswordByPreRequest, opts ...grpc.CallOption) (*ChangePasswordByPreReply, error) {
	out := new(ChangePasswordByPreReply)
	err := c.cc.Invoke(ctx, Auth_ChangePasswordByPre_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetCsrfToken(ctx context.Context, in *GetCsrfTokenRequest, opts ...grpc.CallOption) (*GetCsrfTokenResponse, error) {
	out := new(GetCsrfTokenResponse)
	err := c.cc.Invoke(ctx, Auth_GetCsrfToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RefreshRememberToken(ctx context.Context, in *RefreshRememberTokenRequest, opts ...grpc.CallOption) (*RefreshRememberTokenReply, error) {
	out := new(RefreshRememberTokenReply)
	err := c.cc.Invoke(ctx, Auth_RefreshRememberToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations should embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	Register(context.Context, *RegisterAuthRequest) (*RegisterAuthReply, error)
	Login(context.Context, *LoginAuthRequest) (*LoginAuthReply, error)
	GetLogin(context.Context, *GetLoginRequest) (*GetLoginResponse, error)
	Token(context.Context, *TokenRequest) (*TokenReply, error)
	Refresh(context.Context, *RefreshTokenAuthRequest) (*RefreshTokenAuthReply, error)
	SendPasswordlessToken(context.Context, *PasswordlessTokenAuthRequest) (*PasswordlessTokenAuthReply, error)
	LoginPasswordless(context.Context, *LoginPasswordlessRequest) (*LoginPasswordlessReply, error)
	SendForgetPasswordToken(context.Context, *ForgetPasswordTokenRequest) (*ForgetPasswordTokenReply, error)
	//ForgetPassword
	//verify forget password token
	ForgetPassword(context.Context, *ForgetPasswordRequest) (*ForgetPasswordReply, error)
	ChangePasswordByForget(context.Context, *ChangePasswordByForgetRequest) (*ChangePasswordByForgetReply, error)
	//ValidatePassword
	// server validation for password strength
	ValidatePassword(context.Context, *ValidatePasswordRequest) (*ValidatePasswordReply, error)
	ChangePasswordByPre(context.Context, *ChangePasswordByPreRequest) (*ChangePasswordByPreReply, error)
	GetCsrfToken(context.Context, *GetCsrfTokenRequest) (*GetCsrfTokenResponse, error)
	//RefreshRememberToken internal api for refresh remember token
	RefreshRememberToken(context.Context, *RefreshRememberTokenRequest) (*RefreshRememberTokenReply, error)
}

// UnimplementedAuthServer should be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) Register(context.Context, *RegisterAuthRequest) (*RegisterAuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthServer) Login(context.Context, *LoginAuthRequest) (*LoginAuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServer) GetLogin(context.Context, *GetLoginRequest) (*GetLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogin not implemented")
}
func (UnimplementedAuthServer) Token(context.Context, *TokenRequest) (*TokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Token not implemented")
}
func (UnimplementedAuthServer) Refresh(context.Context, *RefreshTokenAuthRequest) (*RefreshTokenAuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedAuthServer) SendPasswordlessToken(context.Context, *PasswordlessTokenAuthRequest) (*PasswordlessTokenAuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPasswordlessToken not implemented")
}
func (UnimplementedAuthServer) LoginPasswordless(context.Context, *LoginPasswordlessRequest) (*LoginPasswordlessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginPasswordless not implemented")
}
func (UnimplementedAuthServer) SendForgetPasswordToken(context.Context, *ForgetPasswordTokenRequest) (*ForgetPasswordTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendForgetPasswordToken not implemented")
}
func (UnimplementedAuthServer) ForgetPassword(context.Context, *ForgetPasswordRequest) (*ForgetPasswordReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgetPassword not implemented")
}
func (UnimplementedAuthServer) ChangePasswordByForget(context.Context, *ChangePasswordByForgetRequest) (*ChangePasswordByForgetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePasswordByForget not implemented")
}
func (UnimplementedAuthServer) ValidatePassword(context.Context, *ValidatePasswordRequest) (*ValidatePasswordReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatePassword not implemented")
}
func (UnimplementedAuthServer) ChangePasswordByPre(context.Context, *ChangePasswordByPreRequest) (*ChangePasswordByPreReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePasswordByPre not implemented")
}
func (UnimplementedAuthServer) GetCsrfToken(context.Context, *GetCsrfTokenRequest) (*GetCsrfTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCsrfToken not implemented")
}
func (UnimplementedAuthServer) RefreshRememberToken(context.Context, *RefreshRememberTokenRequest) (*RefreshRememberTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshRememberToken not implemented")
}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Register(ctx, req.(*RegisterAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*LoginAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetLogin(ctx, req.(*GetLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Token(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Token_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Token(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Refresh(ctx, req.(*RefreshTokenAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SendPasswordlessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordlessTokenAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SendPasswordlessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_SendPasswordlessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SendPasswordlessToken(ctx, req.(*PasswordlessTokenAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_LoginPasswordless_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginPasswordlessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).LoginPasswordless(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_LoginPasswordless_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).LoginPasswordless(ctx, req.(*LoginPasswordlessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SendForgetPasswordToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgetPasswordTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SendForgetPasswordToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_SendForgetPasswordToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SendForgetPasswordToken(ctx, req.(*ForgetPasswordTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ForgetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ForgetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_ForgetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ForgetPassword(ctx, req.(*ForgetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ChangePasswordByForget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordByForgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ChangePasswordByForget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_ChangePasswordByForget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ChangePasswordByForget(ctx, req.(*ChangePasswordByForgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ValidatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ValidatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_ValidatePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ValidatePassword(ctx, req.(*ValidatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ChangePasswordByPre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordByPreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ChangePasswordByPre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_ChangePasswordByPre_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ChangePasswordByPre(ctx, req.(*ChangePasswordByPreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetCsrfToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCsrfTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetCsrfToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetCsrfToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetCsrfToken(ctx, req.(*GetCsrfTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RefreshRememberToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRememberTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RefreshRememberToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_RefreshRememberToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RefreshRememberToken(ctx, req.(*RefreshRememberTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.api.auth.v1.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Auth_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
		{
			MethodName: "GetLogin",
			Handler:    _Auth_GetLogin_Handler,
		},
		{
			MethodName: "Token",
			Handler:    _Auth_Token_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _Auth_Refresh_Handler,
		},
		{
			MethodName: "SendPasswordlessToken",
			Handler:    _Auth_SendPasswordlessToken_Handler,
		},
		{
			MethodName: "LoginPasswordless",
			Handler:    _Auth_LoginPasswordless_Handler,
		},
		{
			MethodName: "SendForgetPasswordToken",
			Handler:    _Auth_SendForgetPasswordToken_Handler,
		},
		{
			MethodName: "ForgetPassword",
			Handler:    _Auth_ForgetPassword_Handler,
		},
		{
			MethodName: "ChangePasswordByForget",
			Handler:    _Auth_ChangePasswordByForget_Handler,
		},
		{
			MethodName: "ValidatePassword",
			Handler:    _Auth_ValidatePassword_Handler,
		},
		{
			MethodName: "ChangePasswordByPre",
			Handler:    _Auth_ChangePasswordByPre_Handler,
		},
		{
			MethodName: "GetCsrfToken",
			Handler:    _Auth_GetCsrfToken_Handler,
		},
		{
			MethodName: "RefreshRememberToken",
			Handler:    _Auth_RefreshRememberToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/api/auth/v1/auth.proto",
}
