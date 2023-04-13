package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	cloudtracepropagator "github.com/GoogleCloudPlatform/opentelemetry-operations-go/propagator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/justinas/alice"
	v1 "github.com/p1ass/id/backend/generated/oidc/v1"
	"github.com/p1ass/id/backend/oidc"
	"github.com/p1ass/id/backend/pkg/config"
	"github.com/p1ass/id/backend/pkg/zerologgcloud"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// OpenTelemetry Configuration
	tp := trace.NewTracerProvider()
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Error().Err(err)
		}
	}()
	otel.SetTracerProvider(tp)
	propagator := propagation.NewCompositeTextMapPropagator(
		cloudtracepropagator.CloudTraceOneWayPropagator{},
		propagation.TraceContext{},
		propagation.Baggage{},
	)
	otel.SetTextMapPropagator(propagator)

	cfg := config.New()

	c := alice.New()

	// Initialize logger
	if cfg.Env == "local" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
		c = c.Append(zerologgcloud.NewRequestLoggingHandler())
	} else {
		// For Cloud Run
		zerologgcloud.SetCloudLoggingFieldFormat()
	}

	c = c.Append(cors.Default().Handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	gwMux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(func(s string) (string, bool) {
		return s, true
	}))
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := v1.RegisterOIDCPrivateServiceHandlerFromEndpoint(context.Background(), gwMux, fmt.Sprintf("127.0.0.1:%s", port), opts); err != nil {
		log.Error().Err(err)
		return
	}

	mux := http.NewServeMux()
	mux.Handle(oidc.NewServiceHandler(cfg))
	mux.Handle("/", gwMux)

	log.Info().Msg("Starting server...")

	const timeout = 10 * time.Second
	s := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", port),
		Handler:      c.Then(h2c.NewHandler(mux, &http2.Server{})),
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Info().Err(err)
	}
}
