package www

import (
	"github.com/labstack/echo"
)

func AboutHandler(c *echo.Context) error {
	c.Set("tmpl", "www/about")
	c.Set("data", map[string]interface{}{
		"title": "About",
	})

	return nil
}
