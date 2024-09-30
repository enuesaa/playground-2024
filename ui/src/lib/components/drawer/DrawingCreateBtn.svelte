<script lang="ts">
	import type { Registry } from '$lib/registry'
	import { goto } from '$app/navigation'
	import { createDrawing } from '$lib/drawing'

	export let registry: Registry

	async function handleClick() {
		const [id, _] = createDrawing(registry)
		await goto(`/drawing/${id}`)
	}
</script>

<button on:click|preventDefault={handleClick}> Save to LocalStorage </button>

<style lang="postcss">
	button {
		@apply text-sm border-solid border-[1px] rounded mb-1 p-1;
	}
</style>
