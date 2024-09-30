<script lang="ts">
	import Slide from './Slide.svelte'
	import { useUpdateDoc } from '$lib/apidoc'

	export let slides: string[]
	let selected = 0
	let content = slides[selected]

	const updateDoc = useUpdateDoc()

	async function handleSave() {
		slides[selected] = content
		await $updateDoc.mutateAsync({slides})
	}

	function handleNext() {
		slides[selected] = content
		selected += 1
		if (slides.length <= selected) {
			slides.push('')
		}
		content = slides[selected]
	}
	function handlePrev() {
		slides[selected] = content
		selected -= 1
		content = slides[selected]
	}
</script>

<button on:click|preventDefault={handleSave}>save</button>
<button on:click|preventDefault={handlePrev}>prev</button>
<button on:click|preventDefault={handleNext}>next</button>

<Slide bind:content={content} />
