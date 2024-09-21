<script lang="ts">
	import type { PageData } from './$types'
	import PageTitle from './PageTitle.svelte'
	import Description from './Description.svelte'
	import VariantNav from './VariantNav.svelte'
	import VariantBody from './VariantBody.svelte'
	import PagePublishedBar from './PagePublishedBar.svelte'

	export let data: PageData
	let showingVariantName = data.variants[0].name
</script>

<svelte:head>
	<title>{data.title}</title>
</svelte:head>

<PageTitle title={data.title} />
<PagePublishedBar published={data.published} />
<Description content={data.description} />

<VariantNav variants={data.variants} bind:showing={showingVariantName} />

{#each data.variants as variant}
	{#if variant.name === showingVariantName}
		<VariantBody {variant} />
	{/if}
{/each}
