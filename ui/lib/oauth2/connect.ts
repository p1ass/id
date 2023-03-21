// OIDCPrivateService is designed as a private API,
// so it is intended to be requested by the Next.js server, not browser.
import { ConnectError, createPromiseClient } from '@bufbuild/connect'
import { createConnectTransport } from '@bufbuild/connect-node'
import { PlainMessage } from '@bufbuild/protobuf'

import { OIDCPrivateService } from '../../generated/oidc/v1/oidc_connect'
import { AuthenticateRequest, AuthenticateResponse } from '../../generated/oidc/v1/oidc_pb'

const baseUri = process.env.API_BASE_URL ?? 'http://local.p1ass.com:8080'

const transport = createConnectTransport({
  httpVersion: '1.1',
  baseUrl: baseUri
})

type AuthenticateResponseOrError =
  | { success: true; response: PlainMessage<AuthenticateResponse> }
  | { success: false; error: ConnectError }

export async function authenticate(
  req: PlainMessage<AuthenticateRequest>
): Promise<AuthenticateResponseOrError> {
  const client = createPromiseClient(OIDCPrivateService, transport)

  try {
    const res = await client.authenticate(req)
    console.log(res.toJsonString())
    return { response: res, success: true }
  } catch (e) {
    if (e instanceof ConnectError) {
      return { error: e, success: false }
    }
    console.error(e)
    return { error: new ConnectError('unknown error'), success: false }
  }
}
