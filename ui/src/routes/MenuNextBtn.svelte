<script lang="ts">
	import { ArrowRightIcon, PlusIcon } from 'svelte-feather-icons'
	import { createDialog, melt } from '@melt-ui/svelte'
	const {
	  elements: { trigger, portalled, content: dialogContent, title: dialogTitle, overlay },
	  states: { open: dialogOpen }
	} = createDialog()

	export let slides: string[]
	export let selected: number
	export let content: string

	function handleAdd() {
		slides[selected] = content
		slides.push('')
		selected += 1
		content = slides[selected]
		dialogOpen.set(false)
	}
	function handleNext() {
		slides[selected] = content
		selected += 1
		content = slides[selected]
	}
</script>

{#if slides.length === selected + 1}
	<button use:melt={$trigger}><PlusIcon /></button>
{:else}
	<button on:click|preventDefault={handleNext}><ArrowRightIcon /></button>
{/if}

{#if $dialogOpen}
	<div use:melt={$portalled}>
		<div use:melt={$overlay} class="overlay" />
		<div use:melt={$dialogContent} class="content">
			<h2 use:melt={$dialogTitle}>Are you sure to create new slide ?</h2>
			<button on:click|preventDefault={handleAdd} class="border rounded-lg mt-2 p-2">Create</button>
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
