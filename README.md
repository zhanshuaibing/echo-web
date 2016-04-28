# Golang Web

##测试
```
./run.sh
```

##打包
```
./build.sh
```

##依赖管理Glide

```
https://github.com/Masterminds/glide

$ glide create                            	# Start a new workspace
$ open glide.yaml                         	# and edit away!
$ glide get github.com/hobo-go/echo-md 		# Get a package and add to glide.yaml
$ glide install                           	# Install packages and dependencies

$ go build                                	# Go tools work normally
$ glide up                                	# Update to newest versions of the package
```
###glide get
```
github.com/labstack/echo
github.com/go-sql-driver/mysql
github.com/jinzhu/gorm
	github.com/jinzhu/inflection
github.com/labstack/gommon
github.com/hobo-go/echo-mw
	github.com/flosch/pongo2
	github.com/gorilla/sessions
	github.com/boj/redistore

github.com/gin-gonic/contrib/cache

<!-- vendor/github.com/valyala/fasthttp/compress.go:9:2: cannot find package "github.com/klauspost/compress/flate" in any of: -->
github.com/klauspost/compress/flate

<!-- vendor/github.com/klauspost/compress/flate/crc32_amd64.go:9:2: cannot find package "github.com/klauspost/cpuid" in any of: -->
github.com/klauspost/cpuid

<!-- vendor/github.com/klauspost/compress/gzip/gunzip.go:17:2: cannot find package "github.com/klauspost/crc32" in any of: -->
github.com/klauspost/crc32

<!-- vendor/github.com/gorilla/sessions/sessions.go:13:2: cannot find package "github.com/gorilla/context" in any of: -->
github.com/gorilla/context

<!-- vendor/github.com/boj/redistore/redistore.go:19:2: cannot find package "github.com/gorilla/securecookie" in any of: -->
github.com/gorilla/securecookie
```

##框架功能

功能 | 描述
:--- | :---
配置 | conf进行统一配置
子域名部署 | 子域名区分模块
缓存 | Redis、Memcached、Memory
Session | Redis、File、Cookie
ORM | gorm，使用示例需完善
模板 | 支持PONGO2，模板文件打包bindata
静态 | 静态资源打包bindata

目标功能 | 描述
:--- | :---
缓存 | 查询缓存
安全 | SQL注入、XSS、表单令牌、验证码等
日志 | 分级
多语言 | 