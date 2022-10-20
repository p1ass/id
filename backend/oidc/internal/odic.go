package internal

import (
	"context"

	"github.com/bufbuild/connect-go"
	oidcv1 "github.com/p1ass/id/backend/gen/oidc/v1"
)

type OIDCServer struct {
}

func (o OIDCServer) Authenticate(ctx context.Context, req *connect.Request[oidcv1.AuthenticateRequest]) (*connect.Response[oidcv1.AuthenticateResponse], error) {
	// TODO implement me
	panic("implement me")
}

func (o OIDCServer) Exchange(ctx context.Context, req *connect.Request[oidcv1.ExchangeRequest]) (*connect.Response[oidcv1.ExchangeResponse], error) {
	// TODO implement me
	panic("implement me")
}
