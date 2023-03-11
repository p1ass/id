package internal

import "fmt"

type (
	// GrantType represents OAuth 2.0 grant_type
	GrantType string
)

const (
	GrantTypeUnknown           GrantType = "unknown"
	GrantTypeAuthorizationCode GrantType = "authorization_code"
	GrantTypeClientCredentials GrantType = "client_credentials"
)

var grantTypeMap = map[string]GrantType{
	string(GrantTypeAuthorizationCode): GrantTypeAuthorizationCode,
	string(GrantTypeClientCredentials): GrantTypeClientCredentials,
}

func NewGrantType(str string) (GrantType, error) {
	if r, ok := grantTypeMap[str]; ok {
		return r, nil
	}
	return GrantTypeUnknown, fmt.Errorf("%s is not valid grant type", str)
}
