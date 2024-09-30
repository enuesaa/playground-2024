<script lang="ts">
	import RectFixed from './RectFixed.svelte'
	import type { Shape } from '$lib/shape'

	export let shapes: Shape[]
</script>

<svg viewBox="0 0 1000 600" class="w-full h-auto">
	{#each shapes as shape}
		{#if shape.tag === 'rect'}
			<RectFixed {shape} />
		{/if}
	{/each}
</svg>
