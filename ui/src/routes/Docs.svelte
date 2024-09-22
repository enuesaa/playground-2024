<script lang="ts">
	import { listDocs } from '$lib/apidocs'

	const docs = listDocs()
</script>

{#if $docs.isSuccess}
	{#each $docs.data.data as item}
		<a href={`/docs/${item.dirName}`}>{item.dirName}</a>
	{/each}
{/if}
