import { createTreeView } from '@melt-ui/svelte'
import { readable, writable, type Writable } from 'svelte/store'
import { setContext, getContext } from 'svelte'

export type TreeData = {
	id: string
	title: string
	children: TreeData[]
	code: string
	language: string
}

export const createTreeViewCtl = () => {
	setContext('treeView', createTreeView())
}
export const getTreeViewCtl = () => {
	const treeView = getContext<ReturnType<typeof createTreeView>>('treeView')
	return treeView.elements
}
export const createViewing = () => {
	setContext('viewing', writable<TreeData | undefined>(undefined))
}
export const getViewing = () => {
	return getContext<Writable<TreeData | undefined>>('viewing')
}
