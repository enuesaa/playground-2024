<script lang="ts">
	import { getTreeViewCtl, getViewing, type TreeData } from '$lib/tree'
	import { onMount } from 'svelte'
	import Code from './Code.svelte'
	import CodeTree from './CodeTree.svelte'

	export let treeData: TreeData[]

	const { tree } = getTreeViewCtl()
	const viewing = getViewing()
	onMount(() => {
		viewing.set(treeData[0])
	})
</script>

<section>
	<ul class="w-56 flex-none" {...$tree}>
		<CodeTree {treeData} />
	</ul>
	<div class="flex-auto">
		{#key $viewing}
			{#if $viewing !== undefined}
				<Code language={$viewing.language} code={$viewing.code} />
			{/if}
		{/key}
	</div>
</section>

<style lang="postcss">
	section {
		border: solid 1px rgba(0, 0, 0, 0.2);
		box-shadow: 0 1px 1px rgba(255, 255, 255, 0.2);
		@apply flex gap-4 my-5 bg-grayer overflow-y-scroll h-96;
	}	
</style>
