package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	// "github.com/jinzhu/gorm"

	"github.com/hobo-go/echo-web/models"
	"github.com/hobo-go/echo-web/modules/auth"
)

func UserHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		panic(err)
	}

	u := models.GetUserById(id)

	a := auth.Default(c)
	userId := a.User.UniqueId().(uint64)

	c.JSON(http.StatusOK, map[string]interface{}{
		"title":  "User",
		"user":   u,
		"userId": userId,
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
