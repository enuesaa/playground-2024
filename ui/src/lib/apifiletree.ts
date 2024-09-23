import { createMutation, createQuery, getQueryClientContext } from '@tanstack/svelte-query'

export const lookupFiletree = () => createMutation({
    mutationFn: async (): Promise<{data: {tree: string}}> => {
      const res = await fetch(`http://localhost:3000/api/filetree`)
      const body = await res.json()
      return body
    },
  })
