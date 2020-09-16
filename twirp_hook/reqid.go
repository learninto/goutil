package twirp_hook

import (
	"context"
	"time"

	"github.com/learninto/goutil/ctxkit"
	"github.com/learninto/goutil/trace"
	"github.com/learninto/goutil/twirp"
)

// NewRequestID 生成唯一请求标识并记录到 ctx
func NewRequestID() *twirp.ServerHooks {
	return &twirp.ServerHooks{
		RequestReceived: func(ctx context.Context) (context.Context, error) {
			ctx = context.WithValue(ctx, ctxkit.StartTimeKey, time.Now())
			traceID := trace.GetTraceID(ctx)
			_ = twirp.SetHTTPResponseHeader(ctx, "x-trace-id", traceID)

			ctx = ctxkit.WithTraceID(ctx, traceID)
			ctx = twirp.WithAllowGET(ctx, true)

			return ctx, nil
		},
	}
}
