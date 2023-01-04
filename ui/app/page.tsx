import Link from 'next/link'

import { buildAuthorizeUri } from '../pages/api/oauth2/route'

export default function Home() {
  return (
    <main>
      <h2>Endpoint Links for Debug</h2>
      <Link
        href={buildAuthorizeUri({
          client_id: 'dummy_client_id',
          redirect_uri: 'https://localhost:8443/test/a/local/callback',
          scope: 'openid',
          state: 'IcHZqZjjyY',
          response_type: 'code'
        })}
      >
        /api/oauth2/authorize
      </Link>
    </main>
  )
}
