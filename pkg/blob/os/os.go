package os

import (
	"github.com/jace996/multiapp/pkg/blob"
	"github.com/jace996/vfs"
	"github.com/spf13/afero"
	"net/url"
)

func init() {
	blob.Register("os", func(cfg *blob.Config) (vfs.Blob, error) {
		// Initialize the file system
		appfs := afero.NewOsFs()
		if cfg.Os != nil && cfg.Os.Dir != nil {
			appfs = afero.NewBasePathFs(appfs, cfg.Os.Dir.Value)
		}
		public, err := url.Parse(cfg.PublicUrl)
		if err != nil {
			return nil, err
		}
		internal, err := url.Parse(cfg.InternalUrl)
		if err != nil {
			return nil, err
		}
		return vfs.NewOptLinker(appfs, *public, *internal, nil), nil
	})
}
