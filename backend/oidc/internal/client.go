package internal

import (
	"errors"
	"net/url"
)

var (
	ErrNotIdenticalRedirectURI    = errors.New("not identical redirect uri")
	ErrClientNotAuthenticated     = errors.New("client not authenticated")
	ErrClientCredentialNotAllowed = errors.New("client credential not allowed")
)

type ClientType string

const (
	ClientTypeUnknown      ClientType = "unknown"
	ClientTypeConfidential ClientType = "confidential"
	ClientTypePublic       ClientType = "public"
)

// Client represents OAuth 2.0 client.
type Client struct {
	// ID is a unique string  and is exposed to public.
	ID string

	Type ClientType

	// secret is used for HTTP Basic Authentication Scheme [RFC2617].
	// [RFC2617]: https://www.rfc-editor.org/rfc/rfc2617.html
	secret *HashedPassword

	// redirectURIs are absolute URIs.
	// [RFC6749 Section3.1.2]: https://www.rfc-editor.org/rfc/rfc6749#section-3.1.2
	redirectURIs []url.URL
}

// AuthenticatedClient is a user authenticated client made from ClientAuthenticator#Authenticate method.
// It is used for preventing mistakes that we use client without client authentication.
type AuthenticatedClient struct {
	*Client
}

func NewClient(id string, clientType ClientType, hashedPassword *HashedPassword, redirectURIs []url.URL) (*Client, error) {
	c := &Client{
		ID:           id,
		Type:         clientType,
		secret:       hashedPassword,
		redirectURIs: redirectURIs,
	}
	return c, nil
}

func (c *Client) IdenticalRedirectURI(redirectURI url.URL) error {
	for _, uri := range c.redirectURIs {
		if uri.String() == redirectURI.String() {
			return nil
		}
	}
	return ErrNotIdenticalRedirectURI
}
