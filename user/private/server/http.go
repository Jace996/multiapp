package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	sapi "github.com/jace996/multiapp/pkg/api"
	"github.com/jace996/multiapp/pkg/authn/jwt"
	"github.com/jace996/multiapp/pkg/authn/session"
	"github.com/jace996/multiapp/pkg/conf"
	"github.com/jace996/multiapp/pkg/localize"
	"github.com/jace996/multiapp/pkg/logging"
	"github.com/jace996/multiapp/pkg/server"
	"github.com/jace996/multiapp/pkg/server/common"
	kithttp "github.com/jace996/multiapp/pkg/server/http"
	"github.com/jace996/multiapp/user/api"
	v12 "github.com/jace996/multiapp/user/api/user/v1"
	"github.com/jace996/saas"
	shttp "github.com/jace996/saas/http"
	uow2 "github.com/jace996/uow"
	kuow "github.com/jace996/uow/kratos"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Services,
	sCfg *conf.Security,
	tokenizer jwt.Tokenizer,
	uowMgr uow2.Manager,
	mOpt *shttp.WebMultiTenancyOption,
	apiOpt *sapi.Option,
	ts saas.TenantStore,
	reqDecoder khttp.DecodeRequestFunc,
	resEncoder khttp.EncodeResponseFunc,
	errEncoder khttp.EncodeErrorFunc,
	logger log.Logger,
	userTenant *api.UserTenantContrib,
	validator sapi.TrustedContextValidator,
	refreshProvider session.RefreshTokenProvider,
	register []kithttp.ServiceRegister,
) *kithttp.Server {
	var opts []khttp.ServerOption
	cfg := common.GetConf(c, api.ServiceName)
	opts = kithttp.PatchOpts(logger, opts, cfg, sCfg, reqDecoder, resEncoder, errEncoder,
		session.Auth(sCfg, validator),
		session.Refresh(errEncoder, refreshProvider, validator),
	)
	middlewares := []middleware.Middleware{server.Recovery(),
		tracing.Server(),
		logging.Server(logger),
		metrics.Server(),
		validate.Validator(),
		localize.I18N(),
		jwt.ServerExtractAndAuth(tokenizer, logger),
		sapi.ServerPropagation(apiOpt, validator, logger),
		server.Saas(mOpt, ts, validator, func(o *saas.TenantResolveOption) {
			o.AppendContribs(userTenant)
		}),
		kuow.Uow(uowMgr, kuow.WithForceSkipOp(v12.UserInternalService_CreateTenant_FullMethodName)),
	}
	opts = append(opts, []khttp.ServerOption{
		khttp.Middleware(middlewares...),
	}...)

	srv := kithttp.NewServer(cfg, opts...)

	kithttp.ChainServiceRegister(register...).Register(srv.Server, middlewares...)
	return srv
}
