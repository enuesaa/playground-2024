<script lang="ts">
	import Doc from './Doc.svelte'
	import { viewDoc } from '$lib/apidoc'

	const doc = viewDoc()
</script>

{#if $doc.isSuccess}
	<Doc slides={$doc.data.slides} />
{/if}
