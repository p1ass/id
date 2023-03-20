package internal

import "fmt"

type (
	// ResponseType represents OAuth 2.0 Response Type value
	// that determines the authorization processing flow to be used.
	ResponseType string

	ResponseTypes []ResponseType
)

const (
	ResponseUnknown     ResponseType = "unknown"
	ResponseTypeCode    ResponseType = "code"
	ResponseTypeIDToken ResponseType = "id_token"
	ResponseTypeToken   ResponseType = "token"
)

var responseTypeMap = map[string]ResponseType{
	string(ResponseTypeCode):    ResponseTypeCode,
	string(ResponseTypeIDToken): ResponseTypeIDToken,
	string(ResponseTypeToken):   ResponseTypeToken,
}

func NewResponseType(str string) (ResponseType, error) {
	if r, ok := responseTypeMap[str]; ok {
		return r, nil
	}
	return ResponseUnknown, fmt.Errorf("%s is not valid response type", str)
}

func NewResponseTypes(strs []string) (ResponseTypes, error) {
	responseTypes := make([]ResponseType, 0, len(strs))

	for _, str := range strs {
		r, err := NewResponseType(str)
		if err != nil {
			return nil, err
		}
		responseTypes = append(responseTypes, r)
	}
	return responseTypes, nil
}

// ContainsOnlyCode checks if response types contains only ResponseTypeCode
// because we only support the Authorization Code Flow.
//
// [OpenID Connect 1.0 Core Section 3.1.2.1]: https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
func (s ResponseTypes) ContainsOnlyCode() bool {
	if len(s) != 1 {
		return false
	}

	if s[0] == ResponseTypeCode {
		return true
	}

	return false
}
