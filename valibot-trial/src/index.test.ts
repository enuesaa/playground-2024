import { email, maxLength, minLength, object, Output, is, string } from 'valibot'
import { describe, expect, it } from 'vitest'

describe('name', () => {
  const nameSchema = string([minLength(1), maxLength(10)])
  it('normal', () => {
    expect(is(nameSchema, 'aa')).toStrictEqual(true)
  })
  it('min', () => {
    expect(is(nameSchema, '')).toStrictEqual(false)
  })
  it('max', () => {
    expect(is(nameSchema, 'aaaaaaaaaaaaa')).toStrictEqual(false)
  })
  it('type mismatch', () => {
    expect(is(nameSchema, 9)).toStrictEqual(false)
  })
})
