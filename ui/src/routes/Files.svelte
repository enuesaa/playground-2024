<script lang="ts">
	import { listFiles } from '$lib/api'
	import File from './File.svelte'

	export let basepath: string
	export let path: string
	const files = listFiles(basepath + path)
</script>

<ul>
	{#if $files.isSuccess}
		{#each $files.data.data as item}
			<li><File {basepath} {item} /></li>
		{/each}
	{/if}
</ul>
