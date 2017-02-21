package echo

import (
	"context"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"

	"echo-web/conf"
	"echo-web/router"
)

// 子域名部署
func RunSubdomains() {
	hosts := router.InitRoutes()

	// Server
	e := echo.New()
	e.Pre(mw.RemoveTrailingSlash())
	e.Logger.SetLevel(conf.LOG_LEVEL)

	// Secure, XSS/CSS HSTS
	e.Use(mw.SecureWithConfig(mw.DefaultSecureConfig))
	mw.MethodOverride()

	// CORS
	e.Use(mw.CORSWithConfig(mw.CORSConfig{
		AllowOrigins: []string{"http://" + conf.DOMAIN_WWW, "http://" + conf.DOMAIN_API},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAuthorization},
	}))

	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()

		u, _err := url.Parse(c.Scheme() + "://" + req.Host)
		if _err != nil {
			e.Logger.Errorf("Request URL parse error:%v", _err)
		}

		host := hosts[u.Hostname()]
		if host == nil {
			e.Logger.Info("Host not found")
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(res, req)
		}

		return
	})

	if !conf.GRACEFUL {
		e.Logger.Fatal(e.Start(conf.SERVER_ADDR))
	} else {
		// Graceful Shutdown
		// Start server
		go func() {
			if err := e.Start(conf.SERVER_ADDR); err != nil {
				e.Logger.Errorf("Shutting down the server with error:%v", err)
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 10 seconds.
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}
}

/**
func Run() {
	// Echo instance
	e := echo.New()

	// Customization
	e.Logger.SetPrefix("Echo")
	e.Logger.SetLevel(log.DEBUG)
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
*/
