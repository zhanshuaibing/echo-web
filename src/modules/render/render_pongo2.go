package render

import (
	"net/http"

	"github.com/labstack/echo"

	"conf"
	"modules/log"
)

func pongo2() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				c.Error(err)
			}

			tmpl, context, err := getContext(c)
			if err == nil {
				c.Render(http.StatusOK, tmpl+conf.TMPL_SUFFIX, context)
			} else {
				log.DebugPrint("Render Error: %v", err)
			}

			return nil
		}
	}
}
