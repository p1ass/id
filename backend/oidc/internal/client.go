package internal

import (
	"errors"
	"net/url"
)

var ErrNotIdenticalRedirectURI = errors.New("not identical redirect uri")

// Client represents OAuth 2.0 client.
type Client struct {
	// ID is a unique string  and is exposed to public.
	ID string
	// hashedPassword is used for HTTP Basic Authentication Scheme [RFC2617].
	// [RFC2617]: https://www.rfc-editor.org/rfc/rfc2617.html
	hashedPassword HashedPassword

	// redirectURIs are absolute URIs.
	// [RFC6749 Section3.1.2]: https://www.rfc-editor.org/rfc/rfc6749#section-3.1.2
	redirectURIs []url.URL
}

func NewClient(ID string, hashedPassword HashedPassword, redirectURIs []url.URL) (*Client, error) {
	c := &Client{
		ID:             ID,
		hashedPassword: hashedPassword,
		redirectURIs:   redirectURIs,
	}
	return c, nil
}

func (c *Client) IdenticalRedirectURI(redirectURI url.URL) error {
	for _, uri := range c.redirectURIs {
		if uri == redirectURI {
			return nil
		}
	}
	return ErrNotIdenticalRedirectURI
}
