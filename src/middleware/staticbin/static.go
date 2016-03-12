// Based on https://github.com/olebedev/staticbin
package staticbin

import (
	"bytes"
	"log"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/labstack/echo"
)

type Options struct {
	// SkipLogging will disable [Static] log messages when a static file is served.
	SkipLogging bool
	// IndexFile defines which file to serve as index if it exists.
	IndexFile string
	// Path prefix
	Dir string
}

func (o *Options) init() {
	if o.IndexFile == "" {
		o.IndexFile = "index.html"
	}
}

// Static returns a middleware handler that serves static files in the given directory.
func Static(asset func(string) ([]byte, error), options ...Options) echo.HandlerFunc {
	if asset == nil {
		panic("asset is nil")
	}

	opt := Options{}
	for _, o := range options {
		opt = o
		break
	}
	opt.init()

	modtime := time.Now()

	return func(c *echo.Context) error {
		if c.Request().Method != "GET" && c.Request().Method != "HEAD" {
			// Request is not correct. Go farther.
			// return echo.NewHTTPError(http.StatusMethodNotAllowed)
			return nil
		}

		url := c.Request().URL.Path
		if !strings.HasPrefix(url, opt.Dir) {
			// return echo.NewHTTPError(http.StatusUnsupportedMediaType)
			return nil
		}
		file := strings.TrimPrefix(
			strings.TrimPrefix(url, opt.Dir),
			"/",
		)
		b, err := asset(file)

		if err != nil {
			// Try to serve the index file.
			b, err = asset(path.Join(file, opt.IndexFile))

			if err != nil {
				// Go farther if the asset could not be found.
				return nil
			}
		}

		if !opt.SkipLogging {
			log.Println("[Static] Serving " + url)
		}

		// http.ServeContent(c.Writer, c.Request(), url, modtime, bytes.NewReader(b))
		// c.Abort()

		http.ServeContent(c.Response(), c.Request(), url, modtime, bytes.NewReader(b))

		return nil
	}
}
