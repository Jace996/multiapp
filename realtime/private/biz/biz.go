package biz

import (
	"github.com/jace996/multiapp/pkg/dal"
	kitdi "github.com/jace996/multiapp/pkg/di"
)

// ProviderSet is biz providers.
var ProviderSet = kitdi.NewSet()

const ConnName dal.ConnName = "realtime"
