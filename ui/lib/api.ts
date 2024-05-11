
type FileResponse = {
  files: FileResponseItem[]
}
type FileResponseItem = {
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
  const body = await res.json()
  return body?.files ?? []
}
