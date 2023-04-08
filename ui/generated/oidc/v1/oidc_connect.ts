// @generated by protoc-gen-connect-es v0.8.3 with parameter "target=ts,import_extension="
// @generated from file oidc/v1/oidc.proto (package oidc.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AuthenticateRequest, AuthenticateResponse, ExchangeRequest, ExchangeResponse } from "./oidc_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * OIDCPrivateService provides APIs to finish OpenID Connect flow.
 * It is designed as a private API, so it is intended to be requested by the Next.js server, not browser.
 *
 * @generated from service oidc.v1.OIDCPrivateService
 */
export const OIDCPrivateService = {
  typeName: "oidc.v1.OIDCPrivateService",
  methods: {
    /**
     * Authenticate authenticates the end user and generates OAuth2.0 Authorization Code
     * Possible error code (defined by OAuth2.0 or OpenID Connect):
     * - InvalidArgument: "invalid_scope"
     * - InvalidArgument: "invalid_request_uri"
     * - InvalidArgument: "unsupported_response_type"
     * - InvalidArgument: "invalid_request"
     * - PermissionDenied: "unauthorized_client"
     * - PermissionDenied: "consent_required"
     * Possible error code (defined by Self):
     * - InvalidArgument: "invalid_client_id"
     * - InvalidArgument: "invalid_redirect_uri"
     *
     * @generated from rpc oidc.v1.OIDCPrivateService.Authenticate
     */
    authenticate: {
      name: "Authenticate",
      I: AuthenticateRequest,
      O: AuthenticateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Exchange exchanges authorization code into access token and ID Token
     * Spec: [OpenID Connect Core 1.0 Section 3.1.3.](http://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#TokenEndpoint)
     * Possible error code (defined by OAuth2.0 or OpenID Connect):
     * - InvalidArgument: "invalid_request"
     * - InvalidArgument: "unsupported_grant_type"
     * - InvalidArgument: "invalid_grant"
     * - Unauthenticated: "invalid_client"
     * Possible error code (defined by Self):
     * - InvalidArgument: "invalid_redirect_uri"
     *
     * @generated from rpc oidc.v1.OIDCPrivateService.Exchange
     */
    exchange: {
      name: "Exchange",
      I: ExchangeRequest,
      O: ExchangeResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

