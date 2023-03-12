package internal

import (
	"errors"
	"net/url"
	"time"

	"github.com/Songmu/flextime"
	"github.com/p1ass/id/backend/pkg/randgenerator"
)

// AuthorizationCode is a authorization code defined by [RFC 6749 Section 4.1.2].
// For security consideration in [RFC 6819 Section 5.2.4], authorization code binds to clientID and redirectURI.
//
// [RFC 6749 Section 4.1.2]: https://www.rfc-editor.org/rfc/rfc6749#section-4.1.2
// [RFC 6819 Section 5.2.4]: https://www.rfc-editor.org/rfc/rfc6819#section-5.2.4
type AuthorizationCode struct {
	Code        string
	clientID    string
	redirectURI url.URL
	issued      time.Time
	expiry      time.Time
}

type UsedAuthorizationCode struct {
}

const (
	authorizationCodeByteLength = 32
	// The authorization Code MUST expire shortly after it is issued to mitigate the risk of leaks.
	// A maximum authorization code lifetime of 10 minutes is RECOMMENDED.
	// In this system, lifetime sets 30 seconds.
	authorizationCodeExpiration = 30 * time.Second
)

var (
	ErrCodeIsExpired = errors.New("code is expired")
)

func NewAuthorizationCode(client *Client, redirectURI url.URL) *AuthorizationCode {

	now := flextime.Now().UTC()
	return &AuthorizationCode{
		Code:        randgenerator.MustGenerateToString(authorizationCodeByteLength),
		clientID:    client.ID,
		redirectURI: redirectURI,
		issued:      now,
		expiry:      now.Add(authorizationCodeExpiration),
	}
}

func (c *AuthorizationCode) Expired() bool {
	now := flextime.Now().UTC()
	return now.After(c.expiry)
}

// TODO: implement
func (c *AuthorizationCode) Use() (*UsedAuthorizationCode, error) {
	if c.Expired() {
		return nil, ErrCodeIsExpired
	}

	// TODO: If an authorization code is used more than
	//         once, the authorization server MUST deny the request and SHOULD
	//         revoke (when possible) all tokens previously issued based on
	//         that authorization code.
	// https://www.rfc-editor.org/rfc/rfc6749#section-4.1.2

	return &UsedAuthorizationCode{}, nil
}
