import { PlainMessage } from '@bufbuild/protobuf'
import type { NextApiRequest, NextApiResponse } from 'next'

import { AuthenticateRequest } from '../../../generated/oidc/v1/oidc_pb'
import { authenticate } from '../../../lib/oauth2/connect'
import { AuthorizeRequestSchema } from '../../../lib/oauth2/types'
import { convertToSearchParam } from '../../../lib/searchParam'

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
  const parsed = AuthorizeRequestSchema.safeParse(req.query)
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
  const parsed = AuthorizeRequestSchema.safeParse(req.body)
  if (!parsed.success) {
    console.error(parsed.error)
    return res.status(400).send({ error: 'invalid_request' })
  }
  const parameter = parsed.data

  // TODO: consented is not always true
  const consented = true

  const authenticateReq: PlainMessage<AuthenticateRequest> = {
    scopes: parameter.scope.split(' '),
    clientId: parameter.client_id,
    responseTypes: parameter.response_type.split(' '),
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

    // If the resource owner denies the access request or if the request
    // fails for reasons other than a missing or invalid redirection URI,
    // the authorization server informs the client by adding the
    // parameters to the query component of the redirection URI using the
    // "application/x-www-form-urlencoded" format.
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
