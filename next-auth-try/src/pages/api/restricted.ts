import { getServerSession } from 'next-auth/next'
import { authOptions } from './auth/[...nextauth]'
import type { NextApiRequest, NextApiResponse } from 'next'

export default async (req: NextApiRequest, res: NextApiResponse) => {
  const session = await getServerSession(req, res, authOptions)

  if (session !== null) {
    res.send({
      content: 'this is a restricted content',
    })
    return;
  }

  res.send({
    error: 'please sign in',
  })
}
