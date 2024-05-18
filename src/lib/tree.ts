import { createTreeView } from '@melt-ui/svelte'
import { readable, writable } from 'svelte/store'

export type TreeData = {
	id: string
	title: string
	children: TreeData[]
  code: string
  language: string
}

export const treeData: TreeData[] = [
	{
		id: '/index.js',
		title: 'index.js',
		children: [],
    code: 'const a = "b"',
    language: 'js',
	},
	{
		id: '/lib',
		title: 'lib',
    code: '',
    language: '',
		children: [
			{
				id: '/lib/tree',
				title: 'tree',
        code: '',
        language: '',
				children: [
					{
						id: '/lib/tree/Tree.svelte',
						title: 'Tree.svelte',
						children: [],
            code: '<script lang="ts"></script>',
            language: 'html',
					},
					{
						id: '/lib/tree/TreeView.svelte',
						title: 'TreeView.svelte',
						children: [],
            code: '',
            language: 'html',
					},
				],
			},
			{
				id: '/lib/main.go',
				title: 'main.go',
				children: [],
        code: 'package main',
        language: 'go',
			},
		],
	},
]

const treeView = createTreeView()
export const treeViewCtl = readable(treeView.elements)
export const viewing = writable<TreeData>(treeData[0])
