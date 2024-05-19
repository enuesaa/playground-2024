import fs from 'node:fs/promises'
import type { PageServerLoad } from './$types'

type Project = {
	name: string
}
type Data = {
	projects: Project[]
}

export const load: PageServerLoad<Data> = async () => {
	const list: Project[] = []
	const files = await fs.readdir('./data', { withFileTypes: true })

	for (const file of files) {
		list.push({ name: file.name })
	}

	return {
		projects: list,
	}
}
