<script lang="ts">
	import Slide from './Slide.svelte'
	import { useUpdateDoc } from '$lib/apidoc'

	export let slides: string[]
	let selected: number = 0

	const updateDoc = useUpdateDoc()

	async function handleSave() {
		await $updateDoc.mutateAsync({slides})
	}

	function handleNext() {
		selected += 1
	}
</script>

<button on:click|preventDefault={handleSave}>save</button>
<button on:click|preventDefault={handleNext}>next</button>

{#key selected}
	<Slide bind:slides={slides} {selected} />
{/key}

<style lang="postcss">
	textarea {
		@apply block h-full bg-[rgba(0,0,0,0)];
		@apply font-normal w-full p-3 text-black text-lg outline-none;
	}
</style>
