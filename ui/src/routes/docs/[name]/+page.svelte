<script lang="ts">
	import Trailer from './Trailer.svelte'
	import { viewDoc } from '$lib/apidocs'
	import { page } from '$app/stores'

	const dirName = $page.params.name
	const doc = viewDoc(dirName)
</script>

{#if $doc.isSuccess}
	<Trailer {dirName} content={$doc.data.content} />
{/if}
