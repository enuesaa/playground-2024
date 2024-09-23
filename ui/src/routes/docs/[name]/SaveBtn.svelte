<script lang="ts">
	import { useUpdateDoc } from '$lib/apidocs'

	export let dirName: string
	export let content: string

	const updateDoc = useUpdateDoc()

	async function handleClick() {
		await $updateDoc.mutateAsync({
			dirName,
			content,
		})
	}
</script>

<button on:click|preventDefault={handleClick}>
	Save
</button>

<style lang="postcss">
	button {
		@apply block text-center rounded-md bg-white py-1 px-2 m-2 hover:bg-graywhite;
	}
</style>
