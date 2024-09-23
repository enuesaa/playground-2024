<script lang="ts">
	import { lookupFiletree } from '$lib/apifiletree'

	const filetree = lookupFiletree()

	export let value: string

	async function handleClick() {
		const res = await $filetree.mutateAsync()
		value += '\n'
		value += '```\n'
		value += res.tree
		value += '```\n'
	}
</script>

<button on:click={handleClick}>tree</button>
