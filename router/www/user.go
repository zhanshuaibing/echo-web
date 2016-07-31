package www

import (
	"strconv"

	"github.com/labstack/echo"

	"github.com/hobo-go/echo-web/model"
)

func UserHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		panic(err)
	}

	var User model.User
	u := User.GetUserById(id)

	c.Set("tmpl", "www/user")
	c.Set("data", map[string]interface{}{
		"title": "User",
		"user":  u,
	})

	return nil
}
