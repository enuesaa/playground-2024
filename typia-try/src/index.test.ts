import { test, expect } from 'vitest'
import { validateUser } from '.'

test('validate user', () => {
  const user = {id: '14647dfb-8006-4eab-a3f4-aecd6c218269', name: 'a', age: 20}
  // if (validateUser(user)) {
  //   // here, user is considered as type User
  //   console.log(user.id)
  // }
  expect(validateUser(user)).toStrictEqual(true)

  expect(validateUser({aa: 'b'})).toStrictEqual(false)
  expect(validateUser({id: 'a'})).toStrictEqual(false)
  expect(validateUser({id: '14647dfb-8006-4eab-a3f4-aecd6c218269'})).toStrictEqual(false)
})
