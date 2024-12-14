import { PUBLIC_API_ENDPOINT_BASE } from '$env/static/public'
import { createQuery } from '@tanstack/svelte-query'

export type QueryOptions = {
	headers: {
		[key: string]: string
	}
}
export const query = <R>(method: string, path: string, options: Partial<QueryOptions> = {}) => {
	const resolvedOptions = { headers: [], ...options }

	return createQuery({
		queryKey: [path],
		queryFn: async (): Promise<R> => {
			const res = await fetch(`${PUBLIC_API_ENDPOINT_BASE}${path}`, {
				method,
				headers: {
					Accept: 'application/json',
					...resolvedOptions.headers,
				},
			})
			const data = await res.json()
			return data
		},
	})
}
