package echo

import (
	// "github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	// "github.com/labstack/gommon/log"

	// "github.com/hobo-go/echo-mw/binder"
	"github.com/hobo-go/echo-mw/staticbin"

	"echo-web/assets"
	"echo-web/conf"
	"echo-web/model"
	"echo-web/module/auth"
	"echo-web/module/cache"
	"echo-web/module/render"
	"echo-web/module/session"
	"echo-web/router"
	"echo-web/router/api"
	"echo-web/router/www"
)

// 子域名部署
func RunSubdomains() {
	hosts := router.InitRoutes()

	// Server
	e := echo.New()
	e.Pre(mw.RemoveTrailingSlash())

	// CORS
	e.Use(mw.CORSWithConfig(mw.CORSConfig{
		AllowOrigins: []string{"http://" + conf.DOMAIN_WWW, "http://" + conf.DOMAIN_API},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAuthorization},
	}))

	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()

		host := hosts[req.Host]

		if host == nil {
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(res, req)
		}

		return
	})

	e.Start(":8080")

	// Graceful Shutdown
	// std := echo.New()
	// std.SetHandler(e)
	// gracehttp.Serve(std.Server)
}

func Run() {
	// Echo instance
	e := echo.New()

	// Customization
	// e.SetLogPrefix("Echo")
	// e.SetLogLevel(log.DEBUG)
	if conf.RELEASE_MODE {
		// e.SetDebug(false)
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
	// e.SetBinder(binder.New())

	// 模板
	// e.SetRenderer(render.LoadTemplates())
	e.Use(render.Render())

	// Session
	e.Use(session.Session())

	// Cache
	e.Use(cache.Cache())

	// Auth
	e.Use(auth.Auth(model.GenerateAnonymousUser))

	// Routers
	e.GET("/", www.HomeHandler)
	e.GET("/login", www.LoginHandler)
	e.GET("/register", www.RegisterHandler)
	e.GET("/logout", www.LogoutHandler)
	e.POST("/login", www.LoginPostHandler)
	e.POST("/register", www.RegisterPostHandler)

	demo := e.Group("/demo")
	{
		demo.GET("", www.DemoHandler)
	}

	user := e.Group("/user")
	user.Use(auth.LoginRequired())
	{
		user.GET("/:id", www.UserHandler)
	}

	about := e.Group("/about")
	{
		about.GET("", www.AboutHandler)
	}

	gApi := e.Group("/api")
	{
		gApi.GET("/user/:id", api.UserHandler)
		gApi.GET("/login", api.UserLoginHandler)
		gApi.GET("/register", api.UserRegisterHandler)

		gApi.GET("/post/save", api.PostSaveHandler)
		gApi.GET("/post/id/:id", api.PostHandler)
		gApi.GET("/posts/:userId/p/:p/s/:s", api.PostsHandler)
	}

	e.Start(":8080")
}
