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


type CreationSchema = {
  dirName: string
}
export const useCreateDocs = () => {
  const queryClient = getQueryClientContext()

  return createMutation({
    mutationFn: async (body: CreationSchema) => {
      const res = await fetch(`http://localhost:3000/api/docs`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(body)
      })
      await res.json()
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['listDocs'] })
    }
  })
}
  