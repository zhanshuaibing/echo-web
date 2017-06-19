package opentracing

import (
	"fmt"
	"net"
	"net/http"
	"net/url"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/opentracing/opentracing-go"
	"sourcegraph.com/sourcegraph/appdash"
	appdashot "sourcegraph.com/sourcegraph/appdash/opentracing"
	"sourcegraph.com/sourcegraph/appdash/traceapp"
)

const (
	DefaultKey = "github.com/hb-go/echo-web/middleware/opentracing"
)

func InitGlobalTracer() {
	// OpenTracing
	// glide get github.com/opentracing/opentracing-go
	// glide get sourcegraph.com/sourcegraph/appdash
	// glide get github.com/gogo/protobuf
	store := appdash.NewMemoryStore()

	// Listen on any available TCP port locally.
	l, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 0})
	if err != nil {
		log.Panic(err)
	}
	collectorPort := l.Addr().(*net.TCPAddr).Port
	collectorAdd := fmt.Sprintf(":%d", collectorPort)

	// Start an Appdash collection server that will listen for spans and
	// annotations and add them to the local collector (stored in-memory).
	cs := appdash.NewServer(l, appdash.NewLocalCollector(store))
	go cs.Start()

	// Print the URL at which the web UI will be running.
	appdashPort := 8700
	appdashURLStr := fmt.Sprintf("http://localhost:%d", appdashPort)
	appdashURL, err := url.Parse(appdashURLStr)
	if err != nil {
		errStr := fmt.Sprintf("Error parsing %s: %s", appdashURLStr, err)
		log.Panic(errStr)
	}
	log.Debugf("To see your traces, go to %s/traces\n", appdashURL)

	// Start the web UI in a separate goroutine.
	tapp, err := traceapp.New(nil, appdashURL)
	if err != nil {
		log.Panic(err)
	}
	tapp.Store = store
	tapp.Queryer = store
	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", appdashPort), tapp))
	}()

	tracer := appdashot.NewTracer(appdash.NewRemoteCollector(collectorAdd))
	opentracing.InitGlobalTracer(tracer)
}

func OpenTracing(comp string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var span opentracing.Span
			opName := comp + ":" + c.Request().URL.Path
			// 监测Header中是否有Trace信息
			wireContext, err := opentracing.GlobalTracer().Extract(
				opentracing.TextMap,
				opentracing.HTTPHeadersCarrier(c.Request().Header))
			if err != nil {
				// 启动新Span
				span = opentracing.StartSpan(opName)
			} else {
				log.Debugf("opentracing span child!")
				span = opentracing.StartSpan(opName, opentracing.ChildOf(wireContext))
			}

			defer span.Finish()
			c.Set(DefaultKey, span)

			span.SetTag("component", comp)
			span.SetTag("span.kind", "server")
			span.SetTag("http.url", c.Request().Host+c.Request().RequestURI)
			span.SetTag("http.method", c.Request().Method)

			if err := next(c); err != nil {
				span.SetTag("error", true)
				c.Error(err)
			}

			span.SetTag("error", false)
			span.SetTag("http.status_code", c.Response().Status)

			return nil
		}
	}
}

func Default(c echo.Context) opentracing.Span {
	ot := c.Get(DefaultKey)
	if ot == nil {
		return nil
	}
	return c.Get(DefaultKey).(opentracing.Span)
}
