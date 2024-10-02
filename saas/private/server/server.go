package server

import (
	"github.com/go-kratos/kratos/v2/transport"
	kapi "github.com/jace996/multiapp/pkg/api"
	"github.com/jace996/multiapp/pkg/authz/authz"
	kitdi "github.com/jace996/multiapp/pkg/di"
	"github.com/jace996/multiapp/pkg/server"
	"github.com/jace996/multiapp/saas/api"
	"github.com/jace996/multiapp/saas/private/biz"
	"github.com/jace996/multiapp/saas/private/data"
	"github.com/jace996/saas"
	"github.com/jace996/saas/seed"
	"github.com/jace996/uow"
	"github.com/goava/di"
)

// ProviderSet is server providers.
var ProviderSet = kitdi.NewSet(
	kitdi.NewProvider(NewHTTPServer, di.As(new(transport.Server))),
	kitdi.NewProvider(NewGRPCServer, di.As(new(transport.Server))),
	kitdi.NewProvider(NewJobServer, di.As(new(transport.Server))),
	kitdi.NewProvider(NewEventServer, di.As(new(transport.Server))),
	kitdi.Value(ClientName),
	kitdi.Value(biz.ConnName),
	NewSeeder,
	NewSeeding,
	NewAuthorizationOption,
)

var ClientName kapi.ClientName = api.ServiceName

func NewSeeding(uow uow.Manager, migrate *data.Migrate) seed.Contrib {
	return server.NewUowContrib(uow, seed.Chain(migrate))
}

func NewSeeder(ts saas.TenantStore, seeds []seed.Contrib) seed.Seeder {
	res := seed.NewDefaultSeeder(server.NewTraceContrib(server.SeedChangeTenant(ts, seeds...)))
	return res
}

func NewAuthorizationOption() *authz.Option {
	return authz.NewAuthorizationOption()
}
