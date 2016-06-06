package api

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	"github.com/hobo-go/echo-mw/binder"

	"github.com/hobo-go/echo-web/conf"
	"github.com/hobo-go/echo-web/models"
	"github.com/hobo-go/echo-web/modules/auth"
	"github.com/hobo-go/echo-web/modules/cache"
	"github.com/hobo-go/echo-web/modules/session"
)

//-----
// API Routers
//-----
func Routers() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Customization
	// e.SetLogPrefix("Echo")
	e.SetLogLevel(log.DEBUG)
	if conf.RELEASE_MODE {
		e.SetDebug(false)
	}

	// Gzip
	e.Use(mw.GzipWithConfig(mw.GzipConfig{
		Level: 5,
	}))

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// Bind
	e.SetBinder(binder.New())

	// 模型
	model := models.Model()
	e.Use(model)

	// Session
	e.Use(session.Session())

	// Cache
	e.Use(cache.Cache())

	// Auth
	e.Use(auth.Auth(models.GenerateAnonymousUser))

	// Routers
	e.Get("/", ApiHandler)
	e.Get("/:id", ApiHandler)
	e.Get("/user/:id", UserHandler)
	e.Get("/login", UserLoginHandler)
	e.Get("/register", UserRegisterHandler)
	post := e.Group("/post")
	{
		post.Get("/save", PostSaveHandler)
		post.Get("/id/:id", PostHandler)
		post.Get("/:userId/p/:p/s/:s", PostsHandler)
	}

	return e
}
