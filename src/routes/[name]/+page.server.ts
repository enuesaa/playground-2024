import type { TreeData } from '$lib/tree'
import type { PageServerLoad } from './$types'
import { readConfig, type Variant } from '$lib/server/trailer/config'
import { extractVariantFiles } from '$lib/server/trailer/variant'
import { listTrailers } from '$lib/server/trailer/list'

type Data = {
	name: string
	title: string
	description: string
	variants: (Variant & {files: TreeData[]})[]
}
export const load: PageServerLoad<Data> = async ({ params: { name } }) => {
	const config = await readConfig(name)
	const data: Data = {
		name,
		title: config.title,
		description: config.description,
		variants: [],
	}
	for (const variant of config.variants) {
		data.variants.push({
			...variant,
			files: await extractVariantFiles(name, variant.name),
		})
	}
	return data
}

type Entry = {
	name: string
}
export async function entries(): Promise<Entry[]> {
	const list = await listTrailers()
	return list.map(v => ({name: v}))
}
