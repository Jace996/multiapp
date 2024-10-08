package biz

import (
	kitgorm "github.com/jace996/multiapp/pkg/gorm"
	sgorm "github.com/jace996/saas/gorm"
	concurrency "github.com/jace996/gorm-concurrency/v2"
	"time"
)

// ProductSku sku
type ProductSku struct {
	kitgorm.UIDBase
	kitgorm.AuditedModel
	concurrency.HasVersion
	sgorm.MultiTenancy

	ProductId string

	Title string

	MainPic   *ProductMedia `gorm:"foreignKey:MainPicID"`
	MainPicID *string
	Medias    []ProductMedia `gorm:"polymorphic:Owner;polymorphicValue:product_sku"`

	Prices []Price `gorm:"polymorphic:Owner;polymorphicValue:product_sku"`

	Stocks []Stock `gorm:"polymorphic:Owner;polymorphicValue:product_sku"`

	Keywords []Keyword `gorm:"polymorphic:Owner;polymorphicValue:product_sku;comment:商品关键字"`

	SaleableFrom *time.Time
	SaleableTo   *time.Time
	Barcode      string `gorm:"comment:商品条码"`
}
