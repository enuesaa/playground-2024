<script lang="ts">
	import { getTreeViewCtl } from '$lib/tree'
	import type { TreeData } from '$lib/types/tree'
	import { melt } from '@melt-ui/svelte'
	import CodeTreeItemButton from './CodeTreeItemButton.svelte'

	export let treeData: TreeData[]
	const { group } = getTreeViewCtl()
</script>

{#each treeData as data}
	<li class="pl-4">
		<CodeTreeItemButton {data} />

		{#if data.children.length > 0}
			<ul use:melt={$group({ id: data.id })}>
				<svelte:self treeData={data.children} />
			</ul>
		{/if}
	</li>
{/each}
