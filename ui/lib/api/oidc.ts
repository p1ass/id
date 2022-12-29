import { ConnectError } from '@bufbuild/connect-web'
import { codeFromString } from '@bufbuild/connect-web'
import { PlainMessage } from '@bufbuild/protobuf'
import 'server-only'
import { AuthenticateRequest, AuthenticateResponse } from '../../gen/oidc/v1/oidc_pb'

const jsonHeaders = {
  Accept: 'application/json',
  'Content-Type': 'application/json',
}

const baseUri = 'http://local.p1ass.com:8080'

type AuthenticateResponseOrError =
  | { success: true; response: PlainMessage<AuthenticateResponse> }
  | { success: false; error: ConnectError }

export async function authenticate(
  req: PlainMessage<AuthenticateRequest>
): Promise<AuthenticateResponseOrError> {
  const res = await fetch(`${baseUri}/oidc.v1.OIDCPrivateService/Authenticate`, {
    method: 'POST',
    body: JSON.stringify(req),
    headers: jsonHeaders,
    cache: 'no-store',
  })
  const resJson = await res.json()
  console.log(resJson)
  if (res.status !== 200) {
    return {
      error: new ConnectError(resJson.message, codeFromString(resJson.code)),
      success: false,
    }
  }
  return { response: resJson as PlainMessage<AuthenticateResponse>, success: true }
}
