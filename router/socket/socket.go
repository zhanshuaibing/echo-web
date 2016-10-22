package socket

import (
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

func indexHandler(c echo.Context) error {
	c.Set("tmpl", "socket/index")
	c.Set("data", map[string]interface{}{
		"title": "Index",
	})
	return nil
}

func socketHandler() echo.HandlerFunc {
	return echo.WrapHandler(websocket.Handler(func(ws *websocket.Conn) {
		for {
			websocket.Message.Send(ws, "Hello, Client!")
			msg := ""
			websocket.Message.Receive(ws, &msg)
			println(msg)
		}
	}))
}
