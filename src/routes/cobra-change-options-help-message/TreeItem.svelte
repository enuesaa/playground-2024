<script lang="ts">
	import { melt } from '@melt-ui/svelte'
	import { treeViewStore, type TreeItemData } from './tree'
	import Tree from './Tree.svelte'

	export let title: string
	export let id: string
	export let children: TreeItemData[]|undefined

	const hasChildren = !!children?.length
	const { elements: { item, group } } = $treeViewStore
</script>

<li class='pl-4'>
	<button use:melt={$item({id, hasChildren})}>
		<span>{title}</span>
	</button>

	{#if children}
		<ul use:melt={$group({ id })}>
			<Tree treeItems={children} />
		</ul>
	{/if}
</li>

<style lang="postcss">
	button {
		@apply flex items-center gap-1 rounded-md p-1;
	}
	button > span {
		@apply select-none;
	}
</style>
