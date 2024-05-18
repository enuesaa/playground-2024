<script lang="ts">
	import { treeData, treeViewCtl, viewing } from '$lib/tree'
	import Code from './Code.svelte'
	import CodeTree from './CodeTree.svelte'

	const { tree } = $treeViewCtl
</script>

<section>
	<ul class="w-56" {...$tree}>
		<CodeTree {treeData} />
	</ul>
	<div class="flex-auto">
		{#key $viewing}
			<Code language={$viewing.language} code={$viewing.code} />
		{/key}
	</div>
</section>

<style lang="postcss">
	section {
		border: solid 1px rgba(0, 0, 0, 0.2);
		box-shadow: 0 1px 1px rgba(255, 255, 255, 0.2);
		@apply flex gap-4 my-5 bg-grayer;
	}
</style>
