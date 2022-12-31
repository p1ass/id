import Link from 'next/link'

export default function Home() {
  return (
    <main>
      <h2>Endpoint Links for Debug</h2>
      <Link href="/api/oauth2/authorize?client_id=dummy_client_id&redirect_uri=https://localhost:8443/test/a/local/callback&scope=openid&state=IcHZqZjjyY&nonce=uz3nzkvH60&response_type=code">
        /api/oauth2/authorize
      </Link>
    </main>
  )
}
