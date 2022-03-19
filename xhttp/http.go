// Package xhttp 提供基础 http 客户端组件
// 内置以下功能：
// - logging
// - opentracing
// - prometheus
package xhttp

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/learninto/goutil/conf"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/learninto/goutil/errors"
	"github.com/learninto/goutil/log"
	"github.com/learninto/goutil/trace"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// HTTPDurationsSeconds http 调用耗时
var HTTPDurationsSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Namespace:   conf.AppID,
	Name:        "http_durations_seconds",
	Help:        "HTTP latency distributions",
	Buckets:     defBuckets,
	ConstLabels: map[string]string{"app": conf.AppID},
}, []string{"url", "status"})

type myClient struct {
	cli *http.Client
}

// Client http 客户端接口
type Client interface {
	// Do 发送单个 http 请求
	Do(ctx context.Context, req *http.Request) (*http.Response, error)
}

// NewClient 创建 Client 实例
func NewClient(timeout time.Duration) Client {
	c := &myClient{
		cli: &http.Client{
			Timeout: timeout,
		},
	}

	return c
}

// NewInsecureClient 创建不校验证书的 Client 实例
func NewInsecureClient(timeout time.Duration) Client {
	defaultTransport := http.DefaultTransport.(*http.Transport)
	defaultTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	c := &myClient{
		cli: &http.Client{
			Timeout:   timeout,
			Transport: defaultTransport,
		},
	}

	return c
}

var digitsRE = regexp.MustCompile(`\b\d+\b`)

func (c *myClient) Do(ctx context.Context, req *http.Request) (resp *http.Response, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "DoHTTP")
	defer span.Finish()

	req = req.WithContext(ctx)

	trace.InjectTraceHeader(span.Context(), req)

	start := time.Now()
	resp, err = c.cli.Do(req)
	duration := time.Since(start)

	url := fmt.Sprintf("%s%s", req.URL.Host, req.URL.Path)

	status := http.StatusOK
	if err != nil {
		err = errors.Wrap(err)
		status = http.StatusGatewayTimeout
	} else {
		status = resp.StatusCode
	}

	log.Get(ctx).Debugf(
		"[HTTP] method:%s url:%s status:%d query:%s",
		req.Method,
		url,
		status,
		req.URL.RawQuery,
	)

	span.SetTag(string(ext.Component), "http")
	span.SetTag(string(ext.HTTPUrl), url)
	span.SetTag(string(ext.HTTPMethod), req.Method)
	span.SetTag(string(ext.HTTPStatusCode), status)

	// url 中带有的纯数字替换成 %d，不然 prometheus 就炸了
	// /v123/4/56/foo => /v123/%d/%d/foo
	url = digitsRE.ReplaceAllString(url, "%d")

	HTTPDurationsSeconds.WithLabelValues(
		url,
		fmt.Sprint(status),
	).Observe(duration.Seconds())

	return
}
