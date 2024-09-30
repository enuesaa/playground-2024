<script lang="ts">
	import { ArrowLeftIcon } from 'svelte-feather-icons'

	export let slides: string[]
	export let selected: number
	export let content: string

	function handlePrev() {
		slides[selected] = content
		selected -= 1
		content = slides[selected]
	}
</script>

{#if selected === 0}
	<button disabled class="opacity-35"><ArrowLeftIcon /></button>
{:else}
	<button on:click|preventDefault={handlePrev}><ArrowLeftIcon /></button>
{/if}

<style lang="postcss">
	button {
		@apply text-blackgray;
	}
</style>
