# Golang Web

##子域名部署
```
# ./conf/cong.go
DOMAIN_API    = "echo.api.localhost:8080"
DOMAIN_WWW    = "echo.www.localhost:8080"

$ vi /etc/hosts
127.0.0.1       echo.api.localhost
127.0.0.1       echo.www.localhost
```

##测试
```
$ ./run.sh
```

##打包
```
$ ./build.sh 		# 需要go-bindata工具
```
```
Bindata打包工具
https://github.com/jteeuwen/go-bindata
go-bindata -ignore=\\.DS_Store -ignore=assets.go -pkg="assets" -o assets/assets.go assets/...
go-bindata -ignore=\\.DS_Store -ignore=templates.go -pkg="templates" -o templates/templates.go templates/...
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
Session | Redis、File、Cookie，fasthttp模式Session不支持
ORM | gorm，使用示例需完善
模板 | 支持template、PONGO2，template模板支持打包bindata
静态 | 静态资源打包bindata
Socket | Socket演示

目标功能 | 描述
:--- | :---
缓存 | 查询缓存
Flash |
安全 | SQL注入、XSS、表单令牌、验证码等
日志 | 分级
多语言 | 