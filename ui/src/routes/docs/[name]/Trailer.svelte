<script lang="ts">
	import { useUpdateDoc } from '$lib/apidocs'
	import Textarea from '$lib/components/form/Textarea.svelte'
	import FiletreeBtn from './FiletreeBtn.svelte'
	import PromptBtn from './PromptBtn.svelte'
	import SaveBtn from './SaveBtn.svelte'

	export let dirName: string
	export let content: string

	const updateDoc = useUpdateDoc()

	async function save() {
		await $updateDoc.mutateAsync({
			dirName,
			content,
		})
	}
</script>

<FiletreeBtn bind:content={content} />
<PromptBtn bind:content={content} />
<SaveBtn {dirName} {content} />
<Textarea bind:value={content} label='' handleKeyup={save} />
