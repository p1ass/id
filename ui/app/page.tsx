import Link from 'next/link'

import { buildAuthorizeUri } from '../pages/api/oauth2/route'

export default function Home() {
  const vercelEnv = process.env.VERCEL_ENV

  let redirectUri: string
  if (vercelEnv === 'production') {
    redirectUri = 'https://id.p1ass.com/oauth2/callback'
    // preview or development
  } else if (vercelEnv) {
    redirectUri = `https://${process.env.VERCEL_URL}/oauth2/callback`
  } else {
    redirectUri = 'http://local.p1ass.com:3000/oauth2/callback'
  }
  return (
    <main>
      <h2 className="text-3xl">Endpoint Links for Debug</h2>
      <Link
        className="text-blue-900 underline"
        href={buildAuthorizeUri({
          client_id: 'dummy_client_id',
          redirect_uri: redirectUri,
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
