<script lang="ts">
	import { viewDoc, useUpdateDoc } from '$lib/apidocs'
	import { page } from '$app/stores'
	import Textarea from '$lib/components/form/Textarea.svelte'
	import FiletreeBtn from './FiletreeBtn.svelte'
	import PromptBtn from './PromptBtn.svelte'

	const updateDoc = useUpdateDoc()

	let loading = true
	let content = ''

	const name = $page.params.name
	const doc = viewDoc(name)

	$: if ($doc.isSuccess && loading) {
		content = $doc.data.content
		loading = false
	}

	async function save() {
		await $updateDoc.mutateAsync({
			dirName: name,
			content,
		})
	}
</script>

<button>prompt</button>

<FiletreeBtn bind:value={content} />
<PromptBtn bind:value={content} />
<Textarea bind:value={content} label='' handleKeyup={save} />
