import { createMutation, createQuery, getQueryClientContext } from '@tanstack/svelte-query'
import { baseUrl } from './api'

export type SlideSchema = {
  content: string
}

type DocResponse = {slides: SlideSchema[]}
export const viewDoc = () =>
  createQuery({
    queryKey: ['viewDoc'],
    queryFn: async (): Promise<DocResponse> => {
      const res = await fetch(`${baseUrl}/doc`)
      return await res.json()
    }
  })

type UpdateRequest = {
  slides: SlideSchema[]
}
export const useUpdateDoc = () => {
  const queryClient = getQueryClientContext()

  return createMutation({
    mutationFn: async (body: UpdateRequest) => {
      const res = await fetch(`${baseUrl}/doc`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(body)
      })
      await res.json()
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['viewDoc'] })
    }
  })
}
  