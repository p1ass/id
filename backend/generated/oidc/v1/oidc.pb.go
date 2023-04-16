// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: oidc/v1/oidc.proto

package oidcv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Spec: [OpenID Connect Core 1.0 Section 3.1.2.1.](https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest)
type AuthenticateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// scopes MUST contains `openid` scope.
	Scopes []string `protobuf:"bytes,1,rep,name=scopes,proto3" json:"scopes,omitempty"`
	// Supported type is only `code` (Authorization Code Flow)
	ResponseTypes []string `protobuf:"bytes,2,rep,name=response_types,json=responseTypes,proto3" json:"response_types,omitempty"`
	// OAuth 2.0 Client identifier
	ClientId string `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// Redirection URI to which the response will be sent.
	// This URI MUST exactly match one of the Redirection URI values for the Client pre-registered at the OpenID Provider.
	RedirectUri string `protobuf:"bytes,4,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	// Whether user consents to authorize/authenticate the client.
	Consented bool `protobuf:"varint,5,opt,name=consented,proto3" json:"consented,omitempty"`
}

func (x *AuthenticateRequest) Reset() {
	*x = AuthenticateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_v1_oidc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateRequest) ProtoMessage() {}

func (x *AuthenticateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_v1_oidc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return file_oidc_v1_oidc_proto_rawDescGZIP(), []int{0}
}

func (x *AuthenticateRequest) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *AuthenticateRequest) GetResponseTypes() []string {
	if x != nil {
		return x.ResponseTypes
	}
	return nil
}

func (x *AuthenticateRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *AuthenticateRequest) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

func (x *AuthenticateRequest) GetConsented() bool {
	if x != nil {
		return x.Consented
	}
	return false
}

// Spec: [OpenID Connect Core 1.0 Section 3.1.2.6.](http://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#AuthResponse)
type AuthenticateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OAuth 2.0 Authorization Code
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *AuthenticateResponse) Reset() {
	*x = AuthenticateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_v1_oidc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateResponse) ProtoMessage() {}

func (x *AuthenticateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_v1_oidc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateResponse.ProtoReflect.Descriptor instead.
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return file_oidc_v1_oidc_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticateResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

// Spec: [OpenID Connect Core 1.0 Section 3.1.3.1.](http://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#TokenRequest)
type ExchangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Grant type MUST be `authorization_code`
	GrantType string `protobuf:"bytes,1,opt,name=grant_type,json=grantType,proto3" json:"grant_type,omitempty"`
	// OAuth 2.0 Authorization Code
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	// if the "redirect_uri" parameter was included in the authorization request as described in Section 4.1.1, and their values MUST be identical.
	RedirectUri *string `protobuf:"bytes,3,opt,name=redirect_uri,json=redirectUri,proto3,oneof" json:"redirect_uri,omitempty"`
}

func (x *ExchangeRequest) Reset() {
	*x = ExchangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_v1_oidc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRequest) ProtoMessage() {}

func (x *ExchangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_v1_oidc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRequest.ProtoReflect.Descriptor instead.
func (*ExchangeRequest) Descriptor() ([]byte, []int) {
	return file_oidc_v1_oidc_proto_rawDescGZIP(), []int{2}
}

func (x *ExchangeRequest) GetGrantType() string {
	if x != nil {
		return x.GrantType
	}
	return ""
}

func (x *ExchangeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ExchangeRequest) GetRedirectUri() string {
	if x != nil && x.RedirectUri != nil {
		return *x.RedirectUri
	}
	return ""
}

// Spec: [OpenID Connect Core 1.0 Section 3.1.3.3.](http://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#TokenResponse)
type ExchangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Access Token
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// ID Token value associated with the authenticated session
	IdToken string `protobuf:"bytes,2,opt,name=id_token,json=idToken,proto3" json:"id_token,omitempty"`
	// Token type MUST be `Bearer`, as specified in [RFC 6750](https://www.rfc-editor.org/rfc/rfc6750.htm)
	TokenType string `protobuf:"bytes,3,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	// Lifetime in seconds of the access token.
	// Requirement level is MUST (Original spec is RECOMMENDED).
	ExpiresIn uint32 `protobuf:"varint,4,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	// Refresh token
	RefreshToken *string `protobuf:"bytes,5,opt,name=refresh_token,json=refreshToken,proto3,oneof" json:"refresh_token,omitempty"`
}

func (x *ExchangeResponse) Reset() {
	*x = ExchangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_v1_oidc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeResponse) ProtoMessage() {}

