package api

import (
	_ "embed"
	"github.com/jace996/multiapp/pkg/authz/authz"
)

const (
	ResourceClient = "oidc.client"
	ResourceKey    = "oidc.key"
)

//go:embed permission.yaml
var permission []byte

func init() {
	authz.LoadFromYaml(permission)
}
