import type { TreeData } from '$lib/tree'
import fs from 'node:fs/promises'
import path from 'node:path'

type Data = {
	treeData: TreeData[]
}

export async function load({ params }): Promise<Data> {
	const { name } = params
	const treeData = await extract(`./data/${name}/main`)
	return {
		treeData,
	}
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
