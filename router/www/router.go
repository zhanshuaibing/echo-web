package www

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"

	"github.com/hobo-go/echo-mw/captcha"
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

	// Context自定义
	e.Use(NewContext())

	// Customization
	if conf.RELEASE_MODE {
		e.Debug = false
	}
	e.Logger.SetPrefix("Echo")
	e.Logger.SetLevel(conf.LOG_LEVEL)

	// CSRF
	e.Use(mw.CSRFWithConfig(mw.CSRFConfig{
		ContextKey:  "_csrf",
		TokenLookup: "form:_csrf",
	}))

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// 验证码，优先于静态资源
	e.Use(captcha.Captcha(captcha.Config{
		CaptchaPath: "/captcha/",
		SkipLogging: true,
	}))

	// 静态资源
	switch conf.STATIC_TYPE {
	case conf.BINDATA:
		e.Use(staticbin.Static(assets.Asset, staticbin.Options{
			Dir:         "/",
			SkipLogging: true,
		}))
	default:
		e.Static("/assets", "./assets")
	}

	// Gzip，在验证码、静态资源之后
	// 验证码、静态资源使用http.ServeContent()，与Gzip有冲突，Nginx报错，验证码无法访问
	e.Use(mw.GzipWithConfig(mw.GzipConfig{
		Level: 5,
	}))

	// 模板
	e.Renderer = render.LoadTemplates()
	e.Use(render.Render())

	// Session
	e.Use(session.Session())

	// Cache
	e.Use(cache.Cache())

	// Auth
	e.Use(auth.New(model.GenerateAnonymousUser))

	// Routers
	e.GET("/", handler(HomeHandler))
	e.GET("/login", handler(LoginHandler))
	e.GET("/register", handler(RegisterHandler))
	e.GET("/logout", handler(LogoutHandler))
	e.POST("/login", handler(LoginPostHandler))
	e.POST("/register", handler(RegisterPostHandler))

	e.GET("/jwt/tester", handler(JWTTesterHandler))

	demo := e.Group("/demo")
	demo.Use(auth.LoginRequired())
	{
		demo.GET("", handler(DemoHandler))
	}

	user := e.Group("/user")
	user.Use(auth.LoginRequired())
	{
		user.GET("/:id", handler(UserHandler))
	}

	about := e.Group("/about")
	about.Use(auth.LoginRequired())
	{
		about.GET("", handler(AboutHandler))
	}

	return e
}

type (
	HandlerFunc func(*Context) error
)

/**
 * 自定义Context的Handler
 */
func handler(h HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*Context)
		return h(ctx)
	}
}
