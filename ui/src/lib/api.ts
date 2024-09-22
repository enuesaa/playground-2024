import { createMutation, createQuery, getQueryClientContext } from '@tanstack/svelte-query'

type FilesSchema = {
  data: FileItem[]
}
export type FileItem = {
  path: string
  filename: string
  isDir: boolean
}
export const listFiles = (path: string) =>
  createQuery<FilesSchema>({
    queryKey: ['listFiles', path],
    queryFn: async () => {
      const res = await fetch(`http://localhost:3000/api/files?path=${path}`)
      const body = await res.json()
      return body
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
      const res = await fetch(`http://localhost:3000/api/file`, {
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
