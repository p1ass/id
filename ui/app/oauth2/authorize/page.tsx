'use client'
import {
  createConnectTransport,
  createPromiseClient,
} from "@bufbuild/connect-web";
import {OIDCPrivateService} from "../../../gen/oidc/v1/oidc_connectweb";
import {useEffect} from "react";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});

// Here we make the client itself, combining the service
// definition with the transport.
const client = createPromiseClient(OIDCPrivateService, transport);

const AuthorizePage =  () => {
  useEffect(()=>{
    const authenticateAsync = async ()=>{
      const res = await client.authenticate({
      })
      console.log(res)
    }
    authenticateAsync()
  },[])
  return <div>
    <h1>Authorize</h1>
  </div>
}

export default AuthorizePage
