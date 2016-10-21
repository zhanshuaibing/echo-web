package router

import (
	"github.com/labstack/echo"

	"echo-web/conf"
	"echo-web/router/api"
	"echo-web/router/socket"
	"echo-web/router/www"
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
	hosts[conf.DOMAIN_SOCKET] = &Host{socket.Routers()}

	return hosts
}
