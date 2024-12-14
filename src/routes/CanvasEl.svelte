<script lang="ts">
	import { createContextMenu, melt } from '@melt-ui/svelte'
	import type { El } from '$lib/el.svelte'
	import Self from './CanvasEl.svelte'

	type Props = {
		el: El
	}
	let { el }: Props = $props()

	const { elements: { menu, item, trigger }} = createContextMenu()
	function add() {
		el.children = [{
			styles: '',
			classes: 'bg-black/30 w-2/3 h-2/3',
			children: [],
		}]
	}	
</script>

<div use:melt={$trigger} class={el.classes} style={el.styles}>
	{#each el.children as child}
		<Self el={child} />
	{/each}
</div>

<ol use:melt={$menu}>
	<li use:melt={$item} on:m-click={add}>Add</li>
	<li use:melt={$item}>Remove</li>
</ol>
  
<style lang="postcss">
	ol {
		@apply w-20 shadow rounded-lg bg-zinc-100 p-1;
	}
	li {
		@apply px-2 py-1 text-zinc-800 text-sm font-semibold;
	}
</style>
