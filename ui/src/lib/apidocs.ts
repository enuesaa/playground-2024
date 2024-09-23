import { createMutation, createQuery, getQueryClientContext } from '@tanstack/svelte-query'
import { baseUrl } from './api'

type DocsResponse = {path: string; dirName: string}[]
export const listDocs = () =>
  createQuery({
    queryKey: ['listDocs'],
    queryFn: async (): Promise<DocsResponse> => {
      const res = await fetch(`${baseUrl}/docs`)
      const body = await res.json()
      return body
    }
  })

type DocResponse = {path: string; dirName: string, content: string}
export const viewDoc = (name: string) =>
  createQuery({
    queryKey: ['viewDoc', name],
    queryFn: async (): Promise<DocResponse> => {
      const res = await fetch(`${baseUrl}/docs/${name}`)
      return await res.json()
    }
  })

type CreateRequest = {
  dirName: string
}
export const useCreateDocs = () => {
  const queryClient = getQueryClientContext()

  return createMutation({
    mutationFn: async (body: CreateRequest) => {
      const res = await fetch(`${baseUrl}/docs`, {
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

type UpdateRequest = {
  dirName: string
  content: string
}
export const useUpdateDoc = () => {
  const queryClient = getQueryClientContext()

  return createMutation({
    mutationFn: async (body: UpdateRequest) => {
      const res = await fetch(`${baseUrl}/docs/${body.dirName}`, {
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
  