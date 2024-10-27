import { Authenticator } from 'cognito-at-edge'

const authenticator = new Authenticator({
  region: 'us-east-1',
  userPoolId: 'us-east-xxx',
  userPoolAppId: 'xxx',
  userPoolAppSecret: 'xxx',
  userPoolDomain: 'xxx.amazoncognito.com',
  cookiePath: '/',
})

export const handler = async (event) => {
  const res = await authenticator.handle(event)

  // this is normal response of Lambda@Edge
  console.log(res)

  return res
}
