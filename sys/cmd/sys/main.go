package main

import (
	"context"
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport"
	dtmserver "github.com/jace996/multiapp/dtm/server"
	"github.com/jace996/multiapp/event"
	kapi "github.com/jace996/multiapp/pkg/api"
	"github.com/jace996/multiapp/pkg/apisix"
	"github.com/jace996/multiapp/pkg/authn/jwt"
	"github.com/jace996/multiapp/pkg/authz/authz"
	conf2 "github.com/jace996/multiapp/pkg/conf"
	kdal "github.com/jace996/multiapp/pkg/dal"
	kitdi "github.com/jace996/multiapp/pkg/di"
	kitflag "github.com/jace996/multiapp/pkg/flag"
	"github.com/jace996/multiapp/pkg/job"
	"github.com/jace996/multiapp/pkg/logging"
	kitserver "github.com/jace996/multiapp/pkg/server"
	"github.com/jace996/multiapp/pkg/tracers"
	sapi "github.com/jace996/multiapp/saas/api"
	"github.com/jace996/multiapp/sys/private/biz"
	"github.com/jace996/multiapp/sys/private/data"
	"github.com/jace996/multiapp/sys/private/server"
	"github.com/jace996/multiapp/sys/private/service"
	uapi "github.com/jace996/multiapp/user/api"
	"github.com/jace996/saas/seed"
	"github.com/defval/di"
	"github.com/jace996/vfs"
	"github.com/spf13/afero"
	"os"
	"regexp"
	"strings"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/jace996/multiapp/sys/private/conf"

	//This helps to load i18n resources and gateway config in following services
	_ "github.com/jace996/multiapp/realtime/api"
	_ "github.com/jace996/multiapp/saas/api"
	_ "github.com/jace996/multiapp/sys/api"
	_ "github.com/jace996/multiapp/user/api"

	_ "github.com/jace996/multiapp/order/api"
	_ "github.com/jace996/multiapp/payment/api"
	_ "github.com/jace996/multiapp/product/api"

	//This helps to load menu in following services
	_ "github.com/jace996/multiapp/payment/menu"
	_ "github.com/jace996/multiapp/product/menu"
)

// go build -buildvcs=false -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "sys"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf kitflag.ArrayFlags

	ifSyncApisix bool

	ifSeed bool
	id, _  = os.Hostname()
)

func init() {
	flag.Var(&flagconf, "conf", "config path, eg: -conf config.yaml")
	flag.BoolVar(&ifSyncApisix, "apisix.sync", true, "sync with apisix upstreams")
	flag.BoolVar(&ifSeed, "seed", true, "run seeder or not")
}

func newApp(
	logger log.Logger,
	srvs []transport.Server,
	seeder seed.Seeder,
	producer event.Producer,
	r registry.Registrar,
	syncAdmin *apisix.WatchSyncAdmin,
) *kratos.App {
	ctx := event.NewProducerContext(context.Background(), producer)
	if ifSeed {
		if err := seeder.Seed(ctx, seed.AddHost()); err != nil {
			panic(err)
		}
	}

	if ifSyncApisix {
		srvs = append(srvs, syncAdmin)
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
		source = append(source, conf2.NewVfs(v, "/"))
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
		kitdi.Value(bc.Sys),
		kitdi.Value(bc.Data),
		kitdi.Value(bc.App),
		kitdi.Value(bc.Dev),

		kitdi.Value(logger),
		kitdi.Value([]grpc.ClientOption{}),
		authz.ProviderSet, jwt.ProviderSet, kitserver.DefaultProviderSet, kapi.DefaultProviderSet, kdal.DefaultProviderSet,
		job.DefaultProviderSet, dtmserver.DtmProviderSet,
		uapi.GrpcProviderSet,
		sapi.GrpcProviderSet,
		server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet,
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
