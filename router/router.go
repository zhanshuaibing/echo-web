package router

import (
	"github.com/labstack/echo"

	. "echo-web/conf"
	"echo-web/router/api"
	"echo-web/router/socket"
	"echo-web/router/web"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func InitRoutes() map[string]*Host {
	// Hosts
	hosts := make(map[string]*Host)

	hosts[Conf.Server.DomainWeb] = &Host{web.Routers()}
	hosts[Conf.Server.DomainApi] = &Host{api.Routers()}
	hosts[Conf.Server.DomainSocket] = &Host{socket.Routers()}

	return hosts
}
