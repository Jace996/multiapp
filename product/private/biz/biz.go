package biz

import (
	"github.com/jace996/multiapp/pkg/dal"
	kitdi "github.com/jace996/multiapp/pkg/di"
	"github.com/jace996/multiapp/pkg/stripe"
)

// ProviderSet is biz providers.
var ProviderSet = kitdi.NewSet(NewPostSeeder)

var (
	ProductMediaPath = "product/m"
)

type ProductManageProvider string

const (
	ProductManageProviderInternal ProductManageProvider = "internal"
	ProductManageProviderStripe   ProductManageProvider = stripe.ProviderName
)

const ConnName dal.ConnName = "product"
