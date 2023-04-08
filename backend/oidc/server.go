package oidc

import (
	"context"
	"errors"
	"net/url"

	"github.com/bufbuild/connect-go"
	oidcv1 "github.com/p1ass/id/backend/generated/oidc/v1"
	"github.com/p1ass/id/backend/oidc/internal"
	"github.com/rs/zerolog/log"
)

type service struct {
	clientDatastore      internal.ClientDatastore
	clientAuthenticator  internal.ClientAuthenticator
	codeDatastore        internal.CodeDatastore
	accessTokenDatastore internal.AccessTokenDatastore
}

// OAuth 2.0 Authorization Error Responses defined by RFC6749 or OIDC Authentication Error Responses.
//
// [RFC 6749 Section 4.1.2.1]: https://www.rfc-editor.org/rfc/rfc6749#section-4.1.2.1
// [OIDC Core Section 3.1.2.6]: https://openid.net/specs/openid-connect-core-1_0.html#AuthError
var (
	ErrInvalidRequest          = errors.New("invalid_request")
	ErrInvalidScope            = errors.New("invalid_scope")
	ErrUnsupportedResponseType = errors.New("unsupported_response_type")
	ErrUnauthorizedClient      = errors.New("unauthorized_client")
	ErrConsentRequired         = errors.New("consent_required")
)

// OAuth 2.0 Token Error Responses defined by RFC6749.
//
// [RFC 6749 Section 5.2]: https://www.rfc-editor.org/rfc/rfc6749#section-5.2
var (
	ErrInvalidClient        = errors.New("invalid_client")
	ErrInvalidGrant         = errors.New("invalid_grant")
	ErrUnsupportedGrantType = errors.New("unsupported_grant_type")
)

// Self defined error.
var (
	ErrInvalidClientID    = errors.New("invalid_client_id")
	ErrInvalidRedirectURI = errors.New("invalid_redirect_uri")
)

func (s *service) Authenticate(ctx context.Context, req *connect.Request[oidcv1.AuthenticateRequest]) (*connect.Response[oidcv1.AuthenticateResponse], error) {
	client, err := s.clientDatastore.FetchClient(req.Msg.ClientId)
	if err != nil {
		if errors.Is(err, internal.ErrClientNotFound) {
			log.Ctx(ctx).Info().Err(err).Str("clientID", req.Msg.ClientId).Msgf("client is not found")
			return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidClientID)
		}
		log.Ctx(ctx).Error().Err(err).Str("clientID", req.Msg.ClientId).Msgf("failed to fetch client")
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to fetch client"))
	}

	scopes, err := internal.NewScopes(req.Msg.Scopes)
	if err != nil {
		log.Ctx(ctx).Info().Err(err)
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidScope)
	}

	responseTypes, err := internal.NewResponseTypes(req.Msg.ResponseTypes)
	if err != nil {
		log.Ctx(ctx).Info().Err(err)
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidRequest)
	}

	if !scopes.ContainsOpenID() {
		log.Ctx(ctx).Info().Msg("scopes does not contain openid scope")
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidScope)
	}

	if !responseTypes.ContainsOnlyCode() {
		log.Ctx(ctx).Info().Msg("response types does not contain only code")
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrUnsupportedResponseType)
	}

	// TODO: エンドポイントURIはフラグメントコンポーネントを含んではいけない (MUST NOT).
	redirectURI, err := url.Parse(req.Msg.RedirectUri)
	if err != nil {
		log.Ctx(ctx).Info().Err(err).Msgf("redirectURI %s is invalid", req.Msg.RedirectUri)
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidRedirectURI)
	}

	if err := client.IdenticalRedirectURI(*redirectURI); err != nil {
		log.Ctx(ctx).Info().Err(err).Msgf("redirectURI %s is not registered in client %s", redirectURI, req.Msg.ClientId)
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidRedirectURI)
	}

	// We must check user consent after other request parameter validation
	// in order to display the error page before displaying the authorize page.
	if !req.Msg.Consented {
		log.Ctx(ctx).Info().Msg("user not consented")
		return nil, connect.NewError(connect.CodePermissionDenied, ErrConsentRequired)
	}

	code := internal.NewAuthorizationCode(client, *redirectURI)
	if err := s.codeDatastore.Save(code); err != nil {
		log.Ctx(ctx).Info().Err(err).Msgf("failed to save code to datastore")
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to save code to datastore"))
	}

	return connect.NewResponse[oidcv1.AuthenticateResponse](&oidcv1.AuthenticateResponse{
		Code: code.Code,
	}), nil
}

