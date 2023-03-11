package internal

import "fmt"

type (
	// AccessTokenType represents OAuth 2.0 access token type.
	AccessTokenType string

	// AccessToken represents OAUth 2.0 access token.
	AccessToken struct {
	}
)

const (
	AccessTokenTypeUnknown AccessTokenType = "Unknown"
	AccessTokenTypeBearer  AccessTokenType = "Bearer"
)

var accessTokenTypeMap = map[string]AccessTokenType{
	string(AccessTokenTypeBearer): AccessTokenTypeBearer,
}

func NewAccessToken() (*AccessToken, error) {
	return &AccessToken{}, nil
}

func NewAccessTokenType(str string) (AccessTokenType, error) {
	if r, ok := accessTokenTypeMap[str]; ok {
		return r, nil
	}
	return AccessTokenTypeUnknown, fmt.Errorf("%s is not valid access token type", str)
}
