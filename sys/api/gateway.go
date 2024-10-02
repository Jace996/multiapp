package api

import (
	_ "embed"
	"github.com/jace996/multiapp/pkg/apisix"
)

//go:embed gateway.yaml
var gateway []byte

func init() {
	apisix.LoadFromYaml(gateway)
}
