<script lang="ts">
	import { type TreeItemData } from './tree'
	import TreeItem from './TreeItem.svelte'

	export let treeItems: TreeItemData[]
</script>

{#each treeItems as { title, children }, i}
	<TreeItem {title} {children} id={`${title}-${i}`} />
{/each}
