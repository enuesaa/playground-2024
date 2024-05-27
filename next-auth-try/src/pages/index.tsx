import { Header } from '@/components/common/Header'
import { Main } from '@/components/common/Main'
import { LoginBtn } from '@/components/common/LoginBtn'

export default function TopPage() {
  return (
    <>
      <Header />
      <Main>
        <LoginBtn />
      </Main>
    </>
  )
}
