package model

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hobo-go/echo-mw/cache"
	"github.com/hobo-go/gorm"

	"echo-web/conf"
	"echo-web/model/orm"
	"echo-web/module/log"
)

var db *gorm.DB
var dbCacheStore cache.CacheStore

func DB() *gorm.DB {
	if db == nil {
		log.DebugPrint("Model NewDB")
		newDb, err := newDB()
		if err != nil {
			panic(err)
		}
		newDb.DB().SetMaxIdleConns(10)
		newDb.DB().SetMaxOpenConns(100)
		newDb.LogMode(true)
		db = newDb
	}

	return db
}

func newDB() (*gorm.DB, error) {
	sqlConnection := conf.DB_USER_NAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", sqlConnection)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CacheStore() cache.CacheStore {
	if dbCacheStore == nil {
		switch conf.CACHE_STORE {
		case conf.MEMCACHED:
			dbCacheStore = cache.NewMemcachedStore([]string{conf.MEMCACHED_SERVER}, time.Hour)
		case conf.REDIS:
			dbCacheStore = cache.NewRedisCache(conf.REDIS_SERVER, conf.REDIS_PWD, time.Hour)
		default:
			dbCacheStore = cache.NewInMemoryStore(time.Hour)
		}
	}

	return dbCacheStore
}

func Cache(db *gorm.DB) *orm.CacheDB {
	return orm.NewCacheDB(db, CacheStore(), orm.CacheConf{
		Expire: time.Second * 10,
	})
}
