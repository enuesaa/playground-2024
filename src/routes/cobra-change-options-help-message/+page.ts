import type { TreeData } from '$lib/tree'

export type Data = {
  treeData: TreeData[]
}

export async function load(): Promise<Data> {
  const treeData: TreeData[] = [
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

  return {
    treeData,
  }
}
