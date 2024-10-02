package biz

import (
	"github.com/jace996/multiapp/pkg/gorm"
	"github.com/jace996/multiapp/pkg/multilingual"
	"github.com/samber/lo"
)

type Brand struct {
	gorm.UIDBase
	Code string
	Name string
	Logo string
	Url  string
	Desc string

	Trans []*BrandTrans

	OwnedTenantId string
}

type BrandTrans struct {
	gorm.UIDBase
	multilingual.Embed

	BrandId string

	Name        string
	Url         string
	Description string
}

var _ multilingual.Multilingual = (*Brand)(nil)

func (b *Brand) GetTranslations() []interface{} {
	return lo.Map(b.Trans, func(item *BrandTrans, _ int) interface{} {
		return item
	})
}
