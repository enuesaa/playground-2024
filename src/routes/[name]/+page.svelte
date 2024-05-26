<script lang="ts">
	import PageTitle from './PageTitle.svelte'
	import CodeViewer from './CodeViewer.svelte'
	import VariantTitle from './VariantTitle.svelte'
	import Provider from './Provider.svelte'
	import Description from './Description.svelte'
	import type { PageData } from './$types'
	import Output from './Output.svelte'

	export let data: PageData
	let showingVariantName = data.variants[0].name
</script>

<PageTitle title={data.title} />
<Description content={data.description} />

<nav>
	{#each data.variants as variant}
		<button class="m-3"
			style="border: solid 1px #111;"
			on:click|preventDefault={() => showingVariantName = variant.name}
		>{variant.name}</button>
	{/each}
</nav>

{#each data.variants as variant}
	{#if variant.name === showingVariantName}
		<Provider>
			<VariantTitle title={variant.title} />
			<CodeViewer treeData={variant.files} />
			<Output output={variant.output} />
		</Provider>
	{/if}
{/each}
