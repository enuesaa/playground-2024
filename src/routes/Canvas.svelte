<script lang="ts">
	import { browser } from '$app/environment'
	import CanvasEl from './CanvasEl.svelte'
	import CanvasAddButton from './CanvasAddButton.svelte'
	import CanvasConvertImageButton from './CanvasConvertImageButton.svelte'

	let canvas: HTMLDivElement|undefined = $state(undefined)
</script>

{#if browser}
	<script src="https://cdn.tailwindcss.com"></script>
{/if}

<div bind:this={canvas}>
	<CanvasEl id='root' />
</div>

<CanvasAddButton />
<CanvasConvertImageButton {canvas} />

<style lang="postcss">
	div {
		@apply w-full aspect-square bg-zinc-700;
		@apply flex items-center justify-center;
	}
</style>
