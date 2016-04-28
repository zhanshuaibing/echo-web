package socket

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/hobo-go/echo-web/models"
	"github.com/hobo-go/echo-web/modules/auth"
	"github.com/hobo-go/echo-web/modules/cache"
	"github.com/hobo-go/echo-web/modules/render"
	"github.com/hobo-go/echo-web/modules/session"
)

func Routers() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("public"))

	// 模板
	e.SetRenderer(render.LoadTemplates())
	e.Use(render.Render())

	// 模型
	model := models.Model()
	e.Use(model)

	// Session
	e.Use(session.Session())

	// Cache
	e.Use(cache.Cache())

	// Auth
	e.Use(auth.Auth(models.GenerateAnonymousUser))

	e.Get("/", indexHandler)
	e.Get("/ws", socketHandler())

	return e
}
