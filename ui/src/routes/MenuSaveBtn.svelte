<script lang="ts">
	import { useUpdateDoc } from '$lib/api/doc'
	import { SaveIcon } from 'svelte-feather-icons'

	export let slides: string[]
	export let selected: number
	export let content: string

	let saving: boolean = false

	const updateDoc = useUpdateDoc()

	async function handleSave() {
		saving = true
		slides[selected] = content
		await $updateDoc.mutateAsync({slides})
		setTimeout(() => {
			saving = false
		}, 1000)
	}

	async function handleClick(e: KeyboardEvent) {
		// command + s
		if (e.metaKey && e.key === 's') {
			e.preventDefault()
			await handleSave()
		}
	}
</script>

{#if saving}
	<button disabled class="opacity-35"><SaveIcon /></button>
{:else}
	<button on:click|preventDefault={handleSave}><SaveIcon /></button>
{/if}

<svelte:window on:keydown={handleClick} />

<style lang="postcss">
	button {
		@apply text-blackgray;
	}
</style>
