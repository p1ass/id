'use client'
import {
  createConnectTransport,
  createPromiseClient,
} from "@bufbuild/connect-web";
import {OIDCPrivateService} from "../../../gen/oidc/v1/oidc_connectweb";
import {useEffect} from "react";
import {useRouter} from "next/navigation";
import {PartialMessage, PlainMessage} from "@bufbuild/protobuf";
import {AuthenticateRequest} from "../../../gen/oidc/v1/oidc_pb";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});

// Here we make the client itself, combining the service
// definition with the transport.
const client = createPromiseClient(OIDCPrivateService, transport);

const AuthorizePage =  () => {
  const router = useRouter()
  useEffect(()=>{
    const authenticateAsync = async ()=>{
      let req: PlainMessage<AuthenticateRequest> = {
        scopes: ['openid'],
        clientId: 'dummy_client_id',
        state: 'dummy_state',
        responseTypes: ['code'],
        redirectUri: 'http://localhost:3000/oauth2/callback',
        consented:true,
      };
      const res = await client.authenticate(req)
      // TODO: 302 Foundでリダイレクトしたい
      router.push('/oauth2/callback',{
      })
    }
    authenticateAsync()
  },[])
  return <div>
    <h1>Authorize</h1>
  </div>
}

export default AuthorizePage
