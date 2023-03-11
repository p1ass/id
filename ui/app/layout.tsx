import './globals.css'
import type { Metadata } from 'next'

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="ja">
      <body>{children}</body>
    </html>
  )
}

export const metadata: Metadata = {
  title: 'p1ass ID',
  description: 'OAuth 2.0 / OpenID Connect Provider',
  viewport: { width: 'device-width', initialScale: 1 },
  icons: '/favicon.ico'
}
