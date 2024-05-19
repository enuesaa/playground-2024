import type { TreeData } from '$lib/tree'
import type { PageServerLoad } from './$types'
import fs from 'node:fs/promises'
import path from 'node:path'
import toml from 'toml'
import type { Config } from '$lib/config'

type Variant =  {
	output: string
	files: TreeData[]
}
type Data = {
	description: string
	variants: Record<string, Variant>
}
export const load: PageServerLoad<Data> = async ({ params: { name } }) => {
	const configstr = await fs.readFile(`./data/${name}/trailer.toml`, 'utf8')
	const config: Config = toml.parse(configstr)
	const data: Data = {
		description: config.project.description,
		variants: {},
	}
	for (const [variantName, variant] of Object.entries(config.variants)) {
		data.variants[variantName] = {
			output: variant.output,
			files: await extract(`./data/${name}/${variantName}`),
		}
	}
	return data
}

async function extract(dir: string, baseDir: string = ''): Promise<TreeData[]> {
	const data: TreeData[] = []
	const files = await fs.readdir(dir, { withFileTypes: true })

	for (const file of files) {
		const filepath = path.join(dir, file.name)
		const relpath = path.join(baseDir, file.name)

		if (file.isDirectory()) {
			const children = await extract(filepath, relpath)
			data.push({
				id: relpath,
				title: file.name,
				children,
				code: '',
				language: '',
			})
		} else {
			const language = file.name.split('.').at(-1) ?? ''
			const code = await fs.readFile(filepath, 'utf8')
			data.push({
				id: relpath,
				title: file.name,
				children: [],
				code,
				language,
			})
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
