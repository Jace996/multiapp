package menu

import (
	_ "embed"
	"github.com/jace996/multiapp/sys/menu"
)

//go:embed menu.yaml
var menuData []byte

func init() {
	menu.LoadFromYaml(menuData)
}
