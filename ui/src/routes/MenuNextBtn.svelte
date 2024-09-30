<script lang="ts">
	import { ArrowRightIcon, PlusIcon } from 'svelte-feather-icons'

	export let slides: string[]
	export let selected: number
	export let content: string

	function handleAdd() {
		slides[selected] = content
		slides.push('')
		selected += 1
		content = slides[selected]
	}
	function handleNext() {
		slides[selected] = content
		selected += 1
		content = slides[selected]
	}
</script>

{#if slides.length === selected + 1}
	<button on:click|preventDefault={handleAdd}><PlusIcon /></button>
{:else}
	<button on:click|preventDefault={handleNext}><ArrowRightIcon /></button>
{/if}

<style lang="postcss">
	button {
		@apply text-blackgray;
	}
</style>
