package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	sapi "github.com/jace996/multiapp/pkg/api"
	"github.com/jace996/multiapp/pkg/authn/jwt"
	conf2 "github.com/jace996/multiapp/pkg/conf"
	"github.com/jace996/multiapp/pkg/localize"
	"github.com/jace996/multiapp/pkg/logging"
	"github.com/jace996/multiapp/pkg/server"
	"github.com/jace996/multiapp/pkg/server/common"
	kithttp "github.com/jace996/multiapp/pkg/server/http"
	"github.com/jace996/multiapp/sys/api"
	uow2 "github.com/jace996/uow"
	kuow "github.com/jace996/uow/kratos"
)

// NewHTTPServer new a HTTP kithttp.
func NewHTTPServer(c *conf2.Services,
	sCfg *conf2.Security,
	tokenizer jwt.Tokenizer,
	uowMgr uow2.Manager,
	reqDecoder http.DecodeRequestFunc,
	resEncoder http.EncodeResponseFunc,
	errEncoder http.EncodeErrorFunc,
	apiOpt *sapi.Option,
	logger log.Logger,
	validator sapi.TrustedContextValidator,
	register []kithttp.ServiceRegister,
) *kithttp.Server {
	var opts []http.ServerOption
	cfg := common.GetConf(c, api.ServiceName)
	opts = kithttp.PatchOpts(logger, opts, cfg, sCfg, reqDecoder, resEncoder, errEncoder)
	m := []middleware.Middleware{server.Recovery(),
		tracing.Server(),
		logging.Server(logger),
		metrics.Server(),
		validate.Validator(),
		localize.I18N(),
		jwt.ServerExtractAndAuth(tokenizer, logger),
		sapi.ServerPropagation(apiOpt, validator, logger),
		kuow.Uow(uowMgr)}
	opts = append(opts, []http.ServerOption{
		http.Middleware(
			m...,
		),
	}...)
	srv := kithttp.NewServer(cfg, opts...)

	kithttp.ChainServiceRegister(register...).Register(srv.Server, m...)

	return srv
}
