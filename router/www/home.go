package www

import (
	"strconv"

	"echo-web/model"
)

func HomeHandler(c *Context) error {
	id, err := strconv.ParseUint("1", 10, 64)
	if err != nil {
		panic(err)
	}

	var User model.User
	u := User.GetUserById(id)

	c.Set("tmpl", "www/home")
	c.Set("data", map[string]interface{}{
		"title": "Home",
		"user":  u,
	})

	return nil
}
