package cache

import (
	"time"

	gin_cache "github.com/gin-gonic/contrib/cache"
	"github.com/labstack/echo"

	"conf"
)

const (
	DEFAULT    = time.Duration(0)
	FOREVER    = time.Duration(-1)
	DefaultKey = "modules/cache"
)

func Cache() echo.HandlerFunc {
	var store gin_cache.CacheStore

	switch conf.CACHE_STORE {
	case conf.MEMCACHED:
		store = gin_cache.NewMemcachedStore([]string{conf.MEMCACHED_SERVER}, time.Hour)
	default:
		store = gin_cache.NewInMemoryStore(time.Hour)
	}

	return func(c *echo.Context) error {
		c.Set(DefaultKey, store)
		// c.Next();
		return nil
	}
}

// shortcut to get Cache
func Default(c *echo.Context) gin_cache.CacheStore {
	// return c.MustGet(DefaultKey).(gin_cache.CacheStore)
	return c.Get(DefaultKey).(gin_cache.CacheStore)
}
