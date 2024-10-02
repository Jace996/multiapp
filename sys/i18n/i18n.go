package i18n

import (
	"embed"
	_ "github.com/jace996/multiapp/oidc/i18n"
	"github.com/jace996/multiapp/pkg/localize"
)

var (
	//go:embed  embed
	f embed.FS
)

func init() {
	localize.RegisterFileBundle(localize.FileBundle{Fs: f})
}
