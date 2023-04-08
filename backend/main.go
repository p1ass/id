package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	cloudtracepropagator "github.com/GoogleCloudPlatform/opentelemetry-operations-go/propagator"
	"github.com/justinas/alice"
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

	mux := http.NewServeMux()
	mux.Handle(oidc.NewServiceHandler(cfg))

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
