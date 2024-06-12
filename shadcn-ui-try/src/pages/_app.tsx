import type { AppProps } from 'next/app'
import { QueryClient, QueryClientProvider } from 'react-query'
import '@/styles/globals.css'

export default function App({ Component, pageProps }: AppProps) {
  const queryClient = new QueryClient()

  return (
    <QueryClientProvider client={queryClient}>
      <Component {...pageProps} />
    </QueryClientProvider>
  )
}
