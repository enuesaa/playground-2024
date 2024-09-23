<script lang="ts">
	import { useUpdateDoc } from '$lib/apidocs'
	import Actions from './Actions.svelte'

	export let dirName: string
	export let content: string

	const updateDoc = useUpdateDoc()

	let textarea: HTMLTextAreaElement;

	function updateContent(text: string) {
		const position = textarea.selectionStart
		const before = content.substring(0, position)
        const after = content.substring(position, content.length)
        content = before + text + after
		save()
	}

	function save() {
		$updateDoc.mutate({ dirName, content })
	}
</script>

<Actions {updateContent} {dirName} {content} />

<textarea bind:value={content} bind:this={textarea} on:input={save} />

<style lang="postcss">
	textarea {
		@apply font-normal w-full p-5 block text-black min-h-screen;
		@apply text-base outline-none border-black border bg-graywhite;
	}
</style>
