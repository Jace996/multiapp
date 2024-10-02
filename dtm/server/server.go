package server

import (
	"github.com/dtm-labs/dtm/client/dtmcli/dtmimp"
	"github.com/dtm-labs/dtm/client/dtmgrpc"
	"github.com/dtm-labs/dtm/client/workflow"
	driver "github.com/dtm-labs/dtmdriver-kratos"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc/resolver/discovery"
	dtmapi "github.com/jace996/multiapp/dtm/api"
	v1 "github.com/jace996/multiapp/dtm/api/dtm/v1"
	"github.com/jace996/multiapp/dtm/data"
	"github.com/jace996/multiapp/dtm/service"
	sapi "github.com/jace996/multiapp/pkg/api"
	kitdi "github.com/jace996/multiapp/pkg/di"
	"github.com/jace996/multiapp/pkg/gorm"
	kitserver "github.com/jace996/multiapp/pkg/server"
	kitgrpc "github.com/jace996/multiapp/pkg/server/grpc"
	"github.com/defval/di"
	"google.golang.org/grpc/resolver"
	"sync"
)

var DtmProviderSet = kitdi.NewSet(
	NewInit,
	service.NewHelper,
	kitdi.NewProvider(service.NewMsgService, di.As(new(v1.MsgServiceServer))),
	data.NewBarrierMigrator, data.NewStorageMigrator, data.NewMigrator,
	service.NewHttpServerRegister, service.NewGrpcServerRegister,
)

var (
	once sync.Once
)

type Init any

// NewInit init dtm required info
func NewInit(dis registry.Discovery, name kitserver.Name, cache *gorm.SqlDbCache, srv *kitgrpc.Server, opt *sapi.Option) (Init, error) {
	var opts = []discovery.Option{
		discovery.WithInsecure(opt.Insecure),
	}
	if opt.Subset != nil {
		opts = append(opts, discovery.WithSubset(*opt.Subset))
	}
	dtmimp.SetDbCache(cache)
	dtmgrpc.UseDriver(driver.DriverName)
	resolver.Register(discovery.NewBuilder(dis, opts...))
	workflow.InitGrpc(sapi.WithDiscovery(dtmapi.ServiceName), sapi.WithDiscovery(string(name)), srv.Server.Server)
	return "", nil
}
