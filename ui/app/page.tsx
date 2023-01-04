import Link from 'next/link'

import { buildAuthorizeUri } from '../pages/api/oauth2/route'

export default function Home() {
  return (
    <main>
      <h2 className="text-3xl">Endpoint Links for Debug</h2>
      <Link
        className="text-blue-900 underline"
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
