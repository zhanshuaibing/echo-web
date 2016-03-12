package sessions

import (
	"github.com/labstack/echo"
	echo_sessions "github.com/syntaqx/echo-middleware/session"

	"conf"
)

func Sessions() echo.MiddlewareFunc {
	switch conf.SESSION_STORE {
	case conf.REDIS:
		store, err := echo_sessions.NewRedisStore(32, "tcp", conf.REDIS_SERVER, conf.REDIS_PWD, []byte("secret"))
		if err != nil {
			panic(err)
		}
		return echo_sessions.Sessions("mysession", store)
	case "FILE":
		store := echo_sessions.NewFilesystemStore("", []byte("secret-key"))
		return echo_sessions.Sessions("mysession", store)
	default:
		store := echo_sessions.NewCookieStore([]byte("secret"))
		return echo_sessions.Sessions("mysession", store)
	}
}
