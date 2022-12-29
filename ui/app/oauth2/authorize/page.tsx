import { createConnectTransport, createPromiseClient } from '@bufbuild/connect-web'
import { OIDCPrivateService } from '../../../gen/oidc/v1/oidc_connectweb'
import { PartialMessage, PlainMessage } from '@bufbuild/protobuf'
import { AuthenticateRequest } from '../../../gen/oidc/v1/oidc_pb'
import { redirect } from 'next/navigation'

const transport = createConnectTransport({
  baseUrl: 'http://localhost:8080',
})

// Here we make the client itself, combining the service
// definition with the transport.
const client = createPromiseClient(OIDCPrivateService, transport)

type Props = {
  searchParams: {
    client_id?: string
    redirect_uri?: string
    scope?: string
    state?: string
    nonce?: string
    response_type?: string
  }
}

const AuthorizePage = async ({ searchParams }: Props) => {
  console.log(searchParams)
  const req: PlainMessage<AuthenticateRequest> = {
    scopes: searchParams.scope ? [searchParams.scope] : [],
    clientId: searchParams.client_id ?? '',
    state: searchParams.state ?? '',
    responseTypes: searchParams.response_type ? [searchParams.response_type] : [],
    redirectUri: searchParams.redirect_uri ?? '',
    consented: true,
  }
  const headers = {
    Accept: 'application/json',
    'Content-Type': 'application/json',
  }

  const res = await fetch('http://local.p1ass.com:8080/oidc.v1.OIDCPrivateService/Authenticate', {
    method: 'POST',
    body: JSON.stringify(req),
    headers: headers,
    cache: 'no-store',
  })
  const resBody = await res.json()
  console.log(resBody)
  if (res.status !== 200) {
    const errorQuery = new URLSearchParams()
    errorQuery.set('error', resBody.code)
    if (searchParams.state) {
      errorQuery.set('state', searchParams.state)
    }
    redirect(`${searchParams.redirect_uri ?? ''}?${errorQuery.toString()}`)
  }

  const query = new URLSearchParams()
  query.set('code', resBody.code)
  if (searchParams.state) {
    query.set('state', searchParams.state)
  }
  redirect(`${searchParams.redirect_uri ?? ''}?${query.toString()}`)
}

export default AuthorizePage
