<script lang="ts">
	export let text: string

	import { createTooltip, melt } from '@melt-ui/svelte'
	import { fade } from 'svelte/transition'
	import { CopyIcon, CheckIcon } from 'svelte-feather-icons'

	const {
		elements: { trigger, content, arrow },
		states: { open },
	} = createTooltip({
		positioning: {
			placement: 'top',
		},
		openDelay: 0,
		closeOnPointerDown: false,
	})

	let checked: boolean = false

	async function copy() {
		await globalThis.navigator.clipboard.writeText(text)
		checked = true
		setTimeout(() => (checked = false), 3000)
	}
</script>

<button type="button" class="absolute right-4 sm:right-8 top-8" use:melt={$trigger} on:click|preventDefault={copy}>
	{#if checked}
		<CheckIcon />
	{:else}
		<CopyIcon />
	{/if}
</button>

{#if $open}
	<div use:melt={$content} transition:fade={{ duration: 100 }} class="z-10 bg-grayer shadow-xl">
		<div use:melt={$arrow}></div>
		<p class="px-4 py-1">Copy</p>
	</div>
{/if}
