package main

import (
	"context"
	"net/http"

	"github.com/rs/cors"

	"github.com/p1ass/id/backend/log"

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

	err := http.ListenAndServe(
		"localhost:8080",
		corsHandler,
	)
	if err != nil {
		log.Info(context.Background()).Err(err)
	}
}
