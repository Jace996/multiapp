package os

import (
	"github.com/jace996/multiapp/pkg/blob"
	"github.com/jace996/vfs"
	"github.com/spf13/afero"
	"net/url"
)

func init() {
	blob.Register("memory", func(cfg *blob.Config) (vfs.Blob, error) {
		// Initialize the file system
		mm := afero.NewMemMapFs()
		public, err := url.Parse(cfg.PublicUrl)
		if err != nil {
			return nil, err
		}
		internal, err := url.Parse(cfg.InternalUrl)
		if err != nil {
			return nil, err
		}
		return vfs.NewOptLinker(mm, *public, *internal, nil), nil
	})
}
