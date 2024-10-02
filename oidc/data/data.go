package data

import (
	kconf "github.com/jace996/multiapp/pkg/conf"
	kitdi "github.com/jace996/multiapp/pkg/di"
	client "github.com/ory/hydra-client-go/v2"
)

var ProviderSet = kitdi.NewSet(
	NewHydra,
)

func NewHydra(c *kconf.Security) *client.APIClient {
	cfg := client.NewConfiguration()
	cfg.Servers = client.ServerConfigurations{
		{
			URL: c.Oidc.Hydra.AdminUrl,
		},
	}
	return client.NewAPIClient(cfg)
}
