package biz

import (
	"github.com/jace996/multiapp/event"
	"github.com/jace996/multiapp/pkg/dal"
	kitdi "github.com/jace996/multiapp/pkg/di"
	"github.com/defval/di"
)

// ProviderSet is biz providers.
var ProviderSet = kitdi.NewSet(
	NewTenantUserCase,
	kitdi.NewProvider(NewTenantReadyEventHandler, di.As(new(event.ConsumerHandler))),
	kitdi.NewProvider(NewSubscriptionChangedEventHandler, di.As(new(event.ConsumerHandler))),
	kitdi.NewProvider(NewOrderChangedEventHandler, di.As(new(event.ConsumerHandler))),
	NewConfigConnStrGenerator,
)

const ConnName dal.ConnName = "saas"

const TenantLogoPath = "saas/tenant/logo"
