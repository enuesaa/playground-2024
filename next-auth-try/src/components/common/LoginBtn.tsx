import { useSession, signIn, signOut } from 'next-auth/react'
import { Button } from '@radix-ui/themes'

export const LoginBtn = () => {
  const {data: session} = useSession()

  if (session !== null) {
    const email = session?.user?.email ?? ''
    const userImage = session?.user?.image ?? ''

    return (
      <>
        Signed in as {email}
        <img src={userImage} />
        <Button onClick={() => signOut()}>Sign out</Button>
      </>
    )
  }

  return (
    <>
      Not signed in
      <Button onClick={() => signIn()}>Sign in</Button>
    </>
  )
}
