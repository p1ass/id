package oidc

import (
	"net/http"

	"github.com/bufbuild/connect-go"
	otelconnect "github.com/bufbuild/connect-opentelemetry-go"
	"github.com/p1ass/id/backend/generated/oidc/v1/oidcv1connect"
	"github.com/p1ass/id/backend/oidc/internal"
	"github.com/p1ass/id/backend/pkg/config"
	"github.com/p1ass/id/backend/pkg/zerologgcloud"
)

// NewServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
func NewServiceHandler(cfg *config.Config) (string, http.Handler) {
	clientDatastore := internal.NewInMemoryClientDatastore()

	clients := internal.NewClientFixture()
	for _, client := range clients {
		err := clientDatastore.SaveClient(client)
		if err != nil {
			panic(err)
		}
	}
	srv := &server{
		clientDatastore:      clientDatastore,
		codeDatastore:        internal.NewInMemoryCodeDatastore(),
		accessTokenDatastore: internal.NewInMemoryAccessTokenDatastore(),
	}

	return oidcv1connect.NewOIDCPrivateServiceHandler(srv,
		connect.WithInterceptors(
			otelconnect.NewInterceptor(otelconnect.WithTrustRemote()),
			zerologgcloud.NewCloudLoggingTraceContextInterceptor(cfg.GoogleCloudProjectID),
			internal.NewClientAuthenticationInterceptor(internal.NewBasicClientAuthenticator(clientDatastore)),
		),
	)
}
