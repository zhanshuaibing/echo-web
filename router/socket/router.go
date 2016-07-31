package socket

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/hobo-go/echo-web/model"
	"github.com/hobo-go/echo-web/module/auth"
	"github.com/hobo-go/echo-web/module/cache"
	"github.com/hobo-go/echo-web/module/render"
	"github.com/hobo-go/echo-web/module/session"
)

func Routers() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("public"))

	// 模板
	e.SetRenderer(render.LoadTemplates())
	e.Use(render.Render())

	// Session
	e.Use(session.Session())

	// Cache
	e.Use(cache.Cache())

	// Auth
	e.Use(auth.Auth(model.GenerateAnonymousUser))

	e.Get("/", indexHandler)
	e.Get("/ws", socketHandler())

	return e
}
