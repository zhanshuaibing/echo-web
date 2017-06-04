package api

import (
	"github.com/labstack/echo"
	"github.com/opentracing/opentracing-go"

	ot "echo-web/middleware/opentracing"
)

func NewContext() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &Context{c}
			return h(ctx)
		}
	}
}

type Context struct {
	echo.Context
}

/**
 * 以接口参数或后缀返回数据
 * JSONP、JSON/XML
 */
func (c *Context) AutoFMT(code int, i interface{}) (err error) {
	// JSONP
	callback := c.QueryParam("jsonp")
	if len(callback) > 0 {
		c.Logger().Infof("JSONP callback func:%v", callback)
		return c.JSONP(code, callback, i)
	} else {
		return c.JSON(code, i)
	}
}

func (ctx *Context) OpenTracingSpan() opentracing.Span {
	return ot.Default(ctx)
}
