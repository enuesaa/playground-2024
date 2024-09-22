import { createMutation, createQuery, getQueryClientContext } from '@tanstack/svelte-query'

type DocsSchema = {
  data: {path: string; dirName: string}[]
}
export const listDocs = () =>
  createQuery<DocsSchema>({
    queryKey: ['listDocs'],
    queryFn: async () => {
      const res = await fetch(`http://localhost:3000/api/docs`)
      const body = await res.json()
      return body
    }
  })
