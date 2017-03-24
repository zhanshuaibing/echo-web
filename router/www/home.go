package www

import (

)

func HomeHandler(c *Context) error {
	c.Set("tmpl", "www/home")
	c.Set("data", map[string]interface{}{
		"title": "Home",
	})

	return nil
}
