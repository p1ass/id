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
	codeDatastore   internal.CodeDatastore
}

// OAuth 2.0 Error Responses defined by RFC6749 or OIDC Authentication Error Responses.
//
// [RFC 6749 Section 4.1.2.1]: https://www.rfc-editor.org/rfc/rfc6749#section-4.1.2.1
// [OIDC Core Section 3.1.2.6]: https://openid.net/specs/openid-connect-core-1_0.html#AuthError
var (
	ErrInvalidRequest          = errors.New("invalid_request")
	ErrInvalidRequestURI       = errors.New("invalid_request_uri")
	ErrInvalidScope            = errors.New("invalid_scope")
	ErrUnsupportedResponseType = errors.New("unsupported_response_type")
	ErrUnauthorizedClient      = errors.New("unauthorized_client")
	ErrConsentRequired         = errors.New("consent_required")
)

func NewOIDCServer() oidcv1connect.OIDCPrivateServiceHandler {
	return &OIDCServer{
		clientDatastore: internal.NewInMemoryClientDatastore(),
		codeDatastore:   internal.NewInMemoryCodeDatastore(),
	}
}

func (s *OIDCServer) Authenticate(ctx context.Context, req *connect.Request[oidcv1.AuthenticateRequest]) (*connect.Response[oidcv1.AuthenticateResponse], error) {
	client, err := s.clientDatastore.FetchClient(req.Msg.ClientId)
	if err != nil {
		log.Info(ctx).Err(err).Msgf("client id = %s is not found", req.Msg.ClientId)
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidRequest)
	}

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
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrUnsupportedResponseType)
	}

	redirectURI, err := url.Parse(req.Msg.RedirectUri)
	if err != nil {
		log.Info(ctx).Err(err).Msgf("redirectURI %s is invalid", req.Msg.RedirectUri)
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidRequestURI)
	}

	if err := client.IdenticalRedirectURI(*redirectURI); err != nil {
		log.Info(ctx).Err(err).Msgf("redirectURI %s is not registered in client %s", redirectURI, req.Msg.ClientId)
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidRequestURI)
	}

	// We must check user consent after other request parameter validation
	// in order to display the error page before displaying the authorize page.
	if !req.Msg.Consented {
		log.Info(ctx).Msg("user not consented")
		return nil, connect.NewError(connect.CodePermissionDenied, ErrConsentRequired)
	}

	code := internal.NewAuthorizationCode(client, *redirectURI)
	if err := s.codeDatastore.Save(code); err != nil {
		log.Info(ctx).Err(err).Msgf("failed to save code to datastore")
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to save code to datastore"))
	}

	return connect.NewResponse[oidcv1.AuthenticateResponse](&oidcv1.AuthenticateResponse{
		Code:  code.Code,
		State: req.Msg.State,
	}), nil
}

func (s *OIDCServer) Exchange(ctx context.Context, req *connect.Request[oidcv1.ExchangeRequest]) (*connect.Response[oidcv1.ExchangeResponse], error) {
	// TODO implement me
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("exchange is unimplemented"))
}
