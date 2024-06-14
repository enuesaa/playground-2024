import { WebContainer } from '@webcontainer/api'
import { atom, useAtom } from 'jotai'

export const containerAtom = atom(async () => {
  if (typeof window === 'object') {
    return await WebContainer.boot()
  }
})
