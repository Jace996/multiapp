package server

import (
	"github.com/go-kratos/kratos/v2/transport"
	dtmdata "github.com/jace996/multiapp/dtm/data"
	"github.com/jace996/multiapp/pkg/api"
	"github.com/jace996/multiapp/pkg/authz/authz"
	kitdi "github.com/jace996/multiapp/pkg/di"
	"github.com/jace996/multiapp/pkg/server"
	api2 "github.com/jace996/multiapp/user/api"
	"github.com/jace996/multiapp/user/private/biz"
	"github.com/jace996/multiapp/user/private/data"
	"github.com/jace996/multiapp/user/private/service"
	"github.com/jace996/saas"
	"github.com/jace996/saas/seed"
	"github.com/jace996/uow"
	"github.com/goava/di"
)

// ProviderSet is server providers.
var ProviderSet = kitdi.NewSet(
	kitdi.NewProvider(NewHTTPServer, di.As(new(transport.Server))),
	kitdi.NewProvider(NewGRPCServer, di.As(new(transport.Server))),
	//kitdi.NewProvider(NewJobServer, di.As(new(transport.Server))),
	kitdi.NewProvider(NewEventServer, di.As(new(transport.Server))),
	kitdi.Value(ClientName),
	kitdi.Value(biz.ConnName),
	NewSeeding,
	NewSeeder,
	NewAuthorizationOption,
)

var ClientName api.ClientName = api2.ServiceName

// NewSeeding wrap all service migrator into one seed.Contrib, which grants the running sequence of those contribs
func NewSeeding(uow uow.Manager,
	migrate *data.Migrate,
	dtmMigrator *dtmdata.BarrierMigrator, //barrier only
	roleSeed *biz.RoleSeed,
	userSeed *biz.UserSeed,
	p *biz.PermissionSeeder) seed.Contrib {
	return seed.Chain(migrate, dtmMigrator, server.NewUowContrib(uow, roleSeed, userSeed, p))
}

func NewSeeder(ts saas.TenantStore, seeds []seed.Contrib) seed.Seeder {
	res := seed.NewDefaultSeeder(server.NewTraceContrib(server.SeedChangeTenant(ts, seeds...)))
	return res
}

func NewAuthorizationOption(userRole *service.UserRoleContrib) *authz.Option {
	return authz.NewAuthorizationOption(userRole)
}
