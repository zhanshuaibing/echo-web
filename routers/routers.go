package routers

import (
	"github.com/labstack/echo"

	"github.com/hobo-go/echo-web/conf"
	"github.com/hobo-go/echo-web/routers/api"
	"github.com/hobo-go/echo-web/routers/www"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func InitRoutes() map[string]*Host {
	// Hosts
	hosts := make(map[string]*Host)

	hosts[conf.DOMAIN_WWW] = &Host{www.Routers()}
	hosts[conf.DOMAIN_API] = &Host{api.Routers()}

	return hosts
}
