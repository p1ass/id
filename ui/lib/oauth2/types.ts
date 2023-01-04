import { z } from 'zod'

// [RFC6749 The OAuth 2.0 Authorization Framework Section 4.1.1.](https://www.rfc-editor.org/rfc/rfc6749#section-4.1.1)
// [OpenID Connect Core 1.0 Section 3.1.2.1.](https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest)
export const AuthorizeRequestSchema = z.object({
  client_id: z.string().min(1),
  redirect_uri: z.string().url(),
  //The value of the scope parameter is expressed as a list of space-
  // delimited, case-sensitive strings.
  scope: z.string().min(1),
  // Used to maintain state between the request and the callback.
  // This prevents CSRF attack, so MUST be specified.
  state: z.optional(z.string().min(1)),
  response_type: z.string().min(1)
})

export type AuthorizeRequest = z.infer<typeof AuthorizeRequestSchema>
