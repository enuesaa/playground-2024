import type { TreeData } from '$lib/tree'
import fs from 'node:fs/promises'
import path from 'node:path'
import type { PageServerLoad } from './$types'

type Variant = {
	out: string
	files: TreeData[]
}
type Variants = Record<string, Variant>
type Data = {
	variants: Variants;
	readme: string
}
export const load: PageServerLoad<Data> = async ({ params }) => {
	const { name } = params
	const variants: Variants = {}
	const files = await fs.readdir(`./data/${name}`, { withFileTypes: true })

	for (const file of files) {
		if (file.isDirectory()) {
			variants[file.name] = {
				files: await extract(`./data/${name}/${file.name}`),
				out: await fs.readFile(`./data/${name}/${file.name}/out.txt`, 'utf8'),
			}
		}
	}
	const readme = await fs.readFile(`./data/${name}/README.md`, 'utf8')

	return {variants, readme}
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
