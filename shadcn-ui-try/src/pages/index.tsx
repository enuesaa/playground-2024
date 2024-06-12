import { Header } from '@/components/common/Header'
import { Main } from '@/components/common/Main'
import { Calendar } from "@/components/ui/calendar"

export default function TopPage() {
  return (
    <>
      <Header />
      <Main>
        <Calendar mode="single" className="rounded-md border" />
      </Main>
    </>
  )
}
