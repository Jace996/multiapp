// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             (unknown)
// source: saas/api/tenant/v1/tenant.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationTenantServiceChangeTenant = "/saas.api.tenant.v1.TenantService/ChangeTenant"
const OperationTenantServiceCreateTenant = "/saas.api.tenant.v1.TenantService/CreateTenant"
const OperationTenantServiceDeleteTenant = "/saas.api.tenant.v1.TenantService/DeleteTenant"
const OperationTenantServiceGetCurrentTenant = "/saas.api.tenant.v1.TenantService/GetCurrentTenant"
const OperationTenantServiceGetTenant = "/saas.api.tenant.v1.TenantService/GetTenant"
const OperationTenantServiceGetTenantPublic = "/saas.api.tenant.v1.TenantService/GetTenantPublic"
const OperationTenantServiceListTenant = "/saas.api.tenant.v1.TenantService/ListTenant"
const OperationTenantServiceUpdateTenant = "/saas.api.tenant.v1.TenantService/UpdateTenant"
const OperationTenantServiceUserCreateTenant = "/saas.api.tenant.v1.TenantService/UserCreateTenant"

type TenantServiceHTTPServer interface {
	ChangeTenant(context.Context, *ChangeTenantRequest) (*ChangeTenantReply, error)
	// CreateTenantCreateTenant
	//authz: saas.tenant,*,create
	CreateTenant(context.Context, *CreateTenantRequest) (*Tenant, error)
	// DeleteTenantDeleteTenant
	//authz: saas.tenant,{id},delete
	DeleteTenant(context.Context, *DeleteTenantRequest) (*DeleteTenantReply, error)
	// GetCurrentTenantGetCurrentTenant
	GetCurrentTenant(context.Context, *GetCurrentTenantRequest) (*GetCurrentTenantReply, error)
	// GetTenant GetTenant
	// authz: saas.tenant,{id},get
	GetTenant(context.Context, *GetTenantRequest) (*Tenant, error)
	// GetTenantPublic GetTenant
	// authz: saas.tenant,{id},get
	GetTenantPublic(context.Context, *GetTenantPublicRequest) (*TenantInfo, error)
	// ListTenantListTenant
	//authz: saas.tenant,*,list
	ListTenant(context.Context, *ListTenantRequest) (*ListTenantReply, error)
	// UpdateTenantUpdateTenant
	//authz: saas.tenant,{id},update
	UpdateTenant(context.Context, *UpdateTenantRequest) (*Tenant, error)
	// UserCreateTenantCreateTenant
	UserCreateTenant(context.Context, *UserCreateTenantRequest) (*UserCreateTenantReply, error)
}

func RegisterTenantServiceHTTPServer(s *http.Server, srv TenantServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/saas/tenant", _TenantService_CreateTenant0_HTTP_Handler(srv))
	r.PATCH("/v1/saas/tenant/{tenant.id}", _TenantService_UpdateTenant0_HTTP_Handler(srv))
	r.PUT("/v1/saas/tenant/{tenant.id}", _TenantService_UpdateTenant1_HTTP_Handler(srv))
	r.DELETE("/v1/saas/tenant/{id}", _TenantService_DeleteTenant0_HTTP_Handler(srv))
	r.GET("/v1/saas/tenant/{id_or_name}", _TenantService_GetTenant0_HTTP_Handler(srv))
	r.GET("/v1/saas/tenant/{id_or_name}/public", _TenantService_GetTenantPublic0_HTTP_Handler(srv))
	r.POST("/v1/saas/tenant/list", _TenantService_ListTenant0_HTTP_Handler(srv))
	r.GET("/v1/saas/tenants", _TenantService_ListTenant1_HTTP_Handler(srv))
	r.GET("/v1/saas/current", _TenantService_GetCurrentTenant0_HTTP_Handler(srv))
	r.POST("/v1/saas/change-tenant/{id_or_name}", _TenantService_ChangeTenant0_HTTP_Handler(srv))
	r.POST("/v1/saas/user/tenant", _TenantService_UserCreateTenant0_HTTP_Handler(srv))
}

func _TenantService_CreateTenant0_HTTP_Handler(srv TenantServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTenantRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTenantServiceCreateTenant)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateTenant(ctx, req.(*CreateTenantRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Tenant)
		return ctx.Result(200, reply)
	}
}

func _TenantService_UpdateTenant0_HTTP_Handler(srv TenantServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTenantRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTenantServiceUpdateTenant)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateTenant(ctx, req.(*UpdateTenantRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Tenant)
		return ctx.Result(200, reply)
	}
}

func _TenantService_UpdateTenant1_HTTP_Handler(srv TenantServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTenantRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTenantServiceUpdateTenant)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateTenant(ctx, req.(*UpdateTenantRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Tenant)
		return ctx.Result(200, reply)
	}
}

func _TenantService_DeleteTenant0_HTTP_Handler(srv TenantServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteTenantRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTenantServiceDeleteTenant)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteTenant(ctx, req.(*DeleteTenantRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteTenantReply)
		return ctx.Result(200, reply)
	}
}

func _TenantService_GetTenant0_HTTP_Handler(srv TenantServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTenantRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTenantServiceGetTenant)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTenant(ctx, req.(*GetTenantRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Tenant)
		return ctx.Result(200, reply)
	}
}

func _TenantService_GetTenantPublic0_HTTP_Handler(srv TenantServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTenantPublicRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTenantServiceGetTenantPublic)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTenantPublic(ctx, req.(*GetTenantPublicRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TenantInfo)
		return ctx.Result(200, reply)
	}
}

