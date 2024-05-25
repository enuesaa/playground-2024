import type { PageServerLoad } from './$types'
import { listTrailers } from '$lib/server/trailer/list'

type Data = {
	projects: {name: string}[]
}
export const load: PageServerLoad<Data> = async () => {
	const list = await listTrailers()

	return {
		projects: list.map(v => ({name: v})),
	}
}
