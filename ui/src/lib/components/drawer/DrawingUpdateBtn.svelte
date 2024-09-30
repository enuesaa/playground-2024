<script lang="ts">
	import type { Registry } from '$lib/registry'
	import { updateDrawing } from '$lib/drawing'

	export let id: string
	export let registry: Registry

	function handleClick() {
		const [_, updated] = updateDrawing(id, registry)
		registry = updated
	}
</script>

<button on:click|preventDefault={handleClick}> Save to LocalStorage </button>

<style lang="postcss">
	button {
		@apply text-sm border-solid border-[1px] rounded mb-1 p-1;
	}
</style>
