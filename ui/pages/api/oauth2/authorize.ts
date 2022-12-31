import { PlainMessage } from '@bufbuild/protobuf'
import type { NextApiRequest, NextApiResponse } from 'next'
import { URLSearchParams } from 'next/dist/compiled/@edge-runtime/primitives/url'
import { z } from 'zod'

import { AuthenticateRequest } from '../../../gen/oidc/v1/oidc_pb'
import { authenticate } from '../../../lib/api/oidc'

const AuthorizeParameterSchema = z.object({
  client_id: z.string(),
  redirect_uri: z.string().url(),
  scope: z.string(),
  state: z.nullable(z.string()),
  response_type: z.string()
})

type AuthorizeParameter = z.infer<typeof AuthorizeParameterSchema>

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
  if (req.method === 'GET') {
    return getHandler(req, res)
  }
  if (req.method === 'POST') {
    return postHandler(req, res)
  }

  return res.status(405).json('{"message": "method not allowed"}')
}

async function getHandler(req: NextApiRequest, res: NextApiResponse) {
  // TODO: implement. Copy from app/oauth2/authorize
  // If user consent is needed, redirect to consent page.
  // Otherwise, redirect to redirectUri
  const parsed = AuthorizeParameterSchema.safeParse(req.query)
  if (!parsed.success) {
    return res.status(400).send('{"message": "invalid request"}')
  }
  const parameter: AuthorizeParameter = parsed.data

  // TODO: ここの組み立てもう少しうまく型安全にやりたい
  const redirectSearchParam = new URLSearchParams()
  redirectSearchParam.set('client_id', parameter.client_id)
  redirectSearchParam.set('redirect_uri', parameter.redirect_uri)
  redirectSearchParam.set('scope', parameter.scope)
  if (parameter.state) {
    redirectSearchParam.set('state', parameter.state)
  }
  redirectSearchParam.set('response_type', parameter.response_type)
  return res.redirect(302, `/oauth2/authorize/consent?${redirectSearchParam.toString()}`)
}

async function postHandler(req: NextApiRequest, res: NextApiResponse) {
  const parsed = AuthorizeParameterSchema.safeParse(req.body)
  if (!parsed.success) {
    return res.status(400).send('{"message": "invalid request"}')
  }
  const parameter: AuthorizeParameter = parsed.data

  const authenticateReq: PlainMessage<AuthenticateRequest> = {
    scopes: [parameter.scope],
    clientId: parameter.client_id,
    state: parameter.state ?? '',
    responseTypes: [parameter.response_type],
    redirectUri: parameter.redirect_uri,
    consented: true
  }

  const authenticateRes = await authenticate(authenticateReq)

  if (!authenticateRes.success) {
    const errorQuery = new URLSearchParams()
    errorQuery.set('error', authenticateRes.error.rawMessage)
    if (parameter.state) {
      errorQuery.set('state', parameter.state)
    }
    return res.redirect(302, `${parameter.redirect_uri}?${errorQuery.toString()}`)
  }

  const query = new URLSearchParams()
  query.set('code', authenticateRes.response.code)
  if (parameter.state) {
    query.set('state', parameter.state)
  }
  return res.redirect(302, `${parameter.redirect_uri}?${query.toString()}`)
}
