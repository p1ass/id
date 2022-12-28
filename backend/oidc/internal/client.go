package internal

import (
	"context"
	"crypto/sha256"
	"crypto/subtle"
	"errors"
	"net/http"
	"net/url"

	"github.com/p1ass/id/backend/log"
)

var (
	ErrNotIdenticalRedirectURI = errors.New("not identical redirect uri")
	ErrNotAuthenticatedClient  = errors.New("not authenticated client")
)

// Client represents OAuth 2.0 client.
type Client struct {
	// ID is a unique string  and is exposed to public.
	ID string
	// secret is used for HTTP Basic Authentication Scheme [RFC2617].
	// [RFC2617]: https://www.rfc-editor.org/rfc/rfc2617.html
	secret *HashedPassword

	// redirectURIs are absolute URIs.
	// [RFC6749 Section3.1.2]: https://www.rfc-editor.org/rfc/rfc6749#section-3.1.2
	redirectURIs []url.URL
}

// AuthenticatedClient is a user authenticated client made from Client#Authenticate method.
// It is used for preventing mistakes that we use client without client authentication.
type AuthenticatedClient struct {
	*Client
}

func NewClient(ID string, hashedPassword *HashedPassword, redirectURIs []url.URL) (*Client, error) {
	c := &Client{
		ID:           ID,
		secret:       hashedPassword,
		redirectURIs: redirectURIs,
	}
	return c, nil
}

// Authenticate authenticates client using Basic Authentication and returns AuthenticatedClient
func (c *Client) Authenticate(ctx context.Context, header http.Header) (*AuthenticatedClient, error) {
	req := &http.Request{
		Header: header,
	}
	basicClientID, basicClientSecret, ok := req.BasicAuth()
	if !ok {
		log.Info(ctx).Msg("not valid basic auth")
		return nil, ErrNotAuthenticatedClient
	}

	// ref: https://www.alexedwards.net/blog/basic-authentication-in-go
	// Use the subtle.ConstantTimeCompare() function to check if
	// the provided basicClientID hash equal the
	// expected basicClientID hash. ConstantTimeCompare
	// will return 1 if the values are equal, or 0 otherwise.
	basicClientIDHash := sha256.Sum256([]byte(basicClientID))
	expectedClientIDHash := sha256.Sum256([]byte(c.ID))
	clientIDMatched := subtle.ConstantTimeCompare(basicClientIDHash[:], expectedClientIDHash[:]) == 1
	if !clientIDMatched {
		log.Info(ctx).Msg("not authenticated client id")
		return nil, ErrNotAuthenticatedClient
	}

	if err := c.secret.ComparePassword(RawPassword(basicClientSecret)); err != nil {
		log.Info(ctx).Msg("not authenticated password")
		return nil, ErrNotAuthenticatedClient
	}

	return &AuthenticatedClient{c}, nil
}

func (c *Client) IdenticalRedirectURI(redirectURI url.URL) error {
	for _, uri := range c.redirectURIs {
		if uri.String() == redirectURI.String() {
			return nil
		}
	}
	return ErrNotIdenticalRedirectURI
}
