package api

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	grpc2 "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/jace996/multiapp/order/api/order/v1"
	_ "github.com/jace996/multiapp/order/i18n"
	"github.com/jace996/multiapp/pkg/api"
	"github.com/jace996/multiapp/pkg/conf"
	kitdi "github.com/jace996/multiapp/pkg/di"
	"google.golang.org/grpc"
)

type GrpcConn grpc.ClientConnInterface
type HttpClient *http.Client

const ServiceName = "order"

func NewGrpcConn(client *conf.Client, services *conf.Services, dis registry.Discovery, opt *api.Option, tokenMgr api.TokenManager, logger log.Logger, opts []grpc2.ClientOption) (GrpcConn, func()) {
	return api.NewGrpcConn(client, ServiceName, services, dis, opt, tokenMgr, logger, opts)
}

var GrpcProviderSet = kitdi.NewSet(NewGrpcConn, NewOrderGrpcClient, NewOrderInternalGrpcClient)

func NewOrderGrpcClient(conn GrpcConn) v1.OrderServiceServer {
	return v1.NewOrderServiceClientProxy(v1.NewOrderServiceClient(conn))
}

func NewOrderInternalGrpcClient(conn GrpcConn) v1.OrderInternalServiceServer {
	return v1.NewOrderInternalServiceClientProxy(v1.NewOrderInternalServiceClient(conn))
}
