package web

import ()

func DemoHandler(c *Context) error {
	c.Set("tmpl", "web/demo")
	c.Set("data", map[string]interface{}{
		"title": "Demo",
	})

	return nil
}
