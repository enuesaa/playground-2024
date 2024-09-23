import { createMutation, createQuery, getQueryClientContext } from '@tanstack/svelte-query'
import { baseUrl } from './api'

type RunPromptRequest = {
  command: string
}
type RunPromptOutput = {
  data: {
    output: string
  }
}
export const runPrompt = () => createMutation({
    mutationFn: async (body: RunPromptRequest): Promise<RunPromptOutput> => {
      const res = await fetch(`${baseUrl}/prompt`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(body)
      })
      return await res.json()
    },
  })
