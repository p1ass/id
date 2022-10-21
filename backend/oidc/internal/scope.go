package internal

import (
	"fmt"
)

type (
	// Scope represents OAuth 2.0 scope.
	// The authorization server uses the "scope" response parameter to
	// inform the client of the scope of the access token issued.
	// Scope is expressed as a case-sensitive strings.
	//
	// [RFC 6749 Section 3.3]: https://www.rfc-editor.org/rfc/rfc6749#section-3.3
	Scope string

	Scopes []Scope
)

const (
	ScopeUnknown Scope = "unknown"
	ScopeOpenId  Scope = "openid"
	ScopeEmail   Scope = "email"
)

var scopeMap = map[string]Scope{
	string(ScopeOpenId): ScopeOpenId,
	string(ScopeEmail):  ScopeEmail,
}

func NewScope(str string) (Scope, error) {
	if s, ok := scopeMap[str]; ok {
		return s, nil
	}
	return ScopeUnknown, fmt.Errorf("%s is not valid scope", str)
}

func NewScopes(strs []string) (Scopes, error) {
	scopes := make([]Scope, 0, len(strs))

	for _, str := range strs {
		s, err := NewScope(str)
		if err != nil {
			return nil, err
		}
		scopes = append(scopes, s)
	}
	return scopes, nil
}

// ContainsOpenId checks if scopes contains openid scope.
// OpenID Connect requests MUST contain the openid scope value.
// (If no openid scope value is present, the request may still be a valid OAuth 2.0 request,
// but is not an OpenID Connect request.)
//
// [OpenID Connect 1.0 Core Section 3.1.2.1.]: https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
func (scopes Scopes) ContainsOpenId() bool {
	for _, s := range scopes {
		if s == ScopeOpenId {
			return true
		}
	}
	return false
}
