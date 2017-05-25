package web

import (
	"echo-web/module/log"

)

func HomeHandler(c *Context) error {
	log.Debugf("Wrong captcha solution: %v! No robots allowed!")

	c.Set("tmpl", "web/home")
	c.Set("data", map[string]interface{}{
		"title": "Home",
	})

	return nil
}
