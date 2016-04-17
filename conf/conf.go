package conf

import ()

const (
	RELEASE_MODE = false

	// Http
	SERVER_HTTP = "STANDARD" // Default "STANDARD",STANDARD,FASTHTTP

	// Session
	SESSION_STORE = "COOKIE" // Default "COOKIE",COOKIE,REDIS,FILE

	// Cache
	CACHE_STORE = "REDIS"

	// Tmpl
	TMPL_TYPE   = "PONGO2" // Default "",PONGO2,BINDATA
	TMPL_DIR    = "templates/pongo2"
	TMPL_SUFFIX = ".html" // .html,.tpl

	// Static
	STATIC_TYPE = "BINDATA" // Default "",BINDATA

	// MySQL
	DB_NAME      = "goweb_db"
	DB_USER_NAME = "goweb_dba"
	DB_PASSWORD  = "123456"
	DB_HOST      = "127.0.0.1"
	DB_PORT      = "3306"

	// Redis
	REDIS        = "REDIS"
	REDIS_SERVER = "127.0.0.1:6379"
	REDIS_PWD    = "123456"

	// Memcached
	MEMCACHED        = "MEMCACHED"
	MEMCACHED_SERVER = "localhost:11211"

	// Bindata
	BINDATA = "BINDATA"

	// Pongo2
	PONGO2 = "PONGO2"

	// Fasthttp
	FASTHTTP = "FASTHTTP"
)
