package api

import (
	"net/http"
	"strconv"
	. "time"

	"github.com/labstack/echo"
	// "github.com/jinzhu/gorm"

	"github.com/hobo-go/echo-web/models"
)

func PostHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		panic(err)
	}

	model := models.Default(c)
	post := model.GetPostById(id)

	c.JSON(http.StatusOK, map[string]interface{}{
		"title": "Post",
		"post":  post,
	})

	return nil
}

func PostsHandler(c echo.Context) error {

	userIdStr := c.Param("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		panic(err)
	}

	page, err := strconv.Atoi(c.Param("p"))
	if err != nil {
		panic(err)
	}
	size, err := strconv.Atoi(c.Param("s"))
	if err != nil {
		panic(err)
	}

	model := models.Default(c)
	posts := model.GetUserPostsByUserId(userId, page, size)

	c.JSON(http.StatusOK, map[string]interface{}{
		"title": "Post",
		"posts": posts,
	})

	return nil
}

func PostSaveHandler(c echo.Context) error {
	model := models.Default(c)
	model.PostSave()

	c.JSON(http.StatusOK, map[string]interface{}{
		"title": "Post Add",
	})

	return nil
}
