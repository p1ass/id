# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [oidc/v1/oidc.proto](#oidc_v1_oidc-proto)
    - [AuthenticateRequest](#oidc-v1-AuthenticateRequest)
    - [AuthenticateResponse](#oidc-v1-AuthenticateResponse)
    - [ExchangeRequest](#oidc-v1-ExchangeRequest)
    - [ExchangeResponse](#oidc-v1-ExchangeResponse)
  
    - [OIDCPrivateService](#oidc-v1-OIDCPrivateService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="oidc_v1_oidc-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## oidc/v1/oidc.proto



<a name="oidc-v1-AuthenticateRequest"></a>

### AuthenticateRequest
Spec: [OpenID Connect Core 1.0 Section 3.1.2.1.](https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| scopes | [string](#string) | repeated | scopes MUST contains `openid` scope. |
| response_types | [string](#string) | repeated | Supported type is only `code` (Authorization Code Flow) |
| client_id | [string](#string) |  | OAuth 2.0 Client identifier |
| redirect_uri | [string](#string) |  | Redirection URI to which the response will be sent. This URI MUST exactly match one of the Redirection URI values for the Client pre-registered at the OpenID Provider/ |
| consented | [bool](#bool) |  | Whether user consents to authorize/authenticate the client. |






<a name="oidc-v1-AuthenticateResponse"></a>

### AuthenticateResponse
Spec: [OpenID Connect Core 1.0 Section 3.1.2.6.](http://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#AuthResponse)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [string](#string) |  | OAuth 2.0 Authorization Code |






<a name="oidc-v1-ExchangeRequest"></a>

### ExchangeRequest
Spec: [OpenID Connect Core 1.0 Section 3.1.3.1.](http://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#TokenRequest)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| grant_type | [string](#string) |  | Grant type MUST be `authorization_code` |
| code | [string](#string) |  | OAuth 2.0 Authorization Code |
| redirect_uri | [string](#string) |  | redirect_uri MUST be identical authenticate request `redirect_uri` value |






<a name="oidc-v1-ExchangeResponse"></a>

### ExchangeResponse
Spec: [OpenID Connect Core 1.0 Section 3.1.3.3.](http://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#TokenResponse)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_token | [string](#string) |  | Access Token |
| id_token | [string](#string) |  | ID Token value associated with the authenticated session |
| token_type | [string](#string) |  | Token type MUST be `Bearer`, as specified in [RFC 6750](https://www.rfc-editor.org/rfc/rfc6750.htm) |
| expires_in | [uint32](#uint32) |  | Lifetime in seconds of the access token. Requirement level is MUST (Original spec is RECOMMENDED). |
| refresh_token | [string](#string) | optional | Refresh token |





 

 

 


<a name="oidc-v1-OIDCPrivateService"></a>

### OIDCPrivateService
OIDCPrivateService provides APIs to finish OpenID Connect flow.
It is designed as a private API, so it is intended to be requested by the Next.js server, not browser.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Authenticate | [AuthenticateRequest](#oidc-v1-AuthenticateRequest) | [AuthenticateResponse](#oidc-v1-AuthenticateResponse) | Authenticate authenticates the end user and generates OAuth2.0 Authorization Code Possible error code (defined by OAuth2.0 or OpenID Connect): - InvalidArgument: &#34;invalid_scope&#34; - InvalidArgument: &#34;invalid_request_uri&#34; - InvalidArgument: &#34;unsupported_response_type&#34; - InvalidArgument: &#34;invalid_request&#34; - PermissionDenied: &#34;unauthorized_client&#34; - PermissionDenied: &#34;consent_required&#34; Possible error code (defined by Self): - InvalidArgument: &#34;invalid_client_id&#34; - InvalidArgument: &#34;invalid_redirect_uri&#34; |
| Exchange | [ExchangeRequest](#oidc-v1-ExchangeRequest) | [ExchangeResponse](#oidc-v1-ExchangeResponse) | Exchange exchanges authorization code into access token and ID Token Spec: [OpenID Connect Core 1.0 Section 3.1.3.](http://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#TokenEndpoint) |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

