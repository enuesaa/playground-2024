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

type DocSchema = {
  data: {path: string; dirName: string, content: string}
}
export const viewDoc = (name: string) =>
  createQuery<DocSchema>({
    queryKey: ['viewDoc', name],
    queryFn: async () => {
      const res = await fetch(`http://localhost:3000/api/docs/${name}`)
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

type UpdateSchema = {
  dirName: string
  content: string
}
export const useUpdateDoc = () => {
  const queryClient = getQueryClientContext()

  return createMutation({
    mutationFn: async (body: UpdateSchema) => {
      const res = await fetch(`http://localhost:3000/api/docs/${body.dirName}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          content: body.content,
        })
      })
      await res.json()
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['listDocs', 'viewDoc'] })
    }
  })
}
  