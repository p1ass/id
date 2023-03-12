package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/p1ass/id/backend/pkg/log"

	"github.com/rs/cors"

	"github.com/p1ass/id/backend/generated/oidc/v1/oidcv1connect"
	"github.com/p1ass/id/backend/oidc"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	log.Init()

	server := oidc.NewOIDCServer()

	mux := http.NewServeMux()
	path, handler := oidcv1connect.NewOIDCPrivateServiceHandler(server)
	mux.Handle(path, handler)

	// Use h2c so we can serve HTTP/2 without TLS.
	corsHandler := cors.Default().Handler(h2c.NewHandler(mux, &http2.Server{}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Info(context.Background()).Msgf("Starting server...")

	err := http.ListenAndServe(
		fmt.Sprintf("0.0.0.0:%s", port),
		corsHandler,
	)
	if err != nil {
		log.Info(context.Background()).Err(err)
	}
}
