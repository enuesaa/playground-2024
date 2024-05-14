import { writable } from 'svelte/store'

type FileResponse = {
  files: FileResponseItem[]
}
export type FileResponseItem = {
  filename: string
  exists: boolean
  content: string
}
export const fetchFiles = async (): Promise<FileResponseItem[]> => {
  const res = await fetch('http://localhost:3000/api/files', {
    method: 'GET',
    headers: {
      Accept: 'application/json',
    },
  })
  const body = await res.json() as FileResponse
  return body?.files ?? []
}

type FilesState = {
  files: FileResponseItem[],
  open?: string
}
export const filesStore = writable<FilesState>({files: [], open: undefined}, (set) => {
  fetchFiles().then(files => {
    const state = {
      files,
      open: files.length > 0 ? files[0].filename : undefined,
    }
    set(state)
  })
})
