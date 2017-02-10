# Echo Web
Go(Golang) web framework Echo V3 sample. Echo V3 middleware [echo-mw](https://github.com/hobo-go/echo-mw)

##子域名部署
```
# ./conf/conf.go
SERVER_ADDR = ":8080"
DOMAIN_API    = "echo.api.localhost.com"
DOMAIN_WWW    = "echo.www.localhost.com"

$ vi /etc/hosts
127.0.0.1       echo.api.localhost.com
127.0.0.1       echo.www.localhost.com

# Nginx配置
server{

    listen       80;
    server_name  echo.www.localhost.com echo.api.localhost.com;

    charset utf-8;

    location / {
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header Host            $http_host;

        proxy_pass http://127.0.0.1:8080;
    }
}
```

##测试
```
$ glide install
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
https://github.com/dchest/captcha

<!-- vendor/github.com/gorilla/sessions/sessions.go:13:2: cannot find package "github.com/gorilla/context" in any of: -->
github.com/gorilla/context

<!-- vendor/github.com/boj/redistore/redistore.go:19:2: cannot find package "github.com/gorilla/securecookie" in any of: -->
github.com/gorilla/securecookie
```

##框架功能

功能 | 描述
:--- | :---
子域名部署 | 子域名区分模块
缓存 | Redis、Memcached、Memory
Session | Redis、File、Cookie，支持Flash
ORM | gorm，使用示例需完善
模板 | 支持html/template、PONGO2，模板支持打包bindata
静态 | 静态资源，支持打包bindata
安全 | CORS、CSRF、XSS、HSTS、验证码等
其他 | JWT、Socket演示

目标功能 | 描述
:--- | :---
配置 | 配置文件
缓存 | ORM查询缓存
安全 | SQL注入等
日志 | 分级
多语言 | i18n
