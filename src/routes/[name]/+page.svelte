<script lang="ts">
	import type { PageData } from './$types'
	import PageTitle from './PageTitle.svelte'
	import Description from './Description.svelte'
	import VariantNav from './VariantNav.svelte'
	import VariantBody from './VariantBody.svelte'

	export let data: PageData
	let showingVariantName = data.variants[0].name
</script>

<PageTitle title={data.title} />
<Description content={data.description} />

<VariantNav variants={data.variants} bind:showing={showingVariantName} />

{#each data.variants as variant}
	{#if variant.name === showingVariantName}
		<VariantBody {variant} />
	{/if}
{/each}
