package conf

import ()

const (
	// Release
	RELEASE_MODE = false

	// Project
	PROJECT_NAME = "Echo Web"

	// Domains
	DOMAIN_API = "echo.api.localhost:8080"
	DOMAIN_WWW = "echo.www.localhost:8080"

	// Http
	// FASTHTTP不可用，Session模块暂不支持
	SERVER_HTTP = STANDARD // STANDARD,FASTHTTP(STANDARD Default)

	// Session
	SESSION_STORE = REDIS // REDIS,FILE,COOKIE(COOKIE Default)

	// Cache
	CACHE_STORE = REDIS //REDIS,MEMCACHED,IN_MEMORY(IN_MEMORY Default)

	// Tmpl
	TMPL_TYPE   = PONGO2             // PONGO2,BINDATA,FILE(FILE Default)
	TMPL_DIR    = "templates/pongo2" // PONGO2(templates/pongo2),BINDATA/FILE(templates)
	TMPL_SUFFIX = ".html"            // .html,.tpl

	// Static
	STATIC_TYPE = BINDATA // BINDATA,FILE(FILE Default)

	// MySQL
	DB_NAME      = "goweb_db"
	DB_USER_NAME = "goweb_dba"
	DB_PASSWORD  = "123456"
	DB_HOST      = "127.0.0.1"
	DB_PORT      = "3306"

	// Redis
	REDIS_SERVER = "127.0.0.1:6379"
	REDIS_PWD    = "123456"

	// Memcached
	MEMCACHED_SERVER = "localhost:11211"
)

const (
	// Http
	STANDARD = "standard"
	FASTHTTP = "fasthttp" // Fasthttp

	// Pongo2
	PONGO2 = "pongo2"

	// Bindata
	BINDATA = "bindata"

	// File
	FILE = "file"

	// Redis
	REDIS = "redis"

	// Memcached
	MEMCACHED = "memcached"

	// Cookie
	COOKIE = "cookie"

	// In Memory
	IN_MEMORY = "InMemory"
)
