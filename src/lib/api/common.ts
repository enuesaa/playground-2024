import { query, type QueryOptions } from './common-query'
import { mutate, type MutateOptions } from './common-mutate'

export const queryGet = <R>(path: string, options: Partial<QueryOptions> = {}) =>
	query<R>('GET', path, options)

export const mutatePost = <T, R>(path: string, options: Partial<MutateOptions> = {}) =>
	mutate<T, R>('POST', path, options)

export const mutatePut = <T, R>(path: string, options: Partial<MutateOptions> = {}) =>
	mutate<T, R>('PUT', path, options)

export const mutateDelete = <T, R>(path: string, options: Partial<MutateOptions> = {}) =>
	mutate<T, R>('DELETE', path, options)
