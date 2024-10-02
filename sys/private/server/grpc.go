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
	conf2 "github.com/jace996/multiapp/pkg/conf"
	"github.com/jace996/multiapp/pkg/localize"
	"github.com/jace996/multiapp/pkg/logging"
	"github.com/jace996/multiapp/pkg/server"
	"github.com/jace996/multiapp/pkg/server/common"
	kitgrpc "github.com/jace996/multiapp/pkg/server/grpc"
	"github.com/jace996/multiapp/sys/api"
	uow2 "github.com/jace996/uow"
	kuow "github.com/jace996/uow/kratos"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf2.Services,
	tokenizer jwt.Tokenizer,
	uowMgr uow2.Manager,
	apiOpt *sapi.Option,
	logger log.Logger,
	validator sapi.TrustedContextValidator,
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
		kuow.Uow(uowMgr)}
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
