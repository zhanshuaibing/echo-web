package www

import ()

func DemoHandler(c *Context) error {
	c.Set("tmpl", "www/demo")
	c.Set("data", map[string]interface{}{
		"title": "Demo",
	})

	return nil
}
