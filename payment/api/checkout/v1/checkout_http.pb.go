// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             (unknown)
// source: payment/api/checkout/v1/checkout.proto

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

const OperationCheckoutServiceCheckoutNow = "/payment.api.checkout.v1.CheckoutService/CheckoutNow"
const OperationCheckoutServiceCheckoutOrder = "/payment.api.checkout.v1.CheckoutService/CheckoutOrder"

type CheckoutServiceHTTPServer interface {
	CheckoutNow(context.Context, *CheckoutNowRequest) (*CheckoutNowReply, error)
	CheckoutOrder(context.Context, *CheckOutOrderRequest) (*CheckoutOrderReply, error)
}

func RegisterCheckoutServiceHTTPServer(s *http.Server, srv CheckoutServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/payment/checkout/now", _CheckoutService_CheckoutNow0_HTTP_Handler(srv))
	r.POST("/v1/payment/checkout/order/{order_id}", _CheckoutService_CheckoutOrder0_HTTP_Handler(srv))
}

func _CheckoutService_CheckoutNow0_HTTP_Handler(srv CheckoutServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CheckoutNowRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCheckoutServiceCheckoutNow)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CheckoutNow(ctx, req.(*CheckoutNowRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CheckoutNowReply)
		return ctx.Result(200, reply)
	}
}

func _CheckoutService_CheckoutOrder0_HTTP_Handler(srv CheckoutServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CheckOutOrderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCheckoutServiceCheckoutOrder)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CheckoutOrder(ctx, req.(*CheckOutOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CheckoutOrderReply)
		return ctx.Result(200, reply)
	}
}

type CheckoutServiceHTTPClient interface {
	CheckoutNow(ctx context.Context, req *CheckoutNowRequest, opts ...http.CallOption) (rsp *CheckoutNowReply, err error)
	CheckoutOrder(ctx context.Context, req *CheckOutOrderRequest, opts ...http.CallOption) (rsp *CheckoutOrderReply, err error)
}

type CheckoutServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewCheckoutServiceHTTPClient(client *http.Client) CheckoutServiceHTTPClient {
	return &CheckoutServiceHTTPClientImpl{client}
}

func (c *CheckoutServiceHTTPClientImpl) CheckoutNow(ctx context.Context, in *CheckoutNowRequest, opts ...http.CallOption) (*CheckoutNowReply, error) {
	var out CheckoutNowReply
	pattern := "/v1/payment/checkout/now"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCheckoutServiceCheckoutNow))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CheckoutServiceHTTPClientImpl) CheckoutOrder(ctx context.Context, in *CheckOutOrderRequest, opts ...http.CallOption) (*CheckoutOrderReply, error) {
	var out CheckoutOrderReply
	pattern := "/v1/payment/checkout/order/{order_id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCheckoutServiceCheckoutOrder))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
