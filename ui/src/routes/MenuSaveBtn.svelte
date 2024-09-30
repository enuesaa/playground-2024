<script lang="ts">
	import { useUpdateDoc } from '$lib/api/doc'
	import { SaveIcon, ClockIcon } from 'svelte-feather-icons'

	export let slides: string[]
	export let selected: number
	export let content: string

	let saving: boolean = false
	let autoSave: boolean = true

	const updateDoc = useUpdateDoc()

	async function handleSave() {
		saving = true
		slides[selected] = content
		await $updateDoc.mutateAsync({slides})
		setTimeout(() => {
			saving = false
		}, 1000)
		if (autoSave) {
			setTimeout(async() => await handleSave(), 5000)
		}
	}

	async function handleClick(e: KeyboardEvent) {
		// command + s
		if (e.metaKey && e.key === 's') {
			e.preventDefault()
			await handleSave()
		}
	}
	function toggleAutoSave() {
		autoSave = !autoSave
		if (autoSave) {
			setTimeout(async() => await handleSave(), 5000)
		}
	}
</script>

{#if saving}
	<button disabled class="opacity-35"><SaveIcon /></button>
{:else}
	<button on:click|preventDefault={handleSave}><SaveIcon /></button>
{/if}

<button on:click={toggleAutoSave} style={autoSave ? "color: purple;" : "opacity: 0.35;"}><ClockIcon /></button>

<svelte:window on:keydown={handleClick} />

<style lang="postcss">
	button {
		@apply text-blackgray;
	}
</style>
