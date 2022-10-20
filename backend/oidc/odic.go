package oidc

import (
	"github.com/p1ass/id/backend/gen/oidc/v1/oidcv1connect"
	"github.com/p1ass/id/backend/oidc/internal"
)

func NewOIDCServer() oidcv1connect.OIDCPrivateServiceHandler {
	return &internal.OIDCServer{}
}
