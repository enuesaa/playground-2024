<script lang="ts">
	import PageTitle from './PageTitle.svelte'
	import CodeViewer from './CodeViewer.svelte'
	import VariantTitle from './VariantTitle.svelte'
	import Provider from './Provider.svelte'
	import Description from './Description.svelte'
	import type { PageData } from './$types'
	import Output from './Output.svelte'
	import VariantNavButton from './VariantNavButton.svelte'

	export let data: PageData
	let showingVariantName = data.variants[0].name
</script>

<PageTitle title={data.title} />
<Description content={data.description} />

<nav class="m-0">
	{#each data.variants as variant}
		<VariantNavButton bind:showing={showingVariantName} name={variant.name} />
	{/each}
</nav>

<section class="border-2 px-3 py-2 rounded-lg">
	{#each data.variants as variant}
		{#if variant.name === showingVariantName}
			<Provider>
				<VariantTitle title={variant.title} />
				<CodeViewer treeData={variant.files} />
				<Output output={variant.output} />
			</Provider>
		{/if}
	{/each}
</section>
