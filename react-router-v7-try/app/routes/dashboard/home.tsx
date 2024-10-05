import type { MetaFunction } from 'react-router'

export const meta: MetaFunction = () => [
  { title: 'Dashboard Home' },
]

export default function Index() {
  return (
    <div className='flex h-screen items-center justify-center'>
      dashboard home.tsx
    </div>
  )
}
