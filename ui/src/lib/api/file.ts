import { createMutation, getQueryClientContext } from '@tanstack/svelte-query'
import { baseUrl } from './common'

export const useUploadFile = () => {
  const queryClient = getQueryClientContext()

  return createMutation({
    mutationFn: async (file: File) => {
      const filename = file.name
      const formData = new FormData()
      formData.append('file', file)

      const res = await fetch(`${baseUrl}/file/${filename}`, {
        method: 'POST',
        body: formData,
      })
      await res.json()
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['viewDoc'] })
    }
  })
}
  