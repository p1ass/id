package zerologgcloud_test

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/p1ass/id/backend/pkg/zerologgcloud"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

func TestWithCloudLoggingSpanContext(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx       context.Context //nolint:containedctx
		projectID string
	}
	tests := []struct {
		name string
		args args
		want map[string]any
	}{
		{
			name: "when span context is not set, should not contain trace and span id",
			args: args{
				ctx:       context.Background(),
				projectID: "test-project-id",
			},
			want: map[string]any{"level": "info", "message": "test-message"},
		},
		{
			name: "when span context is  set, should  contain trace and span id",
			args: args{
				ctx: trace.ContextWithRemoteSpanContext(context.Background(), trace.NewSpanContext(trace.SpanContextConfig{
					TraceID:    mustNoError(trace.TraceIDFromHex("2e1774ad9898b6d1f6912bdce3d557e0")),
					SpanID:     mustNoError(trace.SpanIDFromHex("970f3b5d49584f74")),
					TraceFlags: 1,
				})),
				projectID: "test-project-id",
			},
			want: map[string]any{
				"level":                                "info",
				"message":                              "test-message",
				"logging.googleapis.com/spanId":        "970f3b5d49584f74",
				"logging.googleapis.com/trace":         "projects/test-project-id/traces/2e1774ad9898b6d1f6912bdce3d557e0",
				"logging.googleapis.com/trace_sampled": true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			log.Logger = zerolog.New(out)

			ctx := zerologgcloud.WithCloudLoggingSpanContext(tt.args.ctx, tt.args.projectID)
			log.Ctx(ctx).Info().Msg("test-message")

			var got map[string]any

			err := json.NewDecoder(out).Decode(&got)
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(got, tt.want) {
				t.Errorf("WithCloudLoggingSpanContext() diff(-got, +want) = %s", cmp.Diff(got, tt.want))
			}
		})
	}
}

func mustNoError[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
