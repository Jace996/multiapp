package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	sapi "github.com/jace996/multiapp/pkg/api"
	"github.com/jace996/multiapp/pkg/authn/jwt"
	"github.com/jace996/multiapp/pkg/conf"
	"github.com/jace996/multiapp/pkg/localize"
	"github.com/jace996/multiapp/pkg/logging"
	"github.com/jace996/multiapp/pkg/server"
	"github.com/jace996/multiapp/pkg/server/common"
	kitgrpc "github.com/jace996/multiapp/pkg/server/grpc"
	"github.com/jace996/multiapp/user/api"
	v12 "github.com/jace996/multiapp/user/api/user/v1"
	"github.com/jace996/saas"
	http2 "github.com/jace996/saas/http"
	uow2 "github.com/jace996/uow"
	kuow "github.com/jace996/uow/kratos"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf.Services,
	tokenizer jwt.Tokenizer,
	ts saas.TenantStore,
	uowMgr uow2.Manager,
	mOpt *http2.WebMultiTenancyOption,
	apiOpt *sapi.Option,
	logger log.Logger,
	validator sapi.TrustedContextValidator,
	userTenant *api.UserTenantContrib,
	register []kitgrpc.ServiceRegister,
) *kitgrpc.Server {
	m := []middleware.Middleware{
		server.Recovery(),
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
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			m...,
		),
	}
	cfg := common.GetConf(c, api.ServiceName)
	opts = kitgrpc.PatchOpts(logger, opts, cfg)
	srv := kitgrpc.NewServer(cfg, opts...)

	kitgrpc.ChainServiceRegister(register...).Register(srv.Server, m...)

	return srv
}
