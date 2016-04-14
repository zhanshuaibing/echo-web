package www

import (
	// "net/http"
	"strconv"

	"github.com/labstack/echo"

	"models"
)

func HomeHandler(c echo.Context) error {
	id, err := strconv.ParseUint("1", 10, 64)
	if err != nil {
		panic(err)
	}
	model := models.Default(c)
	u := model.GetUserById(id)

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
