# Golang Web

# Glide

```
https://github.com/Masterminds/glide

$ glide create                            	# Start a new workspace
$ open glide.yaml                         	# and edit away!
$ glide get github.com/hobo-go/echo-md 		# Get a package and add to glide.yaml
$ glide install                           	# Install packages and dependencies

$ go build                                	# Go tools work normally
$ glide up                                	# Update to newest versions of the package
```
## glide get
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



