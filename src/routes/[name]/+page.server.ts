import type { TreeData } from '$lib/tree'
import fs from 'node:fs/promises'
import path from 'node:path'
import type { PageServerLoad } from './$types'
import toml from 'toml'

type Config = {
	project: {
		description: string,
	},
	variants: Record<string, {console: string}>
}
type Variants = Record<string, {output: string; files: TreeData[]}>
type Data = {
	description: string
	variants: Variants;
}
export const load: PageServerLoad<Data> = async ({ params }) => {
	const { name } = params

	const configstr = await fs.readFile(`./data/${name}/trailer.toml`, 'utf8')
	const config: Config = toml.parse(configstr)
	const data = {
		description: config.project.description,
		variants: {} as Variants,
	}
	for (const variant of Object.keys(config.variants)) {
		data.variants[variant] = {
			output: config.variants[variant].console,
			files: await extract(`./data/${name}/${variant}`),
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

		if (file.name === 'out.txt') {
			continue
		}

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
	const list: Entry[] = []
	const files = await fs.readdir('./data', { withFileTypes: true })

	for (const file of files) {
		list.push({ name: file.name })
	}

	return list
}
