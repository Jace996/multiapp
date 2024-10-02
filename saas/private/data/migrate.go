package data

import (
	"context"
	kitgorm "github.com/jace996/multiapp/pkg/gorm"
	"github.com/jace996/multiapp/saas/private/biz"
	"github.com/jace996/saas/seed"
	"gorm.io/gorm"
)

type Migrate struct {
	data *Data
}

func NewMigrate(data *Data) *Migrate {
	return &Migrate{
		data: data,
	}
}

func (m *Migrate) Seed(ctx context.Context, sCtx *seed.Context) error {
	if len(sCtx.TenantId) > 0 {
		//host only migrate
		return nil
	}
	//make sure database exists
	ctx = kitgorm.NewDbGuardianContext(ctx)
	db := GetDb(ctx, m.data.DbProvider)
	return migrateDb(db)
}

func migrateDb(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&biz.Tenant{},
		&biz.TenantConn{},
		&biz.TenantFeature{},
		&biz.Plan{},
		&biz.PlanFeature{},
	); err != nil {
		return err
	}

	return nil
}
