package zerologgcloud

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

// WithCloudLoggingSpanContext retrieves a span context using OpenTelemetry API and initialize context sub logger.
// Sub logger has two fields that are compatible with [the Cloud Logging format].
//
// [the Cloud Logging format]: https://cloud.google.com/logging/docs/agent/logging/configuration?hl=ja#special-fields
func WithCloudLoggingSpanContext(ctx context.Context, projectID string) context.Context {
	spanContext := trace.SpanContextFromContext(ctx)
	if !spanContext.IsValid() || projectID == "" {
		return log.Logger.WithContext(ctx)
	}

	traceID := fmt.Sprintf("projects/%s/traces/%s", projectID, spanContext.TraceID().String())
	spanID := spanContext.SpanID().String()
	return log.With().
		Str("logging.googleapis.com/trace", traceID).
		Str("logging.googleapis.com/spanId", spanID).
		Bool("logging.googleapis.com/trace_sampled", spanContext.IsSampled()).
		Logger().
		WithContext(ctx)
}
