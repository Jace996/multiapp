package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	kconf "github.com/jace996/multiapp/pkg/conf"
	"github.com/jace996/multiapp/pkg/dal"
	kitdi "github.com/jace996/multiapp/pkg/di"
	"github.com/jace996/multiapp/saas/api"
	"github.com/jace996/multiapp/saas/private/biz"
	g "gorm.io/gorm"

	_ "github.com/jace996/multiapp/pkg/blob/memory"
	_ "github.com/jace996/multiapp/pkg/blob/os"
	_ "github.com/jace996/multiapp/pkg/blob/s3"
)

// ProviderSet is data providers.
var ProviderSet = kitdi.NewSet(
	NewData,
	NewTenantRepo,
	NewMigrate,
	NewPlanRepo,
	api.NewTenantStore,
)

// Data .
type Data struct {
	DbProvider dal.ConstDbProvider
}

func GetDb(ctx context.Context, provider dal.ConstDbProvider) *g.DB {
	db := provider.Get(ctx, string(biz.ConnName))
	return db
}

// NewData .
func NewData(c *kconf.Data, dbProvider dal.ConstDbProvider, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		logger.Log(log.LevelInfo, log.DefaultMessageKey, "closing the data resources")
	}
	return &Data{
		DbProvider: dbProvider,
	}, cleanup, nil
}
