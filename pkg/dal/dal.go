package dal

import (
	"github.com/eko/gocache/v3/cache"
	"github.com/eko/gocache/v3/store"
	"github.com/go-redis/redis/v8"
	"github.com/jace996/multiapp/event"
	"github.com/jace996/multiapp/event/trace"
	"github.com/jace996/multiapp/pkg/blob"
	kitconf "github.com/jace996/multiapp/pkg/conf"
	kitdi "github.com/jace996/multiapp/pkg/di"
	"github.com/jace996/multiapp/pkg/email"
	kitgorm "github.com/jace996/multiapp/pkg/gorm"
	kitredis "github.com/jace996/multiapp/pkg/redis"
	kituow "github.com/jace996/multiapp/pkg/uow"
	"github.com/jace996/saas"
	"github.com/jace996/saas/data"
	sgorm "github.com/jace996/saas/gorm"
	"github.com/goava/di"
	"github.com/goxiaoy/go-eventbus"
	"github.com/goxiaoy/vfs"

	_ "github.com/jace996/multiapp/event/kafka"
	_ "github.com/jace996/multiapp/event/pulsar"

	_ "github.com/jace996/multiapp/pkg/email/log"
	_ "github.com/jace996/multiapp/pkg/email/smtp"
	
	_ "github.com/jace996/multiapp/pkg/registry/etcd"
)

type (
	ConnName        string
	ConstDbProvider sgorm.DbProvider
)

var (
	//DefaultProviderSet shared provider for all data layer
	DefaultProviderSet = kitdi.NewSet(
		NewConnStrResolver,
		NewConstantConnStrResolver,
		kitgorm.NewSqlDbCache,
		kitgorm.NewDbCache,

		kitgorm.NewDbProvider,
		NewConstDbProvider,

		kituow.NewUowManager,

		NewBlob,

		NewRedisUniversalOption,
		NewRedis,
		NewCacheStore,
		kitdi.NewProvider(NewStringCacheManager, di.As(new(cache.CacheInterface[string]))),

		NewEmailer,

		NewEventProducer,
		kitdi.Value(eventbus.Default),
	)
)

func NewConnStrResolver(c *kitconf.Data, ts saas.TenantStore) data.ConnStrResolver {
	return kitgorm.NewConnStrResolver(c.Endpoints, ts)
}

// NewConstantConnStrResolver ignore multi-tenancy
func NewConstantConnStrResolver(c *kitconf.Data) data.ConnStrings {
	// ignore multi-tenancy
	conn := make(data.ConnStrings, 1)
	for k, v := range c.Endpoints.Databases {
		conn[k] = v.Source
	}
	return conn
}

// NewConstDbProvider ignore multi-tenancy
func NewConstDbProvider(cache *kitgorm.DbCache, cs data.ConnStrings, d *kitconf.Data) ConstDbProvider {
	return kitgorm.NewDbProvider(cache, cs, d)
}

func NewBlob(c *kitconf.Data) (vfs.Blob, error) {
	return blob.New(c.Vfs...)
}

func NewEmailer(cfg *kitconf.Data, container *di.Container) (email.Client, error) {
	if cfg == nil || cfg.Endpoints == nil || cfg.Endpoints.Email == nil || cfg.Endpoints.Email == nil {
		return email.NewClient(nil, container)
	}
	return email.NewClient(cfg.Endpoints.Email, container)
}

func NewRedisUniversalOption(c *kitconf.Data, name ConnName) (*redis.UniversalOptions, error) {
	if c.Endpoints.Redis == nil {
		panic("redis endpoints required")
	}
	return kitredis.ResolveRedisEndpointOrDefault(c.Endpoints.Redis, string(name))
}

func NewRedis(opt *redis.UniversalOptions) redis.UniversalClient {
	return kitredis.NewRedisClient(opt)
}

func NewCacheStore(client redis.UniversalClient) store.StoreInterface {
	return store.NewRedis(client)
}

func NewStringCacheManager(store store.StoreInterface) *cache.Cache[string] {
	return cache.New[string](store)
}

func NewEventProducer(c *kitconf.Data, name ConnName, container *di.Container) (event.Producer, func(), error) {
	e := c.Endpoints.GetEventMergedDefault(string(name))
	ret, err := event.NewFactoryProducer(e, container)
	if err != nil {
		return nil, func() {}, err
	}
	ret.Use(trace.Send())
	return ret, func() {
		ret.Close()
	}, nil

}
