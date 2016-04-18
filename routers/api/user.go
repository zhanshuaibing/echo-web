package api

import (
	"net/http"
	"strconv"
	. "time"

	"github.com/labstack/echo"
	// "github.com/jinzhu/gorm"

	"github.com/hobo-go/echo-web/models"
	"github.com/hobo-go/echo-web/modules/cache"
)

func UserHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		panic(err)
	}

	model := models.Default(c)
	u := model.GetUserById(id)

	value := -1
	cacheStore := cache.Default(c)
	if id == 1 {
		value = 0
		cacheStore.Set("userId", 1, Minute)
	} else {
		cacheStore.Get("userId", &value)
	}

	request := c.Request()
	c.JSON(http.StatusOK, map[string]interface{}{
		"title":       "User",
		"user":        u,
		"value":       value,
		"Scheme":      request.Scheme(),
		"Host":        request.Host(),
		"UserAgent":   request.UserAgent(),
		"Method":      request.Method(),
		"URI":         request.URI(),
		"RemoteAddr":  request.RemoteAddress(),
		"Path":        request.URL().Path(),
		"QueryString": request.URL().QueryString(),
		"QueryParams": request.URL().QueryParams(),
		"HeaderKeys":  request.Header().Keys(),
	})

	return nil
}

func UserLoginHandler(c echo.Context) error {

	c.JSON(200, map[string]interface{}{"URI": "api user login"})

	return nil
}

func UserRegisterHandler(c echo.Context) error {

	c.JSON(200, map[string]interface{}{"URI": "api user regist"})

	return nil
}
