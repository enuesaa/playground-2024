import { createTreeView } from '@melt-ui/svelte'
import { readable } from 'svelte/store'

export type TreeItemData = {
  title: string
  children?: TreeItemData[]
}

export const treeItems: TreeItemData[] = [
  { title: 'index.svelte' },
  {
    title: 'lib',
    children: [
      {
        title: 'tree',
        children: [
          {
            title: 'Tree.svelte',
          },
          {
            title: 'TreeView.svelte',
          },
        ],
      },
      {
        title: 'index.js',
      },
    ],
  },
]

const treeView = createTreeView()
export const treeViewStore = readable(treeView)
