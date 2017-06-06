# Echo Web
Go web framework Echo example. 
> Echo middleware [echo-mw](https://github.com/hobo-go/echo-mw)

> Echo中文文档 [go-echo.org](http://go-echo.org/)

> Requires
- go1.8+
- Echo V3

## 环境配置

##### 1.依赖安装
```shell
$ cd echo_web/
$ dep ensure 
```

##### 2.MySQL配置
```shell
# ./conf/conf.toml
[database]
name = "goweb_db"
user_name = "goweb_dba"
pwd  = "123456"
host = "127.0.0.1"
port = "3306"

# 测试数据库SQL脚本
./echo-web/common/db_structure.sql
```

##### 3.Redis、Memcached配置，可选

> 可选需修改session、cache的store配置
- session_store = "FILE"或"COOKIE"
- cache_store = "IN_MEMORY"


```shell
# ./conf/conf.toml
[redis]
server = "127.0.0.1:6379"
pwd = "123456"

[memcached]
server = "localhost:11211"
```

##### 4.子域名
```shell
# ./conf/conf.toml
[server]
addr = ":8080"
domain_api = "echo.api.localhost.com"
domain_web = "echo.www.localhost.com"

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
```shell
# https://github.com/jteeuwen/go-bindata
$ go get -u github.com/jteeuwen/go-bindata/...
```

## 运行
```shell
$ ./run.sh [-a] [-t]        # -a -t 可选(须安装Bindata)，以debug方式更新assets、template的Bindata资源包

# 浏览器访问
http://echo.www.localhost.com      # Nginx代理
http://echo.www.localhost.com:8080 # 无代理

# OpenTracing
http://localhost:8700/traces
```

## 打包
> 打包静态资源及模板文件须安装Bindata

```shell
$ ./build.sh 		    # 默认本机
$ ./build.sh -l		    # 打包Linux平台
```

## 框架功能

功能 | 描述
:--- | :---
配置 | [toml](http://github.com/BurntSushi/toml)配置文件
子域名部署 | 子域名区分模块
缓存 | Redis、Memcached、Memory
Session | Redis、File、Cookie，支持Flash
ORM | Fork [gorm](http://github.com/jinzhu/gorm)，`FirstSQL`、`LastSQL`、`FindSQL`、`CountSQL`支持构造查询SQL
缓存 | 支持`First`、`Last`、`Find`、`Count`的查询缓存
模板 | 支持html/template、[pongo2](http://github.com/flosch/pongo2)，模板支持打包bindata
静态 | 静态资源，支持打包bindata
安全 | CORS、CSRF、XSS、HSTS、验证码等
监控 | [OpenTracing](http://opentracing.io/)，如何在项目中更方便的使用还需要研究，如ORM层
其他 | JWT、Socket演示

目标功能 | 描述
:--- | :---
安全 | SQL注入等
日志 | 分级
多语言 | i18n

## [Dep](https://github.com/golang/dep)依赖管理
```shell
# 安装
$ go get -u github.com/golang/dep/cmd/dep

$ dep -h
dep is a tool for managing dependencies for Go projects

Usage: dep <command>

Commands:

  init    Initialize a new project with manifest and lock files
  status  Report the status of the project's dependencies
  ensure  Ensure a dependency is safely vendored in the project
  prune   Prune the vendor tree of unused packages

Examples:
  dep init                          set up a new project
  dep ensure                        install the project's dependencies
  dep ensure -update                update the locked versions of all dependencies
  dep ensure github.com/pkg/errors  add a dependency to the project

Use "dep help [command]" for more information about a command.
```
