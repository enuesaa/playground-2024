import type { PageServerLoad } from './$types'
import { listTrailers } from '$lib/server/trailer/list'
import { readConfig } from '$lib/server/trailer/config'

type Data = {
	projects: {
		name: string
		title: string
	}[]
}
export const load: PageServerLoad<Data> = async () => {
	const list = await listTrailers()
	const projects = []
	for (const name of list) {
		const config = await readConfig(name)
		projects.push({
			name,
			title: config.title,
		})
	}

	return {
		projects,
	}
}
