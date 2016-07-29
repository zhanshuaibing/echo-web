package www

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	"github.com/hobo-go/echo-mw/binder"
	"github.com/hobo-go/echo-mw/staticbin"

	"github.com/hobo-go/echo-web/assets"
	"github.com/hobo-go/echo-web/conf"
	"github.com/hobo-go/echo-web/models"
	"github.com/hobo-go/echo-web/modules/auth"
	"github.com/hobo-go/echo-web/modules/cache"
	"github.com/hobo-go/echo-web/modules/render"
	"github.com/hobo-go/echo-web/modules/session"
)

//---------
// Website Routers
//---------
func Routers() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Customization
	if conf.RELEASE_MODE {
		e.SetDebug(false)
	}
	// e.SetLogPrefix("Echo")
	e.SetLogLevel(log.DEBUG)

	// CORS
	e.Use(mw.CORSWithConfig(mw.CORSConfig{
		AllowOrigins: []string{"http://echo.www.localhost:8080", "http://echo.api.localhost:8080"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

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

	// Bind
	e.SetBinder(binder.New())

	// 模板
	e.SetRenderer(render.LoadTemplates())
	e.Use(render.Render())

	// Session
	e.Use(session.Session())

	// Cache
	e.Use(cache.Cache())

	// Auth
	e.Use(auth.Auth(models.GenerateAnonymousUser))

	// Routers
	e.Get("/", HomeHandler)
	e.Get("/login", LoginHandler)
	e.Get("/register", RegisterHandler)
	e.Get("/logout", LogoutHandler)
	e.Post("/login", LoginPostHandler)
	e.Post("/register", RegisterPostHandler)

	demo := e.Group("/demo")
	demo.Use(auth.LoginRequired())
	{
		demo.Get("", DemoHandler)
	}

	user := e.Group("/user")
	user.Use(auth.LoginRequired())
	{
		user.Get("/:id", UserHandler)
	}

	about := e.Group("/about")
	about.Use(auth.LoginRequired())
	{
		about.Get("", AboutHandler)
	}

	return e
}
