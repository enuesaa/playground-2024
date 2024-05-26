<script lang="ts">
	import { getTreeViewCtl, getViewing } from '$lib/tree'
	import type { TreeData } from '$lib/types/tree'
	import { onMount } from 'svelte'
	import Code from './Code.svelte'
	import CodeTree from './CodeTree.svelte'

	export let treeData: TreeData[]
	export let firstOpen: string

	const { tree } = getTreeViewCtl()
	const viewing = getViewing()
	onMount(() => {
		for (const data of treeData) {
			if (data.id === firstOpen) {
				viewing.set(data)
				break
			}
		}
	})
</script>

<section>
	<ul class="w-28 flex-none" style="border-right: solid 1.5px rgba(255,255,255,0.3);" {...$tree}>
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
		border: solid 2px rgba(0, 0, 0, 0.2);
		box-shadow: 0 1px 2px rgba(255, 255, 255, 0.2);
		background: #22272e;
		@apply flex my-5 text-grayblack overflow-y-scroll h-96;
	}
</style>
