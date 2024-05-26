<script lang="ts">
	import { melt } from '@melt-ui/svelte'
	import { getTreeViewCtl, getViewing } from '$lib/tree'
	import type { TreeData } from '$lib/types/tree'

	export let data: TreeData
	const hasChildren = data.children.length > 0

	const { item } = getTreeViewCtl()
	const viewing = getViewing()

	function hanldeClick() {
		if (hasChildren) {
			return
		}
		viewing.set(data)
	}
</script>

<button
	use:melt={$item({
		id: data.id,
		hasChildren,
	})}
	on:click|preventDefault={hanldeClick}
>
	{data.title}
</button>

<style lang="postcss">
	button {
		@apply flex items-center gap-1 rounded-md pt-1 select-none;
		@apply text-base;
	}
</style>
