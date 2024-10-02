package main

import (
	"github.com/go-kratos/kratos/v2/log"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/jace996/multiapp/gateway/apisix/cmd/go-runner/plugins"
	"github.com/jace996/multiapp/pkg/api"
	"github.com/jace996/multiapp/pkg/authn/jwt"
	"github.com/jace996/multiapp/pkg/authn/session"
	"github.com/jace996/multiapp/pkg/authz/authz"
	"github.com/jace996/multiapp/pkg/conf"
	kitdi "github.com/jace996/multiapp/pkg/di"
	"github.com/jace996/multiapp/pkg/server"
	uapi "github.com/jace996/multiapp/user/api"
	"github.com/jace996/saas"
	shttp "github.com/jace996/saas/http"
	"github.com/goava/di"
)

type App struct {
	tenantStore     saas.TenantStore
	tokenizer       jwt.Tokenizer
	tokenManager    api.TokenManager
	services        *conf.Services
	clientName      api.ClientName
	logger          klog.Logger
	tenantCfg       *shttp.WebMultiTenancyOption
	security        *conf.Security
	userTenant      *uapi.UserTenantContrib
	refreshProvider session.RefreshTokenProvider
	authService     authz.Service
	subjectResolver authz.SubjectResolver
}

func newApp(tenantStore saas.TenantStore,
	userTenant *uapi.UserTenantContrib,
	t jwt.Tokenizer,
	tmr api.TokenManager,
	services *conf.Services,
	clientName api.ClientName,
	tenantCfg *shttp.WebMultiTenancyOption,
	security *conf.Security,
	refreshProvider session.RefreshTokenProvider,
	authService authz.Service,
	subjectResolver authz.SubjectResolver,

	logger klog.Logger) (*App, error) {
	ret := &App{tenantStore: tenantStore,
		userTenant:      userTenant,
		tokenizer:       t,
		tokenManager:    tmr,
		services:        services,
		clientName:      clientName,
		tenantCfg:       tenantCfg,
		security:        security,
		refreshProvider: refreshProvider,
		authService:     authService,
		subjectResolver: subjectResolver,
		logger:          logger}
	return ret, ret.load()
}

func (a *App) load() error {
	if err := plugins.Init(
		a.tokenizer,
		a.tokenManager,
		a.tenantCfg,
		a.clientName,
		a.services,
		a.security,
		a.userTenant,
		a.tenantStore,
		a.refreshProvider,
		a.authService,
		a.subjectResolver,
		a.logger,
	); err != nil {
		return err
	}
	return nil
}

func NewSelfClientOption(logger log.Logger) *api.Option {
	return api.NewOption(
		false,
		api.NewSaasPropagator(logger),
		api.NewUserPropagator(logger),
		//do not propagate client
		api.NewClientPropagator(true, logger),
	).WithInsecure()
}

func NewAuthorizationOption() *authz.Option {
	return authz.NewAuthorizationOption()
}

var ProviderSet = kitdi.NewSet(
	api.NewClientCfg,
	kitdi.NewProvider(api.NewInMemoryTokenManager, di.As(new(api.TokenManager))),
	api.NewInMemoryTokenManager,
	NewSelfClientOption,
	NewAuthorizationOption,
	server.NewWebMultiTenancyOption,
	api.NewDiscovery)
