import { PlainMessage } from '@bufbuild/protobuf'
import type { NextApiRequest, NextApiResponse } from 'next'
import { z } from 'zod'

import { AuthenticateRequest } from '../../../gen/oidc/v1/oidc_pb'
import { authenticate } from '../../../lib/api/oidc'
import { convertToSearchParam } from '../../../lib/searchParam'

const AuthorizeParameterSchema = z.object({
  client_id: z.string().min(1),
  redirect_uri: z.string().url(),
  // Used to maintain state between the request and the callback.
  // This prevents CSRF attack, so MUST be specified.
  scope: z.string().min(1),
  state: z.optional(z.string().min(1)),
  response_type: z.string().min(1)
})

// ErrorJson is only used when redirectUri is invalid.
// TODO: should render error html not json.
type ErrorJson = {
  error: 'invalid_request' | 'method_not_allowed'
}

export default async function handler(req: NextApiRequest, res: NextApiResponse<ErrorJson>) {
  if (req.method === 'GET') {
    return getHandler(req, res)
  }
  if (req.method === 'POST') {
    return postHandler(req, res)
  }

  return res.status(405).json({ error: 'method_not_allowed' })
}

async function getHandler(req: NextApiRequest, res: NextApiResponse<ErrorJson>) {
  const parsed = AuthorizeParameterSchema.safeParse(req.query)
  if (!parsed.success) {
    console.error(parsed.error)
    return res.status(400).send({ error: 'invalid_request' })
  }

  // TODO: consented is not always false
  const consented = false

  if (!consented) {
    const redirectSearchParam = convertToSearchParam(parsed.data)
    return res.redirect(302, `/oauth2/authorize/consent?${redirectSearchParam.toString()}`)
  }
}

async function postHandler(req: NextApiRequest, res: NextApiResponse<ErrorJson>) {
  const parsed = AuthorizeParameterSchema.safeParse(req.body)
  if (!parsed.success) {
    console.error(parsed.error)
    return res.status(400).send({ error: 'invalid_request' })
  }
  const parameter = parsed.data

  // TODO: consented is not always true
  const consented = true

  const authenticateReq: PlainMessage<AuthenticateRequest> = {
    scopes: [parameter.scope],
    clientId: parameter.client_id,
    responseTypes: [parameter.response_type],
    redirectUri: parameter.redirect_uri,
    consented: consented
  }

  const authenticateRes = await authenticate(authenticateReq)

  if (!authenticateRes.success) {
    // The authorization server MUST NOT automatically redirect the user-agent to the
    // invalid redirection URI.
    if (['invalid_client_id', 'invalid_redirect_uri'].includes(authenticateRes.error.rawMessage)) {
      console.error(authenticateRes.error)
      return res.status(400).send({ error: 'invalid_request' })
    }

    const errorQuery = convertToSearchParam({
      error: authenticateRes.error.rawMessage,
      state: parameter.state
    })
    return res.redirect(302, `${parameter.redirect_uri}?${errorQuery.toString()}`)
  }

  const query = convertToSearchParam({
    code: authenticateRes.response.code,
    state: parameter.state
  })
  return res.redirect(302, `${parameter.redirect_uri}?${query.toString()}`)
}
