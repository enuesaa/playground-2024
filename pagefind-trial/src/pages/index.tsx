import { Header } from '@/components/common/Header'
import { Main } from '@/components/common/Main'
import Script from 'next/script'

export default function TopPage() {
  return (
    <>
      <Header />
      <Main>
        <link href="/pagefind/pagefind-ui.css" rel="stylesheet" />
        <div id="search"></div>
        <Script src='/pagefind/pagefind-ui.js' onLoad={() => {
          new PagefindUI({ element: "#search", showSubResults: true });
        }} />
      </Main>
    </>
  )
}
