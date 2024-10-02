package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/jace996/multiapp/pkg/api"
	"github.com/jace996/multiapp/pkg/conf"
	"github.com/jace996/saas"
	shttp "github.com/jace996/saas/http"
	ksaas "github.com/jace996/saas/kratos"
)

func Saas(hmtOpt *shttp.WebMultiTenancyOption, ts saas.TenantStore, trustedContextValidator api.TrustedContextValidator, options ...saas.ResolveOption) middleware.Middleware {
	return selector.Server(ksaas.Server(
		ts,
		ksaas.WithMultiTenancyOption(hmtOpt),
		ksaas.WithResolveOption(options...),
	)).Match(func(ctx context.Context, operation string) bool {
		ok, _ := trustedContextValidator.Trusted(ctx)
		return !ok
	}).Build()
}

func NewWebMultiTenancyOption(opt *conf.AppConfig) *shttp.WebMultiTenancyOption {
	ret := shttp.NewDefaultWebMultiTenancyOption()
	if opt == nil {
		return ret
	}
	if opt.TenantKey != nil {
		ret.TenantKey = opt.TenantKey.Value
	}
	if opt.DomainFormat != nil {
		ret.DomainFormat = opt.DomainFormat.Value
	}
	return ret
}
