package api

import (
	"github.com/labstack/echo"
)

type Ctx struct {
	echo.Context
}

func (c Ctx) AutoFMT(code int, i interface{}) (err error) {
	callback := c.QueryParam("jsonp")
	c.Logger().Debugf("callback %v", callback)
	if len(callback) > 0 {
		return c.JSONP(code, callback, i)
	} else {
		return c.JSON(code, i)
	}
}
