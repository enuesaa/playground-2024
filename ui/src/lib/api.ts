import { createMutation, createQuery, getQueryClientContext } from '@tanstack/svelte-query'

export const baseUrl = 'http://localhost:3000/api'

type FilesResponse = {path: string; filename: string; isDir: boolean}[]
export const listFiles = (path: string) =>
  createQuery({
    queryKey: ['listFiles', path],
    queryFn: async (): Promise<FilesResponse> => {
      const res = await fetch(`${baseUrl}/files?path=${path}`)
      return await res.json()
    }
  })

type CreateFileRequest = {
  path: string
  content: string
}
export const useCreateFile = () => {
  const queryClient = getQueryClientContext()

  return createMutation({
    mutationFn: async (body: CreateFileRequest) => {
      const res = await fetch(`${baseUrl}/file`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(body)
      })
      await res.json()
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['listFiles'] })
    }
  })
}
