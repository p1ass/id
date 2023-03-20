package zerologgcloud

import (
	"context"

	"github.com/bufbuild/connect-go"
)

// SpanContextConnectInterceptor is a connect-go interceptor which sets zerolog sub logger into context.Context.
type SpanContextConnectInterceptor struct {
	googleCloudProjectID string
}

var _ connect.Interceptor = &SpanContextConnectInterceptor{}

// NewCloudLoggingTraceContextInterceptor initializes a pointer of SpanContextConnectInterceptor.
func NewCloudLoggingTraceContextInterceptor(googleCloudProjectID string) *SpanContextConnectInterceptor {
	return &SpanContextConnectInterceptor{
		googleCloudProjectID: googleCloudProjectID,
	}
}

func (i *SpanContextConnectInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
		ctx = WithCloudLoggingSpanContext(ctx, i.googleCloudProjectID)
		return next(ctx, request)
	}
}

func (i *SpanContextConnectInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return func(ctx context.Context, spec connect.Spec) connect.StreamingClientConn {
		ctx = WithCloudLoggingSpanContext(ctx, i.googleCloudProjectID)
		return next(ctx, spec)
	}
}

func (i *SpanContextConnectInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return func(ctx context.Context, conn connect.StreamingHandlerConn) error {
		ctx = WithCloudLoggingSpanContext(ctx, i.googleCloudProjectID)
		return next(ctx, conn)
	}
}
