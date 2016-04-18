package cache

import (
	"time"

	gc "github.com/gin-gonic/contrib/cache"
	"github.com/labstack/echo"

	"github.com/hobo-go/echo-web/conf"
)

const (
	DefaultExpiration = 3600
	DEFAULT           = time.Duration(0)
	FOREVER           = time.Duration(-1)
	DefaultKey        = "github.com/hobo-go/echo-web/modules/cache"
)

func Cache() echo.MiddlewareFunc {
	var store gc.CacheStore

	switch conf.CACHE_STORE {
	case conf.MEMCACHED:
		store = gc.NewMemcachedStore([]string{conf.MEMCACHED_SERVER}, time.Hour)
	case conf.REDIS:
		store = gc.NewRedisCache(conf.REDIS_SERVER, conf.REDIS_PWD, DefaultExpiration)
	default:
		store = gc.NewInMemoryStore(time.Hour)
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(DefaultKey, store)

			return next(c)
		}
	}
}

// shortcut to get Cache
func Default(c echo.Context) gc.CacheStore {
	// return c.MustGet(DefaultKey).(gc.CacheStore)
	return c.Get(DefaultKey).(gc.CacheStore)
}
