<script lang="ts">
	import PageTitle from './PageTitle.svelte'
	import CodeViewer from './CodeViewer.svelte'
	import VariantTitle from './VariantTitle.svelte'
	import Provider from './Provider.svelte'
	import Description from './Description.svelte'
	import type { PageData } from './$types'
	import Output from './Output.svelte'

	export let data: PageData
</script>

<PageTitle title={data.name} />
<Description content={data.description} />

{#each data.variants as variant}
	<Provider>
		<VariantTitle title={variant.title} />
		<CodeViewer treeData={variant.files} />
		<Output output={variant.output} />
	</Provider>
{/each}
