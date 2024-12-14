<script lang="ts">
	import { browser } from '$app/environment'
	import CanvasEl from './CanvasEl.svelte'
	import type { EventHandler } from 'svelte/elements'
	import { elstore, elcurrent } from '$lib/el.svelte'
	import { nanoid } from 'nanoid'

	const handleClick: EventHandler = (e) => {
		e.preventDefault()

		if (elcurrent.id !== undefined) {
			const id = nanoid()
			elstore[id] = {
				styles: '',
				classes: 'bg-black/30 w-2/3 h-2/3',
				children: [],
			}
			elstore[elcurrent.id].children.push(id)
		}
	}
</script>

{#if browser}
	<script src="https://cdn.tailwindcss.com"></script>
{/if}

<section>
	<CanvasEl id='root' />
</section>

<button onclick={handleClick}>add</button>

<style lang="postcss">
	section {
		@apply w-full aspect-square bg-zinc-700;
		@apply flex items-center justify-center;
	}
</style>
