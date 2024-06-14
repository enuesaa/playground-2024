import { globalStyle } from '@/styles/global'
import { baseTheme } from '@/styles/theme'
import { Global, ThemeProvider } from '@emotion/react'
import type { AppProps } from 'next/app'

export default function App({ Component, pageProps }: AppProps) {
  return (
    <>
      <Global styles={globalStyle} />
      <ThemeProvider theme={baseTheme}>
        <Component {...pageProps} />
      </ThemeProvider>
    </>
  )
}
