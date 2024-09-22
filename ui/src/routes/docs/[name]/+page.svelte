<script lang="ts">
	import { viewDoc } from '$lib/apidocs'
	import { page } from '$app/stores'
	import Code from '$lib/components/editor/Code.svelte'

	const name = $page.params.name
	const doc = viewDoc(name)
</script>

{#if $doc.isSuccess}
	<Code language="md" code={$doc.data.data.content} />
{/if}
