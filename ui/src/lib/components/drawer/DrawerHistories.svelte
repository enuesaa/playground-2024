<script lang="ts">
	import { updateDrawing } from '$lib/drawing'
	import type { Registry } from '$lib/registry'
	import type { Shape } from '$lib/shape'
	import DrawerThumbnail from './DrawerThumbnail.svelte'

	export let id: string | undefined = undefined
	export let registry: Registry

	let histories: Shape[][] = []
	$: {
		histories = structuredClone(registry.histories)
		histories.reverse()
	}

	function recover(history: Shape[]) {
		if (id === undefined) {
			return
		}
		registry.shapes = structuredClone(history)
		updateDrawing(id, registry)
		window.location.reload()
	}
</script>

Histories

{#each histories as history}
	<div class="border-solid border-[1px] mb-2 relative">
		<DrawerThumbnail shapes={history} />
		<button on:click={() => recover(history)}> Recover </button>
	</div>
{/each}

<style lang="postcss">
	button {
		@apply absolute top-1 right-1 text-sm border-solid border-[1px] rounded bg-background my-1 px-1;
	}
</style>
