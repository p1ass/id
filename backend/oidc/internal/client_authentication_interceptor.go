package internal

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/rs/zerolog/log"
)

// NewClientAuthenticationInterceptor returns connect UnaryInterceptorFunc which authenticates client and
// embeds AuthenticatedClient into context.
// Even if authentication fails, proceeds next function, not returns error.
// This behavior aims to be able to handle the error in rpc methods.
func NewClientAuthenticationInterceptor(authenticator ClientAuthenticator) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			if req.Spec().IsClient {
				// Do noting because authentication is executed only server.
				return next(ctx, req)
			}

			authenticatedClient, err := authenticator.Authenticate(ctx, req.Header())
			if err != nil {
				log.Ctx(ctx).Info().Err(err).Msg("failed to authenticate")
				return next(ctx, req)
			}

			newCtx := ContextWithAuthenticatedClient(ctx, authenticatedClient)

			return next(newCtx, req)
		}
	}
	return interceptor
}
