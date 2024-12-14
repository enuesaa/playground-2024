<script lang="ts">
	import { browser } from '$app/environment'
</script>

{#if browser}
	<script src="https://cdn.tailwindcss.com"></script>
{/if}

<section>
	a
	<div></div>
</section>

<style lang="postcss">
	section {
		@apply w-full aspect-square bg-zinc-700;
	}
</style>