func (s *service) Exchange(ctx context.Context, req *connect.Request[oidcv1.ExchangeRequest]) (*connect.Response[oidcv1.ExchangeResponse], error) {
	authenticatedClient, err := s.clientAuthenticator.Authenticate(ctx, req.Header())
	if err != nil {
		log.Ctx(ctx).Info().Err(err).Msgf("failed to authenticate client")
		connectErr := connect.NewError(connect.CodeUnauthenticated, ErrInvalidClient)
		connectErr.Meta().Set("WWW-Authenticate", "Basic")
		return nil, connectErr
	}

	grantType, err := internal.NewGrantType(req.Msg.GrantType)
	if err != nil {
		log.Ctx(ctx).Info().Err(err).Str("grantType", req.Msg.GrantType).Msgf("invalid grant type")
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidRequest)
	}
	// When authorization code grant, value MUST be set to "authorization_code".
	if grantType != internal.GrantTypeAuthorizationCode {
		log.Ctx(ctx).Info().Str("grantType", string(grantType)).Msgf("grant type should be authorization code")
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrUnsupportedGrantType)
	}

	// TODO: エンドポイントURIはフラグメントコンポーネントを含んではいけない (MUST NOT).
	redirectURI, err := url.Parse(req.Msg.RedirectUri)
	if err != nil {
		log.Ctx(ctx).Info().Err(err).Str("redirectURI", req.Msg.RedirectUri).Msgf("redirectURI is invalid")
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidRedirectURI)
	}

	// verify that the authorization code is valid, and ensure that the "redirect_uri" parameter is present and identical if the
	// "redirect_uri" parameter was included in the initial authorization request
	code, err := s.codeDatastore.Fetch(req.Msg.Code, authenticatedClient.ID, *redirectURI)
	if err != nil {
		log.Ctx(ctx).Info().Err(err).Str("code", req.Msg.Code).Msgf("code which is satisfied requirements not found")
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidRedirectURI)
	}
	if code.Expired() {
		log.Ctx(ctx).Info().Msgf("code is expired")
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidGrant)
	}
	if _, err = code.Use(); err != nil {
		log.Ctx(ctx).Error().Err(err).Msgf("failed to use code")
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidGrant)
	}
	// TODO: Verify that the Authorization Code used was issued in response to an OpenID Connect Authentication Request (so that an ID Token will be returned from the Token Endpoint).

	// TODO: pass correct sub and scopes
	const sub = "dummy_sub"
	accessToken, err := internal.NewAccessToken(sub, authenticatedClient, []internal.Scope{internal.ScopeOpenID})
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msgf("failed to initiate access token")
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to initiate access token"))
	}

	if err := s.accessTokenDatastore.Save(accessToken); err != nil {
		log.Ctx(ctx).Error().Err(err).Msgf("failed to save access token")
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to store access token"))
	}

	idToken, err := internal.NewSignedIDToken(sub, authenticatedClient)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msgf("failed to initiate id token")
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to initiate id token"))
	}

	res := connect.NewResponse[oidcv1.ExchangeResponse](&oidcv1.ExchangeResponse{
		AccessToken: accessToken.Token,
		IdToken:     idToken.Token(),
		TokenType:   string(accessToken.TokenType),
		ExpiresIn:   accessToken.ExpiresInSec(),
		// TODO: implement
		RefreshToken: nil,
	})

	// The authorization server MUST include the HTTP "Cache-Control"
	//   response header field [RFC2616] with a value of "no-store" in any
	//   response containing tokens, credentials, or other sensitive
	//   information, as well as the "Pragma" response header field [RFC2616]
	//   with a value of "no-cache".
	res.Trailer().Set("Cache-Control", "no-store")
	res.Trailer().Set("Pragma", "no-store")

	return res, nil
}
