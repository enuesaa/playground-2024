import type { TreeData } from '$lib/tree'
import fs from 'node:fs/promises'
import path from 'node:path'
import type { PageServerLoad } from './$types'

type Variants = Record<string, TreeData[]>
export const load: PageServerLoad<Variants> = async ({ params }) => {
	const { name } = params
	const variants: Variants = {}
	const files = await fs.readdir(`./data/${name}`, { withFileTypes: true })

	for (const file of files) {
		if (file.isDirectory()) {
			variants[file.name] = await extract(`./data/${name}/${file.name}`)
		}
	}

	return variants
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
	const list: Entry[] = []
	const files = await fs.readdir('./data', { withFileTypes: true })

	for (const file of files) {
		list.push({name: file.name})
	}

	return list
}
