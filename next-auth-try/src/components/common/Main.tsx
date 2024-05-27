import { Container } from '@radix-ui/themes'
import { ReactNode } from 'react'

type Props = {
  children: ReactNode
}
export const Main = ({ children }: Props) => {
  return <Container size='4'>{children}</Container>
}
