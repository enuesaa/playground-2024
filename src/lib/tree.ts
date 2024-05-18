import { createTreeView } from '@melt-ui/svelte'
import { readable, writable } from 'svelte/store'

export type TreeData = {
	id: string
	title: string
	children: TreeData[]
  code: string
  language: string
}

const treeView = createTreeView()
export const treeViewCtl = readable(treeView.elements)
export const viewing = writable<TreeData|undefined>(undefined)
