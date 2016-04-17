package session

import (
	"github.com/labstack/echo"

	es "github.com/hobo-go/echo-mw/session" //"github.com/syntaqx/echo-middleware/session"
	"github.com/hobo-go/echo-web/conf"
)

func Session() echo.MiddlewareFunc {
	switch conf.SESSION_STORE {
	case conf.REDIS:
		store, err := es.NewRedisStore(32, "tcp", conf.REDIS_SERVER, conf.REDIS_PWD, []byte("secret"))
		if err != nil {
			panic(err)
		}
		return es.New("mysession", store)
	case "FILE":
		store := es.NewFilesystemStore("", []byte("secret-key"))
		return es.New("mysession", store)
	default:
		store := es.NewCookieStore([]byte("secret"))
		return es.New("mysession", store)
	}
}
