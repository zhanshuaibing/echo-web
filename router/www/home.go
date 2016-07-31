package www

import (
	// "net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/hobo-go/echo-web/model"
)

func HomeHandler(c echo.Context) error {
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

	// return c.Render(http.StatusOK, "www/home", map[string]interface{}{
	// 	"title": "Home",
	// 	"user":  u})

}