func _TenantService_ListTenant0_HTTP_Handler(srv TenantServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTenantRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTenantServiceListTenant)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTenant(ctx, req.(*ListTenantRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTenantReply)
		return ctx.Result(200, reply)
	}
}

func _TenantService_ListTenant1_HTTP_Handler(srv TenantServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTenantRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTenantServiceListTenant)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTenant(ctx, req.(*ListTenantRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTenantReply)
		return ctx.Result(200, reply)
	}
}

func _TenantService_GetCurrentTenant0_HTTP_Handler(srv TenantServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCurrentTenantRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTenantServiceGetCurrentTenant)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCurrentTenant(ctx, req.(*GetCurrentTenantRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCurrentTenantReply)
		return ctx.Result(200, reply)
	}
}

func _TenantService_ChangeTenant0_HTTP_Handler(srv TenantServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ChangeTenantRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTenantServiceChangeTenant)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ChangeTenant(ctx, req.(*ChangeTenantRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ChangeTenantReply)
		return ctx.Result(200, reply)
	}
}

func _TenantService_UserCreateTenant0_HTTP_Handler(srv TenantServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserCreateTenantRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTenantServiceUserCreateTenant)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserCreateTenant(ctx, req.(*UserCreateTenantRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserCreateTenantReply)
		return ctx.Result(200, reply)
	}
}

type TenantServiceHTTPClient interface {
	ChangeTenant(ctx context.Context, req *ChangeTenantRequest, opts ...http.CallOption) (rsp *ChangeTenantReply, err error)
	CreateTenant(ctx context.Context, req *CreateTenantRequest, opts ...http.CallOption) (rsp *Tenant, err error)
	DeleteTenant(ctx context.Context, req *DeleteTenantRequest, opts ...http.CallOption) (rsp *DeleteTenantReply, err error)
	GetCurrentTenant(ctx context.Context, req *GetCurrentTenantRequest, opts ...http.CallOption) (rsp *GetCurrentTenantReply, err error)
	GetTenant(ctx context.Context, req *GetTenantRequest, opts ...http.CallOption) (rsp *Tenant, err error)
	GetTenantPublic(ctx context.Context, req *GetTenantPublicRequest, opts ...http.CallOption) (rsp *TenantInfo, err error)
	ListTenant(ctx context.Context, req *ListTenantRequest, opts ...http.CallOption) (rsp *ListTenantReply, err error)
	UpdateTenant(ctx context.Context, req *UpdateTenantRequest, opts ...http.CallOption) (rsp *Tenant, err error)
	UserCreateTenant(ctx context.Context, req *UserCreateTenantRequest, opts ...http.CallOption) (rsp *UserCreateTenantReply, err error)
}

type TenantServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewTenantServiceHTTPClient(client *http.Client) TenantServiceHTTPClient {
	return &TenantServiceHTTPClientImpl{client}
}

func (c *TenantServiceHTTPClientImpl) ChangeTenant(ctx context.Context, in *ChangeTenantRequest, opts ...http.CallOption) (*ChangeTenantReply, error) {
	var out ChangeTenantReply
	pattern := "/v1/saas/change-tenant/{id_or_name}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTenantServiceChangeTenant))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TenantServiceHTTPClientImpl) CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...http.CallOption) (*Tenant, error) {
	var out Tenant
	pattern := "/v1/saas/tenant"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTenantServiceCreateTenant))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TenantServiceHTTPClientImpl) DeleteTenant(ctx context.Context, in *DeleteTenantRequest, opts ...http.CallOption) (*DeleteTenantReply, error) {
	var out DeleteTenantReply
	pattern := "/v1/saas/tenant/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTenantServiceDeleteTenant))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TenantServiceHTTPClientImpl) GetCurrentTenant(ctx context.Context, in *GetCurrentTenantRequest, opts ...http.CallOption) (*GetCurrentTenantReply, error) {
	var out GetCurrentTenantReply
	pattern := "/v1/saas/current"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTenantServiceGetCurrentTenant))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TenantServiceHTTPClientImpl) GetTenant(ctx context.Context, in *GetTenantRequest, opts ...http.CallOption) (*Tenant, error) {
	var out Tenant
	pattern := "/v1/saas/tenant/{id_or_name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTenantServiceGetTenant))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TenantServiceHTTPClientImpl) GetTenantPublic(ctx context.Context, in *GetTenantPublicRequest, opts ...http.CallOption) (*TenantInfo, error) {
	var out TenantInfo
	pattern := "/v1/saas/tenant/{id_or_name}/public"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTenantServiceGetTenantPublic))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TenantServiceHTTPClientImpl) ListTenant(ctx context.Context, in *ListTenantRequest, opts ...http.CallOption) (*ListTenantReply, error) {
	var out ListTenantReply
	pattern := "/v1/saas/tenants"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTenantServiceListTenant))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TenantServiceHTTPClientImpl) UpdateTenant(ctx context.Context, in *UpdateTenantRequest, opts ...http.CallOption) (*Tenant, error) {
	var out Tenant
	pattern := "/v1/saas/tenant/{tenant.id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTenantServiceUpdateTenant))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TenantServiceHTTPClientImpl) UserCreateTenant(ctx context.Context, in *UserCreateTenantRequest, opts ...http.CallOption) (*UserCreateTenantReply, error) {
	var out UserCreateTenantReply
	pattern := "/v1/saas/user/tenant"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTenantServiceUserCreateTenant))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
