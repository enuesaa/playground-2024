<script lang="ts">
	import type { FileItem } from '$lib/api'
	import { ArrowRightIcon } from 'svelte-feather-icons'

	export let item: FileItem
</script>

<div class="p-1 my-1 flex">
	<span class="flex-none">
		{item.name}
	</span>

	{#if item.isDir}
		<a href={`/?path=${item.name}`} data-sveltekit-reload class="mx-3 px-1 bg-slate-600">
			<ArrowRightIcon />
		</a>
	{/if}
</div>
