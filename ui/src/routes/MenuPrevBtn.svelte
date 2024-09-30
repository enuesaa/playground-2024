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

	function handleClick(e: KeyboardEvent) {
		if (e.key === 'ArrowLeft') {
			if (selected === 0) {
				return
			}
			handlePrev()
		}
	}
</script>

{#if selected === 0}
	<button disabled class="opacity-35"><ArrowLeftIcon /></button>
{:else}
	<button on:click|preventDefault={handlePrev}><ArrowLeftIcon /></button>
{/if}

<svelte:window on:keydown|preventDefault={handleClick} />

<style lang="postcss">
	button {
		@apply text-blackgray;
	}
</style>
