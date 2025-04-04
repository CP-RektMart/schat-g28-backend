import type { Metadata } from 'next'

import './globals.css'

export const metadata: Metadata = {
  title: 'Socket chat',
  description: 'Group 28 created simple chat',
}

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <html lang='en'>
      <body>{children}</body>
    </html>
  )
}
