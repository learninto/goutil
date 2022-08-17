package goutil

import (
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/learninto/goutil/conf"
	"github.com/learninto/goutil/ctxkit"
	"github.com/learninto/goutil/mc"
	"github.com/learninto/goutil/trace"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/learninto/goutil/conf" // init conf
	"github.com/learninto/goutil/db"
	"github.com/learninto/goutil/log"
)

func PrometheusHandleFunc(pattern string) {
	if len(pattern) == 0 {
		pattern = "/metrics"
	}

	metricsHandler := promhttp.Handler()
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		// GatherMetrics 收集一些被动指标
		GatherMetrics()

		metricsHandler.ServeHTTP(w, r)
	})
}

func Ping(pattern string) {
	if len(pattern) == 0 {
		pattern = "/monitor/ping"
	}

	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})
}

// Reset all utils
func Reset() {
	log.Reset()
	db.Reset()
}

// Stop all utils
func Stop() {
}

// PanicHandler
type PanicHandler struct {
	Handler http.Handler
}

// ServeHTTP
func (s PanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	r, span := trace.StartSpanServerHTTP(r, "ServeHTTP") // 开始链路
	defer func() {
		if rec := recover(); rec != nil {
			ctx := r.Context()
			ctx = ctxkit.WithTraceID(ctx, trace.GetTraceID(ctx))
			log.Get(ctx).Error(rec, string(debug.Stack()))
		}
		span.Finish()
	}()

	origin := r.Header.Get("Origin")
	suffix := conf.Get("CORS_ORIGIN_SUFFIX")

	if origin != "" && suffix != "" && strings.HasSuffix(origin, suffix) {
		w.Header().Add("Access-Control-Allow-Origin", origin)
		w.Header().Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Origin,No-Cache,X-Requested-With,If-Modified-Since,Pragma,Last-Modified,Cache-Control,Expires,Content-Type,Access-Control-Allow-Credentials,DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Cache-Webcdn,Content-Length")
	}

	if r.Method == http.MethodOptions {
		return
	}

	s.Handler.ServeHTTP(w, r)
}

// GatherMetrics 收集一些被动指标
func GatherMetrics() {
	mc.GatherMetrics()
}
