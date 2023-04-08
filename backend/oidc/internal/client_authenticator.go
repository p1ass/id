package internal

import (
	"context"
	"crypto/sha256"
	"crypto/subtle"
	"net/http"

	"github.com/rs/zerolog/log"
)

type ClientAuthenticator interface {
	// Authenticate authenticates client and returns AuthenticatedClient or error.
	// If authentication fails, return ErrClientNotAuthenticated error.
	// if ClientType is not ClientTypeConfidential, return ErrClientCredentialNotAllowed error.
	Authenticate(ctx context.Context, header http.Header) (*AuthenticatedClient, error)
}

type BasicClientAuthenticator struct {
	datastore ClientDatastore
}

func NewBasicClientAuthenticator(datastore ClientDatastore) *BasicClientAuthenticator {
	return &BasicClientAuthenticator{datastore: datastore}
}

// Authenticate authenticates client using Basic Authentication.
func (a *BasicClientAuthenticator) Authenticate(ctx context.Context, header http.Header) (*AuthenticatedClient, error) {
	req := &http.Request{Header: header}
	basicClientID, basicClientSecret, ok := req.BasicAuth()
	if !ok {
		log.Ctx(ctx).Info().Msg("not valid basic auth")
		return nil, ErrClientNotAuthenticated
	}

	unauthenticatedClient, err := a.datastore.FetchClient(basicClientID)
	if err != nil {
		log.Ctx(ctx).Info().Msg("client not found")
		return nil, ErrClientNotAuthenticated
	}

	if unauthenticatedClient.Type != ClientTypeConfidential {
		return nil, ErrClientCredentialNotAllowed
	}

	// ref: https://www.alexedwards.net/blog/basic-authentication-in-go
	// Use the subtle.ConstantTimeCompare() function to check if
	// the provided basicClientID hash equal the
	// expected basicClientID hash. ConstantTimeCompare
	// will return 1 if the values are equal, or 0 otherwise.
	basicClientIDHash := sha256.Sum256([]byte(basicClientID))
	expectedClientIDHash := sha256.Sum256([]byte(unauthenticatedClient.ID))
	clientIDMatched := subtle.ConstantTimeCompare(basicClientIDHash[:], expectedClientIDHash[:]) == 1
	if !clientIDMatched {
		log.Ctx(ctx).Info().Msg("not authenticated client id")
		return nil, ErrClientNotAuthenticated
	}

	if err := unauthenticatedClient.secret.ComparePassword(RawPassword(basicClientSecret)); err != nil {
		log.Ctx(ctx).Info().Msg("not authenticated password")
		return nil, ErrClientNotAuthenticated
	}

	return &AuthenticatedClient{unauthenticatedClient}, nil
}
