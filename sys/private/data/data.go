package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	oidcdata "github.com/jace996/multiapp/oidc/data"
	_ "github.com/jace996/multiapp/pkg/blob/memory"
	_ "github.com/jace996/multiapp/pkg/blob/os"
	_ "github.com/jace996/multiapp/pkg/blob/s3"
	kconf "github.com/jace996/multiapp/pkg/conf"
	kitdi "github.com/jace996/multiapp/pkg/di"
	"github.com/jace996/multiapp/sys/private/biz"
	"github.com/jace996/saas/gorm"
	g "gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = kitdi.NewSet(
	NewData,
	NewMigrate,
	NewMenuRepo,
	oidcdata.ProviderSet,
)

// Data .
type Data struct {
	DbProvider gorm.DbProvider
}

func GetDb(ctx context.Context, provider gorm.DbProvider) *g.DB {
	db := provider.Get(ctx, string(biz.ConnName))
	return db
}

// NewData .
func NewData(c *kconf.Data, dbProvider gorm.DbProvider, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		logger.Log(log.LevelInfo, log.DefaultMessageKey, "closing the data resources")
	}
	return &Data{
		DbProvider: dbProvider,
	}, cleanup, nil
}
