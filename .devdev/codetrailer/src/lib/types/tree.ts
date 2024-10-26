import type { Variant } from './config'

export type TreeData = {
	id: string
	title: string
	children: TreeData[]
	code: string
	language: string
}

export type VariantWithTreeData = Variant & { files: TreeData[] }
