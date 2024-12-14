import { PUBLIC_API_ENDPOINT_BASE } from '$env/static/public'
import { createMutation, getQueryClientContext } from '@tanstack/svelte-query'

export type MutateOptions = {
	headers: {
		[key: string]: string
	}
	invalidate: string[]
}
export const mutate = <T, R>(
	method: string,
	path: string,
	options: Partial<MutateOptions> = {}
) => {
	const resolvedOptions = { headers: [], invalidate: [], ...options }
	const queryClient = getQueryClientContext()

	return createMutation({
		mutationFn: async (body: T): Promise<R> => {
			const res = await fetch(`${PUBLIC_API_ENDPOINT_BASE}${path}`, {
				method,
				headers: {
					'Content-Type': 'application/json',
					...resolvedOptions.headers,
				},
				body: JSON.stringify(body),
			})
			return await res.json()
		},
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: resolvedOptions.invalidate })
		},
	})
}
