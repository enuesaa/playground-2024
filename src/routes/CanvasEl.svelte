<script lang="ts">
	import type { El } from '$lib/el.svelte'
	import Self from './CanvasEl.svelte'
	import { elcurrent, elstore } from '$lib/el.svelte'
	import type { EventHandler } from 'svelte/elements'

	type Props = {
		id: string
	}
	let { id }: Props = $props()

	let el: El|undefined = $state(undefined)

	$effect(() => {
		if (elstore.hasOwnProperty(id)) {
			el = elstore[id]
		} else {
			el = undefined
		}
	})

	const handleClick: EventHandler = (e) => {
		e.preventDefault()
		e.stopPropagation()
		elcurrent.update(id)
	}
</script>

{#if el !== undefined}
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<div
		class={elcurrent.id === id ? `${el.classes} outlined` :  el.classes}
		style={el.styles}
		onclick={handleClick}
	>
		{#each el.children as childId}
			<Self id={childId} />
		{/each}
	</div>
{/if}

<style lang="postcss">
	.outlined {
		@apply outline-dashed outline-zinc-300 outline-1 outline-offset-2;
	}
</style>
