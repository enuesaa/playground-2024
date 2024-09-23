<script lang="ts">
	import { lookupFiletree } from '$lib/apifiletree'

	export let updateContent: (text: string) => void

	const filetree = lookupFiletree()

	async function handleClick() {
		const res = await $filetree.mutateAsync()
		let text = '\n'
		text += '```\n'
		text += res.tree
		text += '```\n'
		updateContent(text)
	}
</script>

<button on:click|preventDefault={handleClick}>
	tree
</button>

<style lang="postcss">
	button {
		@apply block text-center rounded-md bg-white py-1 px-2 m-2 hover:bg-graywhite;
	}
</style>
