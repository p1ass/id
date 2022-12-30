import { PlainMessage } from '@bufbuild/protobuf'
import { AuthenticateRequest } from '../../../gen/oidc/v1/oidc_pb'
import { notFound, redirect } from 'next/navigation'
import { authenticate } from '../../../lib/api/oidc'

type PageProps = {
  // Workaround: If we remove ? from searchParams, we get compile error
  searchParams?: {
    client_id?: string
    redirect_uri?: string
    scope?: string
    state?: string
    nonce?: string
    response_type?: string
  }
}

const AuthorizePage = async ({ searchParams }: PageProps) => {
  if (!searchParams) {
    return notFound()
  }
  console.log(searchParams)
  const req: PlainMessage<AuthenticateRequest> = {
    // TODO: scopeはスペース区切りになっている
    scopes: searchParams.scope ? [searchParams.scope] : [],
    clientId: searchParams.client_id ?? '',
    state: searchParams.state ?? '',
    // TODO: responseTypesはスペース区切りになってる
    responseTypes: searchParams.response_type ? [searchParams.response_type] : [],
    redirectUri: searchParams.redirect_uri ?? '',
    consented: true,
  }

  const res = await authenticate(req)

  if (!res.success) {
    const errorQuery = new URLSearchParams()
    errorQuery.set('error', res.error.rawMessage)
    if (searchParams.state) {
      errorQuery.set('state', searchParams.state)
    }
    // TODO: 正しいリダイレクト
    redirect(`${searchParams.redirect_uri ?? ''}?${errorQuery.toString()}`)
  }

  const query = new URLSearchParams()
  query.set('code', res.response.code)
  if (searchParams.state) {
    query.set('state', searchParams.state)
  }
  redirect(`${searchParams.redirect_uri ?? ''}?${query.toString()}`)
}

export default AuthorizePage
