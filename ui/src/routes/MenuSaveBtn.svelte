<script lang="ts">
	import { useUpdateDoc } from '$lib/api/doc'
	import { SaveIcon } from 'svelte-feather-icons'

	export let slides: string[]
	export let selected: number
	export let content: string

	const updateDoc = useUpdateDoc()

	async function handleClick() {
		slides[selected] = content
		await $updateDoc.mutateAsync({slides})
	}
</script>

<button on:click|preventDefault={handleClick}>
	<SaveIcon />
</button>

<style lang="postcss">
	button {
		@apply bg-blackgray text-gray rounded-md;
	}
</style>
