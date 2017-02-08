package captcha

import (
	"strings"

	"github.com/dchest/captcha"
	"github.com/labstack/echo"
)

type Options struct {
	SkipLogging bool
	CaptchaPath string
}

func (o *Options) init() {
	if o.CaptchaPath == "" {
		o.CaptchaPath = "/captcha/"
	}
}

func Captcha(options ...Options) echo.MiddlewareFunc {
	opt := Options{}
	for _, o := range options {
		opt = o
		break
	}
	opt.init()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			request := c.Request()
			method := request.Method
			if method != "GET" {
				return next(c)
			}

			url := request.URL.Path
			if strings.HasPrefix(url, opt.CaptchaPath) {
				if !opt.SkipLogging {
					c.Logger().Debugf("Captcha server, url:%v", url)
				}
				captcha.Server(captcha.StdWidth, captcha.StdHeight).ServeHTTP(c.Response(), c.Request())
				return nil
			} else {
				return next(c)
			}

		}
	}
}