func (x *ExchangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_v1_oidc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeResponse.ProtoReflect.Descriptor instead.
func (*ExchangeResponse) Descriptor() ([]byte, []int) {
	return file_oidc_v1_oidc_proto_rawDescGZIP(), []int{3}
}

func (x *ExchangeResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *ExchangeResponse) GetIdToken() string {
	if x != nil {
		return x.IdToken
	}
	return ""
}

func (x *ExchangeResponse) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *ExchangeResponse) GetExpiresIn() uint32 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

func (x *ExchangeResponse) GetRefreshToken() string {
	if x != nil && x.RefreshToken != nil {
		return *x.RefreshToken
	}
	return ""
}

var File_oidc_v1_oidc_proto protoreflect.FileDescriptor

var file_oidc_v1_oidc_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x76, 0x31, 0x22, 0xb2, 0x01,
	0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x25, 0x0a,
	0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72,
	0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x55, 0x72, 0x69, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x65,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74,
	0x65, 0x64, 0x22, 0x2a, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x7d,
	0x0a, 0x0f, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x5f, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x72, 0x65,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x69, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d,
	0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x22, 0xca, 0x01,
	0x0a, 0x10, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x12, 0x28,
	0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xa6, 0x01, 0x0a, 0x12, 0x4f,
	0x49, 0x44, 0x43, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x12, 0x1c, 0x2e, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x41, 0x0a, 0x08, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x18, 0x2e, 0x6f,
	0x69, 0x64, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x8b, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x69, 0x64, 0x63,
	0x2e, 0x76, 0x31, 0x42, 0x09, 0x4f, 0x69, 0x64, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x31, 0x61,
	0x73, 0x73, 0x2f, 0x69, 0x64, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x76, 0x31, 0x3b,
	0x6f, 0x69, 0x64, 0x63, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x4f,
	0x69, 0x64, 0x63, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x4f, 0x69, 0x64, 0x63, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x13, 0x4f, 0x69, 0x64, 0x63, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x4f, 0x69, 0x64, 0x63, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oidc_v1_oidc_proto_rawDescOnce sync.Once
	file_oidc_v1_oidc_proto_rawDescData = file_oidc_v1_oidc_proto_rawDesc
)

func file_oidc_v1_oidc_proto_rawDescGZIP() []byte {
	file_oidc_v1_oidc_proto_rawDescOnce.Do(func() {
		file_oidc_v1_oidc_proto_rawDescData = protoimpl.X.CompressGZIP(file_oidc_v1_oidc_proto_rawDescData)
	})
	return file_oidc_v1_oidc_proto_rawDescData
}

var file_oidc_v1_oidc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_oidc_v1_oidc_proto_goTypes = []interface{}{
	(*AuthenticateRequest)(nil),  // 0: oidc.v1.AuthenticateRequest
	(*AuthenticateResponse)(nil), // 1: oidc.v1.AuthenticateResponse
	(*ExchangeRequest)(nil),      // 2: oidc.v1.ExchangeRequest
	(*ExchangeResponse)(nil),     // 3: oidc.v1.ExchangeResponse
}
var file_oidc_v1_oidc_proto_depIdxs = []int32{
	0, // 0: oidc.v1.OIDCPrivateService.Authenticate:input_type -> oidc.v1.AuthenticateRequest
	2, // 1: oidc.v1.OIDCPrivateService.Exchange:input_type -> oidc.v1.ExchangeRequest
	1, // 2: oidc.v1.OIDCPrivateService.Authenticate:output_type -> oidc.v1.AuthenticateResponse
	3, // 3: oidc.v1.OIDCPrivateService.Exchange:output_type -> oidc.v1.ExchangeResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_oidc_v1_oidc_proto_init() }
func file_oidc_v1_oidc_proto_init() {
	if File_oidc_v1_oidc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oidc_v1_oidc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_oidc_v1_oidc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_oidc_v1_oidc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_oidc_v1_oidc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_oidc_v1_oidc_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_oidc_v1_oidc_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_oidc_v1_oidc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oidc_v1_oidc_proto_goTypes,
		DependencyIndexes: file_oidc_v1_oidc_proto_depIdxs,
		MessageInfos:      file_oidc_v1_oidc_proto_msgTypes,
	}.Build()
	File_oidc_v1_oidc_proto = out.File
	file_oidc_v1_oidc_proto_rawDesc = nil
	file_oidc_v1_oidc_proto_goTypes = nil
	file_oidc_v1_oidc_proto_depIdxs = nil
}
