package api

import (
	// "time"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	"github.com/hobo-go/echo-mw/binder"

	"echo-web/conf"
	"echo-web/module/cache"
	"echo-web/module/session"
)

//-----
// API Routers
//-----
func Routers() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Customization
	e.Logger.SetPrefix("Echo")
	e.Logger.SetLevel(log.DEBUG)

	if conf.RELEASE_MODE {
		// e.SetDebug(false)
	}

	// CSRF
	e.Use(mw.CSRFWithConfig(mw.CSRFConfig{
		TokenLookup: "form:X-XSRF-TOKEN",
	}))

	// Gzip
	e.Use(mw.GzipWithConfig(mw.GzipConfig{
		Level: 5,
	}))

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.Static("/favicon.ico", "./assets/img/favicon.ico")

	// Bind
	// e.SetBinder(binder.New())
	e.Binder = binder.New()

	// Session
	e.Use(session.Session())

	// Cache
	e.Use(cache.Cache())

	// e.Use(ec.SiteCache(ec.NewMemcachedStore([]string{conf.MEMCACHED_SERVER}, time.Hour), time.Minute))
	// e.GET("/user/:id", ec.CachePage(ec.NewMemcachedStore([]string{conf.MEMCACHED_SERVER}, time.Hour), time.Minute, UserHandler))

	// Routers
	e.GET("/login", UserLoginHandler)
	e.GET("/register", UserRegisterHandler)

	// JWT
	r := e.Group("")
	r.Use(mw.JWTWithConfig(mw.JWTConfig{
		SigningKey:  []byte("secret"),
		ContextKey:  "_user",
		TokenLookup: "header:" + echo.HeaderAuthorization,
	}))

	r.GET("/", ApiHandler)

	// curl http://echo.api.localhost:8080/restricted/user -H "Authorization: Bearer XXX"
	r.GET("/user", UserHandler)

	post := r.Group("/post")
	{
		post.GET("/save", PostSaveHandler)
		post.GET("/id/:id", PostHandler)
		post.GET("/:userId/p/:p/s/:s", PostsHandler)
	}

	return e
}
