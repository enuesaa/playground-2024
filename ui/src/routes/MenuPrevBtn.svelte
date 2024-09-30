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
	<button disabled class="opacity-35" />
{:else}
	<button on:click|preventDefault={handlePrev} />
{/if}

<style lang="postcss">
	button {
		@apply text-blackgray w-full h-full bg-[rgba(0,0,0,0.01)];
	}
</style>
