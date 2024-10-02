package api

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	grpc2 "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/jace996/multiapp/pkg/api"
	"github.com/jace996/multiapp/pkg/conf"
	kitdi "github.com/jace996/multiapp/pkg/di"
	v1 "github.com/jace996/multiapp/saas/api/tenant/v1"
	_ "github.com/jace996/multiapp/saas/i18n"
	"google.golang.org/grpc"
)

type GrpcConn grpc.ClientConnInterface
type HttpClient *http.Client

const ServiceName = "saas"

func NewGrpcConn(
	client *conf.Client,
	services *conf.Services,
	dis registry.Discovery,
	opt *api.Option,
	tokenMgr api.TokenManager,
	logger log.Logger,
	opts []grpc2.ClientOption,
) (GrpcConn, func()) {
	return api.NewGrpcConn(client, ServiceName, services, dis, opt, tokenMgr, logger, opts)
}

var GrpcProviderSet = kitdi.NewSet(NewTenantStore, NewGrpcConn, NewTenantGrpcClient, NewTenantInternalGrpcClient)

func NewTenantGrpcClient(conn GrpcConn) v1.TenantServiceServer {
	return v1.NewTenantServiceClientProxy(v1.NewTenantServiceClient(conn))
}

func NewTenantInternalGrpcClient(conn GrpcConn) v1.TenantInternalServiceServer {
	return v1.NewTenantInternalServiceClientProxy(v1.NewTenantInternalServiceClient(conn))
}
