<script lang="ts">
	import Doc from './Doc.svelte'
	import { viewDoc } from '$lib/api/doc'
	import { page } from '$app/stores'
	import DocPreview from './DocPreview.svelte'

	const isPreview = $page.url.searchParams.has('preview')
	const doc = viewDoc()
</script>

{#if $doc.isSuccess}
	{#if isPreview}
		<DocPreview slides={$doc.data.slides} />
	{:else}
		<Doc slides={$doc.data.slides} />
	{/if}
{/if}
