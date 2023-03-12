package internal

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"time"

	"github.com/Songmu/flextime"
)

type (
	// AccessTokenType represents OAuth 2.0 access token type.
	AccessTokenType string

	// AccessToken represents OAUth 2.0 access token.
	AccessToken struct {
		Token string
		// TokenType MUST be Bearer, as specified in Bearer Token Usage [RFC6750]
		//
		// [RFC6750]: https://www.rfc-editor.org/rfc/rfc6750
		TokenType AccessTokenType

		// sub is a subject identifier (user id)
		sub string
		// aud is that this AccessToken is intended for. (client id)
		aud    string
		expiry time.Time

		scopes Scopes
	}
)

const (
	accessTokenByteLength = 32
	accessTokenExpiration = 15 * time.Minute
)

const (
	AccessTokenTypeUnknown AccessTokenType = "Unknown"
	AccessTokenTypeBearer  AccessTokenType = "Bearer"
)

func NewAccessToken(sub string, aud *Client, scopes Scopes) (*AccessToken, error) {
	token := make([]byte, accessTokenByteLength)
	_, err := io.ReadFull(rand.Reader, token)
	if err != nil {
		panic(err)
	}

	now := flextime.Now().UTC()
	return &AccessToken{
		Token:     base64.RawURLEncoding.WithPadding(base64.NoPadding).EncodeToString(token),
		TokenType: AccessTokenTypeBearer,
		sub:       sub,
		aud:       aud.ID,
		expiry:    now.Add(accessTokenExpiration),
		scopes:    scopes,
	}, nil
}

func (t *AccessToken) Expired() bool {
	now := flextime.Now().UTC()
	return now.After(t.expiry)
}

func (t *AccessToken) ExpiresInSec() uint32 {
	now := flextime.Now().UTC()
	seconds := t.expiry.Sub(now).Seconds()
	if seconds < 0 {
		return 0
	}
	return uint32(seconds)
}
