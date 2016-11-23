package socket

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"echo-web/model"
	"echo-web/module/auth"
	"echo-web/module/cache"
	"echo-web/module/render"
	"echo-web/module/session"
)

func Routers() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 模板
	e.Renderer = render.LoadTemplates()
	e.Use(render.Render())

	// Session
	e.Use(session.Session())

	// Cache
	e.Use(cache.Cache())

	// Auth
	e.Use(auth.New(model.GenerateAnonymousUser))

	e.GET("/", indexHandler)
	e.GET("/ws", socketHandler())

	return e
}
