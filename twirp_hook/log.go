package twirp_hook

import (
	"context"
	"time"

	"github.com/learninto/goutil/conf"
	"github.com/learninto/goutil/ctxkit"
	"github.com/learninto/goutil/log"
	"github.com/learninto/goutil/metrics"
	"github.com/learninto/goutil/twirp"

	opentracing "github.com/opentracing/opentracing-go"
)

type bizResponse interface {
	GetCode() int32
	GetMsg() string
}

type ctxKeyType int

const (
	sendRespKey ctxKeyType = iota
)

// NewLog 统一记录请求日志
func NewLog() *twirp.ServerHooks {
	return &twirp.ServerHooks{
		ResponsePrepared: func(ctx context.Context) context.Context {
			span, ctx := opentracing.StartSpanFromContext(ctx, "SendResp")
			ctx = context.WithValue(ctx, sendRespKey, span)
			return ctx
		},
		ResponseSent: func(ctx context.Context) {
			if span, ok := ctx.Value(sendRespKey).(opentracing.Span); ok {
				defer span.Finish()
			}

			span, ctx := opentracing.StartSpanFromContext(ctx, "LogReq")
			defer span.Finish()

			status, _ := twirp.StatusCode(ctx)
			req, _ := twirp.HttpRequest(ctx)
			resp, _ := twirp.Response(ctx)

			var bizCode int32
			var bizMsg string
			if br, ok := resp.(bizResponse); ok {
				bizCode = br.GetCode()
				bizMsg = br.GetMsg()
			}

			start := ctx.Value(ctxkit.StartTimeKey).(time.Time)
			duration := time.Since(start)

			if _, ok := ctx.Deadline(); ok {
				if ctx.Err() != nil {
					status = "503"
				}
			}

			path := req.URL.Path

			// 外部爬接口脚本会请求任意 API
			// 导致 prometheus 无法展示数据
			if status != "404" {
				metrics.RPCDurationsSeconds.WithLabelValues(
					path,
					status,
				).Observe(duration.Seconds())
			}

			form := req.Form
			// 移除日志中的敏感信息
			if conf.IsProdEnv {
				form.Del("access_key")
				form.Del("appkey")
				form.Del("sign")
			}

			logger := log.Get(ctx)
			logger.WithFields(log.Fields{
				"path":     path,
				"status":   status,
				"params":   form.Encode(),
				"cost":     duration.Seconds(),
				"biz_code": bizCode,
				"biz_msg":  bizMsg,
			}).Info("new rpc")
		},
		Error: func(ctx context.Context, err twirp.Error) context.Context {
			c := twirp.ServerHTTPStatusFromErrorCode(err.Code())

			logger := log.Get(ctx)
			if c >= 500 {
				logger.Errorf("%+v", cause(err))
			} else if c >= 400 {
				logger.Warn(err)
			}

			return ctx
		},
	}
}

func cause(err twirp.Error) error {
	// https://github.com/pkg/errors#retrieving-the-cause-of-an-error
	type causer interface {
		Cause() error
	}
	if c, ok := err.(causer); ok {
		return c.Cause()
	}

	return err
}
