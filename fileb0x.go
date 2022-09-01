package static

import (
	staticFiles "github.com/go-swagno/swagno-files/static"
	"golang.org/x/net/webdav"
)

func NewHandler() *webdav.Handler {
	return &webdav.Handler{
		FileSystem: staticFiles.FS,
		LockSystem: webdav.NewMemLS(),
	}
}
