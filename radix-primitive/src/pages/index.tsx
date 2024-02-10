import { Header } from '@/components/common/Header'
import { Main } from '@/components/common/Main'
import * as Popover from '@radix-ui/react-popover'

export default function TopPage() {
  return (
    <>
      <Header />
      <Main>
        <Popover.Root>
          <Popover.Trigger className="PopoverTrigger">More info</Popover.Trigger>
          <Popover.Portal>
            <Popover.Content className="PopoverContent" sideOffset={5}>
              Some more info…
              <Popover.Arrow className="PopoverArrow" />
            </Popover.Content>
          </Popover.Portal>
        </Popover.Root>
        a
      </Main>
    </>
  )
}
