<script lang="ts">
	import { useCompress, type FileItem } from '$lib/api'
	import { ArrowRightIcon } from 'svelte-feather-icons'

	export let item: FileItem

	const compress = useCompress()
	async function handleConvert () {
		await $compress.mutateAsync(item.name)
	}
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

	{#if item.isCompressable}
		{#if $compress.data?.success}
		<div class="mx-3 px-2 bg-slate-600" style="border-radius: 10px;">
			compressed!
		</div>
		{:else}
		<button class="mx-3 px-2 bg-slate-600" style="border-radius: 10px;" on:click={handleConvert}>
			compress
		</button>
		{/if}
	{/if}
</div>
