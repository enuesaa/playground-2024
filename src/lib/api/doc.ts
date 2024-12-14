import { queryGet, mutatePut } from './common'

type GetResponse = {
	note: string
}
export const getDoc = () => queryGet<GetResponse>('/doc')

type UpdateRequest = {
	note: string
}
export const updateDoc = () => mutatePut<UpdateRequest, {}>('/doc')
