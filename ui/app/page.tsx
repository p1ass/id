import Link from "next/link";

export default function Home() {
  return (
    <main>
      <h2>Endpoint Links for Debug</h2>
      <Link  href="/oauth2/authorize">/oauth2/authorize</Link>
    </main>
  )
}
