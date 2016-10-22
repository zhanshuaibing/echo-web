package www

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	"github.com/hobo-go/echo-mw/binder"
	"github.com/hobo-go/echo-mw/staticbin"

	"echo-web/assets"
	"echo-web/conf"
	"echo-web/model"
	"echo-web/module/auth"
	"echo-web/module/cache"
	"echo-web/module/render"
	"echo-web/module/session"
)

//---------
// Website Routers
//---------
func Routers() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Customization
	if conf.RELEASE_MODE {
		// e.SetDebug(false)
	}
	e.Logger.SetPrefix("Echo")
	e.Logger.SetLevel(log.DEBUG)

	// CSRF
	e.Use(mw.CSRFWithConfig(mw.CSRFConfig{
		ContextKey:  "_csrf",
		TokenLookup: "form:_csrf",
	}))

	// Gzip
	e.Use(mw.GzipWithConfig(mw.GzipConfig{
		Level: 5,
	}))

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// 静态资源
	switch conf.STATIC_TYPE {
	case conf.BINDATA:
		e.Use(staticbin.Static(assets.Asset, staticbin.Options{
			Dir: "/",
		}))
	default:
		e.Static("/assets", "./assets")
	}

	// Binder
	e.Binder = binder.New()

	// 模板
	e.Renderer = render.LoadTemplates()
	e.Use(render.Render())

	// Session
	e.Use(session.Session())

	// Cache
	e.Use(cache.Cache())

	// Auth
	e.Use(auth.Auth(model.GenerateAnonymousUser))

	// Routers
	e.GET("/", HomeHandler)
	e.GET("/login", LoginHandler)
	e.GET("/register", RegisterHandler)
	e.GET("/logout", LogoutHandler)
	e.POST("/login", LoginPostHandler)
	e.POST("/register", RegisterPostHandler)

	e.GET("/jwt/tester", JWTTesterHandler)

	demo := e.Group("/demo")
	demo.Use(auth.LoginRequired())
	{
		demo.GET("", DemoHandler)
	}

	user := e.Group("/user")
	user.Use(auth.LoginRequired())
	{
		user.GET("/:id", UserHandler)
	}

	about := e.Group("/about")
	about.Use(auth.LoginRequired())
	{
		about.GET("", AboutHandler)
	}

	return e
}
