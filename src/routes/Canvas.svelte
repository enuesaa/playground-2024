<script lang="ts">
	import { browser } from '$app/environment'
	import CanvasEl from './CanvasEl.svelte'
</script>

{#if browser}
	<script src="https://cdn.tailwindcss.com"></script>
{/if}

<section>
	<CanvasEl />
</section>

<style lang="postcss">
	section {
		@apply w-full aspect-square bg-zinc-700;
		@apply flex items-center justify-center;
	}
</style>
