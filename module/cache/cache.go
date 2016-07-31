package cache

import (
	"time"

	ec "github.com/hobo-go/echo-mw/cache"
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
	var store ec.CacheStore

	switch conf.CACHE_STORE {
	case conf.MEMCACHED:
		store = ec.NewMemcachedStore([]string{conf.MEMCACHED_SERVER}, time.Hour)
	case conf.REDIS:
		store = ec.NewRedisCache(conf.REDIS_SERVER, conf.REDIS_PWD, DefaultExpiration)
	default:
		store = ec.NewInMemoryStore(time.Hour)
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(DefaultKey, store)

			return next(c)
		}
	}
}

// shortcut to get Cache
func Default(c echo.Context) ec.CacheStore {
	// return c.MustGet(DefaultKey).(ec.CacheStore)
	return c.Get(DefaultKey).(ec.CacheStore)
}
