<script lang="ts">
	import { listFiles } from '$lib/api'
	import File from './File.svelte'

	export let path: string
	const files = listFiles(path)
</script>

<ul>
	{#if $files.isSuccess}
		{#each $files.data.data as item}
			<li><File {item} /></li>
		{/each}
	{/if}
</ul>
