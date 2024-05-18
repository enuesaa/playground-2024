<script lang="ts">
	import { treeViewCtl, type TreeData } from '$lib/tree'
	import { melt } from '@melt-ui/svelte'
	import CodeTreeItemButton from './CodeTreeItemButton.svelte'

	export let treeData: TreeData[]
	const { group } = $treeViewCtl
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
