package main

import (
	"context"
	"flag"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport"
	dtmserver "github.com/jace996/multiapp/dtm/server"
	"github.com/jace996/multiapp/event"
	eventserver "github.com/jace996/multiapp/event/server"
	oidcdata "github.com/jace996/multiapp/oidc/data"
	kapi "github.com/jace996/multiapp/pkg/api"
	"github.com/jace996/multiapp/pkg/authn/jwt"
	"github.com/jace996/multiapp/pkg/authz/authz"
	"github.com/jace996/multiapp/pkg/authz/casbin"
	kitconf "github.com/jace996/multiapp/pkg/conf"
	kdal "github.com/jace996/multiapp/pkg/dal"
	kitdi "github.com/jace996/multiapp/pkg/di"
	kitflag "github.com/jace996/multiapp/pkg/flag"
	"github.com/jace996/multiapp/pkg/job"
	"github.com/jace996/multiapp/pkg/logging"
	kitserver "github.com/jace996/multiapp/pkg/server"
	"github.com/jace996/multiapp/pkg/tracers"
	sapi "github.com/jace996/multiapp/saas/api"
	"github.com/jace996/multiapp/user/private/biz"
	"github.com/jace996/multiapp/user/private/data"
	"github.com/jace996/multiapp/user/private/server"
	"github.com/jace996/multiapp/user/private/service"
	"github.com/jace996/saas/seed"
	"github.com/goava/di"
	"github.com/goxiaoy/vfs"
	"github.com/spf13/afero"
	"os"
	"regexp"
	"strings"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/jace996/multiapp/user/private/conf"

	//this helps to load permission definitions in following services
	_ "github.com/jace996/multiapp/realtime/api"
	_ "github.com/jace996/multiapp/saas/api"
	_ "github.com/jace996/multiapp/sys/api"

	_ "github.com/jace996/multiapp/order/api"
	_ "github.com/jace996/multiapp/payment/api"
	_ "github.com/jace996/multiapp/product/api"
)

// go build -buildvcs=false -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "user"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf kitflag.ArrayFlags
	ifSeed   bool

	id, _ = os.Hostname()
)

func init() {
	flag.Var(&flagconf, "conf", "config path, eg: -conf config.yaml")
	flag.BoolVar(&ifSeed, "seed", true, "run seeder or not")
}

func newApp(
	c *conf.UserConf,
	logger log.Logger,
	seeder seed.Seeder,
	producer event.Producer,
	_ dtmserver.Init,
	r registry.Registrar,
	srvs []transport.Server,
) *kratos.App {
	ctx := event.NewProducerContext(context.Background(), producer)
	if ifSeed {
		if err := seeder.Seed(ctx,
			seed.AddHost(),
			seed.SetExtra(biz.AdminUsernameKey, c.Admin.GetUsername()),
			seed.SetExtra(biz.AdminPasswordKey, c.Admin.GetPassword()),
		); err != nil {
			panic(err)
		}
	}

	return kratos.New(
		kratos.Context(ctx),
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Registrar(r),
		kratos.Server(
			srvs...,
		),
	)
}

func main() {
	flag.Parse()

	source := []config.Source{
		env.NewSource("KRATOS_"),
	}
	if flagconf == nil {
		flagconf = append(flagconf, "./configs")
	}
	for _, s := range flagconf {
		v := vfs.New()
		v.Mount("/", afero.NewRegexpFs(afero.NewBasePathFs(afero.NewOsFs(), strings.TrimSpace(s)), regexp.MustCompile(`\.(json|proto|xml|yaml)$`)))
		source = append(source, kitconf.NewVfs(v, "/"))
	}

	c := config.New(
		config.WithSource(
			source...,
		),
	)
	defer c.Close()
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	l, lc, err := logging.NewLogger(bc.Logging)
	if err != nil {
		panic(err)
	}
	defer lc()

	logger := log.With(l,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)

	shutdown, err := tracers.SetTracerProvider(context.Background(), bc.Tracing, Name)
	if err != nil {
		log.Error(err)
	}
	defer shutdown()
	di.SetTracer(&di.StdTracer{})
	builder, err := di.New(
		kitdi.Value(kitserver.Name(Name)),
		kitdi.Value(bc.Services),
		kitdi.Value(bc.Security),
		kitdi.Value(bc.App),
		kitdi.Value(bc.Dev),
		kitdi.Value(bc.User),
		kitdi.Value(bc.User.Idp),
		kitdi.Value(bc.Stripe),
		kitdi.Value(bc.Data),
		kitdi.Value(logger),
		kitdi.Value([]grpc.ClientOption{}),
		authz.ProviderSet, kitserver.DefaultProviderSet, jwt.ProviderSet, kapi.DefaultProviderSet, kdal.DefaultProviderSet, oidcdata.ProviderSet,
		job.DefaultProviderSet, dtmserver.DtmProviderSet, eventserver.EventProviderSet,
		sapi.GrpcProviderSet,
		casbin.PermissionProviderSet, server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet,
		kitdi.NewSet(newApp),
	)
	if err != nil {
		panic(err)
	}
	defer builder.Cleanup()
	err = builder.Invoke(func(app *kratos.App) error {
		// start and wait for stop signal
		return app.Run()
	})

	if err != nil {
		panic(err)
	}

}
