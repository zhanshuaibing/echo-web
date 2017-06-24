# Echo Web
Go web framework Echo example. 
> Echo middleware [echo-mw](https://github.com/hb-go/echo-mw)

> Echo中文文档 [go-echo.org](http://go-echo.org/)

> Requires
- go1.8+
- Echo V3

## 环境配置

##### 1.源码下载
```shell
$ cd $GOPATH/src
$ git clone git@github.com:hb-go/echo-web.git
```

##### 2.依赖安装
> [dep工具安装](https://github.com/golang/dep#usage)
```shell
$ cd echo_web/
$ dep ensure
```

##### 3.MySQL配置
```shell
# ./conf/conf.toml
[database]
name = "goweb_db"
user_name = "goweb_dba"
pwd  = "123456"
host = "127.0.0.1"
port = "3306"

# 测试数据库SQL脚本
./echo-web/res/db_structure.sql
```

##### 4.Redis、Memcached配置，可选

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

##### 5.子域名
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

##### 6.Bindata打包工具，可选(运行可选，打包必选)
> [Bindata安装](https://github.com/jteeuwen/go-bindata#installation)

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
> 打包静态资源及模板文件须[安装Bindata](https://github.com/jteeuwen/go-bindata#installation)

```shell
$ ./build.sh 		    # 默认本机
$ ./build.sh -l		    # 打包Linux平台
```
## 目录结构
```sh
assets          Web服务静态资源
conf            项目配置
middleware      中间件
mode            模型，数据库连接&ORM
  └ orm         ORM扩展
module          模块封装
  ├ auth        Auth授权
  ├ cache       缓存
  ├ log         日志
  ├ render      渲染
  ├ session     Session
  └ tmpl        Web模板
res             项目资源
  └ db          数据
router          路由
  └ api         接口路由
    ├context    自定义Context，便于扩展API层扩展
    └router     路由
  ├ socket      socket示范
  └ web         Web鲈鱼
    ├context    自定义Context，便于扩展Webb层扩展
    └router     路由        
template        模板
  └ pongo2      pongo2模板
util            公共工具
  ├ conv        类型转换
  ├ crypt       加/解密
  └ sql         SQL
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
模板 | 支持html/template、[pongo2](http://github.com/flosch/pongo2)，模板支持打包[bindata](https://github.com/jteeuwen/go-bindata#installation)
静态 | 静态资源，支持打包[bindata](https://github.com/jteeuwen/go-bindata#installation)
安全 | CORS、CSRF、XSS、HSTS、验证码等
监控 | [OpenTracing](http://opentracing.io/)，如何在项目中更方便的使用还需要研究，如ORM层
其他 | JWT、Socket演示

目标功能 | 描述
:--- | :---
安全 | SQL注入等
日志 | 分级
多语言 | i18n