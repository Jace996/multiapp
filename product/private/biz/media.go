package biz

import (
	"github.com/jace996/multiapp/pkg/data"
	"github.com/jace996/multiapp/pkg/sortable"
)

type ProductMedia struct {
	ID string `gorm:"primaryKey;size:128"`

	OwnerID string
	// OwnerType product/product_sku
	OwnerType string

	Type      string
	MimeType  string
	Usage     string
	Name      string
	Reference string
	sortable.Embed
}

func NewProductMedia() *ProductMedia {
	return &ProductMedia{}
}

type ProductMediaRepo interface {
	data.Repo[ProductMedia, string, interface{}]
}
