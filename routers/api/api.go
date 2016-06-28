package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	// "github.com/jinzhu/gorm"

	"github.com/hobo-go/echo-mw/session"
	"github.com/hobo-go/echo-web/modules/cache"
)

func ApiHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	// 缓存测试
	value := -1
	if err == nil {
		cacheStore := cache.Default(c)
		if id == 1 {
			value = 0
			cacheStore.Set("userId", 1, time.Minute)
		} else {
			cacheStore.Get("userId", &value)
		}
	}

	// Flash测试
	s := session.Default(c)
	s.AddFlash("0")
	s.AddFlash("1", "key")

	request := c.Request()
	c.JSON(http.StatusOK, map[string]interface{}{
		"title":        "Api Index",
		"CacheValue":   value,
		"Scheme":       request.Scheme(),
		"Host":         request.Host(),
		"UserAgent":    request.UserAgent(),
		"Method":       request.Method(),
		"URI":          request.URI(),
		"RemoteAddr":   request.RemoteAddress(),
		"Path":         request.URL().Path(),
		"QueryString":  request.URL().QueryString(),
		"QueryParams":  request.URL().QueryParams(),
		"HeaderKeys":   request.Header().Keys(),
		"FlashDefault": s.Flashes(),
		"Flash":        s.Flashes("key"),
	})

	return nil
}
