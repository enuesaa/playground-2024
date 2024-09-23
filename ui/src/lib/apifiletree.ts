import { createMutation, createQuery, getQueryClientContext } from '@tanstack/svelte-query'
import { baseUrl } from './api'

type FiletreeResponse = {
  tree: string
}
export const lookupFiletree = () => createMutation({
    mutationFn: async (): Promise<FiletreeResponse> => {
      const res = await fetch(`${baseUrl}filetree`)
      return await res.json()
    },
  })
