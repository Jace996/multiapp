package server

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/jace996/multiapp/pkg/conf"
	kitdi "github.com/jace996/multiapp/pkg/di"
	kregistry "github.com/jace996/multiapp/pkg/registry"
	"github.com/jace996/multiapp/pkg/server/http"
	"github.com/defval/di"
)

var DefaultProviderSet = kitdi.NewSet(
	kitdi.Value(http.ReqDecode),
	kitdi.Value(http.ResEncoder),
	kitdi.Value(http.ErrEncoder),
	NewRegistrar,
	NewWebMultiTenancyOption,
)

// Name server name.
type Name string

func NewRegistrar(services *conf.Services, container *di.Container) (registry.Registrar, error) {
	err := container.Provide(func() *kregistry.Config { return services.Registry })
	if err != nil {
		return nil, err
	}
	r, _, err := kregistry.NewRegister(services.Registry, container)
	return r, err
}
