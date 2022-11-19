package oidc

import (
	"context"
	"errors"
	"net/url"

	"github.com/bufbuild/connect-go"

	oidcv1 "github.com/p1ass/id/backend/gen/oidc/v1"
	"github.com/p1ass/id/backend/gen/oidc/v1/oidcv1connect"
	"github.com/p1ass/id/backend/log"
	"github.com/p1ass/id/backend/oidc/internal"
)

type OIDCServer struct {
	clientDatastore internal.ClientDatastore
}

// OAuth 2.0 Error Responses defined by RFC6749.
//
// [RFC 6749 Section 4.1.2.1]: https://www.rfc-editor.org/rfc/rfc6749#section-4.1.2.1
var (
	ErrInvalidRequest = errors.New("invalid request")
	ErrInvalidScope   = errors.New("invalid scope")
)

func NewOIDCServer() oidcv1connect.OIDCPrivateServiceHandler {
	return &OIDCServer{}
}

func (s *OIDCServer) Authenticate(ctx context.Context, req *connect.Request[oidcv1.AuthenticateRequest]) (*connect.Response[oidcv1.AuthenticateResponse], error) {
	scopes, err := internal.NewScopes(req.Msg.Scopes)
	if err != nil {
		log.Info(ctx).Err(err)
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidScope)
	}

	responseTypes, err := internal.NewResponseTypes(req.Msg.ResponseTypes)
	if err != nil {
		log.Info(ctx).Err(err)
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidRequest)
	}

	if !scopes.ContainsOpenId() {
		log.Info(ctx).Msg("scopes does not contain openid scope")
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidScope)
	}

	if !responseTypes.ContainsOnlyCode() {
		log.Info(ctx).Msg("response types does not contain only code")
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidRequest)
	}

	client, err := s.clientDatastore.FetchClient(req.Msg.ClientId)
	if err != nil {
		log.Info(ctx).Err(err).Msgf("client id = %s is not found", req.Msg.ClientId)
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidRequest)
	}

	redirectURI, err := url.Parse(req.Msg.RedirectUri)
	if err != nil {
		log.Info(ctx).Err(err).Msgf("redirectURI %s is invalid", req.Msg.RedirectUri)
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidRequest)
	}

	if err := client.IdenticalRedirectURI(*redirectURI); err != nil {
		log.Info(ctx).Err(err).Msgf("redirectURI %s is not registered in client %s", redirectURI, req.Msg.ClientId)
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidRequest)
	}

	// TODO implement me
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("authenticate is unimplemented"))
}

func (s *OIDCServer) Exchange(ctx context.Context, req *connect.Request[oidcv1.ExchangeRequest]) (*connect.Response[oidcv1.ExchangeResponse], error) {
	// TODO implement me
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("exchange is unimplemented"))
}
