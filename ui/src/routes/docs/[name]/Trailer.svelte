<script lang="ts">
	import { useUpdateDoc } from '$lib/apidocs'
	import FiletreeBtn from './FiletreeBtn.svelte'
	import PromptBtn from './PromptBtn.svelte'
	import SaveBtn from './SaveBtn.svelte'

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

<FiletreeBtn {updateContent} />
<PromptBtn {updateContent} />
<SaveBtn {dirName} {content} />

<textarea bind:value={content} bind:this={textarea} on:keyup={save} />

<style lang="postcss">
	textarea {
		@apply font-normal w-full rounded px-3 py-2 block h-[80vh] text-black;
		@apply text-lg outline-none border-black border bg-graywhite;
	}
</style>
