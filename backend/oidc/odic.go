package oidc

import (
	"context"
	"errors"
	"log"

	"github.com/p1ass/id/backend/oidc/internal"

	"github.com/bufbuild/connect-go"
	oidcv1 "github.com/p1ass/id/backend/gen/oidc/v1"
	"github.com/p1ass/id/backend/gen/oidc/v1/oidcv1connect"
)

type OIDCServer struct {
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

func (o *OIDCServer) Authenticate(ctx context.Context, req *connect.Request[oidcv1.AuthenticateRequest]) (*connect.Response[oidcv1.AuthenticateResponse], error) {
	scopes, err := internal.NewScopes(req.Msg.Scopes)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidScope)
	}

	responseTypes, err := internal.NewResponseTypes(req.Msg.ResponseTypes)
	if err != nil {
		log.Println(err)
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidRequest)
	}

	if !scopes.ContainsOpenId() {
		log.Println("scopes does not contain openid scope")
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidScope)
	}

	if !responseTypes.ContainsOnlyCode() {
		log.Println("response types does not contain only code")
		return nil, connect.NewError(connect.CodeInvalidArgument, ErrInvalidRequest)
	}

	// TODO implement me
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("authenticate is unimplemented"))
}

func (o *OIDCServer) Exchange(ctx context.Context, req *connect.Request[oidcv1.ExchangeRequest]) (*connect.Response[oidcv1.ExchangeResponse], error) {
	// TODO implement me
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("exchange is unimplemented"))
}
