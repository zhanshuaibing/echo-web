package echo

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/engine/standard"
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
	"github.com/hobo-go/echo-web/routers"
	"github.com/hobo-go/echo-web/routers/api"
	"github.com/hobo-go/echo-web/routers/www"
)

// 子域名部署
func RunSubdomains() {
	hosts := routers.InitRoutes()

	// Server
	e := echo.New()
	e.Pre(mw.RemoveTrailingSlash())

	// CORS
	e.Use(mw.CORSWithConfig(mw.CORSConfig{
		AllowOrigins: []string{"http://" + conf.DOMAIN_WWW, "http://" + conf.DOMAIN_API},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()

		host := hosts[req.Host()]

		if host == nil {
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(req, res)
		}

		return
	})

	switch conf.SERVER_HTTP {
	case conf.FASTHTTP:
		e.Run(fasthttp.New(":8080"))
	default:
		e.Run(standard.New(":8080"))
	}
}

func Run() {
	// Echo instance
	e := echo.New()

	// Customization
	e.SetLogPrefix("Echo")
	e.SetLogLevel(log.DEBUG)
	if conf.RELEASE_MODE {
		e.SetDebug(false)
	}

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
	e.Get("/", www.HomeHandler)
	e.Get("/login", www.LoginHandler)
	e.Get("/register", www.RegisterHandler)
	e.Get("/logout", www.LogoutHandler)
	e.Post("/login", www.LoginPostHandler)
	e.Post("/register", www.RegisterPostHandler)

	demo := e.Group("/demo")
	{
		demo.Get("", www.DemoHandler)
	}

	user := e.Group("/user")
	user.Use(auth.LoginRequired())
	{
		user.Get("/:id", www.UserHandler)
	}

	about := e.Group("/about")
	{
		about.Get("", www.AboutHandler)
	}

	gApi := e.Group("/api")
	{
		gApi.Get("/user/:id", api.UserHandler)
		gApi.Get("/login", api.UserLoginHandler)
		gApi.Get("/register", api.UserRegisterHandler)

		gApi.Get("/post/save", api.PostSaveHandler)
		gApi.Get("/post/id/:id", api.PostHandler)
		gApi.Get("/posts/:userId/p/:p/s/:s", api.PostsHandler)
	}

	switch conf.SERVER_HTTP {
	case conf.FASTHTTP:
		e.Run(fasthttp.New(":8080"))
	default:
		e.Run(standard.New(":8080"))
	}
}
