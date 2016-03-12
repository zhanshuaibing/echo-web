package www

import (
	"github.com/labstack/echo"
)

func DemoHandler(c *echo.Context) error {
	c.Set("tmpl", "www/demo")
	c.Set("data", map[string]interface{}{
		"title": "Demo",
	})

	return nil
}
