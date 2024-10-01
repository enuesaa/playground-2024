<script lang="ts">
	import { DeleteIcon } from 'svelte-feather-icons'
	import { createDialog, melt } from '@melt-ui/svelte'
	const {
	  elements: { trigger, portalled, content: dialogContent, title: dialogTitle, overlay },
	  states: { open: dialogOpen }
	} = createDialog()

	export let slides: string[]
	export let selected: number
	export let content: string

	function handleAdd() {
		slides.splice(selected, 1)
		if (slides.length === selected) {
			slides.push('')
		}
		content = slides[selected]
		dialogOpen.set(false)
	}
</script>

<button use:melt={$trigger}><DeleteIcon /></button>

{#if $dialogOpen}
	<div use:melt={$portalled}>
		<div use:melt={$overlay} class="overlay" />
		<div use:melt={$dialogContent} class="content">
			<h2 use:melt={$dialogTitle}>Are you sure to delete this slide ?</h2>
			<button on:click|preventDefault={handleAdd} class="border rounded-lg mt-2 p-2 block">Delete</button>
		</div>
	</div>
{/if}

<style lang="postcss">
	button {
		@apply text-blackgray;
	}
	.overlay {
		@apply fixed inset-0 z-50 bg-black/50;
	}
	.content {
		@apply fixed left-1/2 top-1/2 z-50 max-h-[85vh] w-[90vw] max-w-[450px];
		@apply -translate-x-1/2 -translate-y-1/2 rounded-xl bg-white p-6 shadow-lg;
	}
</style>
