package data

import (
	"context"
	"github.com/jace996/saas/gorm"
	g "gorm.io/gorm"
)

type Repo struct {
	DbProvider gorm.DbProvider
}

func (r *Repo) GetDb(ctx context.Context) *g.DB {
	return GetDb(ctx, r.DbProvider)
}
