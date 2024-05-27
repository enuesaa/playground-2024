import NextAuth from 'next-auth'
import GithubProvider from 'next-auth/providers/github'

export const authOptions: Parameters<typeof NextAuth>[2] = {
  secret: process.env.SECRET ?? '',
  providers: [
    GithubProvider({
      clientId: process.env.GITHUB_ID ?? '',
      clientSecret: process.env.GITHUB_SECRET ?? '',
    }),
  ],
  callbacks: {
    async jwt({ token, account }) {
      if (account) {
        token.accessToken = account.access_token
      }
      return token
    },
    async session({ session, token, user }) {
      /** @ts-ignore This extends session to pass accessToken to frontend app. */
      session.accessToken = token.accessToken
      return session
    }
  }
}

export default NextAuth(authOptions)
