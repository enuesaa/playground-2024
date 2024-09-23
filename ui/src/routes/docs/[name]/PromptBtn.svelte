<script lang="ts">
	import { createDialog, melt } from '@melt-ui/svelte'
	import { runPrompt } from '$lib/apiprompt'

	export let value: string

	const {
		elements: { trigger, overlay, content, title, portalled },
		states: { open }
	} = createDialog({ forceVisible: true })

	const prompt = runPrompt()
	let command = 'echo a'

	async function handleRun() {
		const res = await $prompt.mutateAsync({
			command,
		})

		value += '\n'
		value += '```console\n'
		value += `$ ${command}\n`
		value += `${res.data.output}\n`
		value += '```\n'

		open.set(false)
	}
</script>

<button use:melt={$trigger} class="openbtn">Prompt</button>

{#if $open}
	<div use:melt={$portalled}>
		<div use:melt={$overlay} class="overlay" />

		<div use:melt={$content} class="content">
			<h2 use:melt={$title}>Create New Doc</h2>

			<input type="text" placeholder="title" class="w-full text-lg" bind:value={command} />

			<button on:click={handleRun} class="savebtn">Run</button>
		</div>
	</div>
{/if}

<style lang="postcss">
	.openbtn {
		@apply inline-flex items-center justify-center rounded-xl bg-white;
		@apply px-4 py-3 font-medium leading-none shadow hover:opacity-75;
	}
	.overlay {
		@apply fixed inset-0 z-50 bg-black/50;
	}
	.content {
		@apply fixed left-1/2 top-1/2 z-50 max-h-[85vh] w-[90vw] max-w-[450px];
		@apply -translate-x-1/2 -translate-y-1/2 rounded-xl bg-white p-6 shadow-lg;
	}

	h2 {
		@apply m-0 text-lg font-medium text-black;
	}

	.savebtn {
		@apply inline-flex h-8 items-center justify-center rounded-sm;
		@apply px-4 font-medium leading-none;
	}
</style>
