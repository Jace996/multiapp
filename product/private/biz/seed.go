package biz

import (
	"context"
	"github.com/jace996/saas/seed"
)

type ProductSeeder struct {
}

var _ seed.Contrib = (*ProductSeeder)(nil)

func NewPostSeeder() *ProductSeeder {
	return &ProductSeeder{}
}

func (p *ProductSeeder) Seed(ctx context.Context, sCtx *seed.Context) error {
	return nil
}
