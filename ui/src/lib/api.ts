import { createQuery } from '@tanstack/svelte-query'

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
