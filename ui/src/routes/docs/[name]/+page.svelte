<script lang="ts">
	import { viewDoc, useUpdateDoc } from '$lib/apidocs'
	import { page } from '$app/stores'
	import Textarea from '$lib/components/form/Textarea.svelte'

	const updateDoc = useUpdateDoc()

	let content = ''

	const name = $page.params.name
	const doc = viewDoc(name)

	$: if ($doc.isSuccess && content === '') {
		content = $doc.data.data.content
	}

	async function save() {
		console.log('save')
		await $updateDoc.mutateAsync({
			dirName: name,
			content,
		})
	}
</script>

<Textarea bind:value={content} label='' handleKeyup={save} />
