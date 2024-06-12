import { ReactNode } from 'react'

type Props = {
  children: ReactNode
}
export const Main = ({ children }: Props) => {
  return <div>{children}</div>
}
