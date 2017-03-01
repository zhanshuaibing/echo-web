# Echo Web
Go(Golang) web framework Echo V3 sample. Echo V3 middleware [echo-mw](https://github.com/hobo-go/echo-mw)

> Requires
- go1.8+
- Echo V3

## 环境配置

##### 1.依赖安装
```
$ cd echo_web/
$ glide install
```

##### 2.MySQL配置
```
# ./conf/conf.go
DB_NAME      = "goweb_db"
DB_USER_NAME = "goweb_dba"
DB_PASSWORD  = "123456"
DB_HOST      = "127.0.0.1"
DB_PORT      = "3306"

# 测试数据库SQL脚本
./echo-web/common/db_structure.sql
```

##### 3.Redis、Memcached配置，可选

> 可选需修改session、cache的store配置
- SESSION_STORE: FILE或COOKIE
- CACHE_STORE: IN_MEMORY

```
# ./conf/conf.go
REDIS_SERVER = "127.0.0.1:6379"
REDIS_PWD    = "123456"

MEMCACHED_SERVER = "localhost:11211"
```

##### 4.子域名
```
# ./conf/conf.go
SERVER_ADDR = ":8080"
DOMAIN_API    = "echo.api.localhost.com"
DOMAIN_WWW    = "echo.www.localhost.com"

# 改host
$ vi /etc/hosts
127.0.0.1       echo.api.localhost.com
127.0.0.1       echo.www.localhost.com

# Nginx配置，可选
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

##### 5.Bindata打包工具，可选(运行可选，打包必选)
```
# https://github.com/jteeuwen/go-bindata
$ go get -u github.com/jteeuwen/go-bindata/...
```

## 运行
```
$ ./run.sh [-a] [-t]        # -a -t 可选(须安装Bindata)，以debug方式更新assets、template的Bindata资源包

# 浏览器访问
echo.www.localhost.com      # Nginx代理
echo.www.localhost.com:8080 # 无代理
```

## 打包
> 打包静态资源及模板文件须安装Bindata

```
$ ./build.sh 		    # 默认本机
$ ./build.sh -l		    # 打包Linux平台
```

## 框架功能

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

## 依赖管理Glide
```
https://github.com/Masterminds/glide

$ glide create                            	# Start a new workspace
$ open glide.yaml                         	# and edit away!
$ glide get github.com/hobo-go/echo-md 		# Get a package and add to glide.yaml
$ glide install                           	# Install packages and dependencies

$ go build                                	# Go tools work normally
$ glide up                                	# Update to newest versions of the package
```
### glide 包
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
github.com/dchest/captcha

<!-- vendor/github.com/gorilla/sessions/sessions.go:13:2: cannot find package "github.com/gorilla/context" in any of: -->
github.com/gorilla/context

<!-- vendor/github.com/boj/redistore/redistore.go:19:2: cannot find package "github.com/gorilla/securecookie" in any of: -->
github.com/gorilla/securecookie
```
