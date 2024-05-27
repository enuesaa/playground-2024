import type { AppProps } from 'next/app'
import '@radix-ui/themes/styles.css'
import { Theme } from '@radix-ui/themes'
import '@/styles/app.css'
import { SessionProvider } from 'next-auth/react'

export default function App({ Component, pageProps }: AppProps) {
  return (
    <SessionProvider session={pageProps.session}>
      <Theme appearance='dark'>
        <Component {...pageProps} />
      </Theme>
      </SessionProvider>
  )
}
