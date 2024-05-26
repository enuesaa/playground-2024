<script lang="ts">
	import type { VariantWithTreeData } from '$lib/types/tree'
	import CodeViewer from './CodeViewer.svelte'
	import VariantOutput from './VariantOutput.svelte'
	import VariantTitle from './VariantTitle.svelte'
	import VariantProvider from './VariantProvider.svelte'

	export let variant: VariantWithTreeData
</script>

<VariantProvider>
	<section>
		<VariantTitle title={variant.title} />
		<CodeViewer treeData={variant.files} />
		<VariantOutput output={variant.output} />
	</section>
</VariantProvider>

<style lang="postcss">
	section {
		box-shadow: 0 -0.5px 1px 1px rgba(0, 0, 0, 0.2);
		@apply px-2 md:px-7 py-3 rounded-lg bg-graywhite w-full;
	}
</style>
