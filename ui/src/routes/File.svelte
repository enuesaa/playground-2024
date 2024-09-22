<script lang="ts">
	import type { FileItem } from '$lib/api'
	import { FolderIcon } from 'svelte-feather-icons'
	import Files from './Files.svelte'

	export let item: FileItem

	let listed: boolean = false
	function handleClick() {
		listed = true
	}
</script>

<div class="p-1 my-1 flex gap-1">
	<span class="w-6">
		{#if item.isDir}
			<FolderIcon />
		{/if}
	</span>

	<span class="flex-none">
		{#if item.isDir}
			<button on:click|preventDefault={handleClick}>{item.filename}</button>
		{:else}
			{item.filename}
		{/if}
		{#if item.isDir && listed}
			<Files path={item.path} />
		{/if}
	</span>
</div>
