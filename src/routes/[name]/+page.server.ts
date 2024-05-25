import type { TreeData } from '$lib/tree'
import type { PageServerLoad } from './$types'
import fs from 'node:fs/promises'
import { readConfig } from '$lib/server/trailer/config'
import { extractVariantFiles } from '$lib/server/trailer/variant'

type Variant =  {
	output: string
	files: TreeData[]
}
type Data = {
	name: string
	description: string
	variants: Record<string, Variant>
}
export const load: PageServerLoad<Data> = async ({ params: { name } }) => {
	const config = await readConfig(name)
	const data: Data = {
		name,
		description: config.description,
		variants: {},
	}
	for (const [variantName, variant] of Object.entries(config.variants)) {
		data.variants[variantName] = {
			output: variant.output,
			files: await extractVariantFiles(name, variantName),
		}
	}

	return data
}

type Entry = {
	name: string
}
export async function entries(): Promise<Entry[]> {
	const files = await fs.readdir('./data', { withFileTypes: true })
	return files.map(f => ({name: f.name}))
}
