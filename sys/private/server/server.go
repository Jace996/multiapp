package server

import (
	"github.com/go-kratos/kratos/v2/transport"
	dtmdata "github.com/jace996/multiapp/dtm/data"
	"github.com/jace996/multiapp/pkg/api"
	"github.com/jace996/multiapp/pkg/authz/authz"
	kitdi "github.com/jace996/multiapp/pkg/di"
	"github.com/jace996/multiapp/pkg/server"
	api2 "github.com/jace996/multiapp/sys/api"
	"github.com/jace996/multiapp/sys/private/biz"
	"github.com/jace996/multiapp/sys/private/data"
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
	NewSeeder,
	kitdi.Value(ClientName),
	kitdi.Value(biz.ConnName),
	NewSeeding,
	NewAuthorizationOption,
)

var ClientName api.ClientName = api2.ServiceName

// NewSeeding sys seeding should migrate dtmsrv and dmtcli
//
// wrap all service migrator into one seed.Contrib, which grants the running sequence of those contribs
func NewSeeding(apisixSeeder *biz.ApisixSeed, uow uow.Manager, dtmMigrator *dtmdata.Migrator, migrate *data.Migrate, menu *biz.MenuSeed) seed.Contrib {
	return seed.Chain(apisixSeeder, dtmMigrator, migrate, server.NewUowContrib(uow, seed.Chain(menu)))
}

func NewSeeder(ts saas.TenantStore, seeds []seed.Contrib) seed.Seeder {
	res := seed.NewDefaultSeeder(server.NewTraceContrib(server.SeedChangeTenant(ts, seeds...)))
	return res
}

func NewAuthorizationOption() *authz.Option {
	return authz.NewAuthorizationOption()
}
